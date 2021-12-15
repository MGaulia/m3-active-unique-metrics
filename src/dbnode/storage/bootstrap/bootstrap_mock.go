// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/bootstrap/types.go

// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package bootstrap is a generated GoMock package.
package bootstrap

import (
	"reflect"

	"github.com/m3db/m3/src/dbnode/namespace"
	"github.com/m3db/m3/src/dbnode/persist/fs"
	"github.com/m3db/m3/src/dbnode/storage/block"
	"github.com/m3db/m3/src/dbnode/storage/bootstrap/result"
	"github.com/m3db/m3/src/dbnode/storage/series"
	"github.com/m3db/m3/src/dbnode/topology"
	"github.com/m3db/m3/src/x/context"
	"github.com/m3db/m3/src/x/ident"
	"github.com/m3db/m3/src/x/instrument"
	"github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockProcessProvider is a mock of ProcessProvider interface.
type MockProcessProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProcessProviderMockRecorder
}

// MockProcessProviderMockRecorder is the mock recorder for MockProcessProvider.
type MockProcessProviderMockRecorder struct {
	mock *MockProcessProvider
}

// NewMockProcessProvider creates a new mock instance.
func NewMockProcessProvider(ctrl *gomock.Controller) *MockProcessProvider {
	mock := &MockProcessProvider{ctrl: ctrl}
	mock.recorder = &MockProcessProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessProvider) EXPECT() *MockProcessProviderMockRecorder {
	return m.recorder
}

// BootstrapperProvider mocks base method.
func (m *MockProcessProvider) BootstrapperProvider() BootstrapperProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BootstrapperProvider")
	ret0, _ := ret[0].(BootstrapperProvider)
	return ret0
}

// BootstrapperProvider indicates an expected call of BootstrapperProvider.
func (mr *MockProcessProviderMockRecorder) BootstrapperProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BootstrapperProvider", reflect.TypeOf((*MockProcessProvider)(nil).BootstrapperProvider))
}

// Provide mocks base method.
func (m *MockProcessProvider) Provide() (Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide.
func (mr *MockProcessProviderMockRecorder) Provide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockProcessProvider)(nil).Provide))
}

// SetBootstrapperProvider mocks base method.
func (m *MockProcessProvider) SetBootstrapperProvider(bootstrapper BootstrapperProvider) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBootstrapperProvider", bootstrapper)
}

// SetBootstrapperProvider indicates an expected call of SetBootstrapperProvider.
func (mr *MockProcessProviderMockRecorder) SetBootstrapperProvider(bootstrapper interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrapperProvider", reflect.TypeOf((*MockProcessProvider)(nil).SetBootstrapperProvider), bootstrapper)
}

// MockProcess is a mock of Process interface.
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess.
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance.
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockProcess) Run(ctx context.Context, start time.UnixNano, namespaces []ProcessNamespace) (NamespaceResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx, start, namespaces)
	ret0, _ := ret[0].(NamespaceResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run.
func (mr *MockProcessMockRecorder) Run(ctx, start, namespaces interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockProcess)(nil).Run), ctx, start, namespaces)
}

// MockHook is a mock of Hook interface.
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook.
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance.
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockHook) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockHookMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHook)(nil).Run))
}

// MockNamespaceDataAccumulator is a mock of NamespaceDataAccumulator interface.
type MockNamespaceDataAccumulator struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceDataAccumulatorMockRecorder
}

// MockNamespaceDataAccumulatorMockRecorder is the mock recorder for MockNamespaceDataAccumulator.
type MockNamespaceDataAccumulatorMockRecorder struct {
	mock *MockNamespaceDataAccumulator
}

// NewMockNamespaceDataAccumulator creates a new mock instance.
func NewMockNamespaceDataAccumulator(ctrl *gomock.Controller) *MockNamespaceDataAccumulator {
	mock := &MockNamespaceDataAccumulator{ctrl: ctrl}
	mock.recorder = &MockNamespaceDataAccumulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceDataAccumulator) EXPECT() *MockNamespaceDataAccumulatorMockRecorder {
	return m.recorder
}

