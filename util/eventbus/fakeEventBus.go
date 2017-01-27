// This file was generated by counterfeiter
package eventbus

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type FakeEventBus struct {
	PublishStub        func(event model.Event)
	publishMutex       sync.RWMutex
	publishArgsForCall []struct {
		event model.Event
	}
	RegisterSubscriberStub        func(filter *model.EventFilter, eventChannel chan model.Event)
	registerSubscriberMutex       sync.RWMutex
	registerSubscriberArgsForCall []struct {
		filter       *model.EventFilter
		eventChannel chan model.Event
	}
	UnregisterSubscriberStub        func(eventChannel chan model.Event)
	unregisterSubscriberMutex       sync.RWMutex
	unregisterSubscriberArgsForCall []struct {
		eventChannel chan model.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventBus) Publish(event model.Event) {
	fake.publishMutex.Lock()
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		event model.Event
	}{event})
	fake.recordInvocation("Publish", []interface{}{event})
	fake.publishMutex.Unlock()
	if fake.PublishStub != nil {
		fake.PublishStub(event)
	}
}

func (fake *FakeEventBus) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakeEventBus) PublishArgsForCall(i int) model.Event {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return fake.publishArgsForCall[i].event
}

func (fake *FakeEventBus) RegisterSubscriber(filter *model.EventFilter, eventChannel chan model.Event) {
	fake.registerSubscriberMutex.Lock()
	fake.registerSubscriberArgsForCall = append(fake.registerSubscriberArgsForCall, struct {
		filter       *model.EventFilter
		eventChannel chan model.Event
	}{filter, eventChannel})
	fake.recordInvocation("RegisterSubscriber", []interface{}{filter, eventChannel})
	fake.registerSubscriberMutex.Unlock()
	if fake.RegisterSubscriberStub != nil {
		fake.RegisterSubscriberStub(filter, eventChannel)
	}
}

func (fake *FakeEventBus) RegisterSubscriberCallCount() int {
	fake.registerSubscriberMutex.RLock()
	defer fake.registerSubscriberMutex.RUnlock()
	return len(fake.registerSubscriberArgsForCall)
}

func (fake *FakeEventBus) RegisterSubscriberArgsForCall(i int) (*model.EventFilter, chan model.Event) {
	fake.registerSubscriberMutex.RLock()
	defer fake.registerSubscriberMutex.RUnlock()
	return fake.registerSubscriberArgsForCall[i].filter, fake.registerSubscriberArgsForCall[i].eventChannel
}

func (fake *FakeEventBus) UnregisterSubscriber(eventChannel chan model.Event) {
	fake.unregisterSubscriberMutex.Lock()
	fake.unregisterSubscriberArgsForCall = append(fake.unregisterSubscriberArgsForCall, struct {
		eventChannel chan model.Event
	}{eventChannel})
	fake.recordInvocation("UnregisterSubscriber", []interface{}{eventChannel})
	fake.unregisterSubscriberMutex.Unlock()
	if fake.UnregisterSubscriberStub != nil {
		fake.UnregisterSubscriberStub(eventChannel)
	}
}

func (fake *FakeEventBus) UnregisterSubscriberCallCount() int {
	fake.unregisterSubscriberMutex.RLock()
	defer fake.unregisterSubscriberMutex.RUnlock()
	return len(fake.unregisterSubscriberArgsForCall)
}

func (fake *FakeEventBus) UnregisterSubscriberArgsForCall(i int) chan model.Event {
	fake.unregisterSubscriberMutex.RLock()
	defer fake.unregisterSubscriberMutex.RUnlock()
	return fake.unregisterSubscriberArgsForCall[i].eventChannel
}

func (fake *FakeEventBus) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	fake.registerSubscriberMutex.RLock()
	defer fake.registerSubscriberMutex.RUnlock()
	fake.unregisterSubscriberMutex.RLock()
	defer fake.unregisterSubscriberMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEventBus) recordInvocation(key string, args []interface{}) {
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

var _ EventBus = new(FakeEventBus)
