// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"io"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node/containerruntime"
)

type FakeContainerRuntime struct {
	DeleteStub        func(context.Context) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteContainerIfExistsStub        func(context.Context, string) error
	deleteContainerIfExistsMutex       sync.RWMutex
	deleteContainerIfExistsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteContainerIfExistsReturns struct {
		result1 error
	}
	deleteContainerIfExistsReturnsOnCall map[int]struct {
		result1 error
	}
	KillStub        func(context.Context) error
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		arg1 context.Context
	}
	killReturns struct {
		result1 error
	}
	killReturnsOnCall map[int]struct {
		result1 error
	}
	RunContainerStub        func(context.Context, chan model.Event, *model.ContainerCall, string, io.WriteCloser, io.WriteCloser) (*int64, error)
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 *model.ContainerCall
		arg4 string
		arg5 io.WriteCloser
		arg6 io.WriteCloser
	}
	runContainerReturns struct {
		result1 *int64
		result2 error
	}
	runContainerReturnsOnCall map[int]struct {
		result1 *int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerRuntime) Delete(arg1 context.Context) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeContainerRuntime) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeContainerRuntime) DeleteCalls(stub func(context.Context) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeContainerRuntime) DeleteArgsForCall(i int) context.Context {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainerRuntime) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) DeleteContainerIfExists(arg1 context.Context, arg2 string) error {
	fake.deleteContainerIfExistsMutex.Lock()
	ret, specificReturn := fake.deleteContainerIfExistsReturnsOnCall[len(fake.deleteContainerIfExistsArgsForCall)]
	fake.deleteContainerIfExistsArgsForCall = append(fake.deleteContainerIfExistsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DeleteContainerIfExists", []interface{}{arg1, arg2})
	fake.deleteContainerIfExistsMutex.Unlock()
	if fake.DeleteContainerIfExistsStub != nil {
		return fake.DeleteContainerIfExistsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteContainerIfExistsReturns
	return fakeReturns.result1
}

func (fake *FakeContainerRuntime) DeleteContainerIfExistsCallCount() int {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	return len(fake.deleteContainerIfExistsArgsForCall)
}

func (fake *FakeContainerRuntime) DeleteContainerIfExistsCalls(stub func(context.Context, string) error) {
	fake.deleteContainerIfExistsMutex.Lock()
	defer fake.deleteContainerIfExistsMutex.Unlock()
	fake.DeleteContainerIfExistsStub = stub
}

func (fake *FakeContainerRuntime) DeleteContainerIfExistsArgsForCall(i int) (context.Context, string) {
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	argsForCall := fake.deleteContainerIfExistsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerRuntime) DeleteContainerIfExistsReturns(result1 error) {
	fake.deleteContainerIfExistsMutex.Lock()
	defer fake.deleteContainerIfExistsMutex.Unlock()
	fake.DeleteContainerIfExistsStub = nil
	fake.deleteContainerIfExistsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) DeleteContainerIfExistsReturnsOnCall(i int, result1 error) {
	fake.deleteContainerIfExistsMutex.Lock()
	defer fake.deleteContainerIfExistsMutex.Unlock()
	fake.DeleteContainerIfExistsStub = nil
	if fake.deleteContainerIfExistsReturnsOnCall == nil {
		fake.deleteContainerIfExistsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteContainerIfExistsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) Kill(arg1 context.Context) error {
	fake.killMutex.Lock()
	ret, specificReturn := fake.killReturnsOnCall[len(fake.killArgsForCall)]
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Kill", []interface{}{arg1})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		return fake.KillStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.killReturns
	return fakeReturns.result1
}

func (fake *FakeContainerRuntime) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeContainerRuntime) KillCalls(stub func(context.Context) error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = stub
}

func (fake *FakeContainerRuntime) KillArgsForCall(i int) context.Context {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	argsForCall := fake.killArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainerRuntime) KillReturns(result1 error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) KillReturnsOnCall(i int, result1 error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = nil
	if fake.killReturnsOnCall == nil {
		fake.killReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.killReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainerRuntime) RunContainer(arg1 context.Context, arg2 chan model.Event, arg3 *model.ContainerCall, arg4 string, arg5 io.WriteCloser, arg6 io.WriteCloser) (*int64, error) {
	fake.runContainerMutex.Lock()
	ret, specificReturn := fake.runContainerReturnsOnCall[len(fake.runContainerArgsForCall)]
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 *model.ContainerCall
		arg4 string
		arg5 io.WriteCloser
		arg6 io.WriteCloser
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("RunContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContainerRuntime) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeContainerRuntime) RunContainerCalls(stub func(context.Context, chan model.Event, *model.ContainerCall, string, io.WriteCloser, io.WriteCloser) (*int64, error)) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = stub
}

func (fake *FakeContainerRuntime) RunContainerArgsForCall(i int) (context.Context, chan model.Event, *model.ContainerCall, string, io.WriteCloser, io.WriteCloser) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	argsForCall := fake.runContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeContainerRuntime) RunContainerReturns(result1 *int64, result2 error) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 *int64
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerRuntime) RunContainerReturnsOnCall(i int, result1 *int64, result2 error) {
	fake.runContainerMutex.Lock()
	defer fake.runContainerMutex.Unlock()
	fake.RunContainerStub = nil
	if fake.runContainerReturnsOnCall == nil {
		fake.runContainerReturnsOnCall = make(map[int]struct {
			result1 *int64
			result2 error
		})
	}
	fake.runContainerReturnsOnCall[i] = struct {
		result1 *int64
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteContainerIfExistsMutex.RLock()
	defer fake.deleteContainerIfExistsMutex.RUnlock()
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerRuntime) recordInvocation(key string, args []interface{}) {
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

var _ containerruntime.ContainerRuntime = new(FakeContainerRuntime)
