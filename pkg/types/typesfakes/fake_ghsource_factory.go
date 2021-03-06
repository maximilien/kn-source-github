// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/maximilien/kn-source-github/pkg/types"
	typesa "github.com/maximilien/kn-source-pkg/pkg/types"
	"k8s.io/client-go/rest"
)

type FakeGHSourceFactory struct {
	CreateGHSourceClientStub        func(*rest.Config, string) types.GHSourceClient
	createGHSourceClientMutex       sync.RWMutex
	createGHSourceClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 string
	}
	createGHSourceClientReturns struct {
		result1 types.GHSourceClient
	}
	createGHSourceClientReturnsOnCall map[int]struct {
		result1 types.GHSourceClient
	}
	CreateGHSourceParamsStub        func() *types.GHSourceParams
	createGHSourceParamsMutex       sync.RWMutex
	createGHSourceParamsArgsForCall []struct {
	}
	createGHSourceParamsReturns struct {
		result1 *types.GHSourceParams
	}
	createGHSourceParamsReturnsOnCall map[int]struct {
		result1 *types.GHSourceParams
	}
	CreateKnSourceClientStub        func(*rest.Config, string) typesa.KnSourceClient
	createKnSourceClientMutex       sync.RWMutex
	createKnSourceClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 string
	}
	createKnSourceClientReturns struct {
		result1 typesa.KnSourceClient
	}
	createKnSourceClientReturnsOnCall map[int]struct {
		result1 typesa.KnSourceClient
	}
	CreateKnSourceParamsStub        func() *typesa.KnSourceParams
	createKnSourceParamsMutex       sync.RWMutex
	createKnSourceParamsArgsForCall []struct {
	}
	createKnSourceParamsReturns struct {
		result1 *typesa.KnSourceParams
	}
	createKnSourceParamsReturnsOnCall map[int]struct {
		result1 *typesa.KnSourceParams
	}
	GHSourceClientStub        func() types.GHSourceClient
	gHSourceClientMutex       sync.RWMutex
	gHSourceClientArgsForCall []struct {
	}
	gHSourceClientReturns struct {
		result1 types.GHSourceClient
	}
	gHSourceClientReturnsOnCall map[int]struct {
		result1 types.GHSourceClient
	}
	GHSourceParamsStub        func() *types.GHSourceParams
	gHSourceParamsMutex       sync.RWMutex
	gHSourceParamsArgsForCall []struct {
	}
	gHSourceParamsReturns struct {
		result1 *types.GHSourceParams
	}
	gHSourceParamsReturnsOnCall map[int]struct {
		result1 *types.GHSourceParams
	}
	KnSourceParamsStub        func() *typesa.KnSourceParams
	knSourceParamsMutex       sync.RWMutex
	knSourceParamsArgsForCall []struct {
	}
	knSourceParamsReturns struct {
		result1 *typesa.KnSourceParams
	}
	knSourceParamsReturnsOnCall map[int]struct {
		result1 *typesa.KnSourceParams
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGHSourceFactory) CreateGHSourceClient(arg1 *rest.Config, arg2 string) types.GHSourceClient {
	fake.createGHSourceClientMutex.Lock()
	ret, specificReturn := fake.createGHSourceClientReturnsOnCall[len(fake.createGHSourceClientArgsForCall)]
	fake.createGHSourceClientArgsForCall = append(fake.createGHSourceClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateGHSourceClient", []interface{}{arg1, arg2})
	fake.createGHSourceClientMutex.Unlock()
	if fake.CreateGHSourceClientStub != nil {
		return fake.CreateGHSourceClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createGHSourceClientReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) CreateGHSourceClientCallCount() int {
	fake.createGHSourceClientMutex.RLock()
	defer fake.createGHSourceClientMutex.RUnlock()
	return len(fake.createGHSourceClientArgsForCall)
}

func (fake *FakeGHSourceFactory) CreateGHSourceClientCalls(stub func(*rest.Config, string) types.GHSourceClient) {
	fake.createGHSourceClientMutex.Lock()
	defer fake.createGHSourceClientMutex.Unlock()
	fake.CreateGHSourceClientStub = stub
}

func (fake *FakeGHSourceFactory) CreateGHSourceClientArgsForCall(i int) (*rest.Config, string) {
	fake.createGHSourceClientMutex.RLock()
	defer fake.createGHSourceClientMutex.RUnlock()
	argsForCall := fake.createGHSourceClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGHSourceFactory) CreateGHSourceClientReturns(result1 types.GHSourceClient) {
	fake.createGHSourceClientMutex.Lock()
	defer fake.createGHSourceClientMutex.Unlock()
	fake.CreateGHSourceClientStub = nil
	fake.createGHSourceClientReturns = struct {
		result1 types.GHSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateGHSourceClientReturnsOnCall(i int, result1 types.GHSourceClient) {
	fake.createGHSourceClientMutex.Lock()
	defer fake.createGHSourceClientMutex.Unlock()
	fake.CreateGHSourceClientStub = nil
	if fake.createGHSourceClientReturnsOnCall == nil {
		fake.createGHSourceClientReturnsOnCall = make(map[int]struct {
			result1 types.GHSourceClient
		})
	}
	fake.createGHSourceClientReturnsOnCall[i] = struct {
		result1 types.GHSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateGHSourceParams() *types.GHSourceParams {
	fake.createGHSourceParamsMutex.Lock()
	ret, specificReturn := fake.createGHSourceParamsReturnsOnCall[len(fake.createGHSourceParamsArgsForCall)]
	fake.createGHSourceParamsArgsForCall = append(fake.createGHSourceParamsArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateGHSourceParams", []interface{}{})
	fake.createGHSourceParamsMutex.Unlock()
	if fake.CreateGHSourceParamsStub != nil {
		return fake.CreateGHSourceParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createGHSourceParamsReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) CreateGHSourceParamsCallCount() int {
	fake.createGHSourceParamsMutex.RLock()
	defer fake.createGHSourceParamsMutex.RUnlock()
	return len(fake.createGHSourceParamsArgsForCall)
}

func (fake *FakeGHSourceFactory) CreateGHSourceParamsCalls(stub func() *types.GHSourceParams) {
	fake.createGHSourceParamsMutex.Lock()
	defer fake.createGHSourceParamsMutex.Unlock()
	fake.CreateGHSourceParamsStub = stub
}

func (fake *FakeGHSourceFactory) CreateGHSourceParamsReturns(result1 *types.GHSourceParams) {
	fake.createGHSourceParamsMutex.Lock()
	defer fake.createGHSourceParamsMutex.Unlock()
	fake.CreateGHSourceParamsStub = nil
	fake.createGHSourceParamsReturns = struct {
		result1 *types.GHSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateGHSourceParamsReturnsOnCall(i int, result1 *types.GHSourceParams) {
	fake.createGHSourceParamsMutex.Lock()
	defer fake.createGHSourceParamsMutex.Unlock()
	fake.CreateGHSourceParamsStub = nil
	if fake.createGHSourceParamsReturnsOnCall == nil {
		fake.createGHSourceParamsReturnsOnCall = make(map[int]struct {
			result1 *types.GHSourceParams
		})
	}
	fake.createGHSourceParamsReturnsOnCall[i] = struct {
		result1 *types.GHSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateKnSourceClient(arg1 *rest.Config, arg2 string) typesa.KnSourceClient {
	fake.createKnSourceClientMutex.Lock()
	ret, specificReturn := fake.createKnSourceClientReturnsOnCall[len(fake.createKnSourceClientArgsForCall)]
	fake.createKnSourceClientArgsForCall = append(fake.createKnSourceClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateKnSourceClient", []interface{}{arg1, arg2})
	fake.createKnSourceClientMutex.Unlock()
	if fake.CreateKnSourceClientStub != nil {
		return fake.CreateKnSourceClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createKnSourceClientReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) CreateKnSourceClientCallCount() int {
	fake.createKnSourceClientMutex.RLock()
	defer fake.createKnSourceClientMutex.RUnlock()
	return len(fake.createKnSourceClientArgsForCall)
}

func (fake *FakeGHSourceFactory) CreateKnSourceClientCalls(stub func(*rest.Config, string) typesa.KnSourceClient) {
	fake.createKnSourceClientMutex.Lock()
	defer fake.createKnSourceClientMutex.Unlock()
	fake.CreateKnSourceClientStub = stub
}

func (fake *FakeGHSourceFactory) CreateKnSourceClientArgsForCall(i int) (*rest.Config, string) {
	fake.createKnSourceClientMutex.RLock()
	defer fake.createKnSourceClientMutex.RUnlock()
	argsForCall := fake.createKnSourceClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGHSourceFactory) CreateKnSourceClientReturns(result1 typesa.KnSourceClient) {
	fake.createKnSourceClientMutex.Lock()
	defer fake.createKnSourceClientMutex.Unlock()
	fake.CreateKnSourceClientStub = nil
	fake.createKnSourceClientReturns = struct {
		result1 typesa.KnSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateKnSourceClientReturnsOnCall(i int, result1 typesa.KnSourceClient) {
	fake.createKnSourceClientMutex.Lock()
	defer fake.createKnSourceClientMutex.Unlock()
	fake.CreateKnSourceClientStub = nil
	if fake.createKnSourceClientReturnsOnCall == nil {
		fake.createKnSourceClientReturnsOnCall = make(map[int]struct {
			result1 typesa.KnSourceClient
		})
	}
	fake.createKnSourceClientReturnsOnCall[i] = struct {
		result1 typesa.KnSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateKnSourceParams() *typesa.KnSourceParams {
	fake.createKnSourceParamsMutex.Lock()
	ret, specificReturn := fake.createKnSourceParamsReturnsOnCall[len(fake.createKnSourceParamsArgsForCall)]
	fake.createKnSourceParamsArgsForCall = append(fake.createKnSourceParamsArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateKnSourceParams", []interface{}{})
	fake.createKnSourceParamsMutex.Unlock()
	if fake.CreateKnSourceParamsStub != nil {
		return fake.CreateKnSourceParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createKnSourceParamsReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) CreateKnSourceParamsCallCount() int {
	fake.createKnSourceParamsMutex.RLock()
	defer fake.createKnSourceParamsMutex.RUnlock()
	return len(fake.createKnSourceParamsArgsForCall)
}

func (fake *FakeGHSourceFactory) CreateKnSourceParamsCalls(stub func() *typesa.KnSourceParams) {
	fake.createKnSourceParamsMutex.Lock()
	defer fake.createKnSourceParamsMutex.Unlock()
	fake.CreateKnSourceParamsStub = stub
}

func (fake *FakeGHSourceFactory) CreateKnSourceParamsReturns(result1 *typesa.KnSourceParams) {
	fake.createKnSourceParamsMutex.Lock()
	defer fake.createKnSourceParamsMutex.Unlock()
	fake.CreateKnSourceParamsStub = nil
	fake.createKnSourceParamsReturns = struct {
		result1 *typesa.KnSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) CreateKnSourceParamsReturnsOnCall(i int, result1 *typesa.KnSourceParams) {
	fake.createKnSourceParamsMutex.Lock()
	defer fake.createKnSourceParamsMutex.Unlock()
	fake.CreateKnSourceParamsStub = nil
	if fake.createKnSourceParamsReturnsOnCall == nil {
		fake.createKnSourceParamsReturnsOnCall = make(map[int]struct {
			result1 *typesa.KnSourceParams
		})
	}
	fake.createKnSourceParamsReturnsOnCall[i] = struct {
		result1 *typesa.KnSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) GHSourceClient() types.GHSourceClient {
	fake.gHSourceClientMutex.Lock()
	ret, specificReturn := fake.gHSourceClientReturnsOnCall[len(fake.gHSourceClientArgsForCall)]
	fake.gHSourceClientArgsForCall = append(fake.gHSourceClientArgsForCall, struct {
	}{})
	fake.recordInvocation("GHSourceClient", []interface{}{})
	fake.gHSourceClientMutex.Unlock()
	if fake.GHSourceClientStub != nil {
		return fake.GHSourceClientStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gHSourceClientReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) GHSourceClientCallCount() int {
	fake.gHSourceClientMutex.RLock()
	defer fake.gHSourceClientMutex.RUnlock()
	return len(fake.gHSourceClientArgsForCall)
}

func (fake *FakeGHSourceFactory) GHSourceClientCalls(stub func() types.GHSourceClient) {
	fake.gHSourceClientMutex.Lock()
	defer fake.gHSourceClientMutex.Unlock()
	fake.GHSourceClientStub = stub
}

func (fake *FakeGHSourceFactory) GHSourceClientReturns(result1 types.GHSourceClient) {
	fake.gHSourceClientMutex.Lock()
	defer fake.gHSourceClientMutex.Unlock()
	fake.GHSourceClientStub = nil
	fake.gHSourceClientReturns = struct {
		result1 types.GHSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) GHSourceClientReturnsOnCall(i int, result1 types.GHSourceClient) {
	fake.gHSourceClientMutex.Lock()
	defer fake.gHSourceClientMutex.Unlock()
	fake.GHSourceClientStub = nil
	if fake.gHSourceClientReturnsOnCall == nil {
		fake.gHSourceClientReturnsOnCall = make(map[int]struct {
			result1 types.GHSourceClient
		})
	}
	fake.gHSourceClientReturnsOnCall[i] = struct {
		result1 types.GHSourceClient
	}{result1}
}

func (fake *FakeGHSourceFactory) GHSourceParams() *types.GHSourceParams {
	fake.gHSourceParamsMutex.Lock()
	ret, specificReturn := fake.gHSourceParamsReturnsOnCall[len(fake.gHSourceParamsArgsForCall)]
	fake.gHSourceParamsArgsForCall = append(fake.gHSourceParamsArgsForCall, struct {
	}{})
	fake.recordInvocation("GHSourceParams", []interface{}{})
	fake.gHSourceParamsMutex.Unlock()
	if fake.GHSourceParamsStub != nil {
		return fake.GHSourceParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gHSourceParamsReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) GHSourceParamsCallCount() int {
	fake.gHSourceParamsMutex.RLock()
	defer fake.gHSourceParamsMutex.RUnlock()
	return len(fake.gHSourceParamsArgsForCall)
}

func (fake *FakeGHSourceFactory) GHSourceParamsCalls(stub func() *types.GHSourceParams) {
	fake.gHSourceParamsMutex.Lock()
	defer fake.gHSourceParamsMutex.Unlock()
	fake.GHSourceParamsStub = stub
}

func (fake *FakeGHSourceFactory) GHSourceParamsReturns(result1 *types.GHSourceParams) {
	fake.gHSourceParamsMutex.Lock()
	defer fake.gHSourceParamsMutex.Unlock()
	fake.GHSourceParamsStub = nil
	fake.gHSourceParamsReturns = struct {
		result1 *types.GHSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) GHSourceParamsReturnsOnCall(i int, result1 *types.GHSourceParams) {
	fake.gHSourceParamsMutex.Lock()
	defer fake.gHSourceParamsMutex.Unlock()
	fake.GHSourceParamsStub = nil
	if fake.gHSourceParamsReturnsOnCall == nil {
		fake.gHSourceParamsReturnsOnCall = make(map[int]struct {
			result1 *types.GHSourceParams
		})
	}
	fake.gHSourceParamsReturnsOnCall[i] = struct {
		result1 *types.GHSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) KnSourceParams() *typesa.KnSourceParams {
	fake.knSourceParamsMutex.Lock()
	ret, specificReturn := fake.knSourceParamsReturnsOnCall[len(fake.knSourceParamsArgsForCall)]
	fake.knSourceParamsArgsForCall = append(fake.knSourceParamsArgsForCall, struct {
	}{})
	fake.recordInvocation("KnSourceParams", []interface{}{})
	fake.knSourceParamsMutex.Unlock()
	if fake.KnSourceParamsStub != nil {
		return fake.KnSourceParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.knSourceParamsReturns
	return fakeReturns.result1
}

func (fake *FakeGHSourceFactory) KnSourceParamsCallCount() int {
	fake.knSourceParamsMutex.RLock()
	defer fake.knSourceParamsMutex.RUnlock()
	return len(fake.knSourceParamsArgsForCall)
}

func (fake *FakeGHSourceFactory) KnSourceParamsCalls(stub func() *typesa.KnSourceParams) {
	fake.knSourceParamsMutex.Lock()
	defer fake.knSourceParamsMutex.Unlock()
	fake.KnSourceParamsStub = stub
}

func (fake *FakeGHSourceFactory) KnSourceParamsReturns(result1 *typesa.KnSourceParams) {
	fake.knSourceParamsMutex.Lock()
	defer fake.knSourceParamsMutex.Unlock()
	fake.KnSourceParamsStub = nil
	fake.knSourceParamsReturns = struct {
		result1 *typesa.KnSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) KnSourceParamsReturnsOnCall(i int, result1 *typesa.KnSourceParams) {
	fake.knSourceParamsMutex.Lock()
	defer fake.knSourceParamsMutex.Unlock()
	fake.KnSourceParamsStub = nil
	if fake.knSourceParamsReturnsOnCall == nil {
		fake.knSourceParamsReturnsOnCall = make(map[int]struct {
			result1 *typesa.KnSourceParams
		})
	}
	fake.knSourceParamsReturnsOnCall[i] = struct {
		result1 *typesa.KnSourceParams
	}{result1}
}

func (fake *FakeGHSourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createGHSourceClientMutex.RLock()
	defer fake.createGHSourceClientMutex.RUnlock()
	fake.createGHSourceParamsMutex.RLock()
	defer fake.createGHSourceParamsMutex.RUnlock()
	fake.createKnSourceClientMutex.RLock()
	defer fake.createKnSourceClientMutex.RUnlock()
	fake.createKnSourceParamsMutex.RLock()
	defer fake.createKnSourceParamsMutex.RUnlock()
	fake.gHSourceClientMutex.RLock()
	defer fake.gHSourceClientMutex.RUnlock()
	fake.gHSourceParamsMutex.RLock()
	defer fake.gHSourceParamsMutex.RUnlock()
	fake.knSourceParamsMutex.RLock()
	defer fake.knSourceParamsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGHSourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ types.GHSourceFactory = new(FakeGHSourceFactory)
