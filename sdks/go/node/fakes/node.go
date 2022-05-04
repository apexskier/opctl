// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"github.com/opctl/opctl/sdks/go/node"
	"sync"

	"github.com/opctl/opctl/sdks/go/data"
	"github.com/opctl/opctl/sdks/go/model"
)

type FakeNode struct {
	LabelStub        func() string
	labelMutex       sync.RWMutex
	labelArgsForCall []struct {
	}
	labelReturns struct {
		result1 string
	}
	labelReturnsOnCall map[int]struct {
		result1 string
	}
	ResolveStub        func(context.Context, string) (data.DataHandle, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	resolveReturns struct {
		result1 data.DataHandle
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 data.DataHandle
		result2 error
	}
	StartOpStub        func(context.Context, chan model.Event, model.StartOpReq) (map[string]*model.Value, error)
	startOpMutex       sync.RWMutex
	startOpArgsForCall []struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 model.StartOpReq
	}
	startOpReturns struct {
		result1 map[string]*model.Value
		result2 error
	}
	startOpReturnsOnCall map[int]struct {
		result1 map[string]*model.Value
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNode) Label() string {
	fake.labelMutex.Lock()
	ret, specificReturn := fake.labelReturnsOnCall[len(fake.labelArgsForCall)]
	fake.labelArgsForCall = append(fake.labelArgsForCall, struct {
	}{})
	fake.recordInvocation("Label", []interface{}{})
	fake.labelMutex.Unlock()
	if fake.LabelStub != nil {
		return fake.LabelStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.labelReturns
	return fakeReturns.result1
}

func (fake *FakeNode) LabelCallCount() int {
	fake.labelMutex.RLock()
	defer fake.labelMutex.RUnlock()
	return len(fake.labelArgsForCall)
}

func (fake *FakeNode) LabelCalls(stub func() string) {
	fake.labelMutex.Lock()
	defer fake.labelMutex.Unlock()
	fake.LabelStub = stub
}

func (fake *FakeNode) LabelReturns(result1 string) {
	fake.labelMutex.Lock()
	defer fake.labelMutex.Unlock()
	fake.LabelStub = nil
	fake.labelReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeNode) LabelReturnsOnCall(i int, result1 string) {
	fake.labelMutex.Lock()
	defer fake.labelMutex.Unlock()
	fake.LabelStub = nil
	if fake.labelReturnsOnCall == nil {
		fake.labelReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.labelReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeNode) Resolve(arg1 context.Context, arg2 string) (data.DataHandle, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Resolve", []interface{}{arg1, arg2})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNode) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *FakeNode) ResolveCalls(stub func(context.Context, string) (data.DataHandle, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *FakeNode) ResolveArgsForCall(i int) (context.Context, string) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNode) ResolveReturns(result1 data.DataHandle, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 data.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNode) ResolveReturnsOnCall(i int, result1 data.DataHandle, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 data.DataHandle
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 data.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeNode) StartOp(arg1 context.Context, arg2 chan model.Event, arg3 model.StartOpReq) (map[string]*model.Value, error) {
	fake.startOpMutex.Lock()
	ret, specificReturn := fake.startOpReturnsOnCall[len(fake.startOpArgsForCall)]
	fake.startOpArgsForCall = append(fake.startOpArgsForCall, struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 model.StartOpReq
	}{arg1, arg2, arg3})
	fake.recordInvocation("StartOp", []interface{}{arg1, arg2, arg3})
	fake.startOpMutex.Unlock()
	if fake.StartOpStub != nil {
		return fake.StartOpStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.startOpReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNode) StartOpCallCount() int {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return len(fake.startOpArgsForCall)
}

func (fake *FakeNode) StartOpCalls(stub func(context.Context, chan model.Event, model.StartOpReq) (map[string]*model.Value, error)) {
	fake.startOpMutex.Lock()
	defer fake.startOpMutex.Unlock()
	fake.StartOpStub = stub
}

func (fake *FakeNode) StartOpArgsForCall(i int) (context.Context, chan model.Event, model.StartOpReq) {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	argsForCall := fake.startOpArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNode) StartOpReturns(result1 map[string]*model.Value, result2 error) {
	fake.startOpMutex.Lock()
	defer fake.startOpMutex.Unlock()
	fake.StartOpStub = nil
	fake.startOpReturns = struct {
		result1 map[string]*model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeNode) StartOpReturnsOnCall(i int, result1 map[string]*model.Value, result2 error) {
	fake.startOpMutex.Lock()
	defer fake.startOpMutex.Unlock()
	fake.StartOpStub = nil
	if fake.startOpReturnsOnCall == nil {
		fake.startOpReturnsOnCall = make(map[int]struct {
			result1 map[string]*model.Value
			result2 error
		})
	}
	fake.startOpReturnsOnCall[i] = struct {
		result1 map[string]*model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeNode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.labelMutex.RLock()
	defer fake.labelMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNode) recordInvocation(key string, args []interface{}) {
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

var _ node.Node = new(FakeNode)