// CheckoutSeriesWithLock mocks base method.
func (m *MockNamespaceDataAccumulator) CheckoutSeriesWithLock(shardID uint32, id ident.ID, tags ident.TagIterator) (CheckoutSeriesResult, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckoutSeriesWithLock", shardID, id, tags)
	ret0, _ := ret[0].(CheckoutSeriesResult)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckoutSeriesWithLock indicates an expected call of CheckoutSeriesWithLock.
func (mr *MockNamespaceDataAccumulatorMockRecorder) CheckoutSeriesWithLock(shardID, id, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckoutSeriesWithLock", reflect.TypeOf((*MockNamespaceDataAccumulator)(nil).CheckoutSeriesWithLock), shardID, id, tags)
}

// CheckoutSeriesWithoutLock mocks base method.
func (m *MockNamespaceDataAccumulator) CheckoutSeriesWithoutLock(shardID uint32, id ident.ID, tags ident.TagIterator) (CheckoutSeriesResult, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckoutSeriesWithoutLock", shardID, id, tags)
	ret0, _ := ret[0].(CheckoutSeriesResult)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckoutSeriesWithoutLock indicates an expected call of CheckoutSeriesWithoutLock.
func (mr *MockNamespaceDataAccumulatorMockRecorder) CheckoutSeriesWithoutLock(shardID, id, tags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckoutSeriesWithoutLock", reflect.TypeOf((*MockNamespaceDataAccumulator)(nil).CheckoutSeriesWithoutLock), shardID, id, tags)
}

// Close mocks base method.
func (m *MockNamespaceDataAccumulator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockNamespaceDataAccumulatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockNamespaceDataAccumulator)(nil).Close))
}

// MockProcessOptions is a mock of ProcessOptions interface.
type MockProcessOptions struct {
	ctrl     *gomock.Controller
	recorder *MockProcessOptionsMockRecorder
}

// MockProcessOptionsMockRecorder is the mock recorder for MockProcessOptions.
type MockProcessOptionsMockRecorder struct {
	mock *MockProcessOptions
}

// NewMockProcessOptions creates a new mock instance.
func NewMockProcessOptions(ctrl *gomock.Controller) *MockProcessOptions {
	mock := &MockProcessOptions{ctrl: ctrl}
	mock.recorder = &MockProcessOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcessOptions) EXPECT() *MockProcessOptionsMockRecorder {
	return m.recorder
}

// CacheSeriesMetadata mocks base method.
func (m *MockProcessOptions) CacheSeriesMetadata() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheSeriesMetadata")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CacheSeriesMetadata indicates an expected call of CacheSeriesMetadata.
func (mr *MockProcessOptionsMockRecorder) CacheSeriesMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheSeriesMetadata", reflect.TypeOf((*MockProcessOptions)(nil).CacheSeriesMetadata))
}

// Origin mocks base method.
func (m *MockProcessOptions) Origin() topology.Host {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Origin")
	ret0, _ := ret[0].(topology.Host)
	return ret0
}

// Origin indicates an expected call of Origin.
func (mr *MockProcessOptionsMockRecorder) Origin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MockProcessOptions)(nil).Origin))
}

// SetCacheSeriesMetadata mocks base method.
func (m *MockProcessOptions) SetCacheSeriesMetadata(value bool) ProcessOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCacheSeriesMetadata", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetCacheSeriesMetadata indicates an expected call of SetCacheSeriesMetadata.
func (mr *MockProcessOptionsMockRecorder) SetCacheSeriesMetadata(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheSeriesMetadata", reflect.TypeOf((*MockProcessOptions)(nil).SetCacheSeriesMetadata), value)
}

// SetOrigin mocks base method.
func (m *MockProcessOptions) SetOrigin(value topology.Host) ProcessOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrigin", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetOrigin indicates an expected call of SetOrigin.
func (mr *MockProcessOptionsMockRecorder) SetOrigin(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrigin", reflect.TypeOf((*MockProcessOptions)(nil).SetOrigin), value)
}

// SetTopologyMapProvider mocks base method.
func (m *MockProcessOptions) SetTopologyMapProvider(value topology.MapProvider) ProcessOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTopologyMapProvider", value)
	ret0, _ := ret[0].(ProcessOptions)
	return ret0
}

// SetTopologyMapProvider indicates an expected call of SetTopologyMapProvider.
func (mr *MockProcessOptionsMockRecorder) SetTopologyMapProvider(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTopologyMapProvider", reflect.TypeOf((*MockProcessOptions)(nil).SetTopologyMapProvider), value)
}

