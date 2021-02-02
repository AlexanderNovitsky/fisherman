package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i fisherman/configuration.Rule -o ./testing/mocks/rule_mock.go

import (
	"fisherman/internal"
	"io"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RuleMock implements configuration.Rule
type RuleMock struct {
	t minimock.Tester

	funcCheck          func(e1 internal.ExecutionContext, w1 io.Writer) (err error)
	inspectFuncCheck   func(e1 internal.ExecutionContext, w1 io.Writer)
	afterCheckCounter  uint64
	beforeCheckCounter uint64
	CheckMock          mRuleMockCheck

	funcGetContition          func() (s1 string)
	inspectFuncGetContition   func()
	afterGetContitionCounter  uint64
	beforeGetContitionCounter uint64
	GetContitionMock          mRuleMockGetContition

	funcGetPosition          func() (b1 byte)
	inspectFuncGetPosition   func()
	afterGetPositionCounter  uint64
	beforeGetPositionCounter uint64
	GetPositionMock          mRuleMockGetPosition

	funcGetType          func() (s1 string)
	inspectFuncGetType   func()
	afterGetTypeCounter  uint64
	beforeGetTypeCounter uint64
	GetTypeMock          mRuleMockGetType
}

// NewRuleMock returns a mock for configuration.Rule
func NewRuleMock(t minimock.Tester) *RuleMock {
	m := &RuleMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CheckMock = mRuleMockCheck{mock: m}
	m.CheckMock.callArgs = []*RuleMockCheckParams{}

	m.GetContitionMock = mRuleMockGetContition{mock: m}

	m.GetPositionMock = mRuleMockGetPosition{mock: m}

	m.GetTypeMock = mRuleMockGetType{mock: m}

	return m
}

type mRuleMockCheck struct {
	mock               *RuleMock
	defaultExpectation *RuleMockCheckExpectation
	expectations       []*RuleMockCheckExpectation

	callArgs []*RuleMockCheckParams
	mutex    sync.RWMutex
}

// RuleMockCheckExpectation specifies expectation struct of the Rule.Check
type RuleMockCheckExpectation struct {
	mock    *RuleMock
	params  *RuleMockCheckParams
	results *RuleMockCheckResults
	Counter uint64
}

// RuleMockCheckParams contains parameters of the Rule.Check
type RuleMockCheckParams struct {
	e1 internal.ExecutionContext
	w1 io.Writer
}

// RuleMockCheckResults contains results of the Rule.Check
type RuleMockCheckResults struct {
	err error
}

// Expect sets up expected params for Rule.Check
func (mmCheck *mRuleMockCheck) Expect(e1 internal.ExecutionContext, w1 io.Writer) *mRuleMockCheck {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("RuleMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &RuleMockCheckExpectation{}
	}

	mmCheck.defaultExpectation.params = &RuleMockCheckParams{e1, w1}
	for _, e := range mmCheck.expectations {
		if minimock.Equal(e.params, mmCheck.defaultExpectation.params) {
			mmCheck.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCheck.defaultExpectation.params)
		}
	}

	return mmCheck
}

// Inspect accepts an inspector function that has same arguments as the Rule.Check
func (mmCheck *mRuleMockCheck) Inspect(f func(e1 internal.ExecutionContext, w1 io.Writer)) *mRuleMockCheck {
	if mmCheck.mock.inspectFuncCheck != nil {
		mmCheck.mock.t.Fatalf("Inspect function is already set for RuleMock.Check")
	}

	mmCheck.mock.inspectFuncCheck = f

	return mmCheck
}

// Return sets up results that will be returned by Rule.Check
func (mmCheck *mRuleMockCheck) Return(err error) *RuleMock {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("RuleMock.Check mock is already set by Set")
	}

	if mmCheck.defaultExpectation == nil {
		mmCheck.defaultExpectation = &RuleMockCheckExpectation{mock: mmCheck.mock}
	}
	mmCheck.defaultExpectation.results = &RuleMockCheckResults{err}
	return mmCheck.mock
}

//Set uses given function f to mock the Rule.Check method
func (mmCheck *mRuleMockCheck) Set(f func(e1 internal.ExecutionContext, w1 io.Writer) (err error)) *RuleMock {
	if mmCheck.defaultExpectation != nil {
		mmCheck.mock.t.Fatalf("Default expectation is already set for the Rule.Check method")
	}

	if len(mmCheck.expectations) > 0 {
		mmCheck.mock.t.Fatalf("Some expectations are already set for the Rule.Check method")
	}

	mmCheck.mock.funcCheck = f
	return mmCheck.mock
}

