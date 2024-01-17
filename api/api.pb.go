// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.11
// source: api.proto

package api

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

type Language int32

const (
	Language_C      Language = 0
	Language_JAVA   Language = 1
	Language_PYTHON Language = 2
	Language_GOLANG Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "C",
		1: "JAVA",
		2: "PYTHON",
		3: "GOLANG",
	}
	Language_value = map[string]int32{
		"C":      0,
		"JAVA":   1,
		"PYTHON": 2,
		"GOLANG": 3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	// 状态，即选举状态。当节点处于该状态时，它会认为当前集群中没有 Leader，因此自己进入选举状态。
	Status_Looking Status = 0
	// 状态，即领导者状态，表示已经选出主，且当前节点为 Leader。
	Status_Leading Status = 1
	// 状态，即跟随者状态，集群中已经选出主后，其他非主节点状态更新为 Following，表示对 Leader 的追随。
	Status_Following Status = 2
	// 状态，即观察者状态，表示当前节点为 Observer，持观望态度，没有投票权和选举权。
	Status_Observing Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Looking",
		1: "Leading",
		2: "Following",
		3: "Observing",
	}
	Status_value = map[string]int32{
		"Looking":   0,
		"Leading":   1,
		"Following": 2,
		"Observing": 3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	In         string   `protobuf:"bytes,2,opt,name=In,proto3" json:"In,omitempty"`
	Code       string   `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
	Type       Language `protobuf:"varint,4,opt,name=type,proto3,enum=api.Language" json:"type,omitempty"`
	CpuTime    uint64   `protobuf:"varint,5,opt,name=CpuTime,proto3" json:"CpuTime,omitempty"`
	Memory     uint64   `protobuf:"varint,6,opt,name=Memory,proto3" json:"Memory,omitempty"`
	SourceIp   string   `protobuf:"bytes,7,opt,name=SourceIp,proto3" json:"SourceIp,omitempty"`
	SourcePort string   `protobuf:"bytes,8,opt,name=SourcePort,proto3" json:"SourcePort,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
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
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *Task) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Task) GetType() Language {
	if x != nil {
		return x.Type
	}
	return Language_C
}

func (x *Task) GetCpuTime() uint64 {
	if x != nil {
		return x.CpuTime
	}
	return 0
}

func (x *Task) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *Task) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *Task) GetSourcePort() string {
	if x != nil {
		return x.SourcePort
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Out     string `protobuf:"bytes,2,opt,name=Out,proto3" json:"Out,omitempty"`
	CpuTime uint64 `protobuf:"varint,3,opt,name=CpuTime,proto3" json:"CpuTime,omitempty"`
	Memory  uint64 `protobuf:"varint,4,opt,name=Memory,proto3" json:"Memory,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Result) GetOut() string {
	if x != nil {
		return x.Out
	}
	return ""
}

func (x *Result) GetCpuTime() uint64 {
	if x != nil {
		return x.CpuTime
	}
	return 0
}

func (x *Result) GetMemory() uint64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

type ACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ACK) Reset() {
	*x = ACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACK) ProtoMessage() {}

func (x *ACK) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACK.ProtoReflect.Descriptor instead.
func (*ACK) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ACK) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Ping) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu                uint64 `protobuf:"varint,1,opt,name=Cpu,proto3" json:"Cpu,omitempty"`                 // cpu使用情况
	MemoryAlloc        uint64 `protobuf:"varint,2,opt,name=MemoryAlloc,proto3" json:"MemoryAlloc,omitempty"` // 程序当前使用内存量
	TotalAlloc         uint64 `protobuf:"varint,3,opt,name=TotalAlloc,proto3" json:"TotalAlloc,omitempty"`   // 程序总分配内存量
	Sys                uint64 `protobuf:"varint,4,opt,name=Sys,proto3" json:"Sys,omitempty"`                 // 系统内存
	NumGC              uint32 `protobuf:"varint,5,opt,name=NumGC,proto3" json:"NumGC,omitempty"`             // GC次数
	WorkDir            string `protobuf:"bytes,10,opt,name=WorkDir,proto3" json:"WorkDir,omitempty"`
	ActiveCount        uint64 `protobuf:"varint,11,opt,name=ActiveCount,proto3" json:"ActiveCount,omitempty"`               // 当前活跃任务数
	CompletedTaskCount uint64 `protobuf:"varint,12,opt,name=CompletedTaskCount,proto3" json:"CompletedTaskCount,omitempty"` // 当前已完成任务数
	WaitCount          uint64 `protobuf:"varint,13,opt,name=WaitCount,proto3" json:"WaitCount,omitempty"`                   // 当前等待任务数
	MaxPoolSize        uint64 `protobuf:"varint,14,opt,name=MaxPoolSize,proto3" json:"MaxPoolSize,omitempty"`               // 最大工作线程
	Status             Status `protobuf:"varint,6,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`          // 节点状态
	Time               int64  `protobuf:"varint,20,opt,name=Time,proto3" json:"Time,omitempty"`                             // 当前系统时间
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Pong) GetCpu() uint64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Pong) GetMemoryAlloc() uint64 {
	if x != nil {
		return x.MemoryAlloc
	}
	return 0
}

func (x *Pong) GetTotalAlloc() uint64 {
	if x != nil {
		return x.TotalAlloc
	}
	return 0
}

func (x *Pong) GetSys() uint64 {
	if x != nil {
		return x.Sys
	}
	return 0
}