// TopologyMapProvider mocks base method.
func (m *MockProcessOptions) TopologyMapProvider() topology.MapProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopologyMapProvider")
	ret0, _ := ret[0].(topology.MapProvider)
	return ret0
}

// TopologyMapProvider indicates an expected call of TopologyMapProvider.
func (mr *MockProcessOptionsMockRecorder) TopologyMapProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopologyMapProvider", reflect.TypeOf((*MockProcessOptions)(nil).TopologyMapProvider))
}

// Validate mocks base method.
func (m *MockProcessOptions) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockProcessOptionsMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockProcessOptions)(nil).Validate))
}

// MockRunOptions is a mock of RunOptions interface.
type MockRunOptions struct {
	ctrl     *gomock.Controller
	recorder *MockRunOptionsMockRecorder
}

// MockRunOptionsMockRecorder is the mock recorder for MockRunOptions.
type MockRunOptionsMockRecorder struct {
	mock *MockRunOptions
}

// NewMockRunOptions creates a new mock instance.
func NewMockRunOptions(ctrl *gomock.Controller) *MockRunOptions {
	mock := &MockRunOptions{ctrl: ctrl}
	mock.recorder = &MockRunOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunOptions) EXPECT() *MockRunOptionsMockRecorder {
	return m.recorder
}

// CacheSeriesMetadata mocks base method.
func (m *MockRunOptions) CacheSeriesMetadata() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CacheSeriesMetadata")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CacheSeriesMetadata indicates an expected call of CacheSeriesMetadata.
func (mr *MockRunOptionsMockRecorder) CacheSeriesMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheSeriesMetadata", reflect.TypeOf((*MockRunOptions)(nil).CacheSeriesMetadata))
}

// InitialTopologyState mocks base method.
func (m *MockRunOptions) InitialTopologyState() *topology.StateSnapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialTopologyState")
	ret0, _ := ret[0].(*topology.StateSnapshot)
	return ret0
}

// InitialTopologyState indicates an expected call of InitialTopologyState.
func (mr *MockRunOptionsMockRecorder) InitialTopologyState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialTopologyState", reflect.TypeOf((*MockRunOptions)(nil).InitialTopologyState))
}

// PersistConfig mocks base method.
func (m *MockRunOptions) PersistConfig() PersistConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersistConfig")
	ret0, _ := ret[0].(PersistConfig)
	return ret0
}

// PersistConfig indicates an expected call of PersistConfig.
func (mr *MockRunOptionsMockRecorder) PersistConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersistConfig", reflect.TypeOf((*MockRunOptions)(nil).PersistConfig))
}

// SetCacheSeriesMetadata mocks base method.
func (m *MockRunOptions) SetCacheSeriesMetadata(value bool) RunOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCacheSeriesMetadata", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetCacheSeriesMetadata indicates an expected call of SetCacheSeriesMetadata.
func (mr *MockRunOptionsMockRecorder) SetCacheSeriesMetadata(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCacheSeriesMetadata", reflect.TypeOf((*MockRunOptions)(nil).SetCacheSeriesMetadata), value)
}

// SetInitialTopologyState mocks base method.
func (m *MockRunOptions) SetInitialTopologyState(value *topology.StateSnapshot) RunOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInitialTopologyState", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetInitialTopologyState indicates an expected call of SetInitialTopologyState.
func (mr *MockRunOptionsMockRecorder) SetInitialTopologyState(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitialTopologyState", reflect.TypeOf((*MockRunOptions)(nil).SetInitialTopologyState), value)
}

// SetPersistConfig mocks base method.
func (m *MockRunOptions) SetPersistConfig(value PersistConfig) RunOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPersistConfig", value)
	ret0, _ := ret[0].(RunOptions)
	return ret0
}

// SetPersistConfig indicates an expected call of SetPersistConfig.
func (mr *MockRunOptionsMockRecorder) SetPersistConfig(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPersistConfig", reflect.TypeOf((*MockRunOptions)(nil).SetPersistConfig), value)
}

// MockBootstrapperProvider is a mock of BootstrapperProvider interface.
type MockBootstrapperProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperProviderMockRecorder
}

// MockBootstrapperProviderMockRecorder is the mock recorder for MockBootstrapperProvider.
type MockBootstrapperProviderMockRecorder struct {
	mock *MockBootstrapperProvider
}