// When sets expectation for the Rule.Check which will trigger the result defined by the following
// Then helper
func (mmCheck *mRuleMockCheck) When(e1 internal.ExecutionContext, w1 io.Writer) *RuleMockCheckExpectation {
	if mmCheck.mock.funcCheck != nil {
		mmCheck.mock.t.Fatalf("RuleMock.Check mock is already set by Set")
	}

	expectation := &RuleMockCheckExpectation{
		mock:   mmCheck.mock,
		params: &RuleMockCheckParams{e1, w1},
	}
	mmCheck.expectations = append(mmCheck.expectations, expectation)
	return expectation
}

// Then sets up Rule.Check return parameters for the expectation previously defined by the When method
func (e *RuleMockCheckExpectation) Then(err error) *RuleMock {
	e.results = &RuleMockCheckResults{err}
	return e.mock
}

// Check implements configuration.Rule
func (mmCheck *RuleMock) Check(e1 internal.ExecutionContext, w1 io.Writer) (err error) {
	mm_atomic.AddUint64(&mmCheck.beforeCheckCounter, 1)
	defer mm_atomic.AddUint64(&mmCheck.afterCheckCounter, 1)

	if mmCheck.inspectFuncCheck != nil {
		mmCheck.inspectFuncCheck(e1, w1)
	}

	mm_params := &RuleMockCheckParams{e1, w1}

	// Record call args
	mmCheck.CheckMock.mutex.Lock()
	mmCheck.CheckMock.callArgs = append(mmCheck.CheckMock.callArgs, mm_params)
	mmCheck.CheckMock.mutex.Unlock()

	for _, e := range mmCheck.CheckMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCheck.CheckMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCheck.CheckMock.defaultExpectation.Counter, 1)
		mm_want := mmCheck.CheckMock.defaultExpectation.params
		mm_got := RuleMockCheckParams{e1, w1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCheck.t.Errorf("RuleMock.Check got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCheck.CheckMock.defaultExpectation.results
		if mm_results == nil {
			mmCheck.t.Fatal("No results are set for the RuleMock.Check")
		}
		return (*mm_results).err
	}
	if mmCheck.funcCheck != nil {
		return mmCheck.funcCheck(e1, w1)
	}
	mmCheck.t.Fatalf("Unexpected call to RuleMock.Check. %v %v", e1, w1)
	return
}

// CheckAfterCounter returns a count of finished RuleMock.Check invocations
func (mmCheck *RuleMock) CheckAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.afterCheckCounter)
}

// CheckBeforeCounter returns a count of RuleMock.Check invocations
func (mmCheck *RuleMock) CheckBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCheck.beforeCheckCounter)
}

// Calls returns a list of arguments used in each call to RuleMock.Check.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCheck *mRuleMockCheck) Calls() []*RuleMockCheckParams {
	mmCheck.mutex.RLock()

	argCopy := make([]*RuleMockCheckParams, len(mmCheck.callArgs))
	copy(argCopy, mmCheck.callArgs)

	mmCheck.mutex.RUnlock()

	return argCopy
}

// MinimockCheckDone returns true if the count of the Check invocations corresponds
// the number of defined expectations
func (m *RuleMock) MinimockCheckDone() bool {
	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheck != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		return false
	}
	return true
}

