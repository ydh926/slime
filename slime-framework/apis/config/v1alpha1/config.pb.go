// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apis/config/v1alpha1/config.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Limiter_RateLimitBackend int32

const (
	Limiter_netEaseLocalFlowControl Limiter_RateLimitBackend = 0
	Limiter_envoyLocalRateLimit     Limiter_RateLimitBackend = 1
)

var Limiter_RateLimitBackend_name = map[int32]string{
	0: "netEaseLocalFlowControl",
	1: "envoyLocalRateLimit",
}

var Limiter_RateLimitBackend_value = map[string]int32{
	"netEaseLocalFlowControl": 0,
	"envoyLocalRateLimit":     1,
}

func (x Limiter_RateLimitBackend) String() string {
	return proto.EnumName(Limiter_RateLimitBackend_name, int32(x))
}

func (Limiter_RateLimitBackend) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{3, 0}
}

type Prometheus_Source_Type int32

const (
	Prometheus_Source_Value Prometheus_Source_Type = 0
	Prometheus_Source_Group Prometheus_Source_Type = 1
)

var Prometheus_Source_Type_name = map[int32]string{
	0: "Value",
	1: "Group",
}

var Prometheus_Source_Type_value = map[string]int32{
	"Value": 0,
	"Group": 1,
}

func (x Prometheus_Source_Type) String() string {
	return proto.EnumName(Prometheus_Source_Type_name, int32(x))
}

func (Prometheus_Source_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{6, 0}
}

type LocalSource struct {
	Mount                string   `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalSource) Reset()         { *m = LocalSource{} }
func (m *LocalSource) String() string { return proto.CompactTextString(m) }
func (*LocalSource) ProtoMessage()    {}
func (*LocalSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{0}
}
func (m *LocalSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalSource.Unmarshal(m, b)
}
func (m *LocalSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalSource.Marshal(b, m, deterministic)
}
func (m *LocalSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalSource.Merge(m, src)
}
func (m *LocalSource) XXX_Size() int {
	return xxx_messageInfo_LocalSource.Size(m)
}
func (m *LocalSource) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalSource.DiscardUnknown(m)
}

var xxx_messageInfo_LocalSource proto.InternalMessageInfo

func (m *LocalSource) GetMount() string {
	if m != nil {
		return m.Mount
	}
	return ""
}

type RemoteSource struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteSource) Reset()         { *m = RemoteSource{} }
func (m *RemoteSource) String() string { return proto.CompactTextString(m) }
func (*RemoteSource) ProtoMessage()    {}
func (*RemoteSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{1}
}
func (m *RemoteSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSource.Unmarshal(m, b)
}
func (m *RemoteSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSource.Marshal(b, m, deterministic)
}
func (m *RemoteSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSource.Merge(m, src)
}
func (m *RemoteSource) XXX_Size() int {
	return xxx_messageInfo_RemoteSource.Size(m)
}
func (m *RemoteSource) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSource.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSource proto.InternalMessageInfo

func (m *RemoteSource) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Plugin struct {
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// Types that are valid to be assigned to WasmSource:
	//	*Plugin_Local
	//	*Plugin_Remote
	WasmSource           isPlugin_WasmSource `protobuf_oneof:"wasm_source"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{2}
}
func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plugin.Unmarshal(m, b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return xxx_messageInfo_Plugin.Size(m)
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

type isPlugin_WasmSource interface {
	isPlugin_WasmSource()
}

type Plugin_Local struct {
	Local *LocalSource `protobuf:"bytes,2,opt,name=local,proto3,oneof"`
}
type Plugin_Remote struct {
	Remote *RemoteSource `protobuf:"bytes,3,opt,name=remote,proto3,oneof"`
}

func (*Plugin_Local) isPlugin_WasmSource()  {}
func (*Plugin_Remote) isPlugin_WasmSource() {}

func (m *Plugin) GetWasmSource() isPlugin_WasmSource {
	if m != nil {
		return m.WasmSource
	}
	return nil
}

func (m *Plugin) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Plugin) GetLocal() *LocalSource {
	if x, ok := m.GetWasmSource().(*Plugin_Local); ok {
		return x.Local
	}
	return nil
}

