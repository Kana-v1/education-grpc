// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: server.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BadgeNumber         int32   `protobuf:"varint,2,opt,name=badgeNumber,proto3" json:"badgeNumber,omitempty"`
	FirstName           string  `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName            string  `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	VacationAccrualRate float32 `protobuf:"fixed32,5,opt,name=vacationAccrualRate,proto3" json:"vacationAccrualRate,omitempty"`
	VacationAccured     float32 `protobuf:"fixed32,6,opt,name=vacationAccured,proto3" json:"vacationAccured,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetBadgeNumber() int32 {
	if x != nil {
		return x.BadgeNumber
	}
	return 0
}

func (x *Employee) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Employee) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Employee) GetVacationAccrualRate() float32 {
	if x != nil {
		return x.VacationAccrualRate
	}
	return 0
}

func (x *Employee) GetVacationAccured() float32 {
	if x != nil {
		return x.VacationAccured
	}
	return 0
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

type GetByBadgeNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BadgeNumber int32 `protobuf:"varint,1,opt,name=badgeNumber,proto3" json:"badgeNumber,omitempty"`
}

func (x *GetByBadgeNumberRequest) Reset() {
	*x = GetByBadgeNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByBadgeNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByBadgeNumberRequest) ProtoMessage() {}

func (x *GetByBadgeNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByBadgeNumberRequest.ProtoReflect.Descriptor instead.
func (*GetByBadgeNumberRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetByBadgeNumberRequest) GetBadgeNumber() int32 {
	if x != nil {
		return x.BadgeNumber
	}
	return 0
}

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *EmployeeRequest) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type EmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Employee *Employee `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
}

func (x *EmployeeResponse) Reset() {
	*x = EmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse) ProtoMessage() {}

func (x *EmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *EmployeeResponse) GetEmployee() *Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

type AddPhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AddPhotoRequest) Reset() {
	*x = AddPhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhotoRequest) ProtoMessage() {}

func (x *AddPhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhotoRequest.ProtoReflect.Descriptor instead.
func (*AddPhotoRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *AddPhotoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type AddPhotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOK bool `protobuf:"varint,1,opt,name=isOK,proto3" json:"isOK,omitempty"`
}

func (x *AddPhotoResponse) Reset() {
	*x = AddPhotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPhotoResponse) ProtoMessage() {}

func (x *AddPhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPhotoResponse.ProtoReflect.Descriptor instead.
func (*AddPhotoResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

func (x *AddPhotoResponse) GetIsOK() bool {
	if x != nil {
		return x.IsOK
	}
	return false
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x01, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x62, 0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x76, 0x61, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x72, 0x75, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x76, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x63, 0x72, 0x75, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x76, 0x61, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x75, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0f, 0x76, 0x61, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x75,
	0x72, 0x65, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x79, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x39, 0x0a, 0x10, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x08, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x69, 0x73, 0x4f, 0x4b, 0x32, 0x95, 0x02, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x42, 0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x42, 0x61, 0x64, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x10, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x10, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x16, 0x5a,
	0x14, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_server_proto_goTypes = []interface{}{
	(*Employee)(nil),                // 0: Employee
	(*GetAllRequest)(nil),           // 1: GetAllRequest
	(*GetByBadgeNumberRequest)(nil), // 2: GetByBadgeNumberRequest
	(*EmployeeRequest)(nil),         // 3: EmployeeRequest
	(*EmployeeResponse)(nil),        // 4: EmployeeResponse
	(*AddPhotoRequest)(nil),         // 5: AddPhotoRequest
	(*AddPhotoResponse)(nil),        // 6: AddPhotoResponse
}
var file_server_proto_depIdxs = []int32{
	0, // 0: EmployeeRequest.employee:type_name -> Employee
	0, // 1: EmployeeResponse.employee:type_name -> Employee
	2, // 2: EmployeeService.GetByBadgeNumber:input_type -> GetByBadgeNumberRequest
	1, // 3: EmployeeService.GetAll:input_type -> GetAllRequest
	3, // 4: EmployeeService.Save:input_type -> EmployeeRequest
	3, // 5: EmployeeService.SaveAll:input_type -> EmployeeRequest
	5, // 6: EmployeeService.AddPhoto:input_type -> AddPhotoRequest
	4, // 7: EmployeeService.GetByBadgeNumber:output_type -> EmployeeResponse
	4, // 8: EmployeeService.GetAll:output_type -> EmployeeResponse
	4, // 9: EmployeeService.Save:output_type -> EmployeeResponse
	4, // 10: EmployeeService.SaveAll:output_type -> EmployeeResponse
	6, // 11: EmployeeService.AddPhoto:output_type -> AddPhotoResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByBadgeNumberRequest); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhotoRequest); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPhotoResponse); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
