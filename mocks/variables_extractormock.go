package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i fisherman/config/hooks.VariablesExtractor -o ./mocks\variables_extractormock.go

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// VariablesExtractorMock implements hooks.VariablesExtractor
type VariablesExtractorMock struct {
	t minimock.Tester

	funcGetFromBranch          func(branchName string) (m1 map[string]interface{}, err error)
	inspectFuncGetFromBranch   func(branchName string)
	afterGetFromBranchCounter  uint64
	beforeGetFromBranchCounter uint64
	GetFromBranchMock          mVariablesExtractorMockGetFromBranch

	funcGetFromTag          func(tag string) (m1 map[string]interface{}, err error)
	inspectFuncGetFromTag   func(tag string)
	afterGetFromTagCounter  uint64
	beforeGetFromTagCounter uint64
	GetFromTagMock          mVariablesExtractorMockGetFromTag
}

// NewVariablesExtractorMock returns a mock for hooks.VariablesExtractor
func NewVariablesExtractorMock(t minimock.Tester) *VariablesExtractorMock {
	m := &VariablesExtractorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetFromBranchMock = mVariablesExtractorMockGetFromBranch{mock: m}
	m.GetFromBranchMock.callArgs = []*VariablesExtractorMockGetFromBranchParams{}

	m.GetFromTagMock = mVariablesExtractorMockGetFromTag{mock: m}
	m.GetFromTagMock.callArgs = []*VariablesExtractorMockGetFromTagParams{}

	return m
}

type mVariablesExtractorMockGetFromBranch struct {
	mock               *VariablesExtractorMock
	defaultExpectation *VariablesExtractorMockGetFromBranchExpectation
	expectations       []*VariablesExtractorMockGetFromBranchExpectation

	callArgs []*VariablesExtractorMockGetFromBranchParams
	mutex    sync.RWMutex
}

// VariablesExtractorMockGetFromBranchExpectation specifies expectation struct of the VariablesExtractor.GetFromBranch
type VariablesExtractorMockGetFromBranchExpectation struct {
	mock    *VariablesExtractorMock
	params  *VariablesExtractorMockGetFromBranchParams
	results *VariablesExtractorMockGetFromBranchResults
	Counter uint64
}

// VariablesExtractorMockGetFromBranchParams contains parameters of the VariablesExtractor.GetFromBranch
type VariablesExtractorMockGetFromBranchParams struct {
	branchName string
}

// VariablesExtractorMockGetFromBranchResults contains results of the VariablesExtractor.GetFromBranch
type VariablesExtractorMockGetFromBranchResults struct {
	m1  map[string]interface{}
	err error
}

// Expect sets up expected params for VariablesExtractor.GetFromBranch
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) Expect(branchName string) *mVariablesExtractorMockGetFromBranch {
	if mmGetFromBranch.mock.funcGetFromBranch != nil {
		mmGetFromBranch.mock.t.Fatalf("VariablesExtractorMock.GetFromBranch mock is already set by Set")
	}

	if mmGetFromBranch.defaultExpectation == nil {
		mmGetFromBranch.defaultExpectation = &VariablesExtractorMockGetFromBranchExpectation{}
	}

	mmGetFromBranch.defaultExpectation.params = &VariablesExtractorMockGetFromBranchParams{branchName}
	for _, e := range mmGetFromBranch.expectations {
		if minimock.Equal(e.params, mmGetFromBranch.defaultExpectation.params) {
			mmGetFromBranch.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetFromBranch.defaultExpectation.params)
		}
	}

	return mmGetFromBranch
}

// Inspect accepts an inspector function that has same arguments as the VariablesExtractor.GetFromBranch
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) Inspect(f func(branchName string)) *mVariablesExtractorMockGetFromBranch {
	if mmGetFromBranch.mock.inspectFuncGetFromBranch != nil {
		mmGetFromBranch.mock.t.Fatalf("Inspect function is already set for VariablesExtractorMock.GetFromBranch")
	}

	mmGetFromBranch.mock.inspectFuncGetFromBranch = f

	return mmGetFromBranch
}

