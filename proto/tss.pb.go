// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: proto/tss.proto

package proto

import (
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

type Action int32

const (
	Action_KEYGEN      Action = 0
	Action_SIGN        Action = 1
	Action_INIT_KEYGEN Action = 3
	Action_INIT_SIGN   Action = 4
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "KEYGEN",
		1: "SIGN",
		3: "INIT_KEYGEN",
		4: "INIT_SIGN",
	}
	Action_value = map[string]int32{
		"KEYGEN":      0,
		"SIGN":        1,
		"INIT_KEYGEN": 3,
		"INIT_SIGN":   4,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_tss_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_proto_tss_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_tss_proto_rawDescGZIP(), []int{0}
}

type TSSMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Payload       []byte                 `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	From          uint32                 `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To            uint32                 `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Broadcast     bool                   `protobuf:"varint,5,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TSSMessage) Reset() {
	*x = TSSMessage{}
	mi := &file_proto_tss_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TSSMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TSSMessage) ProtoMessage() {}

func (x *TSSMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tss_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TSSMessage.ProtoReflect.Descriptor instead.
func (*TSSMessage) Descriptor() ([]byte, []int) {
	return file_proto_tss_proto_rawDescGZIP(), []int{0}
}

func (x *TSSMessage) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TSSMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TSSMessage) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *TSSMessage) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *TSSMessage) GetBroadcast() bool {
	if x != nil {
		return x.Broadcast
	}
	return false
}

type ActionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Parties       []uint32               `protobuf:"varint,2,rep,packed,name=parties,proto3" json:"parties,omitempty"`
	Threshold     uint32                 `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	MsgHash       []byte                 `protobuf:"bytes,4,opt,name=msg_hash,json=msgHash,proto3" json:"msg_hash,omitempty"`
	ShareData     []byte                 `protobuf:"bytes,5,opt,name=share_data,json=shareData,proto3" json:"share_data,omitempty"`
	Action        Action                 `protobuf:"varint,6,opt,name=action,proto3,enum=tss.Action" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	mi := &file_proto_tss_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tss_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_proto_tss_proto_rawDescGZIP(), []int{1}
}

func (x *ActionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ActionRequest) GetParties() []uint32 {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *ActionRequest) GetThreshold() uint32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *ActionRequest) GetMsgHash() []byte {
	if x != nil {
		return x.MsgHash
	}
	return nil
}

func (x *ActionRequest) GetShareData() []byte {
	if x != nil {
		return x.ShareData
	}
	return nil
}

func (x *ActionRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_KEYGEN
}

type ActionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	mi := &file_proto_tss_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tss_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_proto_tss_proto_rawDescGZIP(), []int{2}
}

func (x *ActionResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ActionResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_tss_proto protoreflect.FileDescriptor

var file_proto_tss_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x74, 0x73, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0a, 0x54, 0x53, 0x53, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x22, 0xc5, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x3e, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4b, 0x45, 0x59, 0x47, 0x45, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e,
	0x49, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x47, 0x45, 0x4e, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x49,
	0x4e, 0x49, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x04, 0x32, 0x81, 0x01, 0x0a, 0x0a, 0x4d,
	0x50, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x74, 0x73,
	0x73, 0x2e, 0x54, 0x53, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x74,
	0x73, 0x73, 0x2e, 0x54, 0x53, 0x53, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x65,
	0x74, 0x64, 0x64, 0x75, 0x64, 0x65, 0x2f, 0x74, 0x73, 0x73, 0x2d, 0x69, 0x6d, 0x70, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_tss_proto_rawDescOnce sync.Once
	file_proto_tss_proto_rawDescData []byte
)

func file_proto_tss_proto_rawDescGZIP() []byte {
	file_proto_tss_proto_rawDescOnce.Do(func() {
		file_proto_tss_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_tss_proto_rawDesc), len(file_proto_tss_proto_rawDesc)))
	})
	return file_proto_tss_proto_rawDescData
}

var file_proto_tss_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_tss_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_tss_proto_goTypes = []any{
	(Action)(0),            // 0: tss.Action
	(*TSSMessage)(nil),     // 1: tss.TSSMessage
	(*ActionRequest)(nil),  // 2: tss.ActionRequest
	(*ActionResponse)(nil), // 3: tss.ActionResponse
}
var file_proto_tss_proto_depIdxs = []int32{
	0, // 0: tss.ActionRequest.action:type_name -> tss.Action
	1, // 1: tss.MPCService.StreamMessages:input_type -> tss.TSSMessage
	2, // 2: tss.MPCService.NotifyAction:input_type -> tss.ActionRequest
	1, // 3: tss.MPCService.StreamMessages:output_type -> tss.TSSMessage
	3, // 4: tss.MPCService.NotifyAction:output_type -> tss.ActionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_tss_proto_init() }
func file_proto_tss_proto_init() {
	if File_proto_tss_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_tss_proto_rawDesc), len(file_proto_tss_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tss_proto_goTypes,
		DependencyIndexes: file_proto_tss_proto_depIdxs,
		EnumInfos:         file_proto_tss_proto_enumTypes,
		MessageInfos:      file_proto_tss_proto_msgTypes,
	}.Build()
	File_proto_tss_proto = out.File
	file_proto_tss_proto_goTypes = nil
	file_proto_tss_proto_depIdxs = nil
}