// NewMockBootstrapperProvider creates a new mock instance.
func NewMockBootstrapperProvider(ctrl *gomock.Controller) *MockBootstrapperProvider {
	mock := &MockBootstrapperProvider{ctrl: ctrl}
	mock.recorder = &MockBootstrapperProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBootstrapperProvider) EXPECT() *MockBootstrapperProviderMockRecorder {
	return m.recorder
}

// Provide mocks base method.
func (m *MockBootstrapperProvider) Provide() (Bootstrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide")
	ret0, _ := ret[0].(Bootstrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provide indicates an expected call of Provide.
func (mr *MockBootstrapperProviderMockRecorder) Provide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockBootstrapperProvider)(nil).Provide))
}

// String mocks base method.
func (m *MockBootstrapperProvider) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockBootstrapperProviderMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBootstrapperProvider)(nil).String))
}

// MockBootstrapper is a mock of Bootstrapper interface.
type MockBootstrapper struct {
	ctrl     *gomock.Controller
	recorder *MockBootstrapperMockRecorder
}

// MockBootstrapperMockRecorder is the mock recorder for MockBootstrapper.
type MockBootstrapperMockRecorder struct {
	mock *MockBootstrapper
}

// NewMockBootstrapper creates a new mock instance.
func NewMockBootstrapper(ctrl *gomock.Controller) *MockBootstrapper {
	mock := &MockBootstrapper{ctrl: ctrl}
	mock.recorder = &MockBootstrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBootstrapper) EXPECT() *MockBootstrapperMockRecorder {
	return m.recorder
}

// Bootstrap mocks base method.
func (m *MockBootstrapper) Bootstrap(ctx context.Context, namespaces Namespaces, cache Cache) (NamespaceResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", ctx, namespaces, cache)
	ret0, _ := ret[0].(NamespaceResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap.
func (mr *MockBootstrapperMockRecorder) Bootstrap(ctx, namespaces, cache interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockBootstrapper)(nil).Bootstrap), ctx, namespaces, cache)
}

// String mocks base method.
func (m *MockBootstrapper) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockBootstrapperMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBootstrapper)(nil).String))
}

// MockSource is a mock of Source interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource.
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// AvailableData mocks base method.
func (m *MockSource) AvailableData(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, cache Cache, runOpts RunOptions) (result.ShardTimeRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableData", ns, shardsTimeRanges, cache, runOpts)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailableData indicates an expected call of AvailableData.
func (mr *MockSourceMockRecorder) AvailableData(ns, shardsTimeRanges, cache, runOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableData", reflect.TypeOf((*MockSource)(nil).AvailableData), ns, shardsTimeRanges, cache, runOpts)
}

// AvailableIndex mocks base method.
func (m *MockSource) AvailableIndex(ns namespace.Metadata, shardsTimeRanges result.ShardTimeRanges, cache Cache, opts RunOptions) (result.ShardTimeRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableIndex", ns, shardsTimeRanges, cache, opts)
	ret0, _ := ret[0].(result.ShardTimeRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AvailableIndex indicates an expected call of AvailableIndex.
func (mr *MockSourceMockRecorder) AvailableIndex(ns, shardsTimeRanges, cache, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableIndex", reflect.TypeOf((*MockSource)(nil).AvailableIndex), ns, shardsTimeRanges, cache, opts)
}

// Read mocks base method.
func (m *MockSource) Read(ctx context.Context, namespaces Namespaces, cache Cache) (NamespaceResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, namespaces, cache)
	ret0, _ := ret[0].(NamespaceResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSourceMockRecorder) Read(ctx, namespaces, cache interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSource)(nil).Read), ctx, namespaces, cache)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Evict mocks base method.
func (m *MockCache) Evict() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Evict")
}

// Evict indicates an expected call of Evict.
func (mr *MockCacheMockRecorder) Evict() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evict", reflect.TypeOf((*MockCache)(nil).Evict))
}

