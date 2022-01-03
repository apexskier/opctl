// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/node"
)

type FakeCore struct {
	GetDataStub        func(context.Context, chan model.Event, string, model.GetDataReq) (model.ReadSeekCloser, error)
	getDataMutex       sync.RWMutex
	getDataArgsForCall []struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 model.GetDataReq
	}
	getDataReturns struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	getDataReturnsOnCall map[int]struct {
		result1 model.ReadSeekCloser
		result2 error
	}
	ListDescendantsStub        func(context.Context, chan model.Event, string, model.ListDescendantsReq) ([]*model.DirEntry, error)
	listDescendantsMutex       sync.RWMutex
	listDescendantsArgsForCall []struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 model.ListDescendantsReq
	}
	listDescendantsReturns struct {
		result1 []*model.DirEntry
		result2 error
	}
	listDescendantsReturnsOnCall map[int]struct {
		result1 []*model.DirEntry
		result2 error
	}
	ResolveDataStub        func(context.Context, chan model.Event, string, string) (model.DataHandle, error)
	resolveDataMutex       sync.RWMutex
	resolveDataArgsForCall []struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 string
	}
	resolveDataReturns struct {
		result1 model.DataHandle
		result2 error
	}
	resolveDataReturnsOnCall map[int]struct {
		result1 model.DataHandle
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

func (fake *FakeCore) GetData(arg1 context.Context, arg2 chan model.Event, arg3 string, arg4 model.GetDataReq) (model.ReadSeekCloser, error) {
	fake.getDataMutex.Lock()
	ret, specificReturn := fake.getDataReturnsOnCall[len(fake.getDataArgsForCall)]
	fake.getDataArgsForCall = append(fake.getDataArgsForCall, struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 model.GetDataReq
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetData", []interface{}{arg1, arg2, arg3, arg4})
	fake.getDataMutex.Unlock()
	if fake.GetDataStub != nil {
		return fake.GetDataStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCore) GetDataCallCount() int {
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	return len(fake.getDataArgsForCall)
}

func (fake *FakeCore) GetDataCalls(stub func(context.Context, chan model.Event, string, model.GetDataReq) (model.ReadSeekCloser, error)) {
	fake.getDataMutex.Lock()
	defer fake.getDataMutex.Unlock()
	fake.GetDataStub = stub
}

func (fake *FakeCore) GetDataArgsForCall(i int) (context.Context, chan model.Event, string, model.GetDataReq) {
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	argsForCall := fake.getDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCore) GetDataReturns(result1 model.ReadSeekCloser, result2 error) {
	fake.getDataMutex.Lock()
	defer fake.getDataMutex.Unlock()
	fake.GetDataStub = nil
	fake.getDataReturns = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) GetDataReturnsOnCall(i int, result1 model.ReadSeekCloser, result2 error) {
	fake.getDataMutex.Lock()
	defer fake.getDataMutex.Unlock()
	fake.GetDataStub = nil
	if fake.getDataReturnsOnCall == nil {
		fake.getDataReturnsOnCall = make(map[int]struct {
			result1 model.ReadSeekCloser
			result2 error
		})
	}
	fake.getDataReturnsOnCall[i] = struct {
		result1 model.ReadSeekCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) ListDescendants(arg1 context.Context, arg2 chan model.Event, arg3 string, arg4 model.ListDescendantsReq) ([]*model.DirEntry, error) {
	fake.listDescendantsMutex.Lock()
	ret, specificReturn := fake.listDescendantsReturnsOnCall[len(fake.listDescendantsArgsForCall)]
	fake.listDescendantsArgsForCall = append(fake.listDescendantsArgsForCall, struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 model.ListDescendantsReq
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ListDescendants", []interface{}{arg1, arg2, arg3, arg4})
	fake.listDescendantsMutex.Unlock()
	if fake.ListDescendantsStub != nil {
		return fake.ListDescendantsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listDescendantsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCore) ListDescendantsCallCount() int {
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	return len(fake.listDescendantsArgsForCall)
}

func (fake *FakeCore) ListDescendantsCalls(stub func(context.Context, chan model.Event, string, model.ListDescendantsReq) ([]*model.DirEntry, error)) {
	fake.listDescendantsMutex.Lock()
	defer fake.listDescendantsMutex.Unlock()
	fake.ListDescendantsStub = stub
}

func (fake *FakeCore) ListDescendantsArgsForCall(i int) (context.Context, chan model.Event, string, model.ListDescendantsReq) {
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	argsForCall := fake.listDescendantsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCore) ListDescendantsReturns(result1 []*model.DirEntry, result2 error) {
	fake.listDescendantsMutex.Lock()
	defer fake.listDescendantsMutex.Unlock()
	fake.ListDescendantsStub = nil
	fake.listDescendantsReturns = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) ListDescendantsReturnsOnCall(i int, result1 []*model.DirEntry, result2 error) {
	fake.listDescendantsMutex.Lock()
	defer fake.listDescendantsMutex.Unlock()
	fake.ListDescendantsStub = nil
	if fake.listDescendantsReturnsOnCall == nil {
		fake.listDescendantsReturnsOnCall = make(map[int]struct {
			result1 []*model.DirEntry
			result2 error
		})
	}
	fake.listDescendantsReturnsOnCall[i] = struct {
		result1 []*model.DirEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) ResolveData(arg1 context.Context, arg2 chan model.Event, arg3 string, arg4 string) (model.DataHandle, error) {
	fake.resolveDataMutex.Lock()
	ret, specificReturn := fake.resolveDataReturnsOnCall[len(fake.resolveDataArgsForCall)]
	fake.resolveDataArgsForCall = append(fake.resolveDataArgsForCall, struct {
		arg1 context.Context
		arg2 chan model.Event
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("ResolveData", []interface{}{arg1, arg2, arg3, arg4})
	fake.resolveDataMutex.Unlock()
	if fake.ResolveDataStub != nil {
		return fake.ResolveDataStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveDataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCore) ResolveDataCallCount() int {
	fake.resolveDataMutex.RLock()
	defer fake.resolveDataMutex.RUnlock()
	return len(fake.resolveDataArgsForCall)
}

func (fake *FakeCore) ResolveDataCalls(stub func(context.Context, chan model.Event, string, string) (model.DataHandle, error)) {
	fake.resolveDataMutex.Lock()
	defer fake.resolveDataMutex.Unlock()
	fake.ResolveDataStub = stub
}

func (fake *FakeCore) ResolveDataArgsForCall(i int) (context.Context, chan model.Event, string, string) {
	fake.resolveDataMutex.RLock()
	defer fake.resolveDataMutex.RUnlock()
	argsForCall := fake.resolveDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCore) ResolveDataReturns(result1 model.DataHandle, result2 error) {
	fake.resolveDataMutex.Lock()
	defer fake.resolveDataMutex.Unlock()
	fake.ResolveDataStub = nil
	fake.resolveDataReturns = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) ResolveDataReturnsOnCall(i int, result1 model.DataHandle, result2 error) {
	fake.resolveDataMutex.Lock()
	defer fake.resolveDataMutex.Unlock()
	fake.ResolveDataStub = nil
	if fake.resolveDataReturnsOnCall == nil {
		fake.resolveDataReturnsOnCall = make(map[int]struct {
			result1 model.DataHandle
			result2 error
		})
	}
	fake.resolveDataReturnsOnCall[i] = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) StartOp(arg1 context.Context, arg2 chan model.Event, arg3 model.StartOpReq) (map[string]*model.Value, error) {
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

func (fake *FakeCore) StartOpCallCount() int {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return len(fake.startOpArgsForCall)
}

func (fake *FakeCore) StartOpCalls(stub func(context.Context, chan model.Event, model.StartOpReq) (map[string]*model.Value, error)) {
	fake.startOpMutex.Lock()
	defer fake.startOpMutex.Unlock()
	fake.StartOpStub = stub
}

func (fake *FakeCore) StartOpArgsForCall(i int) (context.Context, chan model.Event, model.StartOpReq) {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	argsForCall := fake.startOpArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCore) StartOpReturns(result1 map[string]*model.Value, result2 error) {
	fake.startOpMutex.Lock()
	defer fake.startOpMutex.Unlock()
	fake.StartOpStub = nil
	fake.startOpReturns = struct {
		result1 map[string]*model.Value
		result2 error
	}{result1, result2}
}

func (fake *FakeCore) StartOpReturnsOnCall(i int, result1 map[string]*model.Value, result2 error) {
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

func (fake *FakeCore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDataMutex.RLock()
	defer fake.getDataMutex.RUnlock()
	fake.listDescendantsMutex.RLock()
	defer fake.listDescendantsMutex.RUnlock()
	fake.resolveDataMutex.RLock()
	defer fake.resolveDataMutex.RUnlock()
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCore) recordInvocation(key string, args []interface{}) {
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

var _ node.Core = new(FakeCore)
