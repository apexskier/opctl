// Code generated by counterfeiter. DO NOT EDIT.
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	GetEventStreamStub        func(req *model.GetEventStreamReq, eventChannel chan *model.Event) error
	getEventStreamMutex       sync.RWMutex
	getEventStreamArgsForCall []struct {
		req          *model.GetEventStreamReq
		eventChannel chan *model.Event
	}
	getEventStreamReturns struct {
		result1 error
	}
	getEventStreamReturnsOnCall map[int]struct {
		result1 error
	}
	KillOpStub        func(req model.KillOpReq)
	killOpMutex       sync.RWMutex
	killOpArgsForCall []struct {
		req model.KillOpReq
	}
	StartOpStub        func(req model.StartOpReq) (callId string, err error)
	startOpMutex       sync.RWMutex
	startOpArgsForCall []struct {
		req model.StartOpReq
	}
	startOpReturns struct {
		result1 string
		result2 error
	}
	startOpReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ResolvePkgStub        func(pkgRef string, pullCreds *model.PullCreds) (model.PkgHandle, error)
	resolvePkgMutex       sync.RWMutex
	resolvePkgArgsForCall []struct {
		pkgRef    string
		pullCreds *model.PullCreds
	}
	resolvePkgReturns struct {
		result1 model.PkgHandle
		result2 error
	}
	resolvePkgReturnsOnCall map[int]struct {
		result1 model.PkgHandle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) GetEventStream(req *model.GetEventStreamReq, eventChannel chan *model.Event) error {
	fake.getEventStreamMutex.Lock()
	ret, specificReturn := fake.getEventStreamReturnsOnCall[len(fake.getEventStreamArgsForCall)]
	fake.getEventStreamArgsForCall = append(fake.getEventStreamArgsForCall, struct {
		req          *model.GetEventStreamReq
		eventChannel chan *model.Event
	}{req, eventChannel})
	fake.recordInvocation("GetEventStream", []interface{}{req, eventChannel})
	fake.getEventStreamMutex.Unlock()
	if fake.GetEventStreamStub != nil {
		return fake.GetEventStreamStub(req, eventChannel)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getEventStreamReturns.result1
}

func (fake *Fake) GetEventStreamCallCount() int {
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	return len(fake.getEventStreamArgsForCall)
}

func (fake *Fake) GetEventStreamArgsForCall(i int) (*model.GetEventStreamReq, chan *model.Event) {
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	return fake.getEventStreamArgsForCall[i].req, fake.getEventStreamArgsForCall[i].eventChannel
}

func (fake *Fake) GetEventStreamReturns(result1 error) {
	fake.GetEventStreamStub = nil
	fake.getEventStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *Fake) GetEventStreamReturnsOnCall(i int, result1 error) {
	fake.GetEventStreamStub = nil
	if fake.getEventStreamReturnsOnCall == nil {
		fake.getEventStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getEventStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Fake) KillOp(req model.KillOpReq) {
	fake.killOpMutex.Lock()
	fake.killOpArgsForCall = append(fake.killOpArgsForCall, struct {
		req model.KillOpReq
	}{req})
	fake.recordInvocation("KillOp", []interface{}{req})
	fake.killOpMutex.Unlock()
	if fake.KillOpStub != nil {
		fake.KillOpStub(req)
	}
}

func (fake *Fake) KillOpCallCount() int {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return len(fake.killOpArgsForCall)
}

func (fake *Fake) KillOpArgsForCall(i int) model.KillOpReq {
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	return fake.killOpArgsForCall[i].req
}

func (fake *Fake) StartOp(req model.StartOpReq) (callId string, err error) {
	fake.startOpMutex.Lock()
	ret, specificReturn := fake.startOpReturnsOnCall[len(fake.startOpArgsForCall)]
	fake.startOpArgsForCall = append(fake.startOpArgsForCall, struct {
		req model.StartOpReq
	}{req})
	fake.recordInvocation("StartOp", []interface{}{req})
	fake.startOpMutex.Unlock()
	if fake.StartOpStub != nil {
		return fake.StartOpStub(req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startOpReturns.result1, fake.startOpReturns.result2
}

func (fake *Fake) StartOpCallCount() int {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return len(fake.startOpArgsForCall)
}

func (fake *Fake) StartOpArgsForCall(i int) model.StartOpReq {
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	return fake.startOpArgsForCall[i].req
}

func (fake *Fake) StartOpReturns(result1 string, result2 error) {
	fake.StartOpStub = nil
	fake.startOpReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) StartOpReturnsOnCall(i int, result1 string, result2 error) {
	fake.StartOpStub = nil
	if fake.startOpReturnsOnCall == nil {
		fake.startOpReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startOpReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *Fake) ResolvePkg(pkgRef string, pullCreds *model.PullCreds) (model.PkgHandle, error) {
	fake.resolvePkgMutex.Lock()
	ret, specificReturn := fake.resolvePkgReturnsOnCall[len(fake.resolvePkgArgsForCall)]
	fake.resolvePkgArgsForCall = append(fake.resolvePkgArgsForCall, struct {
		pkgRef    string
		pullCreds *model.PullCreds
	}{pkgRef, pullCreds})
	fake.recordInvocation("ResolvePkg", []interface{}{pkgRef, pullCreds})
	fake.resolvePkgMutex.Unlock()
	if fake.ResolvePkgStub != nil {
		return fake.ResolvePkgStub(pkgRef, pullCreds)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resolvePkgReturns.result1, fake.resolvePkgReturns.result2
}

func (fake *Fake) ResolvePkgCallCount() int {
	fake.resolvePkgMutex.RLock()
	defer fake.resolvePkgMutex.RUnlock()
	return len(fake.resolvePkgArgsForCall)
}

func (fake *Fake) ResolvePkgArgsForCall(i int) (string, *model.PullCreds) {
	fake.resolvePkgMutex.RLock()
	defer fake.resolvePkgMutex.RUnlock()
	return fake.resolvePkgArgsForCall[i].pkgRef, fake.resolvePkgArgsForCall[i].pullCreds
}

func (fake *Fake) ResolvePkgReturns(result1 model.PkgHandle, result2 error) {
	fake.ResolvePkgStub = nil
	fake.resolvePkgReturns = struct {
		result1 model.PkgHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) ResolvePkgReturnsOnCall(i int, result1 model.PkgHandle, result2 error) {
	fake.ResolvePkgStub = nil
	if fake.resolvePkgReturnsOnCall == nil {
		fake.resolvePkgReturnsOnCall = make(map[int]struct {
			result1 model.PkgHandle
			result2 error
		})
	}
	fake.resolvePkgReturnsOnCall[i] = struct {
		result1 model.PkgHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEventStreamMutex.RLock()
	defer fake.getEventStreamMutex.RUnlock()
	fake.killOpMutex.RLock()
	defer fake.killOpMutex.RUnlock()
	fake.startOpMutex.RLock()
	defer fake.startOpMutex.RUnlock()
	fake.resolvePkgMutex.RLock()
	defer fake.resolvePkgMutex.RUnlock()
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

var _ Core = new(Fake)