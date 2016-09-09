// This file was generated by counterfeiter
package opspec

import (
  "sync"

  "github.com/opspec-io/sdk-golang/models"
)

type fakeTryResolveDefaultCollectionUseCase struct {
  ExecuteStub        func(req models.TryResolveDefaultCollectionReq) (pathToDefaultCollection string, err error)
  executeMutex       sync.RWMutex
  executeArgsForCall []struct {
    req models.TryResolveDefaultCollectionReq
  }
  executeReturns     struct {
                       result1 string
                       result2 error
                     }
  invocations        map[string][][]interface{}
  invocationsMutex   sync.RWMutex
}

func (fake *fakeTryResolveDefaultCollectionUseCase) Execute(req models.TryResolveDefaultCollectionReq) (pathToDefaultCollection string, err error) {
  fake.executeMutex.Lock()
  fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
    req models.TryResolveDefaultCollectionReq
  }{req})
  fake.recordInvocation("Execute", []interface{}{req})
  fake.executeMutex.Unlock()
  if fake.ExecuteStub != nil {
    return fake.ExecuteStub(req)
  } else {
    return fake.executeReturns.result1, fake.executeReturns.result2
  }
}

func (fake *fakeTryResolveDefaultCollectionUseCase) ExecuteCallCount() int {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return len(fake.executeArgsForCall)
}

func (fake *fakeTryResolveDefaultCollectionUseCase) ExecuteArgsForCall(i int) models.TryResolveDefaultCollectionReq {
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.executeArgsForCall[i].req
}

func (fake *fakeTryResolveDefaultCollectionUseCase) ExecuteReturns(result1 string, result2 error) {
  fake.ExecuteStub = nil
  fake.executeReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

func (fake *fakeTryResolveDefaultCollectionUseCase) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.executeMutex.RLock()
  defer fake.executeMutex.RUnlock()
  return fake.invocations
}

func (fake *fakeTryResolveDefaultCollectionUseCase) recordInvocation(key string, args []interface{}) {
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