func (m *Plugin) GetRemote() *RemoteSource {
	if x, ok := m.GetWasmSource().(*Plugin_Remote); ok {
		return x.Remote
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Plugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Plugin_OneofMarshaler, _Plugin_OneofUnmarshaler, _Plugin_OneofSizer, []interface{}{
		(*Plugin_Local)(nil),
		(*Plugin_Remote)(nil),
	}
}

func _Plugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Plugin)
	// wasm_source
	switch x := m.WasmSource.(type) {
	case *Plugin_Local:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Local); err != nil {
			return err
		}
	case *Plugin_Remote:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Remote); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Plugin.WasmSource has unexpected type %T", x)
	}
	return nil
}

func _Plugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Plugin)
	switch tag {
	case 2: // wasm_source.local
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LocalSource)
		err := b.DecodeMessage(msg)
		m.WasmSource = &Plugin_Local{msg}
		return true, err
	case 3: // wasm_source.remote
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RemoteSource)
		err := b.DecodeMessage(msg)
		m.WasmSource = &Plugin_Remote{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Plugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Plugin)
	// wasm_source
	switch x := m.WasmSource.(type) {
	case *Plugin_Local:
		s := proto.Size(x.Local)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Plugin_Remote:
		s := proto.Size(x.Remote)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Limiter struct {
	Enable               bool                     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Multicluster         bool                     `protobuf:"varint,2,opt,name=Multicluster,proto3" json:"Multicluster,omitempty"`
	Backend              Limiter_RateLimitBackend `protobuf:"varint,3,opt,name=backend,proto3,enum=slime.config.v1alpha1.Limiter_RateLimitBackend" json:"backend,omitempty"`
	Refresh              *time.Duration           `protobuf:"bytes,4,opt,name=refresh,proto3,stdduration" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Limiter) Reset()         { *m = Limiter{} }
func (m *Limiter) String() string { return proto.CompactTextString(m) }
func (*Limiter) ProtoMessage()    {}
func (*Limiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{3}
}
func (m *Limiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Limiter.Unmarshal(m, b)
}
func (m *Limiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Limiter.Marshal(b, m, deterministic)
}
func (m *Limiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Limiter.Merge(m, src)
}
func (m *Limiter) XXX_Size() int {
	return xxx_messageInfo_Limiter.Size(m)
}
func (m *Limiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Limiter.DiscardUnknown(m)
}

var xxx_messageInfo_Limiter proto.InternalMessageInfo

func (m *Limiter) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Limiter) GetMulticluster() bool {
	if m != nil {
		return m.Multicluster
	}
	return false
}

func (m *Limiter) GetBackend() Limiter_RateLimitBackend {
	if m != nil {
		return m.Backend
	}
	return Limiter_netEaseLocalFlowControl
}

func (m *Limiter) GetRefresh() *time.Duration {
	if m != nil {
		return m.Refresh
	}
	return nil
}

type Global struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Multicluster         string   `protobuf:"bytes,2,opt,name=multicluster,proto3" json:"multicluster,omitempty"`
	IstioNamespace       string   `protobuf:"bytes,3,opt,name=istioNamespace,proto3" json:"istioNamespace,omitempty"`
	SlimeNamespace       string   `protobuf:"bytes,4,opt,name=slimeNamespace,proto3" json:"slimeNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{4}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (m *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(m, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Global) GetMulticluster() string {
	if m != nil {
		return m.Multicluster
	}
	return ""
}

func (m *Global) GetIstioNamespace() string {
	if m != nil {
		return m.IstioNamespace
	}
	return ""
}

func (m *Global) GetSlimeNamespace() string {
	if m != nil {
		return m.SlimeNamespace
	}
	return ""
}

type Fence struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	WormholePort         []string `protobuf:"bytes,2,rep,name=wormholePort,proto3" json:"wormholePort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fence) Reset()         { *m = Fence{} }
func (m *Fence) String() string { return proto.CompactTextString(m) }
func (*Fence) ProtoMessage()    {}
func (*Fence) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{5}
}
func (m *Fence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fence.Unmarshal(m, b)
}
func (m *Fence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fence.Marshal(b, m, deterministic)
}
func (m *Fence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fence.Merge(m, src)
}
func (m *Fence) XXX_Size() int {
	return xxx_messageInfo_Fence.Size(m)
}
func (m *Fence) XXX_DiscardUnknown() {
	xxx_messageInfo_Fence.DiscardUnknown(m)
}

var xxx_messageInfo_Fence proto.InternalMessageInfo

func (m *Fence) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Fence) GetWormholePort() []string {
	if m != nil {
		return m.WormholePort
	}
	return nil
}

type Prometheus_Source struct {
	Address              string                                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Handlers             map[string]*Prometheus_Source_Handler `protobuf:"bytes,2,rep,name=handlers,proto3" json:"handlers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Prometheus_Source) Reset()         { *m = Prometheus_Source{} }
func (m *Prometheus_Source) String() string { return proto.CompactTextString(m) }
func (*Prometheus_Source) ProtoMessage()    {}
func (*Prometheus_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{6}
}
func (m *Prometheus_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prometheus_Source.Unmarshal(m, b)
}
func (m *Prometheus_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prometheus_Source.Marshal(b, m, deterministic)
}
func (m *Prometheus_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prometheus_Source.Merge(m, src)
}
func (m *Prometheus_Source) XXX_Size() int {
	return xxx_messageInfo_Prometheus_Source.Size(m)
}
func (m *Prometheus_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Prometheus_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Prometheus_Source proto.InternalMessageInfo

func (m *Prometheus_Source) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Prometheus_Source) GetHandlers() map[string]*Prometheus_Source_Handler {
	if m != nil {
		return m.Handlers
	}
	return nil
}

type Prometheus_Source_Handler struct {
	Query                string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Type                 Prometheus_Source_Type `protobuf:"varint,2,opt,name=type,proto3,enum=slime.config.v1alpha1.Prometheus_Source_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Prometheus_Source_Handler) Reset()         { *m = Prometheus_Source_Handler{} }
func (m *Prometheus_Source_Handler) String() string { return proto.CompactTextString(m) }
func (*Prometheus_Source_Handler) ProtoMessage()    {}
func (*Prometheus_Source_Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{6, 0}
}
func (m *Prometheus_Source_Handler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prometheus_Source_Handler.Unmarshal(m, b)
}
func (m *Prometheus_Source_Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prometheus_Source_Handler.Marshal(b, m, deterministic)
}
func (m *Prometheus_Source_Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prometheus_Source_Handler.Merge(m, src)
}
func (m *Prometheus_Source_Handler) XXX_Size() int {
	return xxx_messageInfo_Prometheus_Source_Handler.Size(m)
}
func (m *Prometheus_Source_Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Prometheus_Source_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Prometheus_Source_Handler proto.InternalMessageInfo

func (m *Prometheus_Source_Handler) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Prometheus_Source_Handler) GetType() Prometheus_Source_Type {
	if m != nil {
		return m.Type
	}
	return Prometheus_Source_Value
}

type K8S_Source struct {
	Handlers             []string `protobuf:"bytes,1,rep,name=handlers,proto3" json:"handlers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K8S_Source) Reset()         { *m = K8S_Source{} }
func (m *K8S_Source) String() string { return proto.CompactTextString(m) }
func (*K8S_Source) ProtoMessage()    {}
func (*K8S_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{7}
}
func (m *K8S_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8S_Source.Unmarshal(m, b)
}
func (m *K8S_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8S_Source.Marshal(b, m, deterministic)
}
func (m *K8S_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8S_Source.Merge(m, src)
}
func (m *K8S_Source) XXX_Size() int {
	return xxx_messageInfo_K8S_Source.Size(m)
}
func (m *K8S_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_K8S_Source.DiscardUnknown(m)
}

var xxx_messageInfo_K8S_Source proto.InternalMessageInfo

func (m *K8S_Source) GetHandlers() []string {
	if m != nil {
		return m.Handlers
	}
	return nil
}

type Metric struct {
	Prometheus           *Prometheus_Source `protobuf:"bytes,1,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	K8S                  *K8S_Source        `protobuf:"bytes,2,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{8}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetPrometheus() *Prometheus_Source {
	if m != nil {
		return m.Prometheus
	}
	return nil
}

func (m *Metric) GetK8S() *K8S_Source {
	if m != nil {
		return m.K8S
	}
	return nil
}

type Config struct {
	Plugin               *Plugin  `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Limiter              *Limiter `protobuf:"bytes,2,opt,name=limiter,proto3" json:"limiter,omitempty"`
	Global               *Global  `protobuf:"bytes,3,opt,name=global,proto3" json:"global,omitempty"`
	Fence                *Fence   `protobuf:"bytes,4,opt,name=fence,proto3" json:"fence,omitempty"`
	Metric               *Metric  `protobuf:"bytes,6,opt,name=metric,proto3" json:"metric,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_0545491500b2fa54, []int{9}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetPlugin() *Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (m *Config) GetLimiter() *Limiter {
	if m != nil {
		return m.Limiter
	}
	return nil
}

func (m *Config) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *Config) GetFence() *Fence {
	if m != nil {
		return m.Fence
	}
	return nil
}

func (m *Config) GetMetric() *Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("slime.config.v1alpha1.Limiter_RateLimitBackend", Limiter_RateLimitBackend_name, Limiter_RateLimitBackend_value)
	proto.RegisterEnum("slime.config.v1alpha1.Prometheus_Source_Type", Prometheus_Source_Type_name, Prometheus_Source_Type_value)
	proto.RegisterType((*LocalSource)(nil), "slime.config.v1alpha1.LocalSource")
	proto.RegisterType((*RemoteSource)(nil), "slime.config.v1alpha1.RemoteSource")
	proto.RegisterType((*Plugin)(nil), "slime.config.v1alpha1.Plugin")
	proto.RegisterType((*Limiter)(nil), "slime.config.v1alpha1.Limiter")
	proto.RegisterType((*Global)(nil), "slime.config.v1alpha1.Global")
	proto.RegisterType((*Fence)(nil), "slime.config.v1alpha1.Fence")
	proto.RegisterType((*Prometheus_Source)(nil), "slime.config.v1alpha1.Prometheus_Source")
	proto.RegisterMapType((map[string]*Prometheus_Source_Handler)(nil), "slime.config.v1alpha1.Prometheus_Source.HandlersEntry")
	proto.RegisterType((*Prometheus_Source_Handler)(nil), "slime.config.v1alpha1.Prometheus_Source.Handler")
	proto.RegisterType((*K8S_Source)(nil), "slime.config.v1alpha1.K8S_Source")
	proto.RegisterType((*Metric)(nil), "slime.config.v1alpha1.Metric")
	proto.RegisterType((*Config)(nil), "slime.config.v1alpha1.Config")
}

