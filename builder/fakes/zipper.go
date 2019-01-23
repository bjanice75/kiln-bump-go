// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"io"
	"os"
	"sync"
)

type Zipper struct {
	AddStub        func(string, io.Reader) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 string
		arg2 io.Reader
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	AddWithModeStub        func(string, io.Reader, os.FileMode) error
	addWithModeMutex       sync.RWMutex
	addWithModeArgsForCall []struct {
		arg1 string
		arg2 io.Reader
		arg3 os.FileMode
	}
	addWithModeReturns struct {
		result1 error
	}
	addWithModeReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CreateFolderStub        func(string) error
	createFolderMutex       sync.RWMutex
	createFolderArgsForCall []struct {
		arg1 string
	}
	createFolderReturns struct {
		result1 error
	}
	createFolderReturnsOnCall map[int]struct {
		result1 error
	}
	SetWriterStub        func(io.Writer)
	setWriterMutex       sync.RWMutex
	setWriterArgsForCall []struct {
		arg1 io.Writer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Zipper) Add(arg1 string, arg2 io.Reader) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 string
		arg2 io.Reader
	}{arg1, arg2})
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addReturns
	return fakeReturns.result1
}

func (fake *Zipper) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Zipper) AddCalls(stub func(string, io.Reader) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *Zipper) AddArgsForCall(i int) (string, io.Reader) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Zipper) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddWithMode(arg1 string, arg2 io.Reader, arg3 os.FileMode) error {
	fake.addWithModeMutex.Lock()
	ret, specificReturn := fake.addWithModeReturnsOnCall[len(fake.addWithModeArgsForCall)]
	fake.addWithModeArgsForCall = append(fake.addWithModeArgsForCall, struct {
		arg1 string
		arg2 io.Reader
		arg3 os.FileMode
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddWithMode", []interface{}{arg1, arg2, arg3})
	fake.addWithModeMutex.Unlock()
	if fake.AddWithModeStub != nil {
		return fake.AddWithModeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addWithModeReturns
	return fakeReturns.result1
}

func (fake *Zipper) AddWithModeCallCount() int {
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	return len(fake.addWithModeArgsForCall)
}

func (fake *Zipper) AddWithModeCalls(stub func(string, io.Reader, os.FileMode) error) {
	fake.addWithModeMutex.Lock()
	defer fake.addWithModeMutex.Unlock()
	fake.AddWithModeStub = stub
}

func (fake *Zipper) AddWithModeArgsForCall(i int) (string, io.Reader, os.FileMode) {
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	argsForCall := fake.addWithModeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Zipper) AddWithModeReturns(result1 error) {
	fake.addWithModeMutex.Lock()
	defer fake.addWithModeMutex.Unlock()
	fake.AddWithModeStub = nil
	fake.addWithModeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) AddWithModeReturnsOnCall(i int, result1 error) {
	fake.addWithModeMutex.Lock()
	defer fake.addWithModeMutex.Unlock()
	fake.AddWithModeStub = nil
	if fake.addWithModeReturnsOnCall == nil {
		fake.addWithModeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addWithModeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *Zipper) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Zipper) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *Zipper) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CreateFolder(arg1 string) error {
	fake.createFolderMutex.Lock()
	ret, specificReturn := fake.createFolderReturnsOnCall[len(fake.createFolderArgsForCall)]
	fake.createFolderArgsForCall = append(fake.createFolderArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CreateFolder", []interface{}{arg1})
	fake.createFolderMutex.Unlock()
	if fake.CreateFolderStub != nil {
		return fake.CreateFolderStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createFolderReturns
	return fakeReturns.result1
}

func (fake *Zipper) CreateFolderCallCount() int {
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	return len(fake.createFolderArgsForCall)
}

func (fake *Zipper) CreateFolderCalls(stub func(string) error) {
	fake.createFolderMutex.Lock()
	defer fake.createFolderMutex.Unlock()
	fake.CreateFolderStub = stub
}

func (fake *Zipper) CreateFolderArgsForCall(i int) string {
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	argsForCall := fake.createFolderArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Zipper) CreateFolderReturns(result1 error) {
	fake.createFolderMutex.Lock()
	defer fake.createFolderMutex.Unlock()
	fake.CreateFolderStub = nil
	fake.createFolderReturns = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) CreateFolderReturnsOnCall(i int, result1 error) {
	fake.createFolderMutex.Lock()
	defer fake.createFolderMutex.Unlock()
	fake.CreateFolderStub = nil
	if fake.createFolderReturnsOnCall == nil {
		fake.createFolderReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createFolderReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Zipper) SetWriter(arg1 io.Writer) {
	fake.setWriterMutex.Lock()
	fake.setWriterArgsForCall = append(fake.setWriterArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.recordInvocation("SetWriter", []interface{}{arg1})
	fake.setWriterMutex.Unlock()
	if fake.SetWriterStub != nil {
		fake.SetWriterStub(arg1)
	}
}

func (fake *Zipper) SetWriterCallCount() int {
	fake.setWriterMutex.RLock()
	defer fake.setWriterMutex.RUnlock()
	return len(fake.setWriterArgsForCall)
}

func (fake *Zipper) SetWriterCalls(stub func(io.Writer)) {
	fake.setWriterMutex.Lock()
	defer fake.setWriterMutex.Unlock()
	fake.SetWriterStub = stub
}

func (fake *Zipper) SetWriterArgsForCall(i int) io.Writer {
	fake.setWriterMutex.RLock()
	defer fake.setWriterMutex.RUnlock()
	argsForCall := fake.setWriterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Zipper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.addWithModeMutex.RLock()
	defer fake.addWithModeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.createFolderMutex.RLock()
	defer fake.createFolderMutex.RUnlock()
	fake.setWriterMutex.RLock()
	defer fake.setWriterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Zipper) recordInvocation(key string, args []interface{}) {
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
