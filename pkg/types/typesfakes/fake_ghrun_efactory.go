// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/maximilien/kn-source-github/pkg/types"
	typesa "github.com/maximilien/kn-source-pkg/pkg/types"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
)

type FakeGHRunEFactory struct {
	CreateRunEStub        func() func(cmd *cobra.Command, args []string) error
	createRunEMutex       sync.RWMutex
	createRunEArgsForCall []struct {
	}
	createRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	createRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	DeleteRunEStub        func() func(cmd *cobra.Command, args []string) error
	deleteRunEMutex       sync.RWMutex
	deleteRunEArgsForCall []struct {
	}
	deleteRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	deleteRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	DescribeRunEStub        func() func(cmd *cobra.Command, args []string) error
	describeRunEMutex       sync.RWMutex
	describeRunEArgsForCall []struct {
	}
	describeRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	describeRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	GHSourceClientStub        func(*rest.Config, string) types.GHSourceClient
	gHSourceClientMutex       sync.RWMutex
	gHSourceClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 string
	}
	gHSourceClientReturns struct {
		result1 types.GHSourceClient
	}
	gHSourceClientReturnsOnCall map[int]struct {
		result1 types.GHSourceClient
	}
	GHSourceFactoryStub        func() types.GHSourceFactory
	gHSourceFactoryMutex       sync.RWMutex
	gHSourceFactoryArgsForCall []struct {
	}
	gHSourceFactoryReturns struct {
		result1 types.GHSourceFactory
	}
	gHSourceFactoryReturnsOnCall map[int]struct {
		result1 types.GHSourceFactory
	}
	KnSourceClientStub        func(*rest.Config, string) typesa.KnSourceClient
	knSourceClientMutex       sync.RWMutex
	knSourceClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 string
	}
	knSourceClientReturns struct {
		result1 typesa.KnSourceClient
	}
	knSourceClientReturnsOnCall map[int]struct {
		result1 typesa.KnSourceClient
	}
	KnSourceFactoryStub        func() typesa.KnSourceFactory
	knSourceFactoryMutex       sync.RWMutex
	knSourceFactoryArgsForCall []struct {
	}
	knSourceFactoryReturns struct {
		result1 typesa.KnSourceFactory
	}
	knSourceFactoryReturnsOnCall map[int]struct {
		result1 typesa.KnSourceFactory
	}
	UpdateRunEStub        func() func(cmd *cobra.Command, args []string) error
	updateRunEMutex       sync.RWMutex
	updateRunEArgsForCall []struct {
	}
	updateRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	updateRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGHRunEFactory) CreateRunE() func(cmd *cobra.Command, args []string) error {
	fake.createRunEMutex.Lock()
	ret, specificReturn := fake.createRunEReturnsOnCall[len(fake.createRunEArgsForCall)]
	fake.createRunEArgsForCall = append(fake.createRunEArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateRunE", []interface{}{})
	fake.createRunEMutex.Unlock()
	if fake.CreateRunEStub != nil {
		return fake.CreateRunEStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createRunEReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) CreateRunECallCount() int {
	fake.createRunEMutex.RLock()
	defer fake.createRunEMutex.RUnlock()
	return len(fake.createRunEArgsForCall)
}

func (fake *FakeGHRunEFactory) CreateRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = stub
}

func (fake *FakeGHRunEFactory) CreateRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = nil
	fake.createRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) CreateRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = nil
	if fake.createRunEReturnsOnCall == nil {
		fake.createRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.createRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) DeleteRunE() func(cmd *cobra.Command, args []string) error {
	fake.deleteRunEMutex.Lock()
	ret, specificReturn := fake.deleteRunEReturnsOnCall[len(fake.deleteRunEArgsForCall)]
	fake.deleteRunEArgsForCall = append(fake.deleteRunEArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteRunE", []interface{}{})
	fake.deleteRunEMutex.Unlock()
	if fake.DeleteRunEStub != nil {
		return fake.DeleteRunEStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteRunEReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) DeleteRunECallCount() int {
	fake.deleteRunEMutex.RLock()
	defer fake.deleteRunEMutex.RUnlock()
	return len(fake.deleteRunEArgsForCall)
}

func (fake *FakeGHRunEFactory) DeleteRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = stub
}

func (fake *FakeGHRunEFactory) DeleteRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = nil
	fake.deleteRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) DeleteRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = nil
	if fake.deleteRunEReturnsOnCall == nil {
		fake.deleteRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.deleteRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) DescribeRunE() func(cmd *cobra.Command, args []string) error {
	fake.describeRunEMutex.Lock()
	ret, specificReturn := fake.describeRunEReturnsOnCall[len(fake.describeRunEArgsForCall)]
	fake.describeRunEArgsForCall = append(fake.describeRunEArgsForCall, struct {
	}{})
	fake.recordInvocation("DescribeRunE", []interface{}{})
	fake.describeRunEMutex.Unlock()
	if fake.DescribeRunEStub != nil {
		return fake.DescribeRunEStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.describeRunEReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) DescribeRunECallCount() int {
	fake.describeRunEMutex.RLock()
	defer fake.describeRunEMutex.RUnlock()
	return len(fake.describeRunEArgsForCall)
}

func (fake *FakeGHRunEFactory) DescribeRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = stub
}

