// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.0
// source: protos/data.proto

package user

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

type PodMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName   string `protobuf:"bytes,2,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	Ready     string `protobuf:"bytes,3,opt,name=ready,proto3" json:"ready,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Restart   string `protobuf:"bytes,5,opt,name=restart,proto3" json:"restart,omitempty"`
	Age       string `protobuf:"bytes,6,opt,name=age,proto3" json:"age,omitempty"`
	Ip        string `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	NodeName  string `protobuf:"bytes,8,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *PodMessage) Reset() {
	*x = PodMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodMessage) ProtoMessage() {}

func (x *PodMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodMessage.ProtoReflect.Descriptor instead.
func (*PodMessage) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{0}
}

func (x *PodMessage) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodMessage) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *PodMessage) GetReady() string {
	if x != nil {
		return x.Ready
	}
	return ""
}

func (x *PodMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PodMessage) GetRestart() string {
	if x != nil {
		return x.Restart
	}
	return ""
}

func (x *PodMessage) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *PodMessage) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PodMessage) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type GetPodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameSpace string `protobuf:"bytes,1,opt,name=name_space,json=nameSpace,proto3" json:"name_space,omitempty"`
}

func (x *GetPodRequest) Reset() {
	*x = GetPodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodRequest) ProtoMessage() {}

func (x *GetPodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodRequest.ProtoReflect.Descriptor instead.
func (*GetPodRequest) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetPodRequest) GetNameSpace() string {
	if x != nil {
		return x.NameSpace
	}
	return ""
}

type GetPodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodMessage *PodMessage `protobuf:"bytes,1,opt,name=pod_message,json=podMessage,proto3" json:"pod_message,omitempty"`
}

func (x *GetPodResponse) Reset() {
	*x = GetPodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodResponse) ProtoMessage() {}

func (x *GetPodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodResponse.ProtoReflect.Descriptor instead.
func (*GetPodResponse) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetPodResponse) GetPodMessage() *PodMessage {
	if x != nil {
		return x.PodMessage
	}
	return nil
}

type PostPodMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yamldata string `protobuf:"bytes,1,opt,name=yamldata,proto3" json:"yamldata,omitempty"`
}

func (x *PostPodMessage) Reset() {
	*x = PostPodMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPodMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPodMessage) ProtoMessage() {}

func (x *PostPodMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPodMessage.ProtoReflect.Descriptor instead.
func (*PostPodMessage) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{3}
}

func (x *PostPodMessage) GetYamldata() string {
	if x != nil {
		return x.Yamldata
	}
	return ""
}

type PostPodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YamlData string `protobuf:"bytes,1,opt,name=yaml_data,json=yamlData,proto3" json:"yaml_data,omitempty"`
}

func (x *PostPodRequest) Reset() {
	*x = PostPodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPodRequest) ProtoMessage() {}

func (x *PostPodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPodRequest.ProtoReflect.Descriptor instead.
func (*PostPodRequest) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{4}
}

func (x *PostPodRequest) GetYamlData() string {
	if x != nil {
		return x.YamlData
	}
	return ""
}

type PostPodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostPodMessage *PostPodMessage `protobuf:"bytes,1,opt,name=post_pod_message,json=postPodMessage,proto3" json:"post_pod_message,omitempty"`
}

func (x *PostPodResponse) Reset() {
	*x = PostPodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPodResponse) ProtoMessage() {}

func (x *PostPodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPodResponse.ProtoReflect.Descriptor instead.
func (*PostPodResponse) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{5}
}

func (x *PostPodResponse) GetPostPodMessage() *PostPodMessage {
	if x != nil {
		return x.PostPodMessage
	}
	return nil
}

type JobMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account        string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	NodeName       string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	JobState       string `protobuf:"bytes,3,opt,name=job_state,json=jobState,proto3" json:"job_state,omitempty"`
	JobName        string `protobuf:"bytes,4,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	StartTime      string `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	StandardOutput string `protobuf:"bytes,6,opt,name=standard_output,json=standardOutput,proto3" json:"standard_output,omitempty"`
}

func (x *JobMessage) Reset() {
	*x = JobMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobMessage) ProtoMessage() {}

func (x *JobMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobMessage.ProtoReflect.Descriptor instead.
func (*JobMessage) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{6}
}

