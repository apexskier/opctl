// This file was generated by counterfeiter
package cliparamsatisfier

import (
	"sync"
)

type FakeInputSourcer struct {
	SourceStub        func(inputName string) *string
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
		inputName string
	}
	sourceReturns struct {
		result1 *string
	}
	sourceReturnsOnCall map[int]struct {
		result1 *string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputSourcer) Source(inputName string) *string {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
		inputName string
	}{inputName})
	fake.recordInvocation("Source", []interface{}{inputName})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub(inputName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sourceReturns.result1
}

func (fake *FakeInputSourcer) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeInputSourcer) SourceArgsForCall(i int) string {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return fake.sourceArgsForCall[i].inputName
}

func (fake *FakeInputSourcer) SourceReturns(result1 *string) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeInputSourcer) SourceReturnsOnCall(i int, result1 *string) {
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 *string
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 *string
	}{result1}
}

func (fake *FakeInputSourcer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeInputSourcer) recordInvocation(key string, args []interface{}) {
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

var _ InputSourcer = new(FakeInputSourcer)
