// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/builder"
)

type MetadataBuilder struct {
	BuildStub        func(releaseTarballs, runtimeConfigDirectories, variableDirectories []string, pathToStemcell, pathToMetadata, version, pathToTile, iconPath string) (builder.GeneratedMetadata, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		releaseTarballs          []string
		runtimeConfigDirectories []string
		variableDirectories      []string
		pathToStemcell           string
		pathToMetadata           string
		version                  string
		pathToTile               string
		iconPath                 string
	}
	buildReturns struct {
		result1 builder.GeneratedMetadata
		result2 error
	}
	buildReturnsOnCall map[int]struct {
		result1 builder.GeneratedMetadata
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetadataBuilder) Build(releaseTarballs []string, runtimeConfigDirectories []string, variableDirectories []string, pathToStemcell string, pathToMetadata string, version string, pathToTile string, iconPath string) (builder.GeneratedMetadata, error) {
	var releaseTarballsCopy []string
	if releaseTarballs != nil {
		releaseTarballsCopy = make([]string, len(releaseTarballs))
		copy(releaseTarballsCopy, releaseTarballs)
	}
	var runtimeConfigDirectoriesCopy []string
	if runtimeConfigDirectories != nil {
		runtimeConfigDirectoriesCopy = make([]string, len(runtimeConfigDirectories))
		copy(runtimeConfigDirectoriesCopy, runtimeConfigDirectories)
	}
	var variableDirectoriesCopy []string
	if variableDirectories != nil {
		variableDirectoriesCopy = make([]string, len(variableDirectories))
		copy(variableDirectoriesCopy, variableDirectories)
	}
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		releaseTarballs          []string
		runtimeConfigDirectories []string
		variableDirectories      []string
		pathToStemcell           string
		pathToMetadata           string
		version                  string
		pathToTile               string
		iconPath                 string
	}{releaseTarballsCopy, runtimeConfigDirectoriesCopy, variableDirectoriesCopy, pathToStemcell, pathToMetadata, version, pathToTile, iconPath})
	fake.recordInvocation("Build", []interface{}{releaseTarballsCopy, runtimeConfigDirectoriesCopy, variableDirectoriesCopy, pathToStemcell, pathToMetadata, version, pathToTile, iconPath})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(releaseTarballs, runtimeConfigDirectories, variableDirectories, pathToStemcell, pathToMetadata, version, pathToTile, iconPath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.buildReturns.result1, fake.buildReturns.result2
}

func (fake *MetadataBuilder) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *MetadataBuilder) BuildArgsForCall(i int) ([]string, []string, []string, string, string, string, string, string) {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].releaseTarballs, fake.buildArgsForCall[i].runtimeConfigDirectories, fake.buildArgsForCall[i].variableDirectories, fake.buildArgsForCall[i].pathToStemcell, fake.buildArgsForCall[i].pathToMetadata, fake.buildArgsForCall[i].version, fake.buildArgsForCall[i].pathToTile, fake.buildArgsForCall[i].iconPath
}

func (fake *MetadataBuilder) BuildReturns(result1 builder.GeneratedMetadata, result2 error) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 builder.GeneratedMetadata
		result2 error
	}{result1, result2}
}

func (fake *MetadataBuilder) BuildReturnsOnCall(i int, result1 builder.GeneratedMetadata, result2 error) {
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 builder.GeneratedMetadata
			result2 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 builder.GeneratedMetadata
		result2 error
	}{result1, result2}
}

func (fake *MetadataBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetadataBuilder) recordInvocation(key string, args []interface{}) {
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