func (x *JobMessage) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *JobMessage) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *JobMessage) GetJobState() string {
	if x != nil {
		return x.JobState
	}
	return ""
}

func (x *JobMessage) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

func (x *JobMessage) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *JobMessage) GetStandardOutput() string {
	if x != nil {
		return x.StandardOutput
	}
	return ""
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{7}
}

func (x *GetJobRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type GetJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetJobMessage *JobMessage `protobuf:"bytes,1,opt,name=get_job_message,json=getJobMessage,proto3" json:"get_job_message,omitempty"`
}

func (x *GetJobResponse) Reset() {
	*x = GetJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobResponse) ProtoMessage() {}

func (x *GetJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobResponse.ProtoReflect.Descriptor instead.
func (*GetJobResponse) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{8}
}

func (x *GetJobResponse) GetGetJobMessage() *JobMessage {
	if x != nil {
		return x.GetJobMessage
	}
	return nil
}

type WorkListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Runtime   string `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ready     string `protobuf:"bytes,4,opt,name=ready,proto3" json:"ready,omitempty"`
	Status    string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Restart   string `protobuf:"bytes,6,opt,name=restart,proto3" json:"restart,omitempty"`
	Age       string `protobuf:"bytes,7,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *WorkListMessage) Reset() {
	*x = WorkListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkListMessage) ProtoMessage() {}

func (x *WorkListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkListMessage.ProtoReflect.Descriptor instead.
func (*WorkListMessage) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{9}
}

func (x *WorkListMessage) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *WorkListMessage) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkListMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkListMessage) GetReady() string {
	if x != nil {
		return x.Ready
	}
	return ""
}

func (x *WorkListMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WorkListMessage) GetRestart() string {
	if x != nil {
		return x.Restart
	}
	return ""
}

func (x *WorkListMessage) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

type GetWorkListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetWorkListRequest) Reset() {
	*x = GetWorkListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkListRequest) ProtoMessage() {}

func (x *GetWorkListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkListRequest.ProtoReflect.Descriptor instead.
func (*GetWorkListRequest) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{10}
}

func (x *GetWorkListRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetWorkListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetWorklistMessage *WorkListMessage `protobuf:"bytes,1,opt,name=get_worklist_message,json=getWorklistMessage,proto3" json:"get_worklist_message,omitempty"`
}

func (x *GetWorkListResponse) Reset() {
	*x = GetWorkListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkListResponse) ProtoMessage() {}

func (x *GetWorkListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkListResponse.ProtoReflect.Descriptor instead.
func (*GetWorkListResponse) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{11}
}

func (x *GetWorkListResponse) GetGetWorklistMessage() *WorkListMessage {
	if x != nil {
		return x.GetWorklistMessage
	}
	return nil
}

type DeleteWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yamldata  string `protobuf:"bytes,1,opt,name=yamldata,proto3" json:"yamldata,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteWorkRequest) Reset() {
	*x = DeleteWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkRequest) ProtoMessage() {}

func (x *DeleteWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkRequest) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteWorkRequest) GetYamldata() string {
	if x != nil {
		return x.Yamldata
	}
	return ""
}

func (x *DeleteWorkRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteWorkRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteWorkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deletework string `protobuf:"bytes,1,opt,name=deletework,proto3" json:"deletework,omitempty"`
}

func (x *DeleteWorkResponse) Reset() {
	*x = DeleteWorkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_data_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkResponse) ProtoMessage() {}

func (x *DeleteWorkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_data_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkResponse) Descriptor() ([]byte, []int) {
	return file_protos_data_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteWorkResponse) GetDeletework() string {
	if x != nil {
		return x.Deletework
	}
	return ""
}

var File_protos_data_proto protoreflect.FileDescriptor

var file_protos_data_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0xcc, 0x01, 0x0a,
	0x0a, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x0b, 0x70, 0x6f, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x2d, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x79, 0x61, 0x6d, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x54, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x2c, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4a, 0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x67, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0f, 0x57, 0x6f,
	0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x6f,
	0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x12, 0x67, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x61, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x79, 0x61, 0x6d, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x32, 0xcb, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x50,
	0x6f, 0x64, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76,
	0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x15, 0x5a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_data_proto_rawDescOnce sync.Once
	file_protos_data_proto_rawDescData = file_protos_data_proto_rawDesc
)

func file_protos_data_proto_rawDescGZIP() []byte {
	file_protos_data_proto_rawDescOnce.Do(func() {
		file_protos_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_data_proto_rawDescData)
	})
	return file_protos_data_proto_rawDescData
}

var file_protos_data_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_protos_data_proto_goTypes = []interface{}{
	(*PodMessage)(nil),          // 0: v1.user.PodMessage
	(*GetPodRequest)(nil),       // 1: v1.user.GetPodRequest
	(*GetPodResponse)(nil),      // 2: v1.user.GetPodResponse
	(*PostPodMessage)(nil),      // 3: v1.user.PostPodMessage
	(*PostPodRequest)(nil),      // 4: v1.user.PostPodRequest
	(*PostPodResponse)(nil),     // 5: v1.user.PostPodResponse
	(*JobMessage)(nil),          // 6: v1.user.JobMessage
	(*GetJobRequest)(nil),       // 7: v1.user.GetJobRequest
	(*GetJobResponse)(nil),      // 8: v1.user.GetJobResponse
	(*WorkListMessage)(nil),     // 9: v1.user.WorkListMessage
	(*GetWorkListRequest)(nil),  // 10: v1.user.GetWorkListRequest
	(*GetWorkListResponse)(nil), // 11: v1.user.GetWorkListResponse
	(*DeleteWorkRequest)(nil),   // 12: v1.user.DeleteWorkRequest
	(*DeleteWorkResponse)(nil),  // 13: v1.user.DeleteWorkResponse
}
var file_protos_data_proto_depIdxs = []int32{
	0,  // 0: v1.user.GetPodResponse.pod_message:type_name -> v1.user.PodMessage
	3,  // 1: v1.user.PostPodResponse.post_pod_message:type_name -> v1.user.PostPodMessage
	6,  // 2: v1.user.GetJobResponse.get_job_message:type_name -> v1.user.JobMessage
	9,  // 3: v1.user.GetWorkListResponse.get_worklist_message:type_name -> v1.user.WorkListMessage
	1,  // 4: v1.user.User.GetPod:input_type -> v1.user.GetPodRequest
	4,  // 5: v1.user.User.PostPod:input_type -> v1.user.PostPodRequest
	7,  // 6: v1.user.User.GetJob:input_type -> v1.user.GetJobRequest
	10, // 7: v1.user.User.GetWorkList:input_type -> v1.user.GetWorkListRequest
	12, // 8: v1.user.User.DeleteWork:input_type -> v1.user.DeleteWorkRequest
	2,  // 9: v1.user.User.GetPod:output_type -> v1.user.GetPodResponse
	5,  // 10: v1.user.User.PostPod:output_type -> v1.user.PostPodResponse
	8,  // 11: v1.user.User.GetJob:output_type -> v1.user.GetJobResponse
	11, // 12: v1.user.User.GetWorkList:output_type -> v1.user.GetWorkListResponse
	13, // 13: v1.user.User.DeleteWork:output_type -> v1.user.DeleteWorkResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protos_data_proto_init() }
func file_protos_data_proto_init() {
	if File_protos_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodMessage); i {
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
		file_protos_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodRequest); i {
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
		file_protos_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodResponse); i {
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
		file_protos_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPodMessage); i {
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
		file_protos_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPodRequest); i {
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
		file_protos_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPodResponse); i {
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
		file_protos_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobMessage); i {
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
		file_protos_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
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
		file_protos_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobResponse); i {
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
		file_protos_data_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkListMessage); i {
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
		file_protos_data_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkListRequest); i {
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
		file_protos_data_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkListResponse); i {
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
		file_protos_data_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkRequest); i {
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
		file_protos_data_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkResponse); i {
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
			RawDescriptor: file_protos_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_data_proto_goTypes,
		DependencyIndexes: file_protos_data_proto_depIdxs,
		MessageInfos:      file_protos_data_proto_msgTypes,
	}.Build()
	File_protos_data_proto = out.File
	file_protos_data_proto_rawDesc = nil
	file_protos_data_proto_goTypes = nil
	file_protos_data_proto_depIdxs = nil
}