func init() {
	proto.RegisterFile("pkg/apis/config/v1alpha1/config.proto", fileDescriptor_0545491500b2fa54)
}

var fileDescriptor_0545491500b2fa54 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x8e, 0xf3, 0xe3, 0x6c, 0x4e, 0xda, 0x55, 0x18, 0x7e, 0x6a, 0x42, 0x29, 0x8b, 0x2b, 0x20,
	0x5c, 0xd4, 0xa6, 0xa9, 0x40, 0xa1, 0x12, 0x17, 0x64, 0xe9, 0x36, 0x88, 0x16, 0xad, 0xa6, 0x88,
	0x0b, 0x6e, 0x56, 0x13, 0xe7, 0xc4, 0xb1, 0x32, 0xf6, 0x98, 0xf1, 0x38, 0xab, 0x3c, 0x01, 0x4f,
	0x00, 0xe2, 0x19, 0xb8, 0xe1, 0x91, 0x78, 0x07, 0x9e, 0x00, 0x79, 0x3c, 0xce, 0x26, 0x0b, 0x5e,
	0xb6, 0x77, 0x73, 0xce, 0x7c, 0xdf, 0x7c, 0xdf, 0x9c, 0xe3, 0x33, 0x86, 0x8f, 0xd2, 0x75, 0xe8,
	0xb3, 0x34, 0xca, 0xfc, 0x40, 0x24, 0xcb, 0x28, 0xf4, 0x37, 0x8f, 0x19, 0x4f, 0x57, 0xec, 0xb1,
	0x89, 0xbd, 0x54, 0x0a, 0x25, 0xc8, 0xdb, 0x19, 0x8f, 0x62, 0xf4, 0x4c, 0xae, 0xc2, 0x0c, 0x1f,
	0x85, 0x91, 0x5a, 0xe5, 0x73, 0x2f, 0x10, 0xb1, 0x1f, 0x8a, 0x50, 0xf8, 0x1a, 0x3d, 0xcf, 0x97,
	0x3a, 0xd2, 0x81, 0x5e, 0x95, 0xa7, 0x0c, 0x1f, 0x84, 0x42, 0x84, 0x1c, 0xaf, 0x50, 0x8b, 0x5c,
	0x32, 0x15, 0x89, 0xa4, 0xdc, 0x77, 0x1f, 0x42, 0xff, 0x85, 0x08, 0x18, 0x7f, 0x25, 0x72, 0x19,
	0x20, 0x79, 0x0b, 0x3a, 0xb1, 0xc8, 0x13, 0xe5, 0x58, 0x27, 0xd6, 0xa8, 0x47, 0xcb, 0xc0, 0x1d,
	0xc1, 0x1d, 0x8a, 0xb1, 0x50, 0x68, 0x50, 0x0e, 0x74, 0xd9, 0x62, 0x21, 0x31, 0xcb, 0x0c, 0xae,
	0x0a, 0xdd, 0x3f, 0x2c, 0xb0, 0xcf, 0x79, 0x1e, 0x46, 0x09, 0x79, 0x07, 0x6c, 0x4c, 0xd8, 0x9c,
	0xa3, 0xc6, 0x1c, 0x51, 0x13, 0x91, 0xa7, 0xd0, 0xe1, 0x85, 0xa2, 0xd3, 0x3c, 0xb1, 0x46, 0xfd,
	0xb1, 0xeb, 0xfd, 0xe7, 0x3d, 0xbd, 0x3d, 0x57, 0xb3, 0x06, 0x2d, 0x29, 0xe4, 0x2b, 0xb0, 0xa5,
	0x36, 0xe2, 0xb4, 0x34, 0xf9, 0x61, 0x0d, 0x79, 0xdf, 0xed, 0xac, 0x41, 0x0d, 0x69, 0x7a, 0x17,
	0xfa, 0x97, 0x2c, 0x8b, 0x2f, 0x32, 0xbd, 0xe1, 0xfe, 0xda, 0x84, 0xee, 0x8b, 0x28, 0x8e, 0x14,
	0xca, 0x5a, 0xb7, 0x2e, 0xdc, 0x79, 0x99, 0x73, 0x15, 0x05, 0x3c, 0xcf, 0x14, 0x4a, 0x6d, 0xfa,
	0x88, 0x1e, 0xe4, 0xc8, 0xb7, 0xd0, 0x9d, 0xb3, 0x60, 0x8d, 0xc9, 0x42, 0xdb, 0x3a, 0x1e, 0xfb,
	0x75, 0x77, 0x2a, 0xc5, 0x3c, 0xca, 0x14, 0xea, 0xf5, 0xb4, 0xa4, 0xd1, 0x8a, 0x4f, 0xbe, 0x84,
	0xae, 0xc4, 0xa5, 0xc4, 0x6c, 0xe5, 0xb4, 0xf5, 0x0d, 0xdf, 0xf5, 0xca, 0x06, 0x7a, 0x55, 0x03,
	0xbd, 0x6f, 0x4c, 0x03, 0xa7, 0xed, 0xdf, 0xff, 0xfa, 0xc0, 0xa2, 0x15, 0xde, 0x9d, 0xc1, 0xe0,
	0xfa, 0xb9, 0xe4, 0x3d, 0xb8, 0x97, 0xa0, 0x7a, 0xc6, 0x32, 0xd4, 0xe5, 0x3c, 0xe3, 0xe2, 0xf2,
	0x54, 0x24, 0x4a, 0x0a, 0x3e, 0x68, 0x90, 0x7b, 0xf0, 0x26, 0x26, 0x1b, 0xb1, 0xd5, 0x5b, 0x3b,
	0xea, 0xc0, 0x72, 0x7f, 0xb3, 0xc0, 0x7e, 0xce, 0xc5, 0x9c, 0xf1, 0xa2, 0xd3, 0x19, 0xca, 0x4d,
	0x14, 0x60, 0xd5, 0x69, 0x13, 0x16, 0x85, 0x89, 0xaf, 0x17, 0xa6, 0x47, 0x0f, 0x72, 0xe4, 0x63,
	0x38, 0x8e, 0x32, 0x15, 0x89, 0xef, 0x59, 0x8c, 0x59, 0xca, 0x82, 0xb2, 0x6d, 0x3d, 0x7a, 0x2d,
	0x5b, 0xe0, 0x74, 0xc1, 0xae, 0x70, 0xed, 0x12, 0x77, 0x98, 0x75, 0x4f, 0xa1, 0x73, 0x86, 0x49,
	0x80, 0x37, 0x75, 0xeb, 0x52, 0xc8, 0x78, 0x25, 0x38, 0x9e, 0x0b, 0xa9, 0x9c, 0xe6, 0x49, 0xab,
	0x30, 0xb5, 0x9f, 0x73, 0xff, 0x6e, 0xc2, 0x1b, 0xe7, 0x52, 0xc4, 0xa8, 0x56, 0x98, 0x67, 0x17,
	0xff, 0xf7, 0x49, 0x13, 0x0a, 0x47, 0x2b, 0x96, 0x2c, 0x38, 0xca, 0x4c, 0x9f, 0xd7, 0x1f, 0x7f,
	0x51, 0xd3, 0xde, 0x7f, 0x9d, 0xea, 0xcd, 0x0c, 0xf1, 0x59, 0xa2, 0xe4, 0x96, 0xee, 0xce, 0x19,
	0xce, 0xa1, 0x6b, 0xb6, 0x8a, 0x89, 0xfb, 0x39, 0x47, 0xb9, 0xad, 0x26, 0x4e, 0x07, 0xe4, 0x6b,
	0x68, 0xab, 0x6d, 0x8a, 0xba, 0xaa, 0xc7, 0xe3, 0x47, 0xb7, 0x16, 0xfc, 0x61, 0x9b, 0x22, 0xd5,
	0xd4, 0x61, 0x0c, 0x77, 0x0f, 0xe4, 0xc9, 0x00, 0x5a, 0x6b, 0xac, 0x74, 0x8a, 0x25, 0x39, 0x83,
	0xce, 0x86, 0xf1, 0x1c, 0xcd, 0x28, 0x7e, 0xf6, 0xba, 0xf7, 0xa2, 0x25, 0xfd, 0x69, 0x73, 0x62,
	0xb9, 0xf7, 0xa1, 0x5d, 0x88, 0x93, 0x1e, 0x74, 0x7e, 0x2c, 0x92, 0x83, 0x46, 0xb1, 0x7c, 0x2e,
	0x45, 0x9e, 0x0e, 0x2c, 0x77, 0x04, 0xf0, 0xdd, 0xe4, 0x55, 0x55, 0xec, 0xe1, 0x5e, 0x49, 0x2d,
	0xdd, 0xa2, 0x5d, 0xec, 0xfe, 0x62, 0x81, 0xfd, 0x12, 0x95, 0x8c, 0x02, 0x32, 0x03, 0x48, 0x77,
	0xd2, 0xda, 0x77, 0x7f, 0x3c, 0xba, 0xad, 0x47, 0xba, 0xc7, 0x25, 0x4f, 0xa0, 0xb5, 0x9e, 0x64,
	0xe6, 0x9a, 0x1f, 0xd6, 0x1c, 0x71, 0x65, 0x90, 0x16, 0x68, 0xf7, 0xcf, 0x26, 0xd8, 0xa7, 0x1a,
	0x43, 0x3e, 0x07, 0x3b, 0xd5, 0xaf, 0x9a, 0x71, 0xf1, 0x7e, 0x9d, 0x0b, 0x0d, 0xa2, 0x06, 0x4c,
	0x26, 0xd0, 0xe5, 0xe5, 0xc8, 0x1b, 0xe9, 0x07, 0x37, 0x3f, 0x0c, 0xb4, 0x82, 0x17, 0x82, 0xa1,
	0x9e, 0x40, 0xf3, 0xd0, 0xd5, 0x09, 0x96, 0x63, 0x4a, 0x0d, 0x98, 0x8c, 0xa1, 0xb3, 0x2c, 0x06,
	0xc4, 0x3c, 0x1e, 0xf7, 0x6b, 0x58, 0x7a, 0x88, 0x68, 0x09, 0x2d, 0xa4, 0x62, 0x5d, 0x6f, 0xc7,
	0xbe, 0x51, 0xaa, 0x6c, 0x0a, 0x35, 0x60, 0x42, 0xa0, 0x9d, 0xb0, 0x18, 0x9d, 0x8e, 0xfe, 0x9c,
	0xf4, 0x7a, 0xfa, 0xe9, 0x4f, 0x9f, 0x94, 0xdc, 0x48, 0xf8, 0x7a, 0xe1, 0xd7, 0xfd, 0xea, 0xe6,
	0xb6, 0x7e, 0xcf, 0x9e, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x33, 0x26, 0x7a, 0x2c, 0x0d, 0x07,
	0x00, 0x00,
}