// InfoFilesForNamespace mocks base method.
func (m *MockCache) InfoFilesForNamespace(ns namespace.Metadata) (InfoFileResultsPerShard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfoFilesForNamespace", ns)
	ret0, _ := ret[0].(InfoFileResultsPerShard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfoFilesForNamespace indicates an expected call of InfoFilesForNamespace.
func (mr *MockCacheMockRecorder) InfoFilesForNamespace(ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfoFilesForNamespace", reflect.TypeOf((*MockCache)(nil).InfoFilesForNamespace), ns)
}

// InfoFilesForShard mocks base method.
func (m *MockCache) InfoFilesForShard(ns namespace.Metadata, shard uint32) ([]fs.ReadInfoFileResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InfoFilesForShard", ns, shard)
	ret0, _ := ret[0].([]fs.ReadInfoFileResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InfoFilesForShard indicates an expected call of InfoFilesForShard.
func (mr *MockCacheMockRecorder) InfoFilesForShard(ns, shard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InfoFilesForShard", reflect.TypeOf((*MockCache)(nil).InfoFilesForShard), ns, shard)
}

// ReadInfoFiles mocks base method.
func (m *MockCache) ReadInfoFiles() InfoFilesByNamespace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadInfoFiles")
	ret0, _ := ret[0].(InfoFilesByNamespace)
	return ret0
}

// ReadInfoFiles indicates an expected call of ReadInfoFiles.
func (mr *MockCacheMockRecorder) ReadInfoFiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadInfoFiles", reflect.TypeOf((*MockCache)(nil).ReadInfoFiles))
}

// MockCacheOptions is a mock of CacheOptions interface.
type MockCacheOptions struct {
	ctrl     *gomock.Controller
	recorder *MockCacheOptionsMockRecorder
}

// MockCacheOptionsMockRecorder is the mock recorder for MockCacheOptions.
type MockCacheOptionsMockRecorder struct {
	mock *MockCacheOptions
}

// NewMockCacheOptions creates a new mock instance.
func NewMockCacheOptions(ctrl *gomock.Controller) *MockCacheOptions {
	mock := &MockCacheOptions{ctrl: ctrl}
	mock.recorder = &MockCacheOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheOptions) EXPECT() *MockCacheOptionsMockRecorder {
	return m.recorder
}

// FilesystemOptions mocks base method.
func (m *MockCacheOptions) FilesystemOptions() fs.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilesystemOptions")
	ret0, _ := ret[0].(fs.Options)
	return ret0
}

// FilesystemOptions indicates an expected call of FilesystemOptions.
func (mr *MockCacheOptionsMockRecorder) FilesystemOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilesystemOptions", reflect.TypeOf((*MockCacheOptions)(nil).FilesystemOptions))
}

// InstrumentOptions mocks base method.
func (m *MockCacheOptions) InstrumentOptions() instrument.Options {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstrumentOptions")
	ret0, _ := ret[0].(instrument.Options)
	return ret0
}

// InstrumentOptions indicates an expected call of InstrumentOptions.
func (mr *MockCacheOptionsMockRecorder) InstrumentOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstrumentOptions", reflect.TypeOf((*MockCacheOptions)(nil).InstrumentOptions))
}

// NamespaceDetails mocks base method.
func (m *MockCacheOptions) NamespaceDetails() []NamespaceDetails {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamespaceDetails")
	ret0, _ := ret[0].([]NamespaceDetails)
	return ret0
}

// NamespaceDetails indicates an expected call of NamespaceDetails.
func (mr *MockCacheOptionsMockRecorder) NamespaceDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamespaceDetails", reflect.TypeOf((*MockCacheOptions)(nil).NamespaceDetails))
}

// SetFilesystemOptions mocks base method.
func (m *MockCacheOptions) SetFilesystemOptions(value fs.Options) CacheOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFilesystemOptions", value)
	ret0, _ := ret[0].(CacheOptions)
	return ret0
}

// SetFilesystemOptions indicates an expected call of SetFilesystemOptions.
func (mr *MockCacheOptionsMockRecorder) SetFilesystemOptions(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFilesystemOptions", reflect.TypeOf((*MockCacheOptions)(nil).SetFilesystemOptions), value)
}

// SetInstrumentOptions mocks base method.
func (m *MockCacheOptions) SetInstrumentOptions(value instrument.Options) CacheOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInstrumentOptions", value)
	ret0, _ := ret[0].(CacheOptions)
	return ret0
}