func (x *Pong) GetNumGC() uint32 {
	if x != nil {
		return x.NumGC
	}
	return 0
}

func (x *Pong) GetWorkDir() string {
	if x != nil {
		return x.WorkDir
	}
	return ""
}

func (x *Pong) GetActiveCount() uint64 {
	if x != nil {
		return x.ActiveCount
	}
	return 0
}

func (x *Pong) GetCompletedTaskCount() uint64 {
	if x != nil {
		return x.CompletedTaskCount
	}
	return 0
}

func (x *Pong) GetWaitCount() uint64 {
	if x != nil {
		return x.WaitCount
	}
	return 0
}

func (x *Pong) GetMaxPoolSize() uint64 {
	if x != nil {
		return x.MaxPoolSize
	}
	return 0
}

func (x *Pong) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Looking
}

func (x *Pong) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip     string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port   string `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	Weight uint64 `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Status Status `protobuf:"varint,5,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"` // 节点状态
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *Node) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Node) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Looking
}

type Distro struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxPoolSize uint64 `protobuf:"varint,1,opt,name=MaxPoolSize,proto3" json:"MaxPoolSize,omitempty"` // 最大工作线程
}

func (x *Distro) Reset() {
	*x = Distro{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distro) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distro) ProtoMessage() {}

func (x *Distro) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distro.ProtoReflect.Descriptor instead.
func (*Distro) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *Distro) GetMaxPoolSize() uint64 {
	if x != nil {
		return x.MaxPoolSize
	}
	return 0
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterAddr string   `protobuf:"bytes,1,opt,name=MasterAddr,proto3" json:"MasterAddr,omitempty"`
	Addr       []string `protobuf:"bytes,2,rep,name=Addr,proto3" json:"Addr,omitempty"`
	ClientAddr []string `protobuf:"bytes,3,rep,name=clientAddr,proto3" json:"clientAddr,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *Cluster) GetMasterAddr() string {
	if x != nil {
		return x.MasterAddr
	}
	return ""
}

func (x *Cluster) GetAddr() []string {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *Cluster) GetClientAddr() []string {
	if x != nil {
		return x.ClientAddr
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0xcb, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x43, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x5c,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x70,
	0x75, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x43, 0x70, 0x75,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x15, 0x0a, 0x03,
	0x41, 0x43, 0x4b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0xe7, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x43, 0x70, 0x75, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x53, 0x79, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x4e, 0x75, 0x6d, 0x47, 0x43, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4e,
	0x75, 0x6d, 0x47, 0x43, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x69, 0x72, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x57, 0x61, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x77, 0x0a, 0x04, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2a, 0x0a, 0x06, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5d,
	0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x2a, 0x33, 0x0a,
	0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x43, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56, 0x41, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x59,
	0x54, 0x48, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4c, 0x41, 0x4e, 0x47,
	0x10, 0x03, 0x2a, 0x40, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07,
	0x4c, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x10, 0x03, 0x32, 0xc1, 0x02, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x05, 0x48, 0x65, 0x61, 0x72, 0x74, 0x12, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x09, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x07, 0x56, 0x69, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x22, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x0c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x43, 0x4b, 0x12, 0x20, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12,
	0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x1a, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x08, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x43, 0x4b, 0x12, 0x1f, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x08,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x43, 0x4b, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(Language)(0),   // 0: api.Language
	(Status)(0),     // 1: api.Status
	(*Task)(nil),    // 2: api.Task
	(*Result)(nil),  // 3: api.Result
	(*ACK)(nil),     // 4: api.ACK
	(*Ping)(nil),    // 5: api.Ping
	(*Pong)(nil),    // 6: api.Pong
	(*Node)(nil),    // 7: api.Node
	(*Distro)(nil),  // 8: api.Distro
	(*Cluster)(nil), // 9: api.Cluster
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.Task.type:type_name -> api.Language
	1,  // 1: api.Pong.status:type_name -> api.Status
	1,  // 2: api.Node.status:type_name -> api.Status
	5,  // 3: api.DistroServer.Heart:input_type -> api.Ping
	7,  // 4: api.DistroServer.Join:input_type -> api.Node
	5,  // 5: api.DistroServer.Election:input_type -> api.Ping
	5,  // 6: api.DistroServer.Candidate:input_type -> api.Ping
	7,  // 7: api.DistroServer.Victory:input_type -> api.Node
	9,  // 8: api.DistroServer.Checkout:input_type -> api.Cluster
	8,  // 9: api.DistroServer.Modify:input_type -> api.Distro
	2,  // 10: api.DistroServer.Execute:input_type -> api.Task
	3,  // 11: api.DistroServer.Caller:input_type -> api.Result
	6,  // 12: api.DistroServer.Heart:output_type -> api.Pong
	9,  // 13: api.DistroServer.Join:output_type -> api.Cluster
	9,  // 14: api.DistroServer.Election:output_type -> api.Cluster
	9,  // 15: api.DistroServer.Candidate:output_type -> api.Cluster
	7,  // 16: api.DistroServer.Victory:output_type -> api.Node
	4,  // 17: api.DistroServer.Checkout:output_type -> api.ACK
	6,  // 18: api.DistroServer.Modify:output_type -> api.Pong
	4,  // 19: api.DistroServer.Execute:output_type -> api.ACK
	4,  // 20: api.DistroServer.Caller:output_type -> api.ACK
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACK); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distro); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
