// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package
package mock

import (
	context "context"
	reflect "reflect"

	eth "github.com/INFURA/go-ethlibs/eth"
	jsonrpc "github.com/INFURA/go-ethlibs/jsonrpc"
	node "github.com/INFURA/go-ethlibs/node"
	gomock "github.com/golang/mock/gomock"
)

// MockRequester is a mock of Requester interface.
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterMockRecorder
}

// MockRequesterMockRecorder is the mock recorder for MockRequester.
type MockRequesterMockRecorder struct {
	mock *MockRequester
}

// NewMockRequester creates a new mock instance.
func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequester) EXPECT() *MockRequesterMockRecorder {
	return m.recorder
}

// Request mocks base method.
func (m *MockRequester) Request(ctx context.Context, r *jsonrpc.Request) (*jsonrpc.RawResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", ctx, r)
	ret0, _ := ret[0].(*jsonrpc.RawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockRequesterMockRecorder) Request(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockRequester)(nil).Request), ctx, r)
}

// MockSubscriber is a mock of Subscriber interface.
type MockSubscriber struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriberMockRecorder
}

// MockSubscriberMockRecorder is the mock recorder for MockSubscriber.
type MockSubscriberMockRecorder struct {
	mock *MockSubscriber
}

// NewMockSubscriber creates a new mock instance.
func NewMockSubscriber(ctrl *gomock.Controller) *MockSubscriber {
	mock := &MockSubscriber{ctrl: ctrl}
	mock.recorder = &MockSubscriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriber) EXPECT() *MockSubscriberMockRecorder {
	return m.recorder
}

// Subscribe mocks base method.
func (m *MockSubscriber) Subscribe(ctx context.Context, r *jsonrpc.Request) (node.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, r)
	ret0, _ := ret[0].(node.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscriberMockRecorder) Subscribe(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscriber)(nil).Subscribe), ctx, r)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockClient) BlockByHash(ctx context.Context, hash string, full bool) (*eth.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash, full)
	ret0, _ := ret[0].(*eth.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockClientMockRecorder) BlockByHash(ctx, hash, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockClient)(nil).BlockByHash), ctx, hash, full)
}

// BlockByNumber mocks base method.
func (m *MockClient) BlockByNumber(ctx context.Context, number uint64, full bool) (*eth.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number, full)
	ret0, _ := ret[0].(*eth.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockClientMockRecorder) BlockByNumber(ctx, number, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockClient)(nil).BlockByNumber), ctx, number, full)
}

// BlockByNumberOrTag mocks base method.
func (m *MockClient) BlockByNumberOrTag(ctx context.Context, numberOrTag eth.BlockNumberOrTag, full bool) (*eth.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumberOrTag", ctx, numberOrTag, full)
	ret0, _ := ret[0].(*eth.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumberOrTag indicates an expected call of BlockByNumberOrTag.
func (mr *MockClientMockRecorder) BlockByNumberOrTag(ctx, numberOrTag, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumberOrTag", reflect.TypeOf((*MockClient)(nil).BlockByNumberOrTag), ctx, numberOrTag, full)
}

// BlockNumber mocks base method.
func (m *MockClient) BlockNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockClientMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockClient)(nil).BlockNumber), ctx)
}

// Call mocks base method.
func (m *MockClient) Call(ctx context.Context, msg eth.Transaction) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockClientMockRecorder) Call(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockClient)(nil).EstimateGas), ctx, msg)
}

// GetAccounts mocks base method.
func (m *MockClient) GetAccounts(ctx context.Context) ([]eth.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx)
	ret0, _ := ret[0].([]eth.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of GetAccounts.
func (mr *MockClientMockRecorder) GetAccounts(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockClient)(nil).EstimateGas), ctx, msg)
}

// ChainId mocks base method.
func (m *MockClient) ChainId(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainId", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainId indicates an expected call of ChainId.
func (mr *MockClientMockRecorder) ChainId(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainId", reflect.TypeOf((*MockClient)(nil).ChainId), ctx)
}

// EstimateGas mocks base method.
func (m *MockClient) EstimateGas(ctx context.Context, msg eth.Transaction) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateGas", ctx, msg)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateGas indicates an expected call of EstimateGas.
func (mr *MockClientMockRecorder) EstimateGas(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateGas", reflect.TypeOf((*MockClient)(nil).EstimateGas), ctx, msg)
}

// GasPrice mocks base method.
func (m *MockClient) GasPrice(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice.
func (mr *MockClientMockRecorder) GasPrice(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockClient)(nil).GasPrice), ctx)
}

// GetTransactionCount mocks base method.
func (m *MockClient) GetTransactionCount(ctx context.Context, address eth.Address, numberOrTag eth.BlockNumberOrTag) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCount", ctx, address, numberOrTag)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCount indicates an expected call of GetTransactionCount.
func (mr *MockClientMockRecorder) GetTransactionCount(ctx, address, numberOrTag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCount", reflect.TypeOf((*MockClient)(nil).GetTransactionCount), ctx, address, numberOrTag)
}