// Return sets up results that will be returned by VariablesExtractor.GetFromBranch
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) Return(m1 map[string]interface{}, err error) *VariablesExtractorMock {
	if mmGetFromBranch.mock.funcGetFromBranch != nil {
		mmGetFromBranch.mock.t.Fatalf("VariablesExtractorMock.GetFromBranch mock is already set by Set")
	}

	if mmGetFromBranch.defaultExpectation == nil {
		mmGetFromBranch.defaultExpectation = &VariablesExtractorMockGetFromBranchExpectation{mock: mmGetFromBranch.mock}
	}
	mmGetFromBranch.defaultExpectation.results = &VariablesExtractorMockGetFromBranchResults{m1, err}
	return mmGetFromBranch.mock
}

//Set uses given function f to mock the VariablesExtractor.GetFromBranch method
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) Set(f func(branchName string) (m1 map[string]interface{}, err error)) *VariablesExtractorMock {
	if mmGetFromBranch.defaultExpectation != nil {
		mmGetFromBranch.mock.t.Fatalf("Default expectation is already set for the VariablesExtractor.GetFromBranch method")
	}

	if len(mmGetFromBranch.expectations) > 0 {
		mmGetFromBranch.mock.t.Fatalf("Some expectations are already set for the VariablesExtractor.GetFromBranch method")
	}

	mmGetFromBranch.mock.funcGetFromBranch = f
	return mmGetFromBranch.mock
}

// When sets expectation for the VariablesExtractor.GetFromBranch which will trigger the result defined by the following
// Then helper
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) When(branchName string) *VariablesExtractorMockGetFromBranchExpectation {
	if mmGetFromBranch.mock.funcGetFromBranch != nil {
		mmGetFromBranch.mock.t.Fatalf("VariablesExtractorMock.GetFromBranch mock is already set by Set")
	}

	expectation := &VariablesExtractorMockGetFromBranchExpectation{
		mock:   mmGetFromBranch.mock,
		params: &VariablesExtractorMockGetFromBranchParams{branchName},
	}
	mmGetFromBranch.expectations = append(mmGetFromBranch.expectations, expectation)
	return expectation
}

// Then sets up VariablesExtractor.GetFromBranch return parameters for the expectation previously defined by the When method
func (e *VariablesExtractorMockGetFromBranchExpectation) Then(m1 map[string]interface{}, err error) *VariablesExtractorMock {
	e.results = &VariablesExtractorMockGetFromBranchResults{m1, err}
	return e.mock
}

// GetFromBranch implements hooks.VariablesExtractor
func (mmGetFromBranch *VariablesExtractorMock) GetFromBranch(branchName string) (m1 map[string]interface{}, err error) {
	mm_atomic.AddUint64(&mmGetFromBranch.beforeGetFromBranchCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFromBranch.afterGetFromBranchCounter, 1)

	if mmGetFromBranch.inspectFuncGetFromBranch != nil {
		mmGetFromBranch.inspectFuncGetFromBranch(branchName)
	}

	mm_params := &VariablesExtractorMockGetFromBranchParams{branchName}

	// Record call args
	mmGetFromBranch.GetFromBranchMock.mutex.Lock()
	mmGetFromBranch.GetFromBranchMock.callArgs = append(mmGetFromBranch.GetFromBranchMock.callArgs, mm_params)
	mmGetFromBranch.GetFromBranchMock.mutex.Unlock()

	for _, e := range mmGetFromBranch.GetFromBranchMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmGetFromBranch.GetFromBranchMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFromBranch.GetFromBranchMock.defaultExpectation.Counter, 1)
		mm_want := mmGetFromBranch.GetFromBranchMock.defaultExpectation.params
		mm_got := VariablesExtractorMockGetFromBranchParams{branchName}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetFromBranch.t.Errorf("VariablesExtractorMock.GetFromBranch got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetFromBranch.GetFromBranchMock.defaultExpectation.results
		if mm_results == nil {
			mmGetFromBranch.t.Fatal("No results are set for the VariablesExtractorMock.GetFromBranch")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmGetFromBranch.funcGetFromBranch != nil {
		return mmGetFromBranch.funcGetFromBranch(branchName)
	}
	mmGetFromBranch.t.Fatalf("Unexpected call to VariablesExtractorMock.GetFromBranch. %v", branchName)
	return
}

// GetFromBranchAfterCounter returns a count of finished VariablesExtractorMock.GetFromBranch invocations
func (mmGetFromBranch *VariablesExtractorMock) GetFromBranchAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFromBranch.afterGetFromBranchCounter)
}

