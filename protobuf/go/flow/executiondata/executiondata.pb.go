// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/executiondata/executiondata.proto

package executiondata

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	entities "github.com/onflow/flow/protobuf/go/flow/entities"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request for GetExecutionDataByBlockID
type GetExecutionDataByBlockIDRequest struct {
	// Block ID of the block to get execution data for.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,2,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetExecutionDataByBlockIDRequest) Reset()         { *m = GetExecutionDataByBlockIDRequest{} }
func (m *GetExecutionDataByBlockIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetExecutionDataByBlockIDRequest) ProtoMessage()    {}
func (*GetExecutionDataByBlockIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{0}
}

func (m *GetExecutionDataByBlockIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Unmarshal(m, b)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Marshal(b, m, deterministic)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecutionDataByBlockIDRequest.Merge(m, src)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetExecutionDataByBlockIDRequest.Size(m)
}
func (m *GetExecutionDataByBlockIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecutionDataByBlockIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecutionDataByBlockIDRequest proto.InternalMessageInfo

func (m *GetExecutionDataByBlockIDRequest) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetExecutionDataByBlockIDRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_JSON_CDC_V0
}

// The response for GetExecutionDataByBlockID
type GetExecutionDataByBlockIDResponse struct {
	// BlockExecutionData for the block.
	BlockExecutionData   *entities.BlockExecutionData `protobuf:"bytes,1,opt,name=block_execution_data,json=blockExecutionData,proto3" json:"block_execution_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetExecutionDataByBlockIDResponse) Reset()         { *m = GetExecutionDataByBlockIDResponse{} }
func (m *GetExecutionDataByBlockIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetExecutionDataByBlockIDResponse) ProtoMessage()    {}
func (*GetExecutionDataByBlockIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{1}
}

func (m *GetExecutionDataByBlockIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Unmarshal(m, b)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Marshal(b, m, deterministic)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetExecutionDataByBlockIDResponse.Merge(m, src)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetExecutionDataByBlockIDResponse.Size(m)
}
func (m *GetExecutionDataByBlockIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetExecutionDataByBlockIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetExecutionDataByBlockIDResponse proto.InternalMessageInfo

func (m *GetExecutionDataByBlockIDResponse) GetBlockExecutionData() *entities.BlockExecutionData {
	if m != nil {
		return m.BlockExecutionData
	}
	return nil
}

// The request for SubscribeExecutionData
type SubscribeExecutionDataRequest struct {
	// Block ID of the first block to get execution data for.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockId []byte `protobuf:"bytes,1,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	// Block height of the first block to get execution data for.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,3,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscribeExecutionDataRequest) Reset()         { *m = SubscribeExecutionDataRequest{} }
func (m *SubscribeExecutionDataRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeExecutionDataRequest) ProtoMessage()    {}
func (*SubscribeExecutionDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{2}
}

func (m *SubscribeExecutionDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Unmarshal(m, b)
}
func (m *SubscribeExecutionDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeExecutionDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeExecutionDataRequest.Merge(m, src)
}
func (m *SubscribeExecutionDataRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeExecutionDataRequest.Size(m)
}
func (m *SubscribeExecutionDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeExecutionDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeExecutionDataRequest proto.InternalMessageInfo

func (m *SubscribeExecutionDataRequest) GetStartBlockId() []byte {
	if m != nil {
		return m.StartBlockId
	}
	return nil
}

func (m *SubscribeExecutionDataRequest) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *SubscribeExecutionDataRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_JSON_CDC_V0
}

// The response for SubscribeExecutionData
type SubscribeExecutionDataResponse struct {
	// Block height of the block containing the execution data.
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// BlockExecutionData for the block.
	// Note: The block's ID is included within the BlockExecutionData.
	BlockExecutionData *entities.BlockExecutionData `protobuf:"bytes,2,opt,name=block_execution_data,json=blockExecutionData,proto3" json:"block_execution_data,omitempty"`
	// Timestamp from the block containing the execution data.
	BlockTimestamp       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeExecutionDataResponse) Reset()         { *m = SubscribeExecutionDataResponse{} }
func (m *SubscribeExecutionDataResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeExecutionDataResponse) ProtoMessage()    {}
func (*SubscribeExecutionDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{3}
}

func (m *SubscribeExecutionDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Unmarshal(m, b)
}
func (m *SubscribeExecutionDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeExecutionDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeExecutionDataResponse.Merge(m, src)
}
func (m *SubscribeExecutionDataResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeExecutionDataResponse.Size(m)
}
func (m *SubscribeExecutionDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeExecutionDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeExecutionDataResponse proto.InternalMessageInfo

func (m *SubscribeExecutionDataResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SubscribeExecutionDataResponse) GetBlockExecutionData() *entities.BlockExecutionData {
	if m != nil {
		return m.BlockExecutionData
	}
	return nil
}

func (m *SubscribeExecutionDataResponse) GetBlockTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.BlockTimestamp
	}
	return nil
}

// The request for SubscribeEvents
type SubscribeEventsRequest struct {
	// Block ID of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockId []byte `protobuf:"bytes,1,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	// Block height of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	// Filter to apply to events for each block searched.
	// If no filter is provided, all events are returned.
	Filter *EventFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Interval in block heights at which the server should return a heartbeat
	// message to the client. The heartbeat is a normal SubscribeEventsResponse
	// with no events, and allows clients to track which blocks were searched.
	// Clients can use this information to determine which block to start from
	// when reconnecting.
	//
	// The interval is calculated from the last response returned, which could be
	// either another heartbeat or a response containing events.
	HeartbeatInterval uint64 `protobuf:"varint,4,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,5,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscribeEventsRequest) Reset()         { *m = SubscribeEventsRequest{} }
func (m *SubscribeEventsRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeEventsRequest) ProtoMessage()    {}
func (*SubscribeEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{4}
}

func (m *SubscribeEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeEventsRequest.Unmarshal(m, b)
}
func (m *SubscribeEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeEventsRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeEventsRequest.Merge(m, src)
}
func (m *SubscribeEventsRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeEventsRequest.Size(m)
}
func (m *SubscribeEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeEventsRequest proto.InternalMessageInfo

func (m *SubscribeEventsRequest) GetStartBlockId() []byte {
	if m != nil {
		return m.StartBlockId
	}
	return nil
}

func (m *SubscribeEventsRequest) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *SubscribeEventsRequest) GetFilter() *EventFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *SubscribeEventsRequest) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func (m *SubscribeEventsRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_JSON_CDC_V0
}

// The response for SubscribeEvents
type SubscribeEventsResponse struct {
	// Block ID of the block containing the events.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	// Block height of the block containing the events.
	BlockHeight uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Events matching the EventFilter in the request.
	// The API may return no events which signals a periodic heartbeat. This
	// allows clients to track which blocks were searched. Client can use this
	// information to determine which block to start from when reconnecting.
	Events []*entities.Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
	// Timestamp from the block containing the events.
	BlockTimestamp       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=block_timestamp,json=blockTimestamp,proto3" json:"block_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SubscribeEventsResponse) Reset()         { *m = SubscribeEventsResponse{} }
func (m *SubscribeEventsResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeEventsResponse) ProtoMessage()    {}
func (*SubscribeEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{5}
}

func (m *SubscribeEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeEventsResponse.Unmarshal(m, b)
}
func (m *SubscribeEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeEventsResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeEventsResponse.Merge(m, src)
}
func (m *SubscribeEventsResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeEventsResponse.Size(m)
}
func (m *SubscribeEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeEventsResponse proto.InternalMessageInfo

func (m *SubscribeEventsResponse) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *SubscribeEventsResponse) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SubscribeEventsResponse) GetEvents() []*entities.Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *SubscribeEventsResponse) GetBlockTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.BlockTimestamp
	}
	return nil
}

// EventFilter defines the filter to apply to block events.
// Filters are applied as an OR operation, i.e. any event matching any of the
// filters is returned. If no filters are provided, all events are returned. If
// there are any invalid filters, the API will return an InvalidArgument error.
type EventFilter struct {
	// A list of full event types to include.
	//
	// All events exactly matching any of the provided event types will be
	// returned.
	//
	// Event types have 2 formats:
	//   - Protocol events:
	//     flow.[event name]
	//   - Smart contract events:
	//     A.[contract address].[contract name].[event name]
	EventType []string `protobuf:"bytes,1,rep,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// A list of contracts who's events should be included.
	//
	// All events emitted by any of the provided contracts will be returned.
	//
	// Contracts have the following name formats:
	//   - Protocol events:
	//     flow
	//   - Smart contract events:
	//     A.[contract address].[contract name]
	//
	// This filter matches on the full contract including its address, not just
	// the contract's name.
	Contract []string `protobuf:"bytes,2,rep,name=contract,proto3" json:"contract,omitempty"`
	// A list of addresses who's events should be included.
	//
	// All events emitted by any contract held by any of the provided addresses
	// will be returned.
	//
	// Addresses must be Flow account addresses in hex format and valid for the
	// network the node is connected to. i.e. only a mainnet address is valid for
	// a mainnet node. Addresses may optionally include the 0x prefix.
	Address              []string `protobuf:"bytes,3,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventFilter) Reset()         { *m = EventFilter{} }
func (m *EventFilter) String() string { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()    {}
func (*EventFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{6}
}

func (m *EventFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFilter.Unmarshal(m, b)
}
func (m *EventFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFilter.Marshal(b, m, deterministic)
}
func (m *EventFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFilter.Merge(m, src)
}
func (m *EventFilter) XXX_Size() int {
	return xxx_messageInfo_EventFilter.Size(m)
}
func (m *EventFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFilter.DiscardUnknown(m)
}

var xxx_messageInfo_EventFilter proto.InternalMessageInfo

func (m *EventFilter) GetEventType() []string {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *EventFilter) GetContract() []string {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *EventFilter) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

// request for GetRegisterValues
type GetRegisterValuesRequest struct {
	// Block height of the execution state being queried.
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Register IDs of the Ledger.RegisterID format with an owner and key.
	RegisterIds          []*entities.RegisterID `protobuf:"bytes,2,rep,name=register_ids,json=registerIds,proto3" json:"register_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GetRegisterValuesRequest) Reset()         { *m = GetRegisterValuesRequest{} }
func (m *GetRegisterValuesRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegisterValuesRequest) ProtoMessage()    {}
func (*GetRegisterValuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{7}
}

func (m *GetRegisterValuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegisterValuesRequest.Unmarshal(m, b)
}
func (m *GetRegisterValuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegisterValuesRequest.Marshal(b, m, deterministic)
}
func (m *GetRegisterValuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegisterValuesRequest.Merge(m, src)
}
func (m *GetRegisterValuesRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegisterValuesRequest.Size(m)
}
func (m *GetRegisterValuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegisterValuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegisterValuesRequest proto.InternalMessageInfo

func (m *GetRegisterValuesRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *GetRegisterValuesRequest) GetRegisterIds() []*entities.RegisterID {
	if m != nil {
		return m.RegisterIds
	}
	return nil
}

// response for GetRegisterValues
type GetRegisterValuesResponse struct {
	// raw register values at the given height.
	Values               [][]byte `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegisterValuesResponse) Reset()         { *m = GetRegisterValuesResponse{} }
func (m *GetRegisterValuesResponse) String() string { return proto.CompactTextString(m) }
func (*GetRegisterValuesResponse) ProtoMessage()    {}
func (*GetRegisterValuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{8}
}

func (m *GetRegisterValuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegisterValuesResponse.Unmarshal(m, b)
}
func (m *GetRegisterValuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegisterValuesResponse.Marshal(b, m, deterministic)
}
func (m *GetRegisterValuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegisterValuesResponse.Merge(m, src)
}
func (m *GetRegisterValuesResponse) XXX_Size() int {
	return xxx_messageInfo_GetRegisterValuesResponse.Size(m)
}
func (m *GetRegisterValuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegisterValuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegisterValuesResponse proto.InternalMessageInfo

func (m *GetRegisterValuesResponse) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

// The request for SubscribeAccountStatuses
type SubscribeAccountStatusesRequest struct {
	// Block ID of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockId []byte `protobuf:"bytes,1,opt,name=start_block_id,json=startBlockId,proto3" json:"start_block_id,omitempty"`
	// Block height of the first block to search for events.
	// Only one of start_block_id and start_block_height may be provided,
	// otherwise an InvalidArgument error is returned. If neither are provided,
	// the latest sealed block is used.
	StartBlockHeight uint64 `protobuf:"varint,2,opt,name=start_block_height,json=startBlockHeight,proto3" json:"start_block_height,omitempty"`
	// Filter to apply to events for each block searched.
	// If no filter is provided, all statuses are returned.
	Filter *StatusFilter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Interval in block heights at which the server should return a heartbeat
	// message to the client. The heartbeat is a normal SubscribeAccountStatusesResponse
	// with no events, and allows clients to track which blocks were searched.
	// Clients can use this information to determine which block to start from
	// when reconnecting.
	//
	// The interval is calculated from the last response returned, which could be
	// either another heartbeat or a response containing events.
	HeartbeatInterval uint64 `protobuf:"varint,4,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	// Preferred event encoding version of the block events payload.
	// Possible variants:
	// 1. CCF
	// 2. JSON-CDC
	EventEncodingVersion entities.EventEncodingVersion `protobuf:"varint,5,opt,name=event_encoding_version,json=eventEncodingVersion,proto3,enum=flow.entities.EventEncodingVersion" json:"event_encoding_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscribeAccountStatusesRequest) Reset()         { *m = SubscribeAccountStatusesRequest{} }
func (m *SubscribeAccountStatusesRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeAccountStatusesRequest) ProtoMessage()    {}
func (*SubscribeAccountStatusesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{9}
}

func (m *SubscribeAccountStatusesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeAccountStatusesRequest.Unmarshal(m, b)
}
func (m *SubscribeAccountStatusesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeAccountStatusesRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeAccountStatusesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeAccountStatusesRequest.Merge(m, src)
}
func (m *SubscribeAccountStatusesRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeAccountStatusesRequest.Size(m)
}
func (m *SubscribeAccountStatusesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeAccountStatusesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeAccountStatusesRequest proto.InternalMessageInfo

func (m *SubscribeAccountStatusesRequest) GetStartBlockId() []byte {
	if m != nil {
		return m.StartBlockId
	}
	return nil
}

func (m *SubscribeAccountStatusesRequest) GetStartBlockHeight() uint64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *SubscribeAccountStatusesRequest) GetFilter() *StatusFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *SubscribeAccountStatusesRequest) GetHeartbeatInterval() uint64 {
	if m != nil {
		return m.HeartbeatInterval
	}
	return 0
}

func (m *SubscribeAccountStatusesRequest) GetEventEncodingVersion() entities.EventEncodingVersion {
	if m != nil {
		return m.EventEncodingVersion
	}
	return entities.EventEncodingVersion_JSON_CDC_V0
}

// The response for SubscribeAccountStatuses
type SubscribeAccountStatusesResponse struct {
	// Block ID of the block containing the events.
	BlockId []byte `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeAccountStatusesResponse) Reset()         { *m = SubscribeAccountStatusesResponse{} }
func (m *SubscribeAccountStatusesResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeAccountStatusesResponse) ProtoMessage()    {}
func (*SubscribeAccountStatusesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{10}
}

func (m *SubscribeAccountStatusesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeAccountStatusesResponse.Unmarshal(m, b)
}
func (m *SubscribeAccountStatusesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeAccountStatusesResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeAccountStatusesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeAccountStatusesResponse.Merge(m, src)
}
func (m *SubscribeAccountStatusesResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeAccountStatusesResponse.Size(m)
}
func (m *SubscribeAccountStatusesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeAccountStatusesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeAccountStatusesResponse proto.InternalMessageInfo

func (m *SubscribeAccountStatusesResponse) GetBlockId() []byte {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *SubscribeAccountStatusesResponse) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *SubscribeAccountStatusesResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// StatusesFilter defines the filter to apply to block events.
// Filters are applied as an OR operation, i.e. any event matching any of the
// filters is returned. If no filters are provided, all events are returned. If
// there are any invalid filters, the API will return an InvalidArgument error.
type StatusFilter struct {
	// A list of full event types to include.
	//
	// All events exactly matching any of the provided event types will be
	// returned.
	//
	// Event types have 2 formats:
	//   - Protocol events:
	//     flow.[event name]
	EventType            []string `protobuf:"bytes,1,rep,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusFilter) Reset()         { *m = StatusFilter{} }
func (m *StatusFilter) String() string { return proto.CompactTextString(m) }
func (*StatusFilter) ProtoMessage()    {}
func (*StatusFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_da469632570513fb, []int{11}
}

func (m *StatusFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusFilter.Unmarshal(m, b)
}
func (m *StatusFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusFilter.Marshal(b, m, deterministic)
}
func (m *StatusFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusFilter.Merge(m, src)
}
func (m *StatusFilter) XXX_Size() int {
	return xxx_messageInfo_StatusFilter.Size(m)
}
func (m *StatusFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StatusFilter proto.InternalMessageInfo

func (m *StatusFilter) GetEventType() []string {
	if m != nil {
		return m.EventType
	}
	return nil
}

func init() {
	proto.RegisterType((*GetExecutionDataByBlockIDRequest)(nil), "flow.executiondata.GetExecutionDataByBlockIDRequest")
	proto.RegisterType((*GetExecutionDataByBlockIDResponse)(nil), "flow.executiondata.GetExecutionDataByBlockIDResponse")
	proto.RegisterType((*SubscribeExecutionDataRequest)(nil), "flow.executiondata.SubscribeExecutionDataRequest")
	proto.RegisterType((*SubscribeExecutionDataResponse)(nil), "flow.executiondata.SubscribeExecutionDataResponse")
	proto.RegisterType((*SubscribeEventsRequest)(nil), "flow.executiondata.SubscribeEventsRequest")
	proto.RegisterType((*SubscribeEventsResponse)(nil), "flow.executiondata.SubscribeEventsResponse")
	proto.RegisterType((*EventFilter)(nil), "flow.executiondata.EventFilter")
	proto.RegisterType((*GetRegisterValuesRequest)(nil), "flow.executiondata.GetRegisterValuesRequest")
	proto.RegisterType((*GetRegisterValuesResponse)(nil), "flow.executiondata.GetRegisterValuesResponse")
	proto.RegisterType((*SubscribeAccountStatusesRequest)(nil), "flow.executiondata.SubscribeAccountStatusesRequest")
	proto.RegisterType((*SubscribeAccountStatusesResponse)(nil), "flow.executiondata.SubscribeAccountStatusesResponse")
	proto.RegisterType((*StatusFilter)(nil), "flow.executiondata.StatusFilter")
}

func init() {
	proto.RegisterFile("flow/executiondata/executiondata.proto", fileDescriptor_da469632570513fb)
}

var fileDescriptor_da469632570513fb = []byte{
	// 825 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x96, 0x13, 0x08, 0xe4, 0x24, 0xe2, 0x67, 0x84, 0xb8, 0x8e, 0x75, 0xb9, 0x24, 0xbe, 0x55,
	0x15, 0xb5, 0xe0, 0xb4, 0x01, 0xd4, 0x2e, 0xba, 0x21, 0x85, 0xd2, 0xec, 0x2a, 0x83, 0x90, 0xca,
	0x26, 0xf2, 0xcf, 0xe0, 0x58, 0x0d, 0x9e, 0xd4, 0x33, 0x4e, 0x41, 0x5d, 0x75, 0xd5, 0x4a, 0x7d,
	0x81, 0xbe, 0x45, 0xd7, 0x7d, 0x85, 0x2e, 0xfa, 0x06, 0x7d, 0x97, 0x2a, 0xe3, 0xb1, 0x89, 0x13,
	0xc7, 0x10, 0x21, 0xa4, 0x6e, 0x22, 0x9d, 0x39, 0xdf, 0x99, 0xf3, 0xcd, 0x77, 0x7e, 0x1c, 0x78,
	0x78, 0xde, 0x23, 0x1f, 0x1a, 0xf8, 0x12, 0x5b, 0x01, 0x73, 0x89, 0x67, 0x1b, 0xcc, 0x48, 0x5a,
	0x5a, 0xdf, 0x27, 0x8c, 0x20, 0x34, 0xc4, 0x69, 0x09, 0x8f, 0x52, 0x0f, 0x63, 0x3d, 0xe6, 0x32,
	0x17, 0xd3, 0x86, 0xd9, 0x23, 0xd6, 0xbb, 0x4e, 0x0c, 0xe9, 0x5c, 0x47, 0x2b, 0x95, 0x24, 0x12,
	0x0f, 0xb0, 0xc7, 0x84, 0xeb, 0xdf, 0xa4, 0xcb, 0xc7, 0x8e, 0x4b, 0x19, 0xf6, 0x85, 0x77, 0xd3,
	0x21, 0xc4, 0xe9, 0xe1, 0x06, 0xb7, 0xcc, 0xe0, 0xbc, 0xc1, 0xdc, 0x0b, 0x4c, 0x99, 0x71, 0xd1,
	0x0f, 0x01, 0xea, 0x37, 0x09, 0xaa, 0x47, 0x98, 0x1d, 0x46, 0x59, 0x0f, 0x0c, 0x66, 0xb4, 0xae,
	0x5a, 0x43, 0x2a, 0xed, 0x03, 0x1d, 0xbf, 0x0f, 0x30, 0x65, 0xa8, 0x02, 0x8b, 0x21, 0x39, 0xd7,
	0x96, 0xa5, 0xaa, 0x54, 0x2f, 0xeb, 0x0b, 0xdc, 0x6e, 0xdb, 0xe8, 0x2d, 0xac, 0x73, 0x36, 0x1d,
	0xec, 0x59, 0xc4, 0x76, 0x3d, 0xa7, 0x33, 0xc0, 0x3e, 0x75, 0x89, 0x27, 0xe7, 0xaa, 0x52, 0x7d,
	0xa9, 0xf9, 0xbf, 0x16, 0x3e, 0x5c, 0xf0, 0xd3, 0x0e, 0x87, 0xe0, 0x43, 0x81, 0x3d, 0x0d, 0xa1,
	0xfa, 0x1a, 0x4e, 0x39, 0x55, 0x2f, 0xa1, 0x96, 0xc1, 0x8c, 0xf6, 0x89, 0x47, 0x31, 0x3a, 0x86,
	0xb5, 0x34, 0xdd, 0x38, 0xcd, 0x52, 0xb3, 0x36, 0x96, 0x9d, 0x47, 0x27, 0x6e, 0xd4, 0x91, 0x39,
	0x71, 0xa6, 0xfe, 0x92, 0x60, 0xe3, 0x38, 0x30, 0xa9, 0xe5, 0xbb, 0x26, 0x4e, 0xc2, 0x85, 0x22,
	0x0f, 0x60, 0x89, 0x32, 0xc3, 0x67, 0x9d, 0x31, 0x5d, 0xca, 0xfc, 0xb4, 0x25, 0xc4, 0xd9, 0x02,
	0x34, 0x8a, 0xea, 0x62, 0xd7, 0xe9, 0x32, 0x2e, 0xcc, 0x9c, 0xbe, 0x72, 0x8d, 0x7c, 0xcd, 0xcf,
	0x33, 0xa4, 0xcc, 0xdf, 0x55, 0xca, 0xdf, 0x12, 0xfc, 0x37, 0xed, 0x41, 0x42, 0xc8, 0x1a, 0x94,
	0x13, 0x2c, 0x25, 0xce, 0xb2, 0x64, 0x8e, 0x10, 0x9c, 0xa6, 0x75, 0xee, 0x0e, 0x5a, 0xa3, 0x97,
	0xb0, 0x1c, 0x5e, 0x1a, 0x77, 0x26, 0x7f, 0x6e, 0xa9, 0xa9, 0x68, 0x61, 0xef, 0x6a, 0x51, 0xef,
	0x6a, 0x27, 0x11, 0x42, 0x5f, 0xe2, 0x21, 0xb1, 0xad, 0x7e, 0xcf, 0xc1, 0xfa, 0xf5, 0xfb, 0x86,
	0x0a, 0xd0, 0xfb, 0xac, 0xd4, 0x33, 0x28, 0x9c, 0xbb, 0x3d, 0x86, 0x7d, 0x41, 0x75, 0x53, 0x9b,
	0x9c, 0xee, 0xb0, 0x3c, 0xaf, 0x38, 0x4c, 0x17, 0x70, 0xb4, 0x0d, 0xa8, 0x8b, 0x0d, 0x9f, 0x99,
	0xd8, 0x60, 0x1d, 0xd7, 0x63, 0xd8, 0x1f, 0x18, 0x3d, 0x79, 0x8e, 0xa7, 0x59, 0x8d, 0x3d, 0x6d,
	0xe1, 0xc8, 0xe8, 0x88, 0xf9, 0xbb, 0x76, 0xc4, 0x4f, 0x09, 0xfe, 0x99, 0x50, 0x4c, 0xb4, 0x42,
	0xc6, 0xb8, 0x8f, 0x77, 0x49, 0x6e, 0xb2, 0x4b, 0xb6, 0xa0, 0xc0, 0x33, 0x52, 0x39, 0x5f, 0xcd,
	0xd7, 0x4b, 0xcd, 0xb5, 0x34, 0x92, 0xba, 0xc0, 0xa4, 0x95, 0x7f, 0x6e, 0xe6, 0xf2, 0x9b, 0x50,
	0x1a, 0x51, 0x1b, 0x6d, 0x00, 0x84, 0xb2, 0xb1, 0xab, 0x3e, 0x96, 0xa5, 0x6a, 0xbe, 0x5e, 0xd4,
	0x8b, 0xfc, 0xe4, 0xe4, 0xaa, 0x8f, 0x91, 0x02, 0x8b, 0x16, 0xf1, 0x98, 0x6f, 0x58, 0x43, 0xfe,
	0x43, 0x67, 0x6c, 0x23, 0x19, 0x16, 0x0c, 0xdb, 0xf6, 0x31, 0x0d, 0xd9, 0x17, 0xf5, 0xc8, 0x54,
	0x3f, 0x82, 0x7c, 0x84, 0x99, 0x2e, 0xd6, 0xeb, 0xa9, 0xd1, 0x0b, 0x70, 0xdc, 0x63, 0xb7, 0x98,
	0x9d, 0x17, 0x50, 0x8e, 0x56, 0x73, 0xc7, 0xb5, 0x29, 0x4f, 0x5c, 0x6a, 0x56, 0xc6, 0xb4, 0x89,
	0xae, 0x6f, 0x1f, 0xe8, 0xa5, 0x08, 0xde, 0xb6, 0xa9, 0xba, 0x03, 0x95, 0x94, 0xe4, 0xa2, 0x5c,
	0xeb, 0x50, 0x18, 0xf0, 0x13, 0xfe, 0xd4, 0xb2, 0x2e, 0x2c, 0xf5, 0x47, 0x0e, 0x36, 0xe3, 0x12,
	0xef, 0x5b, 0x16, 0x09, 0x3c, 0x76, 0xcc, 0x0c, 0x16, 0x50, 0x7c, 0xaf, 0xd3, 0xf1, 0x7c, 0x6c,
	0x3a, 0xaa, 0x69, 0xd3, 0x11, 0x12, 0xf9, 0x6b, 0xc7, 0x83, 0x40, 0x75, 0xba, 0x74, 0x37, 0x8f,
	0xc9, 0x48, 0x1b, 0xe5, 0x42, 0x8f, 0x30, 0x87, 0xc5, 0xa2, 0xfc, 0x22, 0x2e, 0x4e, 0x51, 0x17,
	0x96, 0xba, 0x0d, 0xe5, 0x51, 0x49, 0x6e, 0xe8, 0xe1, 0xe6, 0xd7, 0x79, 0x58, 0x49, 0xec, 0xd1,
	0xfd, 0x37, 0x6d, 0xf4, 0x45, 0xe2, 0x6d, 0x92, 0xfe, 0xc5, 0x44, 0xbb, 0x69, 0x65, 0xb8, 0xe9,
	0xd3, 0xaf, 0xec, 0xcd, 0x18, 0x25, 0xb4, 0xf9, 0x24, 0x8d, 0x2e, 0xe4, 0xc4, 0xc2, 0x7f, 0x9a,
	0xda, 0x0e, 0x59, 0x5f, 0x5b, 0xa5, 0x39, 0x4b, 0x48, 0xc8, 0xe0, 0x89, 0x84, 0x7a, 0xb0, 0x3c,
	0xb6, 0xe1, 0xd0, 0xa3, 0xec, 0x8b, 0x46, 0x3f, 0x1c, 0xca, 0xe3, 0x5b, 0x61, 0xe3, 0x6c, 0x7d,
	0x58, 0x9d, 0x18, 0x51, 0xb4, 0x35, 0x45, 0xbd, 0xd4, 0x35, 0xa2, 0x6c, 0xdf, 0x12, 0x2d, 0x34,
	0xfe, 0x2c, 0x81, 0x3c, 0xad, 0x49, 0xd1, 0x4e, 0x26, 0xfb, 0xf4, 0x6d, 0xa0, 0xec, 0xce, 0x16,
	0x14, 0xbd, 0xbd, 0x75, 0x06, 0x35, 0xe2, 0x3b, 0x1a, 0xf1, 0x78, 0x78, 0xbc, 0xb0, 0x13, 0xf7,
	0x9c, 0xed, 0x39, 0x2e, 0xeb, 0x06, 0xa6, 0x66, 0x91, 0x8b, 0x46, 0x88, 0x6c, 0xf0, 0x9f, 0xf8,
	0xaf, 0xa9, 0x43, 0x1a, 0x93, 0x7f, 0xa6, 0xcd, 0x02, 0xf7, 0xef, 0xfc, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0xe5, 0x12, 0xaa, 0x69, 0x0b, 0x00, 0x00,
}
