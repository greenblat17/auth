// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/greenblat17/auth/internal/service.UserSaverProducer -o user_saver_producer_minimock.go -n UserSaverProducerMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/greenblat17/auth/internal/model"
)

// UserSaverProducerMock implements mm_service.UserSaverProducer
type UserSaverProducerMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSend          func(ctx context.Context, userInfo *model.UserInfo) (err error)
	funcSendOrigin    string
	inspectFuncSend   func(ctx context.Context, userInfo *model.UserInfo)
	afterSendCounter  uint64
	beforeSendCounter uint64
	SendMock          mUserSaverProducerMockSend
}

// NewUserSaverProducerMock returns a mock for mm_service.UserSaverProducer
func NewUserSaverProducerMock(t minimock.Tester) *UserSaverProducerMock {
	m := &UserSaverProducerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMock = mUserSaverProducerMockSend{mock: m}
	m.SendMock.callArgs = []*UserSaverProducerMockSendParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserSaverProducerMockSend struct {
	optional           bool
	mock               *UserSaverProducerMock
	defaultExpectation *UserSaverProducerMockSendExpectation
	expectations       []*UserSaverProducerMockSendExpectation

	callArgs []*UserSaverProducerMockSendParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// UserSaverProducerMockSendExpectation specifies expectation struct of the UserSaverProducer.Send
type UserSaverProducerMockSendExpectation struct {
	mock               *UserSaverProducerMock
	params             *UserSaverProducerMockSendParams
	paramPtrs          *UserSaverProducerMockSendParamPtrs
	expectationOrigins UserSaverProducerMockSendExpectationOrigins
	results            *UserSaverProducerMockSendResults
	returnOrigin       string
	Counter            uint64
}

// UserSaverProducerMockSendParams contains parameters of the UserSaverProducer.Send
type UserSaverProducerMockSendParams struct {
	ctx      context.Context
	userInfo *model.UserInfo
}

// UserSaverProducerMockSendParamPtrs contains pointers to parameters of the UserSaverProducer.Send
type UserSaverProducerMockSendParamPtrs struct {
	ctx      *context.Context
	userInfo **model.UserInfo
}

// UserSaverProducerMockSendResults contains results of the UserSaverProducer.Send
type UserSaverProducerMockSendResults struct {
	err error
}

// UserSaverProducerMockSendOrigins contains origins of expectations of the UserSaverProducer.Send
type UserSaverProducerMockSendExpectationOrigins struct {
	origin         string
	originCtx      string
	originUserInfo string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSend *mUserSaverProducerMockSend) Optional() *mUserSaverProducerMockSend {
	mmSend.optional = true
	return mmSend
}

// Expect sets up expected params for UserSaverProducer.Send
func (mmSend *mUserSaverProducerMockSend) Expect(ctx context.Context, userInfo *model.UserInfo) *mUserSaverProducerMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &UserSaverProducerMockSendExpectation{}
	}

	if mmSend.defaultExpectation.paramPtrs != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by ExpectParams functions")
	}

	mmSend.defaultExpectation.params = &UserSaverProducerMockSendParams{ctx, userInfo}
	mmSend.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmSend.expectations {
		if minimock.Equal(e.params, mmSend.defaultExpectation.params) {
			mmSend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSend.defaultExpectation.params)
		}
	}

	return mmSend
}

// ExpectCtxParam1 sets up expected param ctx for UserSaverProducer.Send
func (mmSend *mUserSaverProducerMockSend) ExpectCtxParam1(ctx context.Context) *mUserSaverProducerMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &UserSaverProducerMockSendExpectation{}
	}

	if mmSend.defaultExpectation.params != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Expect")
	}

	if mmSend.defaultExpectation.paramPtrs == nil {
		mmSend.defaultExpectation.paramPtrs = &UserSaverProducerMockSendParamPtrs{}
	}
	mmSend.defaultExpectation.paramPtrs.ctx = &ctx
	mmSend.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmSend
}