// GetFromBranchBeforeCounter returns a count of VariablesExtractorMock.GetFromBranch invocations
func (mmGetFromBranch *VariablesExtractorMock) GetFromBranchBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFromBranch.beforeGetFromBranchCounter)
}

// Calls returns a list of arguments used in each call to VariablesExtractorMock.GetFromBranch.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetFromBranch *mVariablesExtractorMockGetFromBranch) Calls() []*VariablesExtractorMockGetFromBranchParams {
	mmGetFromBranch.mutex.RLock()

	argCopy := make([]*VariablesExtractorMockGetFromBranchParams, len(mmGetFromBranch.callArgs))
	copy(argCopy, mmGetFromBranch.callArgs)

	mmGetFromBranch.mutex.RUnlock()

	return argCopy
}

// MinimockGetFromBranchDone returns true if the count of the GetFromBranch invocations corresponds
// the number of defined expectations
func (m *VariablesExtractorMock) MinimockGetFromBranchDone() bool {
	for _, e := range m.GetFromBranchMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFromBranchMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFromBranchCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFromBranch != nil && mm_atomic.LoadUint64(&m.afterGetFromBranchCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFromBranchInspect logs each unmet expectation
func (m *VariablesExtractorMock) MinimockGetFromBranchInspect() {
	for _, e := range m.GetFromBranchMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to VariablesExtractorMock.GetFromBranch with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFromBranchMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFromBranchCounter) < 1 {
		if m.GetFromBranchMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to VariablesExtractorMock.GetFromBranch")
		} else {
			m.t.Errorf("Expected call to VariablesExtractorMock.GetFromBranch with params: %#v", *m.GetFromBranchMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFromBranch != nil && mm_atomic.LoadUint64(&m.afterGetFromBranchCounter) < 1 {
		m.t.Error("Expected call to VariablesExtractorMock.GetFromBranch")
	}
}

type mVariablesExtractorMockGetFromTag struct {
	mock               *VariablesExtractorMock
	defaultExpectation *VariablesExtractorMockGetFromTagExpectation
	expectations       []*VariablesExtractorMockGetFromTagExpectation

	callArgs []*VariablesExtractorMockGetFromTagParams
	mutex    sync.RWMutex
}

// VariablesExtractorMockGetFromTagExpectation specifies expectation struct of the VariablesExtractor.GetFromTag
type VariablesExtractorMockGetFromTagExpectation struct {
	mock    *VariablesExtractorMock
	params  *VariablesExtractorMockGetFromTagParams
	results *VariablesExtractorMockGetFromTagResults
	Counter uint64
}

// VariablesExtractorMockGetFromTagParams contains parameters of the VariablesExtractor.GetFromTag
type VariablesExtractorMockGetFromTagParams struct {
	tag string
}

// VariablesExtractorMockGetFromTagResults contains results of the VariablesExtractor.GetFromTag
type VariablesExtractorMockGetFromTagResults struct {
	m1  map[string]interface{}
	err error
}

// Expect sets up expected params for VariablesExtractor.GetFromTag
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) Expect(tag string) *mVariablesExtractorMockGetFromTag {
	if mmGetFromTag.mock.funcGetFromTag != nil {
		mmGetFromTag.mock.t.Fatalf("VariablesExtractorMock.GetFromTag mock is already set by Set")
	}

	if mmGetFromTag.defaultExpectation == nil {
		mmGetFromTag.defaultExpectation = &VariablesExtractorMockGetFromTagExpectation{}
	}

	mmGetFromTag.defaultExpectation.params = &VariablesExtractorMockGetFromTagParams{tag}
	for _, e := range mmGetFromTag.expectations {
		if minimock.Equal(e.params, mmGetFromTag.defaultExpectation.params) {
			mmGetFromTag.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetFromTag.defaultExpectation.params)
		}
	}

	return mmGetFromTag
}

// Inspect accepts an inspector function that has same arguments as the VariablesExtractor.GetFromTag
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) Inspect(f func(tag string)) *mVariablesExtractorMockGetFromTag {
	if mmGetFromTag.mock.inspectFuncGetFromTag != nil {
		mmGetFromTag.mock.t.Fatalf("Inspect function is already set for VariablesExtractorMock.GetFromTag")
	}

	mmGetFromTag.mock.inspectFuncGetFromTag = f

	return mmGetFromTag
}