// SetInstrumentOptions indicates an expected call of SetInstrumentOptions.
func (mr *MockCacheOptionsMockRecorder) SetInstrumentOptions(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInstrumentOptions", reflect.TypeOf((*MockCacheOptions)(nil).SetInstrumentOptions), value)
}

// SetNamespaceDetails mocks base method.
func (m *MockCacheOptions) SetNamespaceDetails(value []NamespaceDetails) CacheOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNamespaceDetails", value)
	ret0, _ := ret[0].(CacheOptions)
	return ret0
}

// SetNamespaceDetails indicates an expected call of SetNamespaceDetails.
func (mr *MockCacheOptionsMockRecorder) SetNamespaceDetails(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespaceDetails", reflect.TypeOf((*MockCacheOptions)(nil).SetNamespaceDetails), value)
}

// Validate mocks base method.
func (m *MockCacheOptions) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockCacheOptionsMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockCacheOptions)(nil).Validate))
}

// MockSeriesRef is a mock of SeriesRef interface.
type MockSeriesRef struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesRefMockRecorder
}

// MockSeriesRefMockRecorder is the mock recorder for MockSeriesRef.
type MockSeriesRefMockRecorder struct {
	mock *MockSeriesRef
}

// NewMockSeriesRef creates a new mock instance.
func NewMockSeriesRef(ctrl *gomock.Controller) *MockSeriesRef {
	mock := &MockSeriesRef{ctrl: ctrl}
	mock.recorder = &MockSeriesRefMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeriesRef) EXPECT() *MockSeriesRefMockRecorder {
	return m.recorder
}

// LoadBlock mocks base method.
func (m *MockSeriesRef) LoadBlock(block block.DatabaseBlock, writeType series.WriteType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadBlock", block, writeType)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadBlock indicates an expected call of LoadBlock.
func (mr *MockSeriesRefMockRecorder) LoadBlock(block, writeType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadBlock", reflect.TypeOf((*MockSeriesRef)(nil).LoadBlock), block, writeType)
}

// UniqueIndex mocks base method.
func (m *MockSeriesRef) UniqueIndex() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UniqueIndex")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// UniqueIndex indicates an expected call of UniqueIndex.
func (mr *MockSeriesRefMockRecorder) UniqueIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UniqueIndex", reflect.TypeOf((*MockSeriesRef)(nil).UniqueIndex))
}

// Write mocks base method.
func (m *MockSeriesRef) Write(ctx context.Context, timestamp time.UnixNano, value float64, unit time.Unit, annotation []byte, wOpts series.WriteOptions) (bool, series.WriteType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, timestamp, value, unit, annotation, wOpts)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(series.WriteType)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Write indicates an expected call of Write.
func (mr *MockSeriesRefMockRecorder) Write(ctx, timestamp, value, unit, annotation, wOpts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockSeriesRef)(nil).Write), ctx, timestamp, value, unit, annotation, wOpts)
}

// MockSeriesRefResolver is a mock of SeriesRefResolver interface.
type MockSeriesRefResolver struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesRefResolverMockRecorder
}

// MockSeriesRefResolverMockRecorder is the mock recorder for MockSeriesRefResolver.
type MockSeriesRefResolverMockRecorder struct {
	mock *MockSeriesRefResolver
}

// NewMockSeriesRefResolver creates a new mock instance.
func NewMockSeriesRefResolver(ctrl *gomock.Controller) *MockSeriesRefResolver {
	mock := &MockSeriesRefResolver{ctrl: ctrl}
	mock.recorder = &MockSeriesRefResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeriesRefResolver) EXPECT() *MockSeriesRefResolverMockRecorder {
	return m.recorder
}

// ReleaseRef mocks base method.
func (m *MockSeriesRefResolver) ReleaseRef() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseRef")
}

// ReleaseRef indicates an expected call of ReleaseRef.
func (mr *MockSeriesRefResolverMockRecorder) ReleaseRef() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseRef", reflect.TypeOf((*MockSeriesRefResolver)(nil).ReleaseRef))
}

// SeriesRef mocks base method.
func (m *MockSeriesRefResolver) SeriesRef() (SeriesRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeriesRef")
	ret0, _ := ret[0].(SeriesRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SeriesRef indicates an expected call of SeriesRef.
func (mr *MockSeriesRefResolverMockRecorder) SeriesRef() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeriesRef", reflect.TypeOf((*MockSeriesRefResolver)(nil).SeriesRef))
}
