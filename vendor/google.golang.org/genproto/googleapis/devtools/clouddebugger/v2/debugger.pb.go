// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/clouddebugger/v2/debugger.proto

package clouddebugger

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request to set a breakpoint
type SetBreakpointRequest struct {
	// ID of the debuggee where the breakpoint is to be set.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// Breakpoint specification to set.
	// The field `location` of the breakpoint must be set.
	Breakpoint *Breakpoint `protobuf:"bytes,2,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetBreakpointRequest) Reset()         { *m = SetBreakpointRequest{} }
func (m *SetBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*SetBreakpointRequest) ProtoMessage()    {}
func (*SetBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{0}
}

func (m *SetBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBreakpointRequest.Unmarshal(m, b)
}
func (m *SetBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBreakpointRequest.Marshal(b, m, deterministic)
}
func (m *SetBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBreakpointRequest.Merge(m, src)
}
func (m *SetBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_SetBreakpointRequest.Size(m)
}
func (m *SetBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetBreakpointRequest proto.InternalMessageInfo

func (m *SetBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *SetBreakpointRequest) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

func (m *SetBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for setting a breakpoint.
type SetBreakpointResponse struct {
	// Breakpoint resource.
	// The field `id` is guaranteed to be set (in addition to the echoed fileds).
	Breakpoint           *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetBreakpointResponse) Reset()         { *m = SetBreakpointResponse{} }
func (m *SetBreakpointResponse) String() string { return proto.CompactTextString(m) }
func (*SetBreakpointResponse) ProtoMessage()    {}
func (*SetBreakpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{1}
}

func (m *SetBreakpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetBreakpointResponse.Unmarshal(m, b)
}
func (m *SetBreakpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetBreakpointResponse.Marshal(b, m, deterministic)
}
func (m *SetBreakpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBreakpointResponse.Merge(m, src)
}
func (m *SetBreakpointResponse) XXX_Size() int {
	return xxx_messageInfo_SetBreakpointResponse.Size(m)
}
func (m *SetBreakpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBreakpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetBreakpointResponse proto.InternalMessageInfo

func (m *SetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to get breakpoint information.
type GetBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to get.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to get.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId,proto3" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBreakpointRequest) Reset()         { *m = GetBreakpointRequest{} }
func (m *GetBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*GetBreakpointRequest) ProtoMessage()    {}
func (*GetBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{2}
}

func (m *GetBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreakpointRequest.Unmarshal(m, b)
}
func (m *GetBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreakpointRequest.Marshal(b, m, deterministic)
}
func (m *GetBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreakpointRequest.Merge(m, src)
}
func (m *GetBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_GetBreakpointRequest.Size(m)
}
func (m *GetBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreakpointRequest proto.InternalMessageInfo

func (m *GetBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *GetBreakpointRequest) GetBreakpointId() string {
	if m != nil {
		return m.BreakpointId
	}
	return ""
}

func (m *GetBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for getting breakpoint information.
type GetBreakpointResponse struct {
	// Complete breakpoint state.
	// The fields `id` and `location` are guaranteed to be set.
	Breakpoint           *Breakpoint `protobuf:"bytes,1,opt,name=breakpoint,proto3" json:"breakpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetBreakpointResponse) Reset()         { *m = GetBreakpointResponse{} }
func (m *GetBreakpointResponse) String() string { return proto.CompactTextString(m) }
func (*GetBreakpointResponse) ProtoMessage()    {}
func (*GetBreakpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{3}
}

func (m *GetBreakpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBreakpointResponse.Unmarshal(m, b)
}
func (m *GetBreakpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBreakpointResponse.Marshal(b, m, deterministic)
}
func (m *GetBreakpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBreakpointResponse.Merge(m, src)
}
func (m *GetBreakpointResponse) XXX_Size() int {
	return xxx_messageInfo_GetBreakpointResponse.Size(m)
}
func (m *GetBreakpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBreakpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBreakpointResponse proto.InternalMessageInfo

func (m *GetBreakpointResponse) GetBreakpoint() *Breakpoint {
	if m != nil {
		return m.Breakpoint
	}
	return nil
}

// Request to delete a breakpoint.
type DeleteBreakpointRequest struct {
	// ID of the debuggee whose breakpoint to delete.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// ID of the breakpoint to delete.
	BreakpointId string `protobuf:"bytes,2,opt,name=breakpoint_id,json=breakpointId,proto3" json:"breakpoint_id,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,3,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBreakpointRequest) Reset()         { *m = DeleteBreakpointRequest{} }
func (m *DeleteBreakpointRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBreakpointRequest) ProtoMessage()    {}
func (*DeleteBreakpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{4}
}

func (m *DeleteBreakpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBreakpointRequest.Unmarshal(m, b)
}
func (m *DeleteBreakpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBreakpointRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBreakpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBreakpointRequest.Merge(m, src)
}
func (m *DeleteBreakpointRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBreakpointRequest.Size(m)
}
func (m *DeleteBreakpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBreakpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBreakpointRequest proto.InternalMessageInfo

func (m *DeleteBreakpointRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *DeleteBreakpointRequest) GetBreakpointId() string {
	if m != nil {
		return m.BreakpointId
	}
	return ""
}

func (m *DeleteBreakpointRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Request to list breakpoints.
type ListBreakpointsRequest struct {
	// ID of the debuggee whose breakpoints to list.
	DebuggeeId string `protobuf:"bytes,1,opt,name=debuggee_id,json=debuggeeId,proto3" json:"debuggee_id,omitempty"`
	// When set to `true`, the response includes the list of breakpoints set by
	// any user. Otherwise, it includes only breakpoints set by the caller.
	IncludeAllUsers bool `protobuf:"varint,2,opt,name=include_all_users,json=includeAllUsers,proto3" json:"include_all_users,omitempty"`
	// When set to `true`, the response includes active and inactive
	// breakpoints. Otherwise, it includes only active breakpoints.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive,proto3" json:"include_inactive,omitempty"`
	// When set, the response includes only breakpoints with the specified action.
	Action *ListBreakpointsRequest_BreakpointActionValue `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	// This field is deprecated. The following fields are always stripped out of
	// the result: `stack_frames`, `evaluated_expressions` and `variable_table`.
	StripResults bool `protobuf:"varint,5,opt,name=strip_results,json=stripResults,proto3" json:"strip_results,omitempty"` // Deprecated: Do not use.
	// A wait token that, if specified, blocks the call until the breakpoints
	// list has changed, or a server selected timeout has expired.  The value
	// should be set from the last response. The error code
	// `google.rpc.Code.ABORTED` (RPC) is returned on wait timeout, which
	// should be called again with the same `wait_token`.
	WaitToken string `protobuf:"bytes,6,opt,name=wait_token,json=waitToken,proto3" json:"wait_token,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,8,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBreakpointsRequest) Reset()         { *m = ListBreakpointsRequest{} }
func (m *ListBreakpointsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBreakpointsRequest) ProtoMessage()    {}
func (*ListBreakpointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{5}
}

func (m *ListBreakpointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsRequest.Unmarshal(m, b)
}
func (m *ListBreakpointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsRequest.Marshal(b, m, deterministic)
}
func (m *ListBreakpointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsRequest.Merge(m, src)
}
func (m *ListBreakpointsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsRequest.Size(m)
}
func (m *ListBreakpointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsRequest proto.InternalMessageInfo

func (m *ListBreakpointsRequest) GetDebuggeeId() string {
	if m != nil {
		return m.DebuggeeId
	}
	return ""
}

func (m *ListBreakpointsRequest) GetIncludeAllUsers() bool {
	if m != nil {
		return m.IncludeAllUsers
	}
	return false
}

func (m *ListBreakpointsRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListBreakpointsRequest) GetAction() *ListBreakpointsRequest_BreakpointActionValue {
	if m != nil {
		return m.Action
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListBreakpointsRequest) GetStripResults() bool {
	if m != nil {
		return m.StripResults
	}
	return false
}

func (m *ListBreakpointsRequest) GetWaitToken() string {
	if m != nil {
		return m.WaitToken
	}
	return ""
}

func (m *ListBreakpointsRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Wrapper message for `Breakpoint.Action`. Defines a filter on the action
// field of breakpoints.
type ListBreakpointsRequest_BreakpointActionValue struct {
	// Only breakpoints with the specified action will pass the filter.
	Value                Breakpoint_Action `protobuf:"varint,1,opt,name=value,proto3,enum=google.devtools.clouddebugger.v2.Breakpoint_Action" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListBreakpointsRequest_BreakpointActionValue) Reset() {
	*m = ListBreakpointsRequest_BreakpointActionValue{}
}
func (m *ListBreakpointsRequest_BreakpointActionValue) String() string {
	return proto.CompactTextString(m)
}
func (*ListBreakpointsRequest_BreakpointActionValue) ProtoMessage() {}
func (*ListBreakpointsRequest_BreakpointActionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{5, 0}
}

func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Unmarshal(m, b)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Marshal(b, m, deterministic)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Merge(m, src)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.Size(m)
}
func (m *ListBreakpointsRequest_BreakpointActionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsRequest_BreakpointActionValue proto.InternalMessageInfo

func (m *ListBreakpointsRequest_BreakpointActionValue) GetValue() Breakpoint_Action {
	if m != nil {
		return m.Value
	}
	return Breakpoint_CAPTURE
}

// Response for listing breakpoints.
type ListBreakpointsResponse struct {
	// List of breakpoints matching the request.
	// The fields `id` and `location` are guaranteed to be set on each breakpoint.
	// The fields: `stack_frames`, `evaluated_expressions` and `variable_table`
	// are cleared on each breakpoint regardless of its status.
	Breakpoints []*Breakpoint `protobuf:"bytes,1,rep,name=breakpoints,proto3" json:"breakpoints,omitempty"`
	// A wait token that can be used in the next call to `list` (REST) or
	// `ListBreakpoints` (RPC) to block until the list of breakpoints has changes.
	NextWaitToken        string   `protobuf:"bytes,2,opt,name=next_wait_token,json=nextWaitToken,proto3" json:"next_wait_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBreakpointsResponse) Reset()         { *m = ListBreakpointsResponse{} }
func (m *ListBreakpointsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBreakpointsResponse) ProtoMessage()    {}
func (*ListBreakpointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{6}
}

func (m *ListBreakpointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBreakpointsResponse.Unmarshal(m, b)
}
func (m *ListBreakpointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBreakpointsResponse.Marshal(b, m, deterministic)
}
func (m *ListBreakpointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBreakpointsResponse.Merge(m, src)
}
func (m *ListBreakpointsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBreakpointsResponse.Size(m)
}
func (m *ListBreakpointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBreakpointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBreakpointsResponse proto.InternalMessageInfo

func (m *ListBreakpointsResponse) GetBreakpoints() []*Breakpoint {
	if m != nil {
		return m.Breakpoints
	}
	return nil
}

func (m *ListBreakpointsResponse) GetNextWaitToken() string {
	if m != nil {
		return m.NextWaitToken
	}
	return ""
}

// Request to list debuggees.
type ListDebuggeesRequest struct {
	// Project number of a Google Cloud project whose debuggees to list.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// When set to `true`, the result includes all debuggees. Otherwise, the
	// result includes only debuggees that are active.
	IncludeInactive bool `protobuf:"varint,3,opt,name=include_inactive,json=includeInactive,proto3" json:"include_inactive,omitempty"`
	// The client version making the call.
	// Schema: `domain/type/version` (e.g., `google.com/intellij/v1`).
	ClientVersion        string   `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDebuggeesRequest) Reset()         { *m = ListDebuggeesRequest{} }
func (m *ListDebuggeesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDebuggeesRequest) ProtoMessage()    {}
func (*ListDebuggeesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{7}
}

func (m *ListDebuggeesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDebuggeesRequest.Unmarshal(m, b)
}
func (m *ListDebuggeesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDebuggeesRequest.Marshal(b, m, deterministic)
}
func (m *ListDebuggeesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDebuggeesRequest.Merge(m, src)
}
func (m *ListDebuggeesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDebuggeesRequest.Size(m)
}
func (m *ListDebuggeesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDebuggeesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDebuggeesRequest proto.InternalMessageInfo

func (m *ListDebuggeesRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ListDebuggeesRequest) GetIncludeInactive() bool {
	if m != nil {
		return m.IncludeInactive
	}
	return false
}

func (m *ListDebuggeesRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

// Response for listing debuggees.
type ListDebuggeesResponse struct {
	// List of debuggees accessible to the calling user.
	// The fields `debuggee.id` and `description` are guaranteed to be set.
	// The `description` field is a human readable field provided by agents and
	// can be displayed to users.
	Debuggees            []*Debuggee `protobuf:"bytes,1,rep,name=debuggees,proto3" json:"debuggees,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListDebuggeesResponse) Reset()         { *m = ListDebuggeesResponse{} }
func (m *ListDebuggeesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDebuggeesResponse) ProtoMessage()    {}
func (*ListDebuggeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_682c945d4794189b, []int{8}
}

func (m *ListDebuggeesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDebuggeesResponse.Unmarshal(m, b)
}
func (m *ListDebuggeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDebuggeesResponse.Marshal(b, m, deterministic)
}
func (m *ListDebuggeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDebuggeesResponse.Merge(m, src)
}
func (m *ListDebuggeesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDebuggeesResponse.Size(m)
}
func (m *ListDebuggeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDebuggeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDebuggeesResponse proto.InternalMessageInfo

func (m *ListDebuggeesResponse) GetDebuggees() []*Debuggee {
	if m != nil {
		return m.Debuggees
	}
	return nil
}

func init() {
	proto.RegisterType((*SetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.SetBreakpointRequest")
	proto.RegisterType((*SetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.SetBreakpointResponse")
	proto.RegisterType((*GetBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.GetBreakpointRequest")
	proto.RegisterType((*GetBreakpointResponse)(nil), "google.devtools.clouddebugger.v2.GetBreakpointResponse")
	proto.RegisterType((*DeleteBreakpointRequest)(nil), "google.devtools.clouddebugger.v2.DeleteBreakpointRequest")
	proto.RegisterType((*ListBreakpointsRequest)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest")
	proto.RegisterType((*ListBreakpointsRequest_BreakpointActionValue)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsRequest.BreakpointActionValue")
	proto.RegisterType((*ListBreakpointsResponse)(nil), "google.devtools.clouddebugger.v2.ListBreakpointsResponse")
	proto.RegisterType((*ListDebuggeesRequest)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesRequest")
	proto.RegisterType((*ListDebuggeesResponse)(nil), "google.devtools.clouddebugger.v2.ListDebuggeesResponse")
}

func init() {
	proto.RegisterFile("google/devtools/clouddebugger/v2/debugger.proto", fileDescriptor_682c945d4794189b)
}

var fileDescriptor_682c945d4794189b = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x06, 0x9d, 0xe5, 0xc7, 0xc7, 0x71, 0x92, 0x11, 0xf9, 0x11, 0xbc, 0x3f, 0x43, 0xfb, 0xcb,
	0xb2, 0x41, 0x1a, 0x94, 0x61, 0x4b, 0xb6, 0x9b, 0xc5, 0xcb, 0xe0, 0x18, 0xc8, 0x82, 0x40, 0xdb,
	0x3c, 0x60, 0x08, 0x60, 0xc8, 0x36, 0x23, 0xa8, 0x51, 0x44, 0x55, 0xa4, 0xdc, 0x16, 0x41, 0x6e,
	0x52, 0xa0, 0xf7, 0x45, 0x5f, 0xa0, 0xd7, 0x45, 0x81, 0xbe, 0x40, 0x0b, 0xf4, 0x3a, 0xbd, 0xec,
	0x2b, 0xf4, 0x41, 0x0a, 0x49, 0x64, 0x2c, 0xbb, 0x6a, 0x6d, 0x39, 0x40, 0xee, 0xa8, 0x8f, 0x3c,
	0x87, 0xdf, 0xf7, 0xf1, 0xf0, 0x50, 0xa0, 0xdb, 0x94, 0xda, 0x2e, 0xd1, 0xbb, 0xa4, 0xc7, 0x29,
	0x75, 0x99, 0xde, 0x71, 0x69, 0xd8, 0xed, 0x92, 0x76, 0x68, 0xdb, 0x24, 0xd0, 0x7b, 0x86, 0x2e,
	0xc7, 0x9a, 0x1f, 0x50, 0x4e, 0x71, 0x35, 0x09, 0xd0, 0x64, 0x80, 0x36, 0x10, 0xa0, 0xf5, 0x8c,
	0xca, 0xa7, 0x22, 0xa5, 0xe5, 0x3b, 0xba, 0xe5, 0x79, 0x94, 0x5b, 0xdc, 0xa1, 0x1e, 0x4b, 0xe2,
	0x2b, 0xdf, 0x8f, 0xde, 0xd0, 0xe2, 0x96, 0x58, 0xfc, 0x89, 0x58, 0x1c, 0x7f, 0xb5, 0xc3, 0x63,
	0x9d, 0x9c, 0xfa, 0xfc, 0x5e, 0x32, 0xa9, 0x3e, 0x45, 0xb0, 0xfc, 0x37, 0xe1, 0xb5, 0x80, 0x58,
	0x27, 0x3e, 0x75, 0x3c, 0x6e, 0x92, 0xdb, 0x21, 0x61, 0x1c, 0x7f, 0x01, 0x25, 0x91, 0x8f, 0xb4,
	0x9c, 0xae, 0x82, 0xaa, 0x68, 0xbd, 0x68, 0x82, 0x84, 0x1a, 0x5d, 0xbc, 0x0f, 0xd0, 0xbe, 0x8a,
	0x52, 0x0a, 0x55, 0xb4, 0x5e, 0x32, 0x7e, 0xd0, 0x46, 0x09, 0xd3, 0x52, 0x3b, 0xa5, 0xe2, 0xf1,
	0xd7, 0xb0, 0xd0, 0x71, 0x1d, 0xe2, 0xf1, 0x56, 0x8f, 0x04, 0xcc, 0xa1, 0x9e, 0xf2, 0x51, 0xbc,
	0x63, 0x39, 0x41, 0x9b, 0x09, 0xa8, 0x12, 0x58, 0x19, 0x62, 0xcb, 0x7c, 0xea, 0x31, 0x32, 0xc4,
	0x06, 0x5d, 0x8f, 0x8d, 0x7a, 0x1f, 0xc1, 0x72, 0x7d, 0x22, 0x57, 0xbe, 0x84, 0x72, 0x3f, 0x4f,
	0xb4, 0xa4, 0x10, 0x2f, 0x99, 0xef, 0x83, 0x8d, 0x6e, 0x0e, 0xb1, 0xf5, 0x1b, 0x10, 0xfb, 0x00,
	0xc1, 0xda, 0x2e, 0x71, 0x09, 0x27, 0x37, 0xa7, 0x77, 0x2a, 0x4b, 0xef, 0xab, 0x29, 0x58, 0xdd,
	0x77, 0x58, 0x4a, 0x31, 0x1b, 0x9b, 0xc7, 0x06, 0x7c, 0xec, 0x78, 0x1d, 0x37, 0xec, 0x92, 0x96,
	0xe5, 0xba, 0xad, 0x90, 0x91, 0x80, 0xc5, 0x5c, 0xe6, 0xcc, 0x45, 0x31, 0xb1, 0xe3, 0xba, 0xff,
	0x46, 0x30, 0xfe, 0x0e, 0x96, 0xe4, 0x5a, 0xc7, 0xb3, 0x3a, 0xdc, 0xe9, 0x91, 0x98, 0x50, 0x7f,
	0x69, 0x43, 0xc0, 0xf8, 0x18, 0x66, 0xa2, 0x91, 0x38, 0xa1, 0x92, 0x71, 0x30, 0xda, 0xe5, 0x6c,
	0x05, 0x29, 0xf3, 0x77, 0xe2, 0x84, 0x4d, 0xcb, 0x0d, 0x89, 0x29, 0xb2, 0xe3, 0x6f, 0xa1, 0xcc,
	0x78, 0xe0, 0xf8, 0xad, 0x80, 0xb0, 0xd0, 0xe5, 0x4c, 0x99, 0x8e, 0xf8, 0xd4, 0x0a, 0x0a, 0x32,
	0xe7, 0xe3, 0x09, 0x33, 0xc1, 0xf1, 0x67, 0x00, 0x77, 0x2c, 0x87, 0xb7, 0x38, 0x3d, 0x21, 0x9e,
	0x32, 0x13, 0xfb, 0x50, 0x8c, 0x90, 0x7f, 0x22, 0x20, 0xc3, 0xe9, 0xb9, 0x0c, 0xa7, 0x2b, 0x6d,
	0x58, 0xc9, 0xe4, 0x83, 0x1b, 0x30, 0xdd, 0x8b, 0x06, 0xb1, 0xc3, 0x0b, 0xc6, 0x66, 0x9e, 0xa2,
	0xd2, 0x92, 0x44, 0x66, 0x92, 0x41, 0x7d, 0x88, 0x60, 0xed, 0x1d, 0x2f, 0x44, 0x01, 0x1f, 0x40,
	0xa9, 0x5f, 0x20, 0x4c, 0x41, 0xd5, 0xa9, 0xdc, 0x15, 0x9c, 0x4e, 0x80, 0xbf, 0x81, 0x45, 0x8f,
	0xdc, 0xe5, 0xad, 0x94, 0x35, 0x49, 0x1d, 0x96, 0x23, 0xf8, 0x3f, 0x69, 0x8f, 0x7a, 0x81, 0x60,
	0x39, 0xe2, 0xb4, 0x2b, 0x0a, 0xe7, 0xaa, 0xbe, 0x14, 0x98, 0xf5, 0x03, 0x7a, 0x8b, 0x74, 0xb8,
	0x08, 0x94, 0x9f, 0x79, 0x8a, 0x65, 0xcc, 0x6b, 0x6d, 0xc1, 0xca, 0x10, 0x07, 0xe1, 0xca, 0x1e,
	0x14, 0x65, 0x45, 0x4b, 0x4f, 0x36, 0x46, 0x7b, 0x22, 0xf3, 0x98, 0xfd, 0x60, 0xe3, 0xc5, 0x2c,
	0x14, 0x05, 0x1e, 0x18, 0xf8, 0x12, 0x41, 0x79, 0xa0, 0x6b, 0xe2, 0x9f, 0x47, 0xa7, 0xcd, 0x7a,
	0x14, 0x2a, 0xbf, 0xe4, 0x8e, 0x4b, 0xa4, 0xa9, 0x7b, 0x17, 0xaf, 0xdf, 0x3c, 0x2a, 0xd4, 0xd4,
	0x9f, 0xd2, 0x8f, 0xa1, 0x7e, 0x45, 0x58, 0x3f, 0x4b, 0xdd, 0xee, 0x73, 0x3d, 0x75, 0xb4, 0x3a,
	0x23, 0xfc, 0xd7, 0xf4, 0x43, 0x11, 0x89, 0xa9, 0xe7, 0x15, 0x53, 0x9f, 0x50, 0x4c, 0xfd, 0x43,
	0x62, 0xf0, 0xef, 0xb9, 0xc5, 0x9c, 0x0d, 0xf4, 0xca, 0x73, 0xfc, 0x0c, 0xc1, 0xd2, 0x70, 0xeb,
	0xc5, 0xdb, 0xe3, 0x9c, 0x79, 0x66, 0xbb, 0xae, 0xac, 0xca, 0x50, 0xf9, 0xd6, 0x6b, 0x7f, 0x46,
	0x6f, 0xbd, 0x64, 0xbc, 0x71, 0x7d, 0xc6, 0x2f, 0x11, 0x2c, 0x0e, 0xdd, 0x6a, 0xbc, 0x35, 0x69,
	0x53, 0xac, 0x6c, 0x4f, 0x10, 0x29, 0x0e, 0x61, 0x2b, 0x96, 0x64, 0xe0, 0x1f, 0xf3, 0x4a, 0xc2,
	0x8f, 0x11, 0x94, 0x07, 0x2e, 0xe0, 0x38, 0x15, 0x94, 0xd5, 0x35, 0xc6, 0xa9, 0xa0, 0xcc, 0x9b,
	0xae, 0x7e, 0x1e, 0x93, 0x57, 0xf0, 0x6a, 0x36, 0xf9, 0xda, 0x73, 0x04, 0x5f, 0x75, 0xe8, 0xe9,
	0xc8, 0xf4, 0xb5, 0xb2, 0xbc, 0xe5, 0x87, 0xd1, 0x81, 0x1f, 0xa2, 0xff, 0xff, 0x12, 0x21, 0x36,
	0x75, 0x2d, 0xcf, 0xd6, 0x68, 0x60, 0xeb, 0x36, 0xf1, 0xe2, 0x72, 0x10, 0x7f, 0xa9, 0x96, 0xef,
	0xb0, 0xf7, 0xff, 0x38, 0xfe, 0x36, 0x00, 0x3c, 0x29, 0x28, 0xf5, 0x24, 0xdf, 0x1f, 0x11, 0x2c,
	0x7b, 0x4d, 0xa0, 0x35, 0x8d, 0x4b, 0x39, 0x75, 0x14, 0x4f, 0x1d, 0xc9, 0xa9, 0xa3, 0xa6, 0xd1,
	0x9e, 0x89, 0xf7, 0xdb, 0x7c, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x13, 0xd7, 0xe1, 0x18, 0x0b,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Debugger2Client is the client API for Debugger2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Debugger2Client interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user has access to.
	ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error)
}

type debugger2Client struct {
	cc *grpc.ClientConn
}

func NewDebugger2Client(cc *grpc.ClientConn) Debugger2Client {
	return &debugger2Client{cc}
}

func (c *debugger2Client) SetBreakpoint(ctx context.Context, in *SetBreakpointRequest, opts ...grpc.CallOption) (*SetBreakpointResponse, error) {
	out := new(SetBreakpointResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) GetBreakpoint(ctx context.Context, in *GetBreakpointRequest, opts ...grpc.CallOption) (*GetBreakpointResponse, error) {
	out := new(GetBreakpointResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) DeleteBreakpoint(ctx context.Context, in *DeleteBreakpointRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListBreakpoints(ctx context.Context, in *ListBreakpointsRequest, opts ...grpc.CallOption) (*ListBreakpointsResponse, error) {
	out := new(ListBreakpointsResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugger2Client) ListDebuggees(ctx context.Context, in *ListDebuggeesRequest, opts ...grpc.CallOption) (*ListDebuggeesResponse, error) {
	out := new(ListDebuggeesResponse)
	err := c.cc.Invoke(ctx, "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Debugger2Server is the server API for Debugger2 service.
type Debugger2Server interface {
	// Sets the breakpoint to the debuggee.
	SetBreakpoint(context.Context, *SetBreakpointRequest) (*SetBreakpointResponse, error)
	// Gets breakpoint information.
	GetBreakpoint(context.Context, *GetBreakpointRequest) (*GetBreakpointResponse, error)
	// Deletes the breakpoint from the debuggee.
	DeleteBreakpoint(context.Context, *DeleteBreakpointRequest) (*empty.Empty, error)
	// Lists all breakpoints for the debuggee.
	ListBreakpoints(context.Context, *ListBreakpointsRequest) (*ListBreakpointsResponse, error)
	// Lists all the debuggees that the user has access to.
	ListDebuggees(context.Context, *ListDebuggeesRequest) (*ListDebuggeesResponse, error)
}

func RegisterDebugger2Server(s *grpc.Server, srv Debugger2Server) {
	s.RegisterService(&_Debugger2_serviceDesc, srv)
}

func _Debugger2_SetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).SetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/SetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).SetBreakpoint(ctx, req.(*SetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_GetBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).GetBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/GetBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).GetBreakpoint(ctx, req.(*GetBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_DeleteBreakpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBreakpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/DeleteBreakpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).DeleteBreakpoint(ctx, req.(*DeleteBreakpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListBreakpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBreakpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListBreakpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListBreakpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListBreakpoints(ctx, req.(*ListBreakpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debugger2_ListDebuggees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDebuggeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Debugger2Server).ListDebuggees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouddebugger.v2.Debugger2/ListDebuggees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Debugger2Server).ListDebuggees(ctx, req.(*ListDebuggeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debugger2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouddebugger.v2.Debugger2",
	HandlerType: (*Debugger2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBreakpoint",
			Handler:    _Debugger2_SetBreakpoint_Handler,
		},
		{
			MethodName: "GetBreakpoint",
			Handler:    _Debugger2_GetBreakpoint_Handler,
		},
		{
			MethodName: "DeleteBreakpoint",
			Handler:    _Debugger2_DeleteBreakpoint_Handler,
		},
		{
			MethodName: "ListBreakpoints",
			Handler:    _Debugger2_ListBreakpoints_Handler,
		},
		{
			MethodName: "ListDebuggees",
			Handler:    _Debugger2_ListDebuggees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouddebugger/v2/debugger.proto",
}