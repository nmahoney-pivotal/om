// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type DeleteCertificateAuthorityService struct {
	DeleteCertificateAuthorityStub        func(api.DeleteCertificateAuthorityInput) error
	deleteCertificateAuthorityMutex       sync.RWMutex
	deleteCertificateAuthorityArgsForCall []struct {
		arg1 api.DeleteCertificateAuthorityInput
	}
	deleteCertificateAuthorityReturns struct {
		result1 error
	}
	deleteCertificateAuthorityReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthority(arg1 api.DeleteCertificateAuthorityInput) error {
	fake.deleteCertificateAuthorityMutex.Lock()
	ret, specificReturn := fake.deleteCertificateAuthorityReturnsOnCall[len(fake.deleteCertificateAuthorityArgsForCall)]
	fake.deleteCertificateAuthorityArgsForCall = append(fake.deleteCertificateAuthorityArgsForCall, struct {
		arg1 api.DeleteCertificateAuthorityInput
	}{arg1})
	fake.recordInvocation("DeleteCertificateAuthority", []interface{}{arg1})
	fake.deleteCertificateAuthorityMutex.Unlock()
	if fake.DeleteCertificateAuthorityStub != nil {
		return fake.DeleteCertificateAuthorityStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteCertificateAuthorityReturns
	return fakeReturns.result1
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthorityCallCount() int {
	fake.deleteCertificateAuthorityMutex.RLock()
	defer fake.deleteCertificateAuthorityMutex.RUnlock()
	return len(fake.deleteCertificateAuthorityArgsForCall)
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthorityCalls(stub func(api.DeleteCertificateAuthorityInput) error) {
	fake.deleteCertificateAuthorityMutex.Lock()
	defer fake.deleteCertificateAuthorityMutex.Unlock()
	fake.DeleteCertificateAuthorityStub = stub
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthorityArgsForCall(i int) api.DeleteCertificateAuthorityInput {
	fake.deleteCertificateAuthorityMutex.RLock()
	defer fake.deleteCertificateAuthorityMutex.RUnlock()
	argsForCall := fake.deleteCertificateAuthorityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthorityReturns(result1 error) {
	fake.deleteCertificateAuthorityMutex.Lock()
	defer fake.deleteCertificateAuthorityMutex.Unlock()
	fake.DeleteCertificateAuthorityStub = nil
	fake.deleteCertificateAuthorityReturns = struct {
		result1 error
	}{result1}
}

func (fake *DeleteCertificateAuthorityService) DeleteCertificateAuthorityReturnsOnCall(i int, result1 error) {
	fake.deleteCertificateAuthorityMutex.Lock()
	defer fake.deleteCertificateAuthorityMutex.Unlock()
	fake.DeleteCertificateAuthorityStub = nil
	if fake.deleteCertificateAuthorityReturnsOnCall == nil {
		fake.deleteCertificateAuthorityReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteCertificateAuthorityReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DeleteCertificateAuthorityService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteCertificateAuthorityMutex.RLock()
	defer fake.deleteCertificateAuthorityMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeleteCertificateAuthorityService) recordInvocation(key string, args []interface{}) {
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
