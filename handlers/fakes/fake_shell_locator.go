// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
)

type FakeShellLocator struct {
	ShellPathStub        func() string
	shellPathMutex       sync.RWMutex
	shellPathArgsForCall []struct{}
	shellPathReturns     struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeShellLocator) ShellPath() string {
	fake.shellPathMutex.Lock()
	fake.shellPathArgsForCall = append(fake.shellPathArgsForCall, struct{}{})
	fake.recordInvocation("ShellPath", []interface{}{})
	fake.shellPathMutex.Unlock()
	if fake.ShellPathStub != nil {
		return fake.ShellPathStub()
	} else {
		return fake.shellPathReturns.result1
	}
}

func (fake *FakeShellLocator) ShellPathCallCount() int {
	fake.shellPathMutex.RLock()
	defer fake.shellPathMutex.RUnlock()
	return len(fake.shellPathArgsForCall)
}

func (fake *FakeShellLocator) ShellPathReturns(result1 string) {
	fake.ShellPathStub = nil
	fake.shellPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeShellLocator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.shellPathMutex.RLock()
	defer fake.shellPathMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeShellLocator) recordInvocation(key string, args []interface{}) {
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

var _ handlers.ShellLocator = new(FakeShellLocator)
