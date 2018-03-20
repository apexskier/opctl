// Code generated by counterfeiter. DO NOT EDIT.
package data

import (
	"context"
	"net/url"
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	NewFSProviderStub        func(basePaths ...string) Provider
	newFSProviderMutex       sync.RWMutex
	newFSProviderArgsForCall []struct {
		basePaths []string
	}
	newFSProviderReturns struct {
		result1 Provider
	}
	newFSProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewGitProviderStub        func(basePath string, pullCreds *model.PullCreds) Provider
	newGitProviderMutex       sync.RWMutex
	newGitProviderArgsForCall []struct {
		basePath  string
		pullCreds *model.PullCreds
	}
	newGitProviderReturns struct {
		result1 Provider
	}
	newGitProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	NewNodeProviderStub        func(apiBaseURL url.URL, pullCreds *model.PullCreds) Provider
	newNodeProviderMutex       sync.RWMutex
	newNodeProviderArgsForCall []struct {
		apiBaseURL url.URL
		pullCreds  *model.PullCreds
	}
	newNodeProviderReturns struct {
		result1 Provider
	}
	newNodeProviderReturnsOnCall map[int]struct {
		result1 Provider
	}
	ResolveStub        func(ctx context.Context, dataRef string, providers ...Provider) (model.DataHandle, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		ctx       context.Context
		dataRef   string
		providers []Provider
	}
	resolveReturns struct {
		result1 model.DataHandle
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 model.DataHandle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) NewFSProvider(basePaths ...string) Provider {
	fake.newFSProviderMutex.Lock()
	ret, specificReturn := fake.newFSProviderReturnsOnCall[len(fake.newFSProviderArgsForCall)]
	fake.newFSProviderArgsForCall = append(fake.newFSProviderArgsForCall, struct {
		basePaths []string
	}{basePaths})
	fake.recordInvocation("NewFSProvider", []interface{}{basePaths})
	fake.newFSProviderMutex.Unlock()
	if fake.NewFSProviderStub != nil {
		return fake.NewFSProviderStub(basePaths...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newFSProviderReturns.result1
}

func (fake *Fake) NewFSProviderCallCount() int {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	return len(fake.newFSProviderArgsForCall)
}

func (fake *Fake) NewFSProviderArgsForCall(i int) []string {
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	return fake.newFSProviderArgsForCall[i].basePaths
}

func (fake *Fake) NewFSProviderReturns(result1 Provider) {
	fake.NewFSProviderStub = nil
	fake.newFSProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewFSProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewFSProviderStub = nil
	if fake.newFSProviderReturnsOnCall == nil {
		fake.newFSProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newFSProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProvider(basePath string, pullCreds *model.PullCreds) Provider {
	fake.newGitProviderMutex.Lock()
	ret, specificReturn := fake.newGitProviderReturnsOnCall[len(fake.newGitProviderArgsForCall)]
	fake.newGitProviderArgsForCall = append(fake.newGitProviderArgsForCall, struct {
		basePath  string
		pullCreds *model.PullCreds
	}{basePath, pullCreds})
	fake.recordInvocation("NewGitProvider", []interface{}{basePath, pullCreds})
	fake.newGitProviderMutex.Unlock()
	if fake.NewGitProviderStub != nil {
		return fake.NewGitProviderStub(basePath, pullCreds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newGitProviderReturns.result1
}

func (fake *Fake) NewGitProviderCallCount() int {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	return len(fake.newGitProviderArgsForCall)
}

func (fake *Fake) NewGitProviderArgsForCall(i int) (string, *model.PullCreds) {
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	return fake.newGitProviderArgsForCall[i].basePath, fake.newGitProviderArgsForCall[i].pullCreds
}

func (fake *Fake) NewGitProviderReturns(result1 Provider) {
	fake.NewGitProviderStub = nil
	fake.newGitProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewGitProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewGitProviderStub = nil
	if fake.newGitProviderReturnsOnCall == nil {
		fake.newGitProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newGitProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProvider(apiBaseURL url.URL, pullCreds *model.PullCreds) Provider {
	fake.newNodeProviderMutex.Lock()
	ret, specificReturn := fake.newNodeProviderReturnsOnCall[len(fake.newNodeProviderArgsForCall)]
	fake.newNodeProviderArgsForCall = append(fake.newNodeProviderArgsForCall, struct {
		apiBaseURL url.URL
		pullCreds  *model.PullCreds
	}{apiBaseURL, pullCreds})
	fake.recordInvocation("NewNodeProvider", []interface{}{apiBaseURL, pullCreds})
	fake.newNodeProviderMutex.Unlock()
	if fake.NewNodeProviderStub != nil {
		return fake.NewNodeProviderStub(apiBaseURL, pullCreds)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newNodeProviderReturns.result1
}

func (fake *Fake) NewNodeProviderCallCount() int {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	return len(fake.newNodeProviderArgsForCall)
}

func (fake *Fake) NewNodeProviderArgsForCall(i int) (url.URL, *model.PullCreds) {
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	return fake.newNodeProviderArgsForCall[i].apiBaseURL, fake.newNodeProviderArgsForCall[i].pullCreds
}

func (fake *Fake) NewNodeProviderReturns(result1 Provider) {
	fake.NewNodeProviderStub = nil
	fake.newNodeProviderReturns = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) NewNodeProviderReturnsOnCall(i int, result1 Provider) {
	fake.NewNodeProviderStub = nil
	if fake.newNodeProviderReturnsOnCall == nil {
		fake.newNodeProviderReturnsOnCall = make(map[int]struct {
			result1 Provider
		})
	}
	fake.newNodeProviderReturnsOnCall[i] = struct {
		result1 Provider
	}{result1}
}

func (fake *Fake) Resolve(ctx context.Context, dataRef string, providers ...Provider) (model.DataHandle, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		ctx       context.Context
		dataRef   string
		providers []Provider
	}{ctx, dataRef, providers})
	fake.recordInvocation("Resolve", []interface{}{ctx, dataRef, providers})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(ctx, dataRef, providers...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.resolveReturns.result1, fake.resolveReturns.result2
}

func (fake *Fake) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *Fake) ResolveArgsForCall(i int) (context.Context, string, []Provider) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return fake.resolveArgsForCall[i].ctx, fake.resolveArgsForCall[i].dataRef, fake.resolveArgsForCall[i].providers
}

func (fake *Fake) ResolveReturns(result1 model.DataHandle, result2 error) {
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) ResolveReturnsOnCall(i int, result1 model.DataHandle, result2 error) {
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 model.DataHandle
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 model.DataHandle
		result2 error
	}{result1, result2}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFSProviderMutex.RLock()
	defer fake.newFSProviderMutex.RUnlock()
	fake.newGitProviderMutex.RLock()
	defer fake.newGitProviderMutex.RUnlock()
	fake.newNodeProviderMutex.RLock()
	defer fake.newNodeProviderMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
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

var _ Data = new(Fake)
