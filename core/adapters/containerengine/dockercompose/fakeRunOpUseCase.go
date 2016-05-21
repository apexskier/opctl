// This file was generated by counterfeiter
package dockercompose

import (
	"sync"

	"github.com/open-devops/engine/core/logging"
)

type fakeRunOpUseCase struct {
	ExecuteStub        func(correlationId string, pathToOpDir string, opName string, logger logging.Logger) (exitCode int, err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		correlationId string
		pathToOpDir   string
		opName        string
		logger        logging.Logger
	}
	executeReturns struct {
		result1 int
		result2 error
	}
}

func (fake *fakeRunOpUseCase) Execute(correlationId string, pathToOpDir string, opName string, logger logging.Logger) (exitCode int, err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		correlationId string
		pathToOpDir   string
		opName        string
		logger        logging.Logger
	}{correlationId, pathToOpDir, opName, logger})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(correlationId, pathToOpDir, opName, logger)
	} else {
		return fake.executeReturns.result1, fake.executeReturns.result2
	}
}

func (fake *fakeRunOpUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeRunOpUseCase) ExecuteArgsForCall(i int) (string, string, string, logging.Logger) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].correlationId, fake.executeArgsForCall[i].pathToOpDir, fake.executeArgsForCall[i].opName, fake.executeArgsForCall[i].logger
}

func (fake *fakeRunOpUseCase) ExecuteReturns(result1 int, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}
