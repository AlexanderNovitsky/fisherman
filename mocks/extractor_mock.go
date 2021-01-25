package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i fisherman/internal/configcompiler.Extractor -o ./mocks\extractor_mock.go

import (
	"fisherman/configuration"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ExtractorMock implements configcompiler.Extractor
type ExtractorMock struct {
	t minimock.Tester

	funcVariables          func(section configuration.VariablesConfig) (m1 map[string]interface{}, err error)
	inspectFuncVariables   func(section configuration.VariablesConfig)
	afterVariablesCounter  uint64
	beforeVariablesCounter uint64
	VariablesMock          mExtractorMockVariables
}

// NewExtractorMock returns a mock for configcompiler.Extractor
func NewExtractorMock(t minimock.Tester) *ExtractorMock {
	m := &ExtractorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.VariablesMock = mExtractorMockVariables{mock: m}
	m.VariablesMock.callArgs = []*ExtractorMockVariablesParams{}

	return m
}

type mExtractorMockVariables struct {
	mock               *ExtractorMock
	defaultExpectation *ExtractorMockVariablesExpectation
	expectations       []*ExtractorMockVariablesExpectation

	callArgs []*ExtractorMockVariablesParams
	mutex    sync.RWMutex
}

// ExtractorMockVariablesExpectation specifies expectation struct of the Extractor.Variables
type ExtractorMockVariablesExpectation struct {
	mock    *ExtractorMock
	params  *ExtractorMockVariablesParams
	results *ExtractorMockVariablesResults
	Counter uint64
}

// ExtractorMockVariablesParams contains parameters of the Extractor.Variables
type ExtractorMockVariablesParams struct {
	section configuration.VariablesConfig
}

// ExtractorMockVariablesResults contains results of the Extractor.Variables
type ExtractorMockVariablesResults struct {
	m1  map[string]interface{}
	err error
}

// Expect sets up expected params for Extractor.Variables
func (mmVariables *mExtractorMockVariables) Expect(section configuration.VariablesConfig) *mExtractorMockVariables {
	if mmVariables.mock.funcVariables != nil {
		mmVariables.mock.t.Fatalf("ExtractorMock.Variables mock is already set by Set")
	}

	if mmVariables.defaultExpectation == nil {
		mmVariables.defaultExpectation = &ExtractorMockVariablesExpectation{}
	}

	mmVariables.defaultExpectation.params = &ExtractorMockVariablesParams{section}
	for _, e := range mmVariables.expectations {
		if minimock.Equal(e.params, mmVariables.defaultExpectation.params) {
			mmVariables.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmVariables.defaultExpectation.params)
		}
	}

	return mmVariables
}

// Inspect accepts an inspector function that has same arguments as the Extractor.Variables
func (mmVariables *mExtractorMockVariables) Inspect(f func(section configuration.VariablesConfig)) *mExtractorMockVariables {
	if mmVariables.mock.inspectFuncVariables != nil {
		mmVariables.mock.t.Fatalf("Inspect function is already set for ExtractorMock.Variables")
	}

	mmVariables.mock.inspectFuncVariables = f

	return mmVariables
}

// Return sets up results that will be returned by Extractor.Variables
func (mmVariables *mExtractorMockVariables) Return(m1 map[string]interface{}, err error) *ExtractorMock {
	if mmVariables.mock.funcVariables != nil {
		mmVariables.mock.t.Fatalf("ExtractorMock.Variables mock is already set by Set")
	}

	if mmVariables.defaultExpectation == nil {
		mmVariables.defaultExpectation = &ExtractorMockVariablesExpectation{mock: mmVariables.mock}
	}
	mmVariables.defaultExpectation.results = &ExtractorMockVariablesResults{m1, err}
	return mmVariables.mock
}

//Set uses given function f to mock the Extractor.Variables method
func (mmVariables *mExtractorMockVariables) Set(f func(section configuration.VariablesConfig) (m1 map[string]interface{}, err error)) *ExtractorMock {
	if mmVariables.defaultExpectation != nil {
		mmVariables.mock.t.Fatalf("Default expectation is already set for the Extractor.Variables method")
	}

	if len(mmVariables.expectations) > 0 {
		mmVariables.mock.t.Fatalf("Some expectations are already set for the Extractor.Variables method")
	}

	mmVariables.mock.funcVariables = f
	return mmVariables.mock
}