// MinimockCheckInspect logs each unmet expectation
func (m *RuleMock) MinimockCheckInspect() {
	for _, e := range m.CheckMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RuleMock.Check with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CheckMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		if m.CheckMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RuleMock.Check")
		} else {
			m.t.Errorf("Expected call to RuleMock.Check with params: %#v", *m.CheckMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCheck != nil && mm_atomic.LoadUint64(&m.afterCheckCounter) < 1 {
		m.t.Error("Expected call to RuleMock.Check")
	}
}

type mRuleMockGetContition struct {
	mock               *RuleMock
	defaultExpectation *RuleMockGetContitionExpectation
	expectations       []*RuleMockGetContitionExpectation
}

// RuleMockGetContitionExpectation specifies expectation struct of the Rule.GetContition
type RuleMockGetContitionExpectation struct {
	mock *RuleMock

	results *RuleMockGetContitionResults
	Counter uint64
}

// RuleMockGetContitionResults contains results of the Rule.GetContition
type RuleMockGetContitionResults struct {
	s1 string
}

// Expect sets up expected params for Rule.GetContition
func (mmGetContition *mRuleMockGetContition) Expect() *mRuleMockGetContition {
	if mmGetContition.mock.funcGetContition != nil {
		mmGetContition.mock.t.Fatalf("RuleMock.GetContition mock is already set by Set")
	}

	if mmGetContition.defaultExpectation == nil {
		mmGetContition.defaultExpectation = &RuleMockGetContitionExpectation{}
	}

	return mmGetContition
}

// Inspect accepts an inspector function that has same arguments as the Rule.GetContition
func (mmGetContition *mRuleMockGetContition) Inspect(f func()) *mRuleMockGetContition {
	if mmGetContition.mock.inspectFuncGetContition != nil {
		mmGetContition.mock.t.Fatalf("Inspect function is already set for RuleMock.GetContition")
	}

	mmGetContition.mock.inspectFuncGetContition = f

	return mmGetContition
}

// Return sets up results that will be returned by Rule.GetContition
func (mmGetContition *mRuleMockGetContition) Return(s1 string) *RuleMock {
	if mmGetContition.mock.funcGetContition != nil {
		mmGetContition.mock.t.Fatalf("RuleMock.GetContition mock is already set by Set")
	}

	if mmGetContition.defaultExpectation == nil {
		mmGetContition.defaultExpectation = &RuleMockGetContitionExpectation{mock: mmGetContition.mock}
	}
	mmGetContition.defaultExpectation.results = &RuleMockGetContitionResults{s1}
	return mmGetContition.mock
}

//Set uses given function f to mock the Rule.GetContition method
func (mmGetContition *mRuleMockGetContition) Set(f func() (s1 string)) *RuleMock {
	if mmGetContition.defaultExpectation != nil {
		mmGetContition.mock.t.Fatalf("Default expectation is already set for the Rule.GetContition method")
	}

	if len(mmGetContition.expectations) > 0 {
		mmGetContition.mock.t.Fatalf("Some expectations are already set for the Rule.GetContition method")
	}

	mmGetContition.mock.funcGetContition = f
	return mmGetContition.mock
}

// GetContition implements configuration.Rule
func (mmGetContition *RuleMock) GetContition() (s1 string) {
	mm_atomic.AddUint64(&mmGetContition.beforeGetContitionCounter, 1)
	defer mm_atomic.AddUint64(&mmGetContition.afterGetContitionCounter, 1)

	if mmGetContition.inspectFuncGetContition != nil {
		mmGetContition.inspectFuncGetContition()
	}

	if mmGetContition.GetContitionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetContition.GetContitionMock.defaultExpectation.Counter, 1)

		mm_results := mmGetContition.GetContitionMock.defaultExpectation.results
		if mm_results == nil {
			mmGetContition.t.Fatal("No results are set for the RuleMock.GetContition")
		}
		return (*mm_results).s1
	}
	if mmGetContition.funcGetContition != nil {
		return mmGetContition.funcGetContition()
	}
	mmGetContition.t.Fatalf("Unexpected call to RuleMock.GetContition.")
	return
}

// GetContitionAfterCounter returns a count of finished RuleMock.GetContition invocations
func (mmGetContition *RuleMock) GetContitionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetContition.afterGetContitionCounter)
}

// GetContitionBeforeCounter returns a count of RuleMock.GetContition invocations
func (mmGetContition *RuleMock) GetContitionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetContition.beforeGetContitionCounter)
}