func (fake *FakeGHRunEFactory) DescribeRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = nil
	fake.describeRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) DescribeRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = nil
	if fake.describeRunEReturnsOnCall == nil {
		fake.describeRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.describeRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) GHSourceClient(arg1 *rest.Config, arg2 string) types.GHSourceClient {
	fake.gHSourceClientMutex.Lock()
	ret, specificReturn := fake.gHSourceClientReturnsOnCall[len(fake.gHSourceClientArgsForCall)]
	fake.gHSourceClientArgsForCall = append(fake.gHSourceClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GHSourceClient", []interface{}{arg1, arg2})
	fake.gHSourceClientMutex.Unlock()
	if fake.GHSourceClientStub != nil {
		return fake.GHSourceClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gHSourceClientReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) GHSourceClientCallCount() int {
	fake.gHSourceClientMutex.RLock()
	defer fake.gHSourceClientMutex.RUnlock()
	return len(fake.gHSourceClientArgsForCall)
}

func (fake *FakeGHRunEFactory) GHSourceClientCalls(stub func(*rest.Config, string) types.GHSourceClient) {
	fake.gHSourceClientMutex.Lock()
	defer fake.gHSourceClientMutex.Unlock()
	fake.GHSourceClientStub = stub
}

func (fake *FakeGHRunEFactory) GHSourceClientArgsForCall(i int) (*rest.Config, string) {
	fake.gHSourceClientMutex.RLock()
	defer fake.gHSourceClientMutex.RUnlock()
	argsForCall := fake.gHSourceClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGHRunEFactory) GHSourceClientReturns(result1 types.GHSourceClient) {
	fake.gHSourceClientMutex.Lock()
	defer fake.gHSourceClientMutex.Unlock()
	fake.GHSourceClientStub = nil
	fake.gHSourceClientReturns = struct {
		result1 types.GHSourceClient
	}{result1}
}

func (fake *FakeGHRunEFactory) GHSourceClientReturnsOnCall(i int, result1 types.GHSourceClient) {
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

func (fake *FakeGHRunEFactory) GHSourceFactory() types.GHSourceFactory {
	fake.gHSourceFactoryMutex.Lock()
	ret, specificReturn := fake.gHSourceFactoryReturnsOnCall[len(fake.gHSourceFactoryArgsForCall)]
	fake.gHSourceFactoryArgsForCall = append(fake.gHSourceFactoryArgsForCall, struct {
	}{})
	fake.recordInvocation("GHSourceFactory", []interface{}{})
	fake.gHSourceFactoryMutex.Unlock()
	if fake.GHSourceFactoryStub != nil {
		return fake.GHSourceFactoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.gHSourceFactoryReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) GHSourceFactoryCallCount() int {
	fake.gHSourceFactoryMutex.RLock()
	defer fake.gHSourceFactoryMutex.RUnlock()
	return len(fake.gHSourceFactoryArgsForCall)
}

func (fake *FakeGHRunEFactory) GHSourceFactoryCalls(stub func() types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = stub
}

func (fake *FakeGHRunEFactory) GHSourceFactoryReturns(result1 types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = nil
	fake.gHSourceFactoryReturns = struct {
		result1 types.GHSourceFactory
	}{result1}
}

func (fake *FakeGHRunEFactory) GHSourceFactoryReturnsOnCall(i int, result1 types.GHSourceFactory) {
	fake.gHSourceFactoryMutex.Lock()
	defer fake.gHSourceFactoryMutex.Unlock()
	fake.GHSourceFactoryStub = nil
	if fake.gHSourceFactoryReturnsOnCall == nil {
		fake.gHSourceFactoryReturnsOnCall = make(map[int]struct {
			result1 types.GHSourceFactory
		})
	}
	fake.gHSourceFactoryReturnsOnCall[i] = struct {
		result1 types.GHSourceFactory
	}{result1}
}

func (fake *FakeGHRunEFactory) KnSourceClient(arg1 *rest.Config, arg2 string) typesa.KnSourceClient {
	fake.knSourceClientMutex.Lock()
	ret, specificReturn := fake.knSourceClientReturnsOnCall[len(fake.knSourceClientArgsForCall)]
	fake.knSourceClientArgsForCall = append(fake.knSourceClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("KnSourceClient", []interface{}{arg1, arg2})
	fake.knSourceClientMutex.Unlock()
	if fake.KnSourceClientStub != nil {
		return fake.KnSourceClientStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.knSourceClientReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) KnSourceClientCallCount() int {
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	return len(fake.knSourceClientArgsForCall)
}

func (fake *FakeGHRunEFactory) KnSourceClientCalls(stub func(*rest.Config, string) typesa.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = stub
}

func (fake *FakeGHRunEFactory) KnSourceClientArgsForCall(i int) (*rest.Config, string) {
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	argsForCall := fake.knSourceClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGHRunEFactory) KnSourceClientReturns(result1 typesa.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = nil
	fake.knSourceClientReturns = struct {
		result1 typesa.KnSourceClient
	}{result1}
}

func (fake *FakeGHRunEFactory) KnSourceClientReturnsOnCall(i int, result1 typesa.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = nil
	if fake.knSourceClientReturnsOnCall == nil {
		fake.knSourceClientReturnsOnCall = make(map[int]struct {
			result1 typesa.KnSourceClient
		})
	}
	fake.knSourceClientReturnsOnCall[i] = struct {
		result1 typesa.KnSourceClient
	}{result1}
}

func (fake *FakeGHRunEFactory) KnSourceFactory() typesa.KnSourceFactory {
	fake.knSourceFactoryMutex.Lock()
	ret, specificReturn := fake.knSourceFactoryReturnsOnCall[len(fake.knSourceFactoryArgsForCall)]
	fake.knSourceFactoryArgsForCall = append(fake.knSourceFactoryArgsForCall, struct {
	}{})
	fake.recordInvocation("KnSourceFactory", []interface{}{})
	fake.knSourceFactoryMutex.Unlock()
	if fake.KnSourceFactoryStub != nil {
		return fake.KnSourceFactoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.knSourceFactoryReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) KnSourceFactoryCallCount() int {
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	return len(fake.knSourceFactoryArgsForCall)
}

func (fake *FakeGHRunEFactory) KnSourceFactoryCalls(stub func() typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = stub
}

func (fake *FakeGHRunEFactory) KnSourceFactoryReturns(result1 typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	fake.knSourceFactoryReturns = struct {
		result1 typesa.KnSourceFactory
	}{result1}
}

func (fake *FakeGHRunEFactory) KnSourceFactoryReturnsOnCall(i int, result1 typesa.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	if fake.knSourceFactoryReturnsOnCall == nil {
		fake.knSourceFactoryReturnsOnCall = make(map[int]struct {
			result1 typesa.KnSourceFactory
		})
	}
	fake.knSourceFactoryReturnsOnCall[i] = struct {
		result1 typesa.KnSourceFactory
	}{result1}
}

func (fake *FakeGHRunEFactory) UpdateRunE() func(cmd *cobra.Command, args []string) error {
	fake.updateRunEMutex.Lock()
	ret, specificReturn := fake.updateRunEReturnsOnCall[len(fake.updateRunEArgsForCall)]
	fake.updateRunEArgsForCall = append(fake.updateRunEArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateRunE", []interface{}{})
	fake.updateRunEMutex.Unlock()
	if fake.UpdateRunEStub != nil {
		return fake.UpdateRunEStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateRunEReturns
	return fakeReturns.result1
}

func (fake *FakeGHRunEFactory) UpdateRunECallCount() int {
	fake.updateRunEMutex.RLock()
	defer fake.updateRunEMutex.RUnlock()
	return len(fake.updateRunEArgsForCall)
}

func (fake *FakeGHRunEFactory) UpdateRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = stub
}

func (fake *FakeGHRunEFactory) UpdateRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = nil
	fake.updateRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) UpdateRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = nil
	if fake.updateRunEReturnsOnCall == nil {
		fake.updateRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.updateRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeGHRunEFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRunEMutex.RLock()
	defer fake.createRunEMutex.RUnlock()
	fake.deleteRunEMutex.RLock()
	defer fake.deleteRunEMutex.RUnlock()
	fake.describeRunEMutex.RLock()
	defer fake.describeRunEMutex.RUnlock()
	fake.gHSourceClientMutex.RLock()
	defer fake.gHSourceClientMutex.RUnlock()
	fake.gHSourceFactoryMutex.RLock()
	defer fake.gHSourceFactoryMutex.RUnlock()
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	fake.updateRunEMutex.RLock()
	defer fake.updateRunEMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGHRunEFactory) recordInvocation(key string, args []interface{}) {
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

var _ types.GHRunEFactory = new(FakeGHRunEFactory)