// Return sets up results that will be returned by VariablesExtractor.GetFromTag
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) Return(m1 map[string]interface{}, err error) *VariablesExtractorMock {
	if mmGetFromTag.mock.funcGetFromTag != nil {
		mmGetFromTag.mock.t.Fatalf("VariablesExtractorMock.GetFromTag mock is already set by Set")
	}

	if mmGetFromTag.defaultExpectation == nil {
		mmGetFromTag.defaultExpectation = &VariablesExtractorMockGetFromTagExpectation{mock: mmGetFromTag.mock}
	}
	mmGetFromTag.defaultExpectation.results = &VariablesExtractorMockGetFromTagResults{m1, err}
	return mmGetFromTag.mock
}

//Set uses given function f to mock the VariablesExtractor.GetFromTag method
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) Set(f func(tag string) (m1 map[string]interface{}, err error)) *VariablesExtractorMock {
	if mmGetFromTag.defaultExpectation != nil {
		mmGetFromTag.mock.t.Fatalf("Default expectation is already set for the VariablesExtractor.GetFromTag method")
	}

	if len(mmGetFromTag.expectations) > 0 {
		mmGetFromTag.mock.t.Fatalf("Some expectations are already set for the VariablesExtractor.GetFromTag method")
	}

	mmGetFromTag.mock.funcGetFromTag = f
	return mmGetFromTag.mock
}

// When sets expectation for the VariablesExtractor.GetFromTag which will trigger the result defined by the following
// Then helper
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) When(tag string) *VariablesExtractorMockGetFromTagExpectation {
	if mmGetFromTag.mock.funcGetFromTag != nil {
		mmGetFromTag.mock.t.Fatalf("VariablesExtractorMock.GetFromTag mock is already set by Set")
	}

	expectation := &VariablesExtractorMockGetFromTagExpectation{
		mock:   mmGetFromTag.mock,
		params: &VariablesExtractorMockGetFromTagParams{tag},
	}
	mmGetFromTag.expectations = append(mmGetFromTag.expectations, expectation)
	return expectation
}

// Then sets up VariablesExtractor.GetFromTag return parameters for the expectation previously defined by the When method
func (e *VariablesExtractorMockGetFromTagExpectation) Then(m1 map[string]interface{}, err error) *VariablesExtractorMock {
	e.results = &VariablesExtractorMockGetFromTagResults{m1, err}
	return e.mock
}

