// Code generated by counterfeiter. DO NOT EDIT.
package number

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	InterpretStub        func(scope map[string]*model.Value, expression string, pkgHandle model.PkgHandle) (float64, error)
	interpretMutex       sync.RWMutex
	interpretArgsForCall []struct {
		scope      map[string]*model.Value
		expression string
		pkgHandle  model.PkgHandle
	}
	interpretReturns struct {
		result1 float64
		result2 error
	}
	interpretReturnsOnCall map[int]struct {
		result1 float64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Interpret(scope map[string]*model.Value, expression string, pkgHandle model.PkgHandle) (float64, error) {
	fake.interpretMutex.Lock()
	ret, specificReturn := fake.interpretReturnsOnCall[len(fake.interpretArgsForCall)]
	fake.interpretArgsForCall = append(fake.interpretArgsForCall, struct {
		scope      map[string]*model.Value
		expression string
		pkgHandle  model.PkgHandle
	}{scope, expression, pkgHandle})
	fake.recordInvocation("Interpret", []interface{}{scope, expression, pkgHandle})
	fake.interpretMutex.Unlock()
	if fake.InterpretStub != nil {
		return fake.InterpretStub(scope, expression, pkgHandle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpretReturns.result1, fake.interpretReturns.result2
}

func (fake *Fake) InterpretCallCount() int {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return len(fake.interpretArgsForCall)
}

func (fake *Fake) InterpretArgsForCall(i int) (map[string]*model.Value, string, model.PkgHandle) {
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	return fake.interpretArgsForCall[i].scope, fake.interpretArgsForCall[i].expression, fake.interpretArgsForCall[i].pkgHandle
}

func (fake *Fake) InterpretReturns(result1 float64, result2 error) {
	fake.InterpretStub = nil
	fake.interpretReturns = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *Fake) InterpretReturnsOnCall(i int, result1 float64, result2 error) {
	fake.InterpretStub = nil
	if fake.interpretReturnsOnCall == nil {
		fake.interpretReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 error
		})
	}
	fake.interpretReturnsOnCall[i] = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpretMutex.RLock()
	defer fake.interpretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ Number = new(Fake)