// ExpectUserInfoParam2 sets up expected param userInfo for UserSaverProducer.Send
func (mmSend *mUserSaverProducerMockSend) ExpectUserInfoParam2(userInfo *model.UserInfo) *mUserSaverProducerMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &UserSaverProducerMockSendExpectation{}
	}

	if mmSend.defaultExpectation.params != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Expect")
	}

	if mmSend.defaultExpectation.paramPtrs == nil {
		mmSend.defaultExpectation.paramPtrs = &UserSaverProducerMockSendParamPtrs{}
	}
	mmSend.defaultExpectation.paramPtrs.userInfo = &userInfo
	mmSend.defaultExpectation.expectationOrigins.originUserInfo = minimock.CallerInfo(1)

	return mmSend
}

// Inspect accepts an inspector function that has same arguments as the UserSaverProducer.Send
func (mmSend *mUserSaverProducerMockSend) Inspect(f func(ctx context.Context, userInfo *model.UserInfo)) *mUserSaverProducerMockSend {
	if mmSend.mock.inspectFuncSend != nil {
		mmSend.mock.t.Fatalf("Inspect function is already set for UserSaverProducerMock.Send")
	}

	mmSend.mock.inspectFuncSend = f

	return mmSend
}

// Return sets up results that will be returned by UserSaverProducer.Send
func (mmSend *mUserSaverProducerMockSend) Return(err error) *UserSaverProducerMock {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &UserSaverProducerMockSendExpectation{mock: mmSend.mock}
	}
	mmSend.defaultExpectation.results = &UserSaverProducerMockSendResults{err}
	mmSend.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmSend.mock
}

// Set uses given function f to mock the UserSaverProducer.Send method
func (mmSend *mUserSaverProducerMockSend) Set(f func(ctx context.Context, userInfo *model.UserInfo) (err error)) *UserSaverProducerMock {
	if mmSend.defaultExpectation != nil {
		mmSend.mock.t.Fatalf("Default expectation is already set for the UserSaverProducer.Send method")
	}

	if len(mmSend.expectations) > 0 {
		mmSend.mock.t.Fatalf("Some expectations are already set for the UserSaverProducer.Send method")
	}

	mmSend.mock.funcSend = f
	mmSend.mock.funcSendOrigin = minimock.CallerInfo(1)
	return mmSend.mock
}

// When sets expectation for the UserSaverProducer.Send which will trigger the result defined by the following
// Then helper
func (mmSend *mUserSaverProducerMockSend) When(ctx context.Context, userInfo *model.UserInfo) *UserSaverProducerMockSendExpectation {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("UserSaverProducerMock.Send mock is already set by Set")
	}

	expectation := &UserSaverProducerMockSendExpectation{
		mock:               mmSend.mock,
		params:             &UserSaverProducerMockSendParams{ctx, userInfo},
		expectationOrigins: UserSaverProducerMockSendExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmSend.expectations = append(mmSend.expectations, expectation)
	return expectation
}

// Then sets up UserSaverProducer.Send return parameters for the expectation previously defined by the When method
func (e *UserSaverProducerMockSendExpectation) Then(err error) *UserSaverProducerMock {
	e.results = &UserSaverProducerMockSendResults{err}
	return e.mock
}