// When sets expectation for the Extractor.Variables which will trigger the result defined by the following
// Then helper
func (mmVariables *mExtractorMockVariables) When(section configuration.VariablesConfig) *ExtractorMockVariablesExpectation {
	if mmVariables.mock.funcVariables != nil {
		mmVariables.mock.t.Fatalf("ExtractorMock.Variables mock is already set by Set")
	}

	expectation := &ExtractorMockVariablesExpectation{
		mock:   mmVariables.mock,
		params: &ExtractorMockVariablesParams{section},
	}
	mmVariables.expectations = append(mmVariables.expectations, expectation)
	return expectation
}

// Then sets up Extractor.Variables return parameters for the expectation previously defined by the When method
func (e *ExtractorMockVariablesExpectation) Then(m1 map[string]interface{}, err error) *ExtractorMock {
	e.results = &ExtractorMockVariablesResults{m1, err}
	return e.mock
}

// Variables implements configcompiler.Extractor
func (mmVariables *ExtractorMock) Variables(section configuration.VariablesConfig) (m1 map[string]interface{}, err error) {
	mm_atomic.AddUint64(&mmVariables.beforeVariablesCounter, 1)
	defer mm_atomic.AddUint64(&mmVariables.afterVariablesCounter, 1)

	if mmVariables.inspectFuncVariables != nil {
		mmVariables.inspectFuncVariables(section)
	}

	mm_params := &ExtractorMockVariablesParams{section}

	// Record call args
	mmVariables.VariablesMock.mutex.Lock()
	mmVariables.VariablesMock.callArgs = append(mmVariables.VariablesMock.callArgs, mm_params)
	mmVariables.VariablesMock.mutex.Unlock()

	for _, e := range mmVariables.VariablesMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmVariables.VariablesMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmVariables.VariablesMock.defaultExpectation.Counter, 1)
		mm_want := mmVariables.VariablesMock.defaultExpectation.params
		mm_got := ExtractorMockVariablesParams{section}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmVariables.t.Errorf("ExtractorMock.Variables got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmVariables.VariablesMock.defaultExpectation.results
		if mm_results == nil {
			mmVariables.t.Fatal("No results are set for the ExtractorMock.Variables")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmVariables.funcVariables != nil {
		return mmVariables.funcVariables(section)
	}
	mmVariables.t.Fatalf("Unexpected call to ExtractorMock.Variables. %v", section)
	return
}

// VariablesAfterCounter returns a count of finished ExtractorMock.Variables invocations
func (mmVariables *ExtractorMock) VariablesAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVariables.afterVariablesCounter)
}

// VariablesBeforeCounter returns a count of ExtractorMock.Variables invocations
func (mmVariables *ExtractorMock) VariablesBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVariables.beforeVariablesCounter)
}

// Calls returns a list of arguments used in each call to ExtractorMock.Variables.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmVariables *mExtractorMockVariables) Calls() []*ExtractorMockVariablesParams {
	mmVariables.mutex.RLock()

	argCopy := make([]*ExtractorMockVariablesParams, len(mmVariables.callArgs))
	copy(argCopy, mmVariables.callArgs)

	mmVariables.mutex.RUnlock()

	return argCopy
}

// MinimockVariablesDone returns true if the count of the Variables invocations corresponds
// the number of defined expectations
func (m *ExtractorMock) MinimockVariablesDone() bool {
	for _, e := range m.VariablesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VariablesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVariablesCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVariables != nil && mm_atomic.LoadUint64(&m.afterVariablesCounter) < 1 {
		return false
	}
	return true
}

// MinimockVariablesInspect logs each unmet expectation
func (m *ExtractorMock) MinimockVariablesInspect() {
	for _, e := range m.VariablesMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ExtractorMock.Variables with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VariablesMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVariablesCounter) < 1 {
		if m.VariablesMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ExtractorMock.Variables")
		} else {
			m.t.Errorf("Expected call to ExtractorMock.Variables with params: %#v", *m.VariablesMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVariables != nil && mm_atomic.LoadUint64(&m.afterVariablesCounter) < 1 {
		m.t.Error("Expected call to ExtractorMock.Variables")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ExtractorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockVariablesInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ExtractorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ExtractorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockVariablesDone()
}
