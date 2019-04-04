// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/index (interfaces: QueryResults,AggregateResults,Block,OnIndexSeries)

// Copyright (c) 2019 Uber Technologies, Inc.
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

// Package index is a generated GoMock package.
package index

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/storage/bootstrap/result"
	"github.com/m3db/m3/src/m3ninx/doc"
	"github.com/m3db/m3/src/x/resource"
	"github.com/m3db/m3x/context"
	"github.com/m3db/m3x/ident"
	time0 "github.com/m3db/m3x/time"

	"github.com/golang/mock/gomock"
)

// MockQueryResults is a mock of QueryResults interface
type MockQueryResults struct {
	ctrl     *gomock.Controller
	recorder *MockQueryResultsMockRecorder
}

// MockQueryResultsMockRecorder is the mock recorder for MockQueryResults
type MockQueryResultsMockRecorder struct {
	mock *MockQueryResults
}

// NewMockQueryResults creates a new mock instance
func NewMockQueryResults(ctrl *gomock.Controller) *MockQueryResults {
	mock := &MockQueryResults{ctrl: ctrl}
	mock.recorder = &MockQueryResultsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryResults) EXPECT() *MockQueryResultsMockRecorder {
	return m.recorder
}

// AddDocuments mocks base method
func (m *MockQueryResults) AddDocuments(arg0 []doc.Document) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDocuments", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDocuments indicates an expected call of AddDocuments
func (mr *MockQueryResultsMockRecorder) AddDocuments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDocuments", reflect.TypeOf((*MockQueryResults)(nil).AddDocuments), arg0)
}

// Finalize mocks base method
func (m *MockQueryResults) Finalize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockQueryResultsMockRecorder) Finalize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockQueryResults)(nil).Finalize))
}

// Map mocks base method
func (m *MockQueryResults) Map() *ResultsMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(*ResultsMap)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockQueryResultsMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockQueryResults)(nil).Map))
}

// Namespace mocks base method
func (m *MockQueryResults) Namespace() ident.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockQueryResultsMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockQueryResults)(nil).Namespace))
}

// NoFinalize mocks base method
func (m *MockQueryResults) NoFinalize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NoFinalize")
}

// NoFinalize indicates an expected call of NoFinalize
func (mr *MockQueryResultsMockRecorder) NoFinalize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoFinalize", reflect.TypeOf((*MockQueryResults)(nil).NoFinalize))
}

// Reset mocks base method
func (m *MockQueryResults) Reset(arg0 ident.ID, arg1 QueryResultsOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0, arg1)
}

// Reset indicates an expected call of Reset
func (mr *MockQueryResultsMockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockQueryResults)(nil).Reset), arg0, arg1)
}

// Size mocks base method
func (m *MockQueryResults) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockQueryResultsMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockQueryResults)(nil).Size))
}

// MockAggregateResults is a mock of AggregateResults interface
type MockAggregateResults struct {
	ctrl     *gomock.Controller
	recorder *MockAggregateResultsMockRecorder
}

// MockAggregateResultsMockRecorder is the mock recorder for MockAggregateResults
type MockAggregateResultsMockRecorder struct {
	mock *MockAggregateResults
}

// NewMockAggregateResults creates a new mock instance
func NewMockAggregateResults(ctrl *gomock.Controller) *MockAggregateResults {
	mock := &MockAggregateResults{ctrl: ctrl}
	mock.recorder = &MockAggregateResultsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAggregateResults) EXPECT() *MockAggregateResultsMockRecorder {
	return m.recorder
}

// AddDocuments mocks base method
func (m *MockAggregateResults) AddDocuments(arg0 []doc.Document) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDocuments", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDocuments indicates an expected call of AddDocuments
func (mr *MockAggregateResultsMockRecorder) AddDocuments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDocuments", reflect.TypeOf((*MockAggregateResults)(nil).AddDocuments), arg0)
}

// Finalize mocks base method
func (m *MockAggregateResults) Finalize() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Finalize")
}

// Finalize indicates an expected call of Finalize
func (mr *MockAggregateResultsMockRecorder) Finalize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockAggregateResults)(nil).Finalize))
}

// Map mocks base method
func (m *MockAggregateResults) Map() *AggregateResultsMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map")
	ret0, _ := ret[0].(*AggregateResultsMap)
	return ret0
}

// Map indicates an expected call of Map
func (mr *MockAggregateResultsMockRecorder) Map() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockAggregateResults)(nil).Map))
}

// Namespace mocks base method
func (m *MockAggregateResults) Namespace() ident.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(ident.ID)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockAggregateResultsMockRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockAggregateResults)(nil).Namespace))
}

// Reset mocks base method
func (m *MockAggregateResults) Reset(arg0 ident.ID, arg1 AggregateResultsOptions) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0, arg1)
}

// Reset indicates an expected call of Reset
func (mr *MockAggregateResultsMockRecorder) Reset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockAggregateResults)(nil).Reset), arg0, arg1)
}

// Size mocks base method
func (m *MockAggregateResults) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockAggregateResultsMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockAggregateResults)(nil).Size))
}

// MockBlock is a mock of Block interface
type MockBlock struct {
	ctrl     *gomock.Controller
	recorder *MockBlockMockRecorder
}

// MockBlockMockRecorder is the mock recorder for MockBlock
type MockBlockMockRecorder struct {
	mock *MockBlock
}

