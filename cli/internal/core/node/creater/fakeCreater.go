// Code generated by counterfeiter. DO NOT EDIT.
package creater

import (
	"sync"

	"github.com/opctl/opctl/cli/internal/model"
)

type FakeCreater struct {
	CreateStub        func(model.NodeCreateOpts)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		arg1 model.NodeCreateOpts
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreater) Create(arg1 model.NodeCreateOpts) {
	fake.invokeMutex.Lock()
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		arg1 model.NodeCreateOpts
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.invokeMutex.Unlock()
	if fake.CreateStub != nil {
		fake.CreateStub(arg1)
	}
}

func (fake *FakeCreater) CreateCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeCreater) CreateCalls(stub func(model.NodeCreateOpts)) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeCreater) CreateArgsForCall(i int) model.NodeCreateOpts {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	argsForCall := fake.invokeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCreater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreater) recordInvocation(key string, args []interface{}) {
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

var _ Creater = new(FakeCreater)