// MinimockGetContitionDone returns true if the count of the GetContition invocations corresponds
// the number of defined expectations
func (m *RuleMock) MinimockGetContitionDone() bool {
	for _, e := range m.GetContitionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetContitionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetContitionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetContition != nil && mm_atomic.LoadUint64(&m.afterGetContitionCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetContitionInspect logs each unmet expectation
func (m *RuleMock) MinimockGetContitionInspect() {
	for _, e := range m.GetContitionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RuleMock.GetContition")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetContitionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetContitionCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetContition")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetContition != nil && mm_atomic.LoadUint64(&m.afterGetContitionCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetContition")
	}
}

type mRuleMockGetPosition struct {
	mock               *RuleMock
	defaultExpectation *RuleMockGetPositionExpectation
	expectations       []*RuleMockGetPositionExpectation
}

// RuleMockGetPositionExpectation specifies expectation struct of the Rule.GetPosition
type RuleMockGetPositionExpectation struct {
	mock *RuleMock

	results *RuleMockGetPositionResults
	Counter uint64
}

// RuleMockGetPositionResults contains results of the Rule.GetPosition
type RuleMockGetPositionResults struct {
	b1 byte
}

// Expect sets up expected params for Rule.GetPosition
func (mmGetPosition *mRuleMockGetPosition) Expect() *mRuleMockGetPosition {
	if mmGetPosition.mock.funcGetPosition != nil {
		mmGetPosition.mock.t.Fatalf("RuleMock.GetPosition mock is already set by Set")
	}

	if mmGetPosition.defaultExpectation == nil {
		mmGetPosition.defaultExpectation = &RuleMockGetPositionExpectation{}
	}

	return mmGetPosition
}

// Inspect accepts an inspector function that has same arguments as the Rule.GetPosition
func (mmGetPosition *mRuleMockGetPosition) Inspect(f func()) *mRuleMockGetPosition {
	if mmGetPosition.mock.inspectFuncGetPosition != nil {
		mmGetPosition.mock.t.Fatalf("Inspect function is already set for RuleMock.GetPosition")
	}

	mmGetPosition.mock.inspectFuncGetPosition = f

	return mmGetPosition
}

// Return sets up results that will be returned by Rule.GetPosition
func (mmGetPosition *mRuleMockGetPosition) Return(b1 byte) *RuleMock {
	if mmGetPosition.mock.funcGetPosition != nil {
		mmGetPosition.mock.t.Fatalf("RuleMock.GetPosition mock is already set by Set")
	}

	if mmGetPosition.defaultExpectation == nil {
		mmGetPosition.defaultExpectation = &RuleMockGetPositionExpectation{mock: mmGetPosition.mock}
	}
	mmGetPosition.defaultExpectation.results = &RuleMockGetPositionResults{b1}
	return mmGetPosition.mock
}

//Set uses given function f to mock the Rule.GetPosition method
func (mmGetPosition *mRuleMockGetPosition) Set(f func() (b1 byte)) *RuleMock {
	if mmGetPosition.defaultExpectation != nil {
		mmGetPosition.mock.t.Fatalf("Default expectation is already set for the Rule.GetPosition method")
	}

	if len(mmGetPosition.expectations) > 0 {
		mmGetPosition.mock.t.Fatalf("Some expectations are already set for the Rule.GetPosition method")
	}

	mmGetPosition.mock.funcGetPosition = f
	return mmGetPosition.mock
}

// GetPosition implements configuration.Rule
func (mmGetPosition *RuleMock) GetPosition() (b1 byte) {
	mm_atomic.AddUint64(&mmGetPosition.beforeGetPositionCounter, 1)
	defer mm_atomic.AddUint64(&mmGetPosition.afterGetPositionCounter, 1)

	if mmGetPosition.inspectFuncGetPosition != nil {
		mmGetPosition.inspectFuncGetPosition()
	}

	if mmGetPosition.GetPositionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetPosition.GetPositionMock.defaultExpectation.Counter, 1)

		mm_results := mmGetPosition.GetPositionMock.defaultExpectation.results
		if mm_results == nil {
			mmGetPosition.t.Fatal("No results are set for the RuleMock.GetPosition")
		}
		return (*mm_results).b1
	}
	if mmGetPosition.funcGetPosition != nil {
		return mmGetPosition.funcGetPosition()
	}
	mmGetPosition.t.Fatalf("Unexpected call to RuleMock.GetPosition.")
	return
}

// GetPositionAfterCounter returns a count of finished RuleMock.GetPosition invocations
func (mmGetPosition *RuleMock) GetPositionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPosition.afterGetPositionCounter)
}

// GetPositionBeforeCounter returns a count of RuleMock.GetPosition invocations
func (mmGetPosition *RuleMock) GetPositionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPosition.beforeGetPositionCounter)
}

