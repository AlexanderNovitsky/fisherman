package rules_test

import (
	syserrors "errors"
	. "fisherman/internal/rules"
	"fisherman/internal/validation"
	"fisherman/testing/mocks"
	"io/ioutil"
	"testing"

	"github.com/go-errors/errors"
	"github.com/go-git/go-git/v5"
	"github.com/stretchr/testify/assert"
)

func TestAddToIndex_GetPosition(t *testing.T) {
	rule := AddToIndex{}

	assert.Equal(t, PostScripts, rule.GetPosition())
}

func TestAddToIndex_Compile(t *testing.T) {
	rule := AddToIndex{
		Globs: []Glob{
			{Pattern: "{{var1}}", IsRequired: false},
			{Pattern: "data", IsRequired: false},
		},
	}

	rule.Compile(map[string]interface{}{"var1": "VALUE"})

	assert.Equal(t, AddToIndex{
		Globs: []Glob{
			{Pattern: "VALUE", IsRequired: false},
			{Pattern: "data", IsRequired: false},
		},
	}, rule)
}

func TestAddToIndex_Check(t *testing.T) {
	t.Run("not configured rules", func(t *testing.T) {
		rule := AddToIndex{BaseRule: BaseRule{Type: AddToIndexType}}

		err := rule.Check(mocks.NewExecutionContextMock(t), ioutil.Discard)

		assert.NoError(t, err)
	})

	t.Run("added to index correctly", func(t *testing.T) {
		repo := mocks.NewRepositoryMock(t).
			AddGlobMock.When("glob1/*.go").Then(nil).
			AddGlobMock.When("*.css").Then(nil).
			AddGlobMock.When("mocks").Then(nil)

		ctx := mocks.NewExecutionContextMock(t).RepositoryMock.Return(repo)

		rule := AddToIndex{
			Globs: []Glob{
				{"glob1/*.go", true},
				{"*.css", true},
				{"mocks", true},
			},
		}

		err := rule.Check(ctx, ioutil.Discard)

		assert.NoError(t, err)
	})

	t.Run("failed to add files", func(t *testing.T) {
		repo := mocks.NewRepositoryMock(t).
			AddGlobMock.When("glob1/*.go").Then(nil).
			AddGlobMock.When("*.css").Then(syserrors.New("testError")).
			AddGlobMock.When("mocks").Then(nil)

		ctx := mocks.NewExecutionContextMock(t).RepositoryMock.Return(repo)

		rule := AddToIndex{
			Globs: []Glob{
				{"glob1/*.go", true},
				{"*.css", true},
				{"mocks", true},
			},
		}
		err := rule.Check(ctx, ioutil.Discard)

		assert.EqualError(t, err, "failed to add files matching pattern '*.css' to the index: testError")
		assert.IsType(t, &errors.Error{}, err)
	})

	t.Run("failed to add files in index with optional flag", func(t *testing.T) {
		repo := mocks.NewRepositoryMock(t).
			AddGlobMock.When("glob1/*.go").Then(nil).
			AddGlobMock.When("*.css").Then(git.ErrGlobNoMatches).
			AddGlobMock.When("mocks").Then(nil)

		ctx := mocks.NewExecutionContextMock(t).RepositoryMock.Return(repo)

		tests := []struct {
			name       string
			isRequired bool
		}{
			{name: "optional is true", isRequired: false},
			{name: "optional is false", isRequired: true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				rule := AddToIndex{
					BaseRule: BaseRule{Type: AddToIndexType},
					Globs: []Glob{
						{"glob1/*.go", tt.isRequired},
						{"*.css", tt.isRequired},
						{"mocks", tt.isRequired},
					},
				}

				err := rule.Check(ctx, ioutil.Discard)

				if !tt.isRequired {
					assert.NoError(t, err)
				} else {
					assert.EqualError(t, err, errorMessage(AddToIndexType, "can't add files matching pattern *.css"))
					assert.IsType(t, &validation.Error{}, err)
				}
			})
		}
	})
}
