// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: server/pb/taskx-service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *AddTaskReq) Reset() {
	*x = AddTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pb_taskx_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskReq) ProtoMessage() {}

func (x *AddTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_pb_taskx_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskReq.ProtoReflect.Descriptor instead.
func (*AddTaskReq) Descriptor() ([]byte, []int) {
	return file_server_pb_taskx_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddTaskReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CompleteTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CompleteTaskReq) Reset() {
	*x = CompleteTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pb_taskx_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskReq) ProtoMessage() {}

func (x *CompleteTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_pb_taskx_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskReq.ProtoReflect.Descriptor instead.
func (*CompleteTaskReq) Descriptor() ([]byte, []int) {
	return file_server_pb_taskx_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteTaskReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetTasksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeCompleted bool `protobuf:"varint,1,opt,name=IncludeCompleted,proto3" json:"IncludeCompleted,omitempty"`
}

func (x *GetTasksReq) Reset() {
	*x = GetTasksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pb_taskx_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksReq) ProtoMessage() {}

func (x *GetTasksReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_pb_taskx_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksReq.ProtoReflect.Descriptor instead.
func (*GetTasksReq) Descriptor() ([]byte, []int) {
	return file_server_pb_taskx_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTasksReq) GetIncludeCompleted() bool {
	if x != nil {
		return x.IncludeCompleted
	}
	return false
}

type GetTasksResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
}

func (x *GetTasksResp) Reset() {
	*x = GetTasksResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pb_taskx_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksResp) ProtoMessage() {}

func (x *GetTasksResp) ProtoReflect() protoreflect.Message {
	mi := &file_server_pb_taskx_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksResp.ProtoReflect.Descriptor instead.
func (*GetTasksResp) Descriptor() ([]byte, []int) {
	return file_server_pb_taskx_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTasksResp) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Completed   bool   `protobuf:"varint,3,opt,name=Completed,proto3" json:"Completed,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_pb_taskx_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_server_pb_taskx_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_server_pb_taskx_service_proto_rawDescGZIP(), []int{4}
}

func (x *Task) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

var File_server_pb_taskx_service_proto protoreflect.FileDescriptor

var file_server_pb_taskx_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x78, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x2a, 0x0a, 0x10, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x2e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e,
	0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x56,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xba, 0x01, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x58,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_server_pb_taskx_service_proto_rawDescOnce sync.Once
	file_server_pb_taskx_service_proto_rawDescData = file_server_pb_taskx_service_proto_rawDesc
)

func file_server_pb_taskx_service_proto_rawDescGZIP() []byte {
	file_server_pb_taskx_service_proto_rawDescOnce.Do(func() {
		file_server_pb_taskx_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_pb_taskx_service_proto_rawDescData)
	})
	return file_server_pb_taskx_service_proto_rawDescData
}

var file_server_pb_taskx_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_server_pb_taskx_service_proto_goTypes = []interface{}{
	(*AddTaskReq)(nil),      // 0: pb.AddTaskReq
	(*CompleteTaskReq)(nil), // 1: pb.CompleteTaskReq
	(*GetTasksReq)(nil),     // 2: pb.GetTasksReq
	(*GetTasksResp)(nil),    // 3: pb.GetTasksResp
	(*Task)(nil),            // 4: pb.Task
	(*emptypb.Empty)(nil),   // 5: google.protobuf.Empty
}
var file_server_pb_taskx_service_proto_depIdxs = []int32{
	4, // 0: pb.GetTasksResp.Tasks:type_name -> pb.Task
	0, // 1: pb.TaskXService.AddTask:input_type -> pb.AddTaskReq
	1, // 2: pb.TaskXService.CompleteTask:input_type -> pb.CompleteTaskReq
	5, // 3: pb.TaskXService.GetTasks:input_type -> google.protobuf.Empty
	5, // 4: pb.TaskXService.AddTask:output_type -> google.protobuf.Empty
	5, // 5: pb.TaskXService.CompleteTask:output_type -> google.protobuf.Empty
	3, // 6: pb.TaskXService.GetTasks:output_type -> pb.GetTasksResp
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_pb_taskx_service_proto_init() }
func file_server_pb_taskx_service_proto_init() {
	if File_server_pb_taskx_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_pb_taskx_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_pb_taskx_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_pb_taskx_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_pb_taskx_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_pb_taskx_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_pb_taskx_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_pb_taskx_service_proto_goTypes,
		DependencyIndexes: file_server_pb_taskx_service_proto_depIdxs,
		MessageInfos:      file_server_pb_taskx_service_proto_msgTypes,
	}.Build()
	File_server_pb_taskx_service_proto = out.File
	file_server_pb_taskx_service_proto_rawDesc = nil
	file_server_pb_taskx_service_proto_goTypes = nil
	file_server_pb_taskx_service_proto_depIdxs = nil
}