// MinimockGetPositionDone returns true if the count of the GetPosition invocations corresponds
// the number of defined expectations
func (m *RuleMock) MinimockGetPositionDone() bool {
	for _, e := range m.GetPositionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPositionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPositionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPosition != nil && mm_atomic.LoadUint64(&m.afterGetPositionCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetPositionInspect logs each unmet expectation
func (m *RuleMock) MinimockGetPositionInspect() {
	for _, e := range m.GetPositionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RuleMock.GetPosition")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPositionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPositionCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetPosition")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPosition != nil && mm_atomic.LoadUint64(&m.afterGetPositionCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetPosition")
	}
}

type mRuleMockGetType struct {
	mock               *RuleMock
	defaultExpectation *RuleMockGetTypeExpectation
	expectations       []*RuleMockGetTypeExpectation
}

// RuleMockGetTypeExpectation specifies expectation struct of the Rule.GetType
type RuleMockGetTypeExpectation struct {
	mock *RuleMock

	results *RuleMockGetTypeResults
	Counter uint64
}

// RuleMockGetTypeResults contains results of the Rule.GetType
type RuleMockGetTypeResults struct {
	s1 string
}

// Expect sets up expected params for Rule.GetType
func (mmGetType *mRuleMockGetType) Expect() *mRuleMockGetType {
	if mmGetType.mock.funcGetType != nil {
		mmGetType.mock.t.Fatalf("RuleMock.GetType mock is already set by Set")
	}

	if mmGetType.defaultExpectation == nil {
		mmGetType.defaultExpectation = &RuleMockGetTypeExpectation{}
	}

	return mmGetType
}

// Inspect accepts an inspector function that has same arguments as the Rule.GetType
func (mmGetType *mRuleMockGetType) Inspect(f func()) *mRuleMockGetType {
	if mmGetType.mock.inspectFuncGetType != nil {
		mmGetType.mock.t.Fatalf("Inspect function is already set for RuleMock.GetType")
	}

	mmGetType.mock.inspectFuncGetType = f

	return mmGetType
}

// Return sets up results that will be returned by Rule.GetType
func (mmGetType *mRuleMockGetType) Return(s1 string) *RuleMock {
	if mmGetType.mock.funcGetType != nil {
		mmGetType.mock.t.Fatalf("RuleMock.GetType mock is already set by Set")
	}

	if mmGetType.defaultExpectation == nil {
		mmGetType.defaultExpectation = &RuleMockGetTypeExpectation{mock: mmGetType.mock}
	}
	mmGetType.defaultExpectation.results = &RuleMockGetTypeResults{s1}
	return mmGetType.mock
}

//Set uses given function f to mock the Rule.GetType method
func (mmGetType *mRuleMockGetType) Set(f func() (s1 string)) *RuleMock {
	if mmGetType.defaultExpectation != nil {
		mmGetType.mock.t.Fatalf("Default expectation is already set for the Rule.GetType method")
	}

	if len(mmGetType.expectations) > 0 {
		mmGetType.mock.t.Fatalf("Some expectations are already set for the Rule.GetType method")
	}

	mmGetType.mock.funcGetType = f
	return mmGetType.mock
}

// GetType implements configuration.Rule
func (mmGetType *RuleMock) GetType() (s1 string) {
	mm_atomic.AddUint64(&mmGetType.beforeGetTypeCounter, 1)
	defer mm_atomic.AddUint64(&mmGetType.afterGetTypeCounter, 1)

	if mmGetType.inspectFuncGetType != nil {
		mmGetType.inspectFuncGetType()
	}

	if mmGetType.GetTypeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetType.GetTypeMock.defaultExpectation.Counter, 1)

		mm_results := mmGetType.GetTypeMock.defaultExpectation.results
		if mm_results == nil {
			mmGetType.t.Fatal("No results are set for the RuleMock.GetType")
		}
		return (*mm_results).s1
	}
	if mmGetType.funcGetType != nil {
		return mmGetType.funcGetType()
	}
	mmGetType.t.Fatalf("Unexpected call to RuleMock.GetType.")
	return
}

// GetTypeAfterCounter returns a count of finished RuleMock.GetType invocations
func (mmGetType *RuleMock) GetTypeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetType.afterGetTypeCounter)
}

// GetTypeBeforeCounter returns a count of RuleMock.GetType invocations
func (mmGetType *RuleMock) GetTypeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetType.beforeGetTypeCounter)
}

// MinimockGetTypeDone returns true if the count of the GetType invocations corresponds
// the number of defined expectations
func (m *RuleMock) MinimockGetTypeDone() bool {
	for _, e := range m.GetTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTypeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetType != nil && mm_atomic.LoadUint64(&m.afterGetTypeCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetTypeInspect logs each unmet expectation
func (m *RuleMock) MinimockGetTypeInspect() {
	for _, e := range m.GetTypeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to RuleMock.GetType")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetTypeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetTypeCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetType")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetType != nil && mm_atomic.LoadUint64(&m.afterGetTypeCounter) < 1 {
		m.t.Error("Expected call to RuleMock.GetType")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RuleMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCheckInspect()

		m.MinimockGetContitionInspect()

		m.MinimockGetPositionInspect()

		m.MinimockGetTypeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RuleMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RuleMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCheckDone() &&
		m.MinimockGetContitionDone() &&
		m.MinimockGetPositionDone() &&
		m.MinimockGetTypeDone()
}