// NewMockBlock creates a new mock instance
func NewMockBlock(ctrl *gomock.Controller) *MockBlock {
	mock := &MockBlock{ctrl: ctrl}
	mock.recorder = &MockBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlock) EXPECT() *MockBlockMockRecorder {
	return m.recorder
}

// AddResults mocks base method
func (m *MockBlock) AddResults(arg0 result.IndexBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResults", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResults indicates an expected call of AddResults
func (mr *MockBlockMockRecorder) AddResults(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResults", reflect.TypeOf((*MockBlock)(nil).AddResults), arg0)
}

// Close mocks base method
func (m *MockBlock) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBlockMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBlock)(nil).Close))
}

// EndTime mocks base method
func (m *MockBlock) EndTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// EndTime indicates an expected call of EndTime
func (mr *MockBlockMockRecorder) EndTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndTime", reflect.TypeOf((*MockBlock)(nil).EndTime))
}

// EvictMutableSegments mocks base method
func (m *MockBlock) EvictMutableSegments() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvictMutableSegments")
	ret0, _ := ret[0].(error)
	return ret0
}

// EvictMutableSegments indicates an expected call of EvictMutableSegments
func (mr *MockBlockMockRecorder) EvictMutableSegments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvictMutableSegments", reflect.TypeOf((*MockBlock)(nil).EvictMutableSegments))
}

// IsSealed mocks base method
func (m *MockBlock) IsSealed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSealed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSealed indicates an expected call of IsSealed
func (mr *MockBlockMockRecorder) IsSealed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSealed", reflect.TypeOf((*MockBlock)(nil).IsSealed))
}

// NeedsMutableSegmentsEvicted mocks base method
func (m *MockBlock) NeedsMutableSegmentsEvicted() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsMutableSegmentsEvicted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedsMutableSegmentsEvicted indicates an expected call of NeedsMutableSegmentsEvicted
func (mr *MockBlockMockRecorder) NeedsMutableSegmentsEvicted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsMutableSegmentsEvicted", reflect.TypeOf((*MockBlock)(nil).NeedsMutableSegmentsEvicted))
}

// Query mocks base method
func (m *MockBlock) Query(arg0 *resource.CancellableLifetime, arg1 Query, arg2 QueryOptions, arg3 BaseResults) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockBlockMockRecorder) Query(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockBlock)(nil).Query), arg0, arg1, arg2, arg3)
}

// Seal mocks base method
func (m *MockBlock) Seal() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal")
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal
func (mr *MockBlockMockRecorder) Seal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockBlock)(nil).Seal))
}

// StartTime mocks base method
func (m *MockBlock) StartTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// StartTime indicates an expected call of StartTime
func (mr *MockBlockMockRecorder) StartTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTime", reflect.TypeOf((*MockBlock)(nil).StartTime))
}

// Stats mocks base method
func (m *MockBlock) Stats(arg0 BlockStatsReporter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockBlockMockRecorder) Stats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockBlock)(nil).Stats), arg0)
}

// Tick mocks base method
func (m *MockBlock) Tick(arg0 context.Cancellable, arg1 time.Time) (BlockTickResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick", arg0, arg1)
	ret0, _ := ret[0].(BlockTickResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tick indicates an expected call of Tick
func (mr *MockBlockMockRecorder) Tick(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockBlock)(nil).Tick), arg0, arg1)
}

// WriteBatch mocks base method
func (m *MockBlock) WriteBatch(arg0 *WriteBatch) (WriteBatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteBatch", arg0)
	ret0, _ := ret[0].(WriteBatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteBatch indicates an expected call of WriteBatch
func (mr *MockBlockMockRecorder) WriteBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteBatch", reflect.TypeOf((*MockBlock)(nil).WriteBatch), arg0)
}

// MockOnIndexSeries is a mock of OnIndexSeries interface
type MockOnIndexSeries struct {
	ctrl     *gomock.Controller
	recorder *MockOnIndexSeriesMockRecorder
}

// MockOnIndexSeriesMockRecorder is the mock recorder for MockOnIndexSeries
type MockOnIndexSeriesMockRecorder struct {
	mock *MockOnIndexSeries
}

// NewMockOnIndexSeries creates a new mock instance
func NewMockOnIndexSeries(ctrl *gomock.Controller) *MockOnIndexSeries {
	mock := &MockOnIndexSeries{ctrl: ctrl}
	mock.recorder = &MockOnIndexSeriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOnIndexSeries) EXPECT() *MockOnIndexSeriesMockRecorder {
	return m.recorder
}

// OnIndexFinalize mocks base method
func (m *MockOnIndexSeries) OnIndexFinalize(arg0 time0.UnixNano) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnIndexFinalize", arg0)
}

// OnIndexFinalize indicates an expected call of OnIndexFinalize
func (mr *MockOnIndexSeriesMockRecorder) OnIndexFinalize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIndexFinalize", reflect.TypeOf((*MockOnIndexSeries)(nil).OnIndexFinalize), arg0)
}

// OnIndexSuccess mocks base method
func (m *MockOnIndexSeries) OnIndexSuccess(arg0 time0.UnixNano) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnIndexSuccess", arg0)
}

// OnIndexSuccess indicates an expected call of OnIndexSuccess
func (mr *MockOnIndexSeriesMockRecorder) OnIndexSuccess(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIndexSuccess", reflect.TypeOf((*MockOnIndexSeries)(nil).OnIndexSuccess), arg0)
}