// GetFromTag implements hooks.VariablesExtractor
func (mmGetFromTag *VariablesExtractorMock) GetFromTag(tag string) (m1 map[string]interface{}, err error) {
	mm_atomic.AddUint64(&mmGetFromTag.beforeGetFromTagCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFromTag.afterGetFromTagCounter, 1)

	if mmGetFromTag.inspectFuncGetFromTag != nil {
		mmGetFromTag.inspectFuncGetFromTag(tag)
	}

	mm_params := &VariablesExtractorMockGetFromTagParams{tag}

	// Record call args
	mmGetFromTag.GetFromTagMock.mutex.Lock()
	mmGetFromTag.GetFromTagMock.callArgs = append(mmGetFromTag.GetFromTagMock.callArgs, mm_params)
	mmGetFromTag.GetFromTagMock.mutex.Unlock()

	for _, e := range mmGetFromTag.GetFromTagMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.m1, e.results.err
		}
	}

	if mmGetFromTag.GetFromTagMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFromTag.GetFromTagMock.defaultExpectation.Counter, 1)
		mm_want := mmGetFromTag.GetFromTagMock.defaultExpectation.params
		mm_got := VariablesExtractorMockGetFromTagParams{tag}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetFromTag.t.Errorf("VariablesExtractorMock.GetFromTag got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetFromTag.GetFromTagMock.defaultExpectation.results
		if mm_results == nil {
			mmGetFromTag.t.Fatal("No results are set for the VariablesExtractorMock.GetFromTag")
		}
		return (*mm_results).m1, (*mm_results).err
	}
	if mmGetFromTag.funcGetFromTag != nil {
		return mmGetFromTag.funcGetFromTag(tag)
	}
	mmGetFromTag.t.Fatalf("Unexpected call to VariablesExtractorMock.GetFromTag. %v", tag)
	return
}

// GetFromTagAfterCounter returns a count of finished VariablesExtractorMock.GetFromTag invocations
func (mmGetFromTag *VariablesExtractorMock) GetFromTagAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFromTag.afterGetFromTagCounter)
}

// GetFromTagBeforeCounter returns a count of VariablesExtractorMock.GetFromTag invocations
func (mmGetFromTag *VariablesExtractorMock) GetFromTagBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFromTag.beforeGetFromTagCounter)
}

// Calls returns a list of arguments used in each call to VariablesExtractorMock.GetFromTag.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetFromTag *mVariablesExtractorMockGetFromTag) Calls() []*VariablesExtractorMockGetFromTagParams {
	mmGetFromTag.mutex.RLock()

	argCopy := make([]*VariablesExtractorMockGetFromTagParams, len(mmGetFromTag.callArgs))
	copy(argCopy, mmGetFromTag.callArgs)

	mmGetFromTag.mutex.RUnlock()

	return argCopy
}

// MinimockGetFromTagDone returns true if the count of the GetFromTag invocations corresponds
// the number of defined expectations
func (m *VariablesExtractorMock) MinimockGetFromTagDone() bool {
	for _, e := range m.GetFromTagMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFromTagMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFromTagCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFromTag != nil && mm_atomic.LoadUint64(&m.afterGetFromTagCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFromTagInspect logs each unmet expectation
func (m *VariablesExtractorMock) MinimockGetFromTagInspect() {
	for _, e := range m.GetFromTagMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to VariablesExtractorMock.GetFromTag with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFromTagMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFromTagCounter) < 1 {
		if m.GetFromTagMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to VariablesExtractorMock.GetFromTag")
		} else {
			m.t.Errorf("Expected call to VariablesExtractorMock.GetFromTag with params: %#v", *m.GetFromTagMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFromTag != nil && mm_atomic.LoadUint64(&m.afterGetFromTagCounter) < 1 {
		m.t.Error("Expected call to VariablesExtractorMock.GetFromTag")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *VariablesExtractorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetFromBranchInspect()

		m.MinimockGetFromTagInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *VariablesExtractorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *VariablesExtractorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetFromBranchDone() &&
		m.MinimockGetFromTagDone()
}
