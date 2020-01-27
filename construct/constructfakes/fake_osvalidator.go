// Code generated by counterfeiter. DO NOT EDIT.
package constructfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/stembuild/construct"
)

type FakeOSValidator struct {
	ValidateStub        func(string) error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		arg1 string
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOSValidator) Validate(arg1 string) error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Validate", []interface{}{arg1})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateReturns
	return fakeReturns.result1
}

func (fake *FakeOSValidator) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeOSValidator) ValidateCalls(stub func(string) error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *FakeOSValidator) ValidateArgsForCall(i int) string {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	argsForCall := fake.validateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOSValidator) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOSValidator) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOSValidator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOSValidator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ construct.OSValidator = new(FakeOSValidator)