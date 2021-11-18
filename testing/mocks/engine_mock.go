package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i fisherman/internal/expression.Engine -o ./testing/mocks/engine_mock.go -n EngineMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// EngineMock implements expression.Engine
type EngineMock struct {
	t minimock.Tester

	funcEval          func(expression string, vars map[string]interface{}) (b1 bool, err error)
	inspectFuncEval   func(expression string, vars map[string]interface{})
	afterEvalCounter  uint64
	beforeEvalCounter uint64
	EvalMock          mEngineMockEval
}

// NewEngineMock returns a mock for expression.Engine
func NewEngineMock(t minimock.Tester) *EngineMock {
	m := &EngineMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.EvalMock = mEngineMockEval{mock: m}
	m.EvalMock.callArgs = []*EngineMockEvalParams{}

	return m
}

type mEngineMockEval struct {
	mock               *EngineMock
	defaultExpectation *EngineMockEvalExpectation
	expectations       []*EngineMockEvalExpectation

	callArgs []*EngineMockEvalParams
	mutex    sync.RWMutex
}

// EngineMockEvalExpectation specifies expectation struct of the Engine.Eval
type EngineMockEvalExpectation struct {
	mock    *EngineMock
	params  *EngineMockEvalParams
	results *EngineMockEvalResults
	Counter uint64
}

// EngineMockEvalParams contains parameters of the Engine.Eval
type EngineMockEvalParams struct {
	expression string
	vars       map[string]interface{}
}

// EngineMockEvalResults contains results of the Engine.Eval
type EngineMockEvalResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for Engine.Eval
func (mmEval *mEngineMockEval) Expect(expression string, vars map[string]interface{}) *mEngineMockEval {
	if mmEval.mock.funcEval != nil {
		mmEval.mock.t.Fatalf("EngineMock.Eval mock is already set by Set")
	}

	if mmEval.defaultExpectation == nil {
		mmEval.defaultExpectation = &EngineMockEvalExpectation{}
	}

	mmEval.defaultExpectation.params = &EngineMockEvalParams{expression, vars}
	for _, e := range mmEval.expectations {
		if minimock.Equal(e.params, mmEval.defaultExpectation.params) {
			mmEval.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmEval.defaultExpectation.params)
		}
	}

	return mmEval
}

// Inspect accepts an inspector function that has same arguments as the Engine.Eval
func (mmEval *mEngineMockEval) Inspect(f func(expression string, vars map[string]interface{})) *mEngineMockEval {
	if mmEval.mock.inspectFuncEval != nil {
		mmEval.mock.t.Fatalf("Inspect function is already set for EngineMock.Eval")
	}

	mmEval.mock.inspectFuncEval = f

	return mmEval
}

// Return sets up results that will be returned by Engine.Eval
func (mmEval *mEngineMockEval) Return(b1 bool, err error) *EngineMock {
	if mmEval.mock.funcEval != nil {
		mmEval.mock.t.Fatalf("EngineMock.Eval mock is already set by Set")
	}

	if mmEval.defaultExpectation == nil {
		mmEval.defaultExpectation = &EngineMockEvalExpectation{mock: mmEval.mock}
	}
	mmEval.defaultExpectation.results = &EngineMockEvalResults{b1, err}
	return mmEval.mock
}

//Set uses given function f to mock the Engine.Eval method
func (mmEval *mEngineMockEval) Set(f func(expression string, vars map[string]interface{}) (b1 bool, err error)) *EngineMock {
	if mmEval.defaultExpectation != nil {
		mmEval.mock.t.Fatalf("Default expectation is already set for the Engine.Eval method")
	}

	if len(mmEval.expectations) > 0 {
		mmEval.mock.t.Fatalf("Some expectations are already set for the Engine.Eval method")
	}

	mmEval.mock.funcEval = f
	return mmEval.mock
}

// When sets expectation for the Engine.Eval which will trigger the result defined by the following
// Then helper
func (mmEval *mEngineMockEval) When(expression string, vars map[string]interface{}) *EngineMockEvalExpectation {
	if mmEval.mock.funcEval != nil {
		mmEval.mock.t.Fatalf("EngineMock.Eval mock is already set by Set")
	}

	expectation := &EngineMockEvalExpectation{
		mock:   mmEval.mock,
		params: &EngineMockEvalParams{expression, vars},
	}
	mmEval.expectations = append(mmEval.expectations, expectation)
	return expectation
}

// Then sets up Engine.Eval return parameters for the expectation previously defined by the When method
func (e *EngineMockEvalExpectation) Then(b1 bool, err error) *EngineMock {
	e.results = &EngineMockEvalResults{b1, err}
	return e.mock
}

// Eval implements expression.Engine
func (mmEval *EngineMock) Eval(expression string, vars map[string]interface{}) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmEval.beforeEvalCounter, 1)
	defer mm_atomic.AddUint64(&mmEval.afterEvalCounter, 1)

	if mmEval.inspectFuncEval != nil {
		mmEval.inspectFuncEval(expression, vars)
	}

	mm_params := &EngineMockEvalParams{expression, vars}

	// Record call args
	mmEval.EvalMock.mutex.Lock()
	mmEval.EvalMock.callArgs = append(mmEval.EvalMock.callArgs, mm_params)
	mmEval.EvalMock.mutex.Unlock()

	for _, e := range mmEval.EvalMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmEval.EvalMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEval.EvalMock.defaultExpectation.Counter, 1)
		mm_want := mmEval.EvalMock.defaultExpectation.params
		mm_got := EngineMockEvalParams{expression, vars}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmEval.t.Errorf("EngineMock.Eval got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmEval.EvalMock.defaultExpectation.results
		if mm_results == nil {
			mmEval.t.Fatal("No results are set for the EngineMock.Eval")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmEval.funcEval != nil {
		return mmEval.funcEval(expression, vars)
	}
	mmEval.t.Fatalf("Unexpected call to EngineMock.Eval. %v %v", expression, vars)
	return
}

// EvalAfterCounter returns a count of finished EngineMock.Eval invocations
func (mmEval *EngineMock) EvalAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEval.afterEvalCounter)
}

// EvalBeforeCounter returns a count of EngineMock.Eval invocations
func (mmEval *EngineMock) EvalBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEval.beforeEvalCounter)
}

// Calls returns a list of arguments used in each call to EngineMock.Eval.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmEval *mEngineMockEval) Calls() []*EngineMockEvalParams {
	mmEval.mutex.RLock()

	argCopy := make([]*EngineMockEvalParams, len(mmEval.callArgs))
	copy(argCopy, mmEval.callArgs)

	mmEval.mutex.RUnlock()

	return argCopy
}

// MinimockEvalDone returns true if the count of the Eval invocations corresponds
// the number of defined expectations
func (m *EngineMock) MinimockEvalDone() bool {
	for _, e := range m.EvalMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EvalMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEvalCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEval != nil && mm_atomic.LoadUint64(&m.afterEvalCounter) < 1 {
		return false
	}
	return true
}

// MinimockEvalInspect logs each unmet expectation
func (m *EngineMock) MinimockEvalInspect() {
	for _, e := range m.EvalMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to EngineMock.Eval with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EvalMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEvalCounter) < 1 {
		if m.EvalMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to EngineMock.Eval")
		} else {
			m.t.Errorf("Expected call to EngineMock.Eval with params: %#v", *m.EvalMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEval != nil && mm_atomic.LoadUint64(&m.afterEvalCounter) < 1 {
		m.t.Error("Expected call to EngineMock.Eval")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *EngineMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockEvalInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *EngineMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *EngineMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockEvalDone()
}