// Times sets number of times UserSaverProducer.Send should be invoked
func (mmSend *mUserSaverProducerMockSend) Times(n uint64) *mUserSaverProducerMockSend {
	if n == 0 {
		mmSend.mock.t.Fatalf("Times of UserSaverProducerMock.Send mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSend.expectedInvocations, n)
	mmSend.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmSend
}

func (mmSend *mUserSaverProducerMockSend) invocationsDone() bool {
	if len(mmSend.expectations) == 0 && mmSend.defaultExpectation == nil && mmSend.mock.funcSend == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSend.mock.afterSendCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSend.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Send implements mm_service.UserSaverProducer
func (mmSend *UserSaverProducerMock) Send(ctx context.Context, userInfo *model.UserInfo) (err error) {
	mm_atomic.AddUint64(&mmSend.beforeSendCounter, 1)
	defer mm_atomic.AddUint64(&mmSend.afterSendCounter, 1)

	mmSend.t.Helper()

	if mmSend.inspectFuncSend != nil {
		mmSend.inspectFuncSend(ctx, userInfo)
	}

	mm_params := UserSaverProducerMockSendParams{ctx, userInfo}

	// Record call args
	mmSend.SendMock.mutex.Lock()
	mmSend.SendMock.callArgs = append(mmSend.SendMock.callArgs, &mm_params)
	mmSend.SendMock.mutex.Unlock()

	for _, e := range mmSend.SendMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSend.SendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSend.SendMock.defaultExpectation.Counter, 1)
		mm_want := mmSend.SendMock.defaultExpectation.params
		mm_want_ptrs := mmSend.SendMock.defaultExpectation.paramPtrs

		mm_got := UserSaverProducerMockSendParams{ctx, userInfo}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSend.t.Errorf("UserSaverProducerMock.Send got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSend.SendMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.userInfo != nil && !minimock.Equal(*mm_want_ptrs.userInfo, mm_got.userInfo) {
				mmSend.t.Errorf("UserSaverProducerMock.Send got unexpected parameter userInfo, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmSend.SendMock.defaultExpectation.expectationOrigins.originUserInfo, *mm_want_ptrs.userInfo, mm_got.userInfo, minimock.Diff(*mm_want_ptrs.userInfo, mm_got.userInfo))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSend.t.Errorf("UserSaverProducerMock.Send got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmSend.SendMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSend.SendMock.defaultExpectation.results
		if mm_results == nil {
			mmSend.t.Fatal("No results are set for the UserSaverProducerMock.Send")
		}
		return (*mm_results).err
	}
	if mmSend.funcSend != nil {
		return mmSend.funcSend(ctx, userInfo)
	}
	mmSend.t.Fatalf("Unexpected call to UserSaverProducerMock.Send. %v %v", ctx, userInfo)
	return
}

// SendAfterCounter returns a count of finished UserSaverProducerMock.Send invocations
func (mmSend *UserSaverProducerMock) SendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.afterSendCounter)
}

// SendBeforeCounter returns a count of UserSaverProducerMock.Send invocations
func (mmSend *UserSaverProducerMock) SendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.beforeSendCounter)
}

// Calls returns a list of arguments used in each call to UserSaverProducerMock.Send.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSend *mUserSaverProducerMockSend) Calls() []*UserSaverProducerMockSendParams {
	mmSend.mutex.RLock()

	argCopy := make([]*UserSaverProducerMockSendParams, len(mmSend.callArgs))
	copy(argCopy, mmSend.callArgs)

	mmSend.mutex.RUnlock()

	return argCopy
}

// MinimockSendDone returns true if the count of the Send invocations corresponds
// the number of defined expectations
func (m *UserSaverProducerMock) MinimockSendDone() bool {
	if m.SendMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendMock.invocationsDone()
}

// MinimockSendInspect logs each unmet expectation
func (m *UserSaverProducerMock) MinimockSendInspect() {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserSaverProducerMock.Send at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterSendCounter := mm_atomic.LoadUint64(&m.afterSendCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && afterSendCounter < 1 {
		if m.SendMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to UserSaverProducerMock.Send at\n%s", m.SendMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to UserSaverProducerMock.Send at\n%s with params: %#v", m.SendMock.defaultExpectation.expectationOrigins.origin, *m.SendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && afterSendCounter < 1 {
		m.t.Errorf("Expected call to UserSaverProducerMock.Send at\n%s", m.funcSendOrigin)
	}

	if !m.SendMock.invocationsDone() && afterSendCounter > 0 {
		m.t.Errorf("Expected %d calls to UserSaverProducerMock.Send at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.SendMock.expectedInvocations), m.SendMock.expectedInvocationsOrigin, afterSendCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserSaverProducerMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserSaverProducerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserSaverProducerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendDone()
}