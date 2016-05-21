// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/open-devops/engine/core/models"
)

type fakeRunOpUseCase struct {
	ExecuteStub        func(req models.RunOpReq) (opRunId string, correlationId string, err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		req models.RunOpReq
	}
	executeReturns struct {
		result1 string
		result2 string
		result3 error
	}
}

func (fake *fakeRunOpUseCase) Execute(req models.RunOpReq) (opRunId string, correlationId string, err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		req models.RunOpReq
	}{req})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(req)
	} else {
		return fake.executeReturns.result1, fake.executeReturns.result2, fake.executeReturns.result3
	}
}

func (fake *fakeRunOpUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeRunOpUseCase) ExecuteArgsForCall(i int) models.RunOpReq {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].req
}

func (fake *fakeRunOpUseCase) ExecuteReturns(result1 string, result2 string, result3 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}