// IsBidirectional mocks base method.
func (m *MockClient) IsBidirectional() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBidirectional")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBidirectional indicates an expected call of IsBidirectional.
func (mr *MockClientMockRecorder) IsBidirectional() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBidirectional", reflect.TypeOf((*MockClient)(nil).IsBidirectional))
}

// Logs mocks base method.
func (m *MockClient) Logs(ctx context.Context, filter eth.LogFilter) ([]eth.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, filter)
	ret0, _ := ret[0].([]eth.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockClientMockRecorder) Logs(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockClient)(nil).Logs), ctx, filter)
}

// MaxPriorityFeePerGas mocks base method.
func (m *MockClient) MaxPriorityFeePerGas(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPriorityFeePerGas", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxPriorityFeePerGas indicates an expected call of MaxPriorityFeePerGas.
func (mr *MockClientMockRecorder) MaxPriorityFeePerGas(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPriorityFeePerGas", reflect.TypeOf((*MockClient)(nil).MaxPriorityFeePerGas), ctx)
}

// NetVersion mocks base method.
func (m *MockClient) NetVersion(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetVersion", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetVersion indicates an expected call of NetVersion.
func (mr *MockClientMockRecorder) NetVersion(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetVersion", reflect.TypeOf((*MockClient)(nil).NetVersion), ctx)
}

// Request mocks base method.
func (m *MockClient) Request(ctx context.Context, r *jsonrpc.Request) (*jsonrpc.RawResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", ctx, r)
	ret0, _ := ret[0].(*jsonrpc.RawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockClientMockRecorder) Request(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockClient)(nil).Request), ctx, r)
}

// SendRawTransaction mocks base method.
func (m *MockClient) SendRawTransaction(ctx context.Context, msg string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRawTransaction", ctx, msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawTransaction indicates an expected call of SendRawTransaction.
func (mr *MockClientMockRecorder) SendRawTransaction(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawTransaction", reflect.TypeOf((*MockClient)(nil).SendRawTransaction), ctx, msg)
}

// Subscribe mocks base method.
func (m *MockClient) Subscribe(ctx context.Context, r *jsonrpc.Request) (node.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, r)
	ret0, _ := ret[0].(node.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockClientMockRecorder) Subscribe(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), ctx, r)
}

// SubscribeNewHeads mocks base method.
func (m *MockClient) SubscribeNewHeads(ctx context.Context) (node.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewHeads", ctx)
	ret0, _ := ret[0].(node.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewHeads indicates an expected call of SubscribeNewHeads.
func (mr *MockClientMockRecorder) SubscribeNewHeads(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewHeads", reflect.TypeOf((*MockClient)(nil).SubscribeNewHeads), ctx)
}

// SubscribeNewPendingTransactions mocks base method.
func (m *MockClient) SubscribeNewPendingTransactions(ctx context.Context) (node.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewPendingTransactions", ctx)
	ret0, _ := ret[0].(node.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeNewPendingTransactions indicates an expected call of SubscribeNewPendingTransactions.
func (mr *MockClientMockRecorder) SubscribeNewPendingTransactions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewPendingTransactions", reflect.TypeOf((*MockClient)(nil).SubscribeNewPendingTransactions), ctx)
}

// TransactionByHash mocks base method.
func (m *MockClient) TransactionByHash(ctx context.Context, hash string) (*eth.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(*eth.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockClientMockRecorder) TransactionByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockClient)(nil).TransactionByHash), ctx, hash)
}

// TransactionReceipt mocks base method.
func (m *MockClient) TransactionReceipt(ctx context.Context, hash string) (*eth.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, hash)
	ret0, _ := ret[0].(*eth.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockClientMockRecorder) TransactionReceipt(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockClient)(nil).TransactionReceipt), ctx, hash)
}

// URL mocks base method.
func (m *MockClient) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockClientMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockClient)(nil).URL))
}

// MockSubscription is a mock of Subscription interface.
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription.
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance.
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// Ch mocks base method.
func (m *MockSubscription) Ch() <-chan *jsonrpc.Notification {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ch")
	ret0, _ := ret[0].(<-chan *jsonrpc.Notification)
	return ret0
}

// Ch indicates an expected call of Ch.
func (mr *MockSubscriptionMockRecorder) Ch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ch", reflect.TypeOf((*MockSubscription)(nil).Ch))
}

// ID mocks base method.
func (m *MockSubscription) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSubscriptionMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSubscription)(nil).ID))
}

// Response mocks base method.
func (m *MockSubscription) Response() *jsonrpc.RawResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Response")
	ret0, _ := ret[0].(*jsonrpc.RawResponse)
	return ret0
}

// Response indicates an expected call of Response.
func (mr *MockSubscriptionMockRecorder) Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockSubscription)(nil).Response))
}

// Unsubscribe mocks base method.
func (m *MockSubscription) Unsubscribe(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockSubscriptionMockRecorder) Unsubscribe(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockSubscription)(nil).Unsubscribe), ctx)
}
