// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.12.4
// source: agent.proto

package api

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 双向消息
type AgentMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Body:
	//
	//	*AgentMessage_FunctionRequest
	//	*AgentMessage_FunctionResult
	//	*AgentMessage_Heartbeat
	//	*AgentMessage_CancelTask
	Body          isAgentMessage_Body `protobuf_oneof:"body"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AgentMessage) Reset() {
	*x = AgentMessage{}
	mi := &file_agent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentMessage) ProtoMessage() {}

func (x *AgentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentMessage.ProtoReflect.Descriptor instead.
func (*AgentMessage) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AgentMessage) GetBody() isAgentMessage_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *AgentMessage) GetFunctionRequest() *FunctionRequest {
	if x != nil {
		if x, ok := x.Body.(*AgentMessage_FunctionRequest); ok {
			return x.FunctionRequest
		}
	}
	return nil
}

func (x *AgentMessage) GetFunctionResult() *FunctionResult {
	if x != nil {
		if x, ok := x.Body.(*AgentMessage_FunctionResult); ok {
			return x.FunctionResult
		}
	}
	return nil
}

func (x *AgentMessage) GetHeartbeat() *Heartbeat {
	if x != nil {
		if x, ok := x.Body.(*AgentMessage_Heartbeat); ok {
			return x.Heartbeat
		}
	}
	return nil
}

func (x *AgentMessage) GetCancelTask() *CancelTask {
	if x != nil {
		if x, ok := x.Body.(*AgentMessage_CancelTask); ok {
			return x.CancelTask
		}
	}
	return nil
}

type isAgentMessage_Body interface {
	isAgentMessage_Body()
}

type AgentMessage_FunctionRequest struct {
	FunctionRequest *FunctionRequest `protobuf:"bytes,1,opt,name=function_request,json=functionRequest,proto3,oneof"`
}

type AgentMessage_FunctionResult struct {
	FunctionResult *FunctionResult `protobuf:"bytes,2,opt,name=function_result,json=functionResult,proto3,oneof"`
}

type AgentMessage_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

type AgentMessage_CancelTask struct {
	CancelTask *CancelTask `protobuf:"bytes,4,opt,name=cancel_task,json=cancelTask,proto3,oneof"`
}

func (*AgentMessage_FunctionRequest) isAgentMessage_Body() {}

func (*AgentMessage_FunctionResult) isAgentMessage_Body() {}

func (*AgentMessage_Heartbeat) isAgentMessage_Body() {}

func (*AgentMessage_CancelTask) isAgentMessage_Body() {}

// 函数调用请求
type FunctionRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RequestId      string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`                 // 请求ID，用于跟踪请求
	FunctionName   string                 `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`        // 要调用的函数名，比如 "train"、"predict"
	Parameters     *_struct.Struct        `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`                                // 参数，支持嵌套复杂结构
	TimeoutSeconds int64                  `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"` // 超时时间，单位：秒（可选）
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FunctionRequest) Reset() {
	*x = FunctionRequest{}
	mi := &file_agent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FunctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionRequest) ProtoMessage() {}

func (x *FunctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionRequest.ProtoReflect.Descriptor instead.
func (*FunctionRequest) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FunctionRequest) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *FunctionRequest) GetParameters() *_struct.Struct {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *FunctionRequest) GetTimeoutSeconds() int64 {
	if x != nil {
		return x.TimeoutSeconds
	}
	return 0
}

// 函数调用结果
type FunctionResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestId     string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`          // 对应请求ID
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`                              // 是否成功
	Result        *_struct.Struct        `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`                                 // 返回的数据（任何格式）
	ErrorMessage  string                 `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // 错误信息（如果失败）
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FunctionResult) Reset() {
	*x = FunctionResult{}
	mi := &file_agent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FunctionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionResult) ProtoMessage() {}

func (x *FunctionResult) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionResult.ProtoReflect.Descriptor instead.
func (*FunctionResult) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionResult) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FunctionResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FunctionResult) GetResult() *_struct.Struct {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *FunctionResult) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// 心跳
type Heartbeat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AgentId       string                 `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	AgentType     string                 `protobuf:"bytes,2,opt,name=agent_type,json=agentType,proto3" json:"agent_type,omitempty"` // 代理类型，用于区分不同应用
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                 // 时间戳
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	mi := &file_agent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{3}
}

func (x *Heartbeat) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Heartbeat) GetAgentType() string {
	if x != nil {
		return x.AgentType
	}
	return ""
}

func (x *Heartbeat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 取消任务
type CancelTask struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestId     string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // 要取消的 request_id
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelTask) Reset() {
	*x = CancelTask{}
	mi := &file_agent_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelTask) ProtoMessage() {}

func (x *CancelTask) ProtoReflect() protoreflect.Message {
	mi := &file_agent_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelTask.ProtoReflect.Descriptor instead.
func (*CancelTask) Descriptor() ([]byte, []int) {
	return file_agent_proto_rawDescGZIP(), []int{4}
}

func (x *CancelTask) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_agent_proto protoreflect.FileDescriptor

var file_agent_proto_rawDesc = string([]byte{
	0x0a, 0x0b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x01, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x41, 0x0a, 0x10, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x63, 0x0a, 0x09,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x2b, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x32, 0x47,
	0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37,
	0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_agent_proto_rawDescOnce sync.Once
	file_agent_proto_rawDescData []byte
)

func file_agent_proto_rawDescGZIP() []byte {
	file_agent_proto_rawDescOnce.Do(func() {
		file_agent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_agent_proto_rawDesc), len(file_agent_proto_rawDesc)))
	})
	return file_agent_proto_rawDescData
}

var file_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agent_proto_goTypes = []any{
	(*AgentMessage)(nil),    // 0: api.AgentMessage
	(*FunctionRequest)(nil), // 1: api.FunctionRequest
	(*FunctionResult)(nil),  // 2: api.FunctionResult
	(*Heartbeat)(nil),       // 3: api.Heartbeat
	(*CancelTask)(nil),      // 4: api.CancelTask
	(*_struct.Struct)(nil),  // 5: google.protobuf.Struct
}
var file_agent_proto_depIdxs = []int32{
	1, // 0: api.AgentMessage.function_request:type_name -> api.FunctionRequest
	2, // 1: api.AgentMessage.function_result:type_name -> api.FunctionResult
	3, // 2: api.AgentMessage.heartbeat:type_name -> api.Heartbeat
	4, // 3: api.AgentMessage.cancel_task:type_name -> api.CancelTask
	5, // 4: api.FunctionRequest.parameters:type_name -> google.protobuf.Struct
	5, // 5: api.FunctionResult.result:type_name -> google.protobuf.Struct
	0, // 6: api.AgentService.AgentStream:input_type -> api.AgentMessage
	0, // 7: api.AgentService.AgentStream:output_type -> api.AgentMessage
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_agent_proto_init() }
func file_agent_proto_init() {
	if File_agent_proto != nil {
		return
	}
	file_agent_proto_msgTypes[0].OneofWrappers = []any{
		(*AgentMessage_FunctionRequest)(nil),
		(*AgentMessage_FunctionResult)(nil),
		(*AgentMessage_Heartbeat)(nil),
		(*AgentMessage_CancelTask)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_agent_proto_rawDesc), len(file_agent_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_proto_goTypes,
		DependencyIndexes: file_agent_proto_depIdxs,
		MessageInfos:      file_agent_proto_msgTypes,
	}.Build()
	File_agent_proto = out.File
	file_agent_proto_goTypes = nil
	file_agent_proto_depIdxs = nil
}
