// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/opctl/opctl/sdks/go/internal/uniquestring"
)

type FakeUniqueStringFactory struct {
	ConstructStub        func() (string, error)
	constructMutex       sync.RWMutex
	constructArgsForCall []struct {
	}
	constructReturns struct {
		result1 string
		result2 error
	}
	constructReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUniqueStringFactory) Construct() (string, error) {
	fake.constructMutex.Lock()
	ret, specificReturn := fake.constructReturnsOnCall[len(fake.constructArgsForCall)]
	fake.constructArgsForCall = append(fake.constructArgsForCall, struct {
	}{})
	fake.recordInvocation("Construct", []interface{}{})
	fake.constructMutex.Unlock()
	if fake.ConstructStub != nil {
		return fake.ConstructStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.constructReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUniqueStringFactory) ConstructCallCount() int {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return len(fake.constructArgsForCall)
}

func (fake *FakeUniqueStringFactory) ConstructCalls(stub func() (string, error)) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = stub
}

func (fake *FakeUniqueStringFactory) ConstructReturns(result1 string, result2 error) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	fake.constructReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUniqueStringFactory) ConstructReturnsOnCall(i int, result1 string, result2 error) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	if fake.constructReturnsOnCall == nil {
		fake.constructReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.constructReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUniqueStringFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUniqueStringFactory) recordInvocation(key string, args []interface{}) {
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

var _ uniquestring.UniqueStringFactory = new(FakeUniqueStringFactory)
