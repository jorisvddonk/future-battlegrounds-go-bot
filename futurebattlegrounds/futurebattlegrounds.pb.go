// Code generated by protoc-gen-go. DO NOT EDIT.
// source: futurebattlegrounds.proto

package futurebattlegroundsRPC

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ShipSpawnRequest struct {
	IFF                  string   `protobuf:"bytes,1,opt,name=IFF,proto3" json:"IFF,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipSpawnRequest) Reset()         { *m = ShipSpawnRequest{} }
func (m *ShipSpawnRequest) String() string { return proto.CompactTextString(m) }
func (*ShipSpawnRequest) ProtoMessage()    {}
func (*ShipSpawnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{0}
}

func (m *ShipSpawnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipSpawnRequest.Unmarshal(m, b)
}
func (m *ShipSpawnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipSpawnRequest.Marshal(b, m, deterministic)
}
func (m *ShipSpawnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipSpawnRequest.Merge(m, src)
}
func (m *ShipSpawnRequest) XXX_Size() int {
	return xxx_messageInfo_ShipSpawnRequest.Size(m)
}
func (m *ShipSpawnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipSpawnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShipSpawnRequest proto.InternalMessageInfo

func (m *ShipSpawnRequest) GetIFF() string {
	if m != nil {
		return m.IFF
	}
	return ""
}

type ShipSpawnReply struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipSpawnReply) Reset()         { *m = ShipSpawnReply{} }
func (m *ShipSpawnReply) String() string { return proto.CompactTextString(m) }
func (*ShipSpawnReply) ProtoMessage()    {}
func (*ShipSpawnReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{1}
}

func (m *ShipSpawnReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipSpawnReply.Unmarshal(m, b)
}
func (m *ShipSpawnReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipSpawnReply.Marshal(b, m, deterministic)
}
func (m *ShipSpawnReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipSpawnReply.Merge(m, src)
}
func (m *ShipSpawnReply) XXX_Size() int {
	return xxx_messageInfo_ShipSpawnReply.Size(m)
}
func (m *ShipSpawnReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipSpawnReply.DiscardUnknown(m)
}

var xxx_messageInfo_ShipSpawnReply proto.InternalMessageInfo

func (m *ShipSpawnReply) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type ShipActionStateRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Thrust               float32  `protobuf:"fixed32,2,opt,name=thrust,proto3" json:"thrust,omitempty"`
	Rotate               float32  `protobuf:"fixed32,3,opt,name=rotate,proto3" json:"rotate,omitempty"`
	Shooting             bool     `protobuf:"varint,4,opt,name=shooting,proto3" json:"shooting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipActionStateRequest) Reset()         { *m = ShipActionStateRequest{} }
func (m *ShipActionStateRequest) String() string { return proto.CompactTextString(m) }
func (*ShipActionStateRequest) ProtoMessage()    {}
func (*ShipActionStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{2}
}

func (m *ShipActionStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipActionStateRequest.Unmarshal(m, b)
}
func (m *ShipActionStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipActionStateRequest.Marshal(b, m, deterministic)
}
func (m *ShipActionStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipActionStateRequest.Merge(m, src)
}
func (m *ShipActionStateRequest) XXX_Size() int {
	return xxx_messageInfo_ShipActionStateRequest.Size(m)
}
func (m *ShipActionStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipActionStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShipActionStateRequest proto.InternalMessageInfo

func (m *ShipActionStateRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *ShipActionStateRequest) GetThrust() float32 {
	if m != nil {
		return m.Thrust
	}
	return 0
}

func (m *ShipActionStateRequest) GetRotate() float32 {
	if m != nil {
		return m.Rotate
	}
	return 0
}

func (m *ShipActionStateRequest) GetShooting() bool {
	if m != nil {
		return m.Shooting
	}
	return false
}

type BaseReply struct {
	OK                   bool     `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseReply) Reset()         { *m = BaseReply{} }
func (m *BaseReply) String() string { return proto.CompactTextString(m) }
func (*BaseReply) ProtoMessage()    {}
func (*BaseReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{3}
}

func (m *BaseReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseReply.Unmarshal(m, b)
}
func (m *BaseReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseReply.Marshal(b, m, deterministic)
}
func (m *BaseReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseReply.Merge(m, src)
}
func (m *BaseReply) XXX_Size() int {
	return xxx_messageInfo_BaseReply.Size(m)
}
func (m *BaseReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseReply.DiscardUnknown(m)
}

var xxx_messageInfo_BaseReply proto.InternalMessageInfo

func (m *BaseReply) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Vector2D struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector2D) Reset()         { *m = Vector2D{} }
func (m *Vector2D) String() string { return proto.CompactTextString(m) }
func (*Vector2D) ProtoMessage()    {}
func (*Vector2D) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{5}
}

func (m *Vector2D) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector2D.Unmarshal(m, b)
}
func (m *Vector2D) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector2D.Marshal(b, m, deterministic)
}
func (m *Vector2D) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector2D.Merge(m, src)
}
func (m *Vector2D) XXX_Size() int {
	return xxx_messageInfo_Vector2D.Size(m)
}
func (m *Vector2D) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector2D.DiscardUnknown(m)
}

var xxx_messageInfo_Vector2D proto.InternalMessageInfo

func (m *Vector2D) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector2D) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Ship struct {
	Position             *Vector2D `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	MovementVector       *Vector2D `protobuf:"bytes,2,opt,name=movementVector,proto3" json:"movementVector,omitempty"`
	RotationVector       *Vector2D `protobuf:"bytes,3,opt,name=rotationVector,proto3" json:"rotationVector,omitempty"`
	RemainingLifetime    float64   `protobuf:"fixed64,4,opt,name=remainingLifetime,proto3" json:"remainingLifetime,omitempty"`
	IFF                  string    `protobuf:"bytes,5,opt,name=IFF,proto3" json:"IFF,omitempty"`
	Hull                 float64   `protobuf:"fixed64,6,opt,name=hull,proto3" json:"hull,omitempty"`
	Battery              float64   `protobuf:"fixed64,7,opt,name=battery,proto3" json:"battery,omitempty"`
	UUID                 string    `protobuf:"bytes,8,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ship) Reset()         { *m = Ship{} }
func (m *Ship) String() string { return proto.CompactTextString(m) }
func (*Ship) ProtoMessage()    {}
func (*Ship) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{6}
}

func (m *Ship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ship.Unmarshal(m, b)
}
func (m *Ship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ship.Marshal(b, m, deterministic)
}
func (m *Ship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ship.Merge(m, src)
}
func (m *Ship) XXX_Size() int {
	return xxx_messageInfo_Ship.Size(m)
}
func (m *Ship) XXX_DiscardUnknown() {
	xxx_messageInfo_Ship.DiscardUnknown(m)
}

var xxx_messageInfo_Ship proto.InternalMessageInfo

func (m *Ship) GetPosition() *Vector2D {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Ship) GetMovementVector() *Vector2D {
	if m != nil {
		return m.MovementVector
	}
	return nil
}

func (m *Ship) GetRotationVector() *Vector2D {
	if m != nil {
		return m.RotationVector
	}
	return nil
}

func (m *Ship) GetRemainingLifetime() float64 {
	if m != nil {
		return m.RemainingLifetime
	}
	return 0
}

func (m *Ship) GetIFF() string {
	if m != nil {
		return m.IFF
	}
	return ""
}

func (m *Ship) GetHull() float64 {
	if m != nil {
		return m.Hull
	}
	return 0
}

func (m *Ship) GetBattery() float64 {
	if m != nil {
		return m.Battery
	}
	return 0
}

func (m *Ship) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type Bullet struct {
	Position             *Vector2D `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	MovementVector       *Vector2D `protobuf:"bytes,2,opt,name=movementVector,proto3" json:"movementVector,omitempty"`
	RotationVector       *Vector2D `protobuf:"bytes,3,opt,name=rotationVector,proto3" json:"rotationVector,omitempty"`
	RemainingLifetime    float64   `protobuf:"fixed64,4,opt,name=remainingLifetime,proto3" json:"remainingLifetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Bullet) Reset()         { *m = Bullet{} }
func (m *Bullet) String() string { return proto.CompactTextString(m) }
func (*Bullet) ProtoMessage()    {}
func (*Bullet) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{7}
}

func (m *Bullet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bullet.Unmarshal(m, b)
}
func (m *Bullet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bullet.Marshal(b, m, deterministic)
}
func (m *Bullet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bullet.Merge(m, src)
}
func (m *Bullet) XXX_Size() int {
	return xxx_messageInfo_Bullet.Size(m)
}
func (m *Bullet) XXX_DiscardUnknown() {
	xxx_messageInfo_Bullet.DiscardUnknown(m)
}

var xxx_messageInfo_Bullet proto.InternalMessageInfo

func (m *Bullet) GetPosition() *Vector2D {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Bullet) GetMovementVector() *Vector2D {
	if m != nil {
		return m.MovementVector
	}
	return nil
}

func (m *Bullet) GetRotationVector() *Vector2D {
	if m != nil {
		return m.RotationVector
	}
	return nil
}

func (m *Bullet) GetRemainingLifetime() float64 {
	if m != nil {
		return m.RemainingLifetime
	}
	return 0
}

type Battleground struct {
	Ships                []*Ship   `protobuf:"bytes,1,rep,name=ships,proto3" json:"ships,omitempty"`
	Bullets              []*Bullet `protobuf:"bytes,2,rep,name=bullets,proto3" json:"bullets,omitempty"`
	Timestamp            float64   `protobuf:"fixed64,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Battleground) Reset()         { *m = Battleground{} }
func (m *Battleground) String() string { return proto.CompactTextString(m) }
func (*Battleground) ProtoMessage()    {}
func (*Battleground) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fa39c494238874a, []int{8}
}

func (m *Battleground) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Battleground.Unmarshal(m, b)
}
func (m *Battleground) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Battleground.Marshal(b, m, deterministic)
}
func (m *Battleground) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Battleground.Merge(m, src)
}
func (m *Battleground) XXX_Size() int {
	return xxx_messageInfo_Battleground.Size(m)
}
func (m *Battleground) XXX_DiscardUnknown() {
	xxx_messageInfo_Battleground.DiscardUnknown(m)
}

var xxx_messageInfo_Battleground proto.InternalMessageInfo

func (m *Battleground) GetShips() []*Ship {
	if m != nil {
		return m.Ships
	}
	return nil
}

func (m *Battleground) GetBullets() []*Bullet {
	if m != nil {
		return m.Bullets
	}
	return nil
}

func (m *Battleground) GetTimestamp() float64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*ShipSpawnRequest)(nil), "futurebattlegroundsRPC.ShipSpawnRequest")
	proto.RegisterType((*ShipSpawnReply)(nil), "futurebattlegroundsRPC.ShipSpawnReply")
	proto.RegisterType((*ShipActionStateRequest)(nil), "futurebattlegroundsRPC.ShipActionStateRequest")
	proto.RegisterType((*BaseReply)(nil), "futurebattlegroundsRPC.BaseReply")
	proto.RegisterType((*Empty)(nil), "futurebattlegroundsRPC.Empty")
	proto.RegisterType((*Vector2D)(nil), "futurebattlegroundsRPC.Vector2d")
	proto.RegisterType((*Ship)(nil), "futurebattlegroundsRPC.Ship")
	proto.RegisterType((*Bullet)(nil), "futurebattlegroundsRPC.Bullet")
	proto.RegisterType((*Battleground)(nil), "futurebattlegroundsRPC.Battleground")
}

func init() { proto.RegisterFile("futurebattlegrounds.proto", fileDescriptor_1fa39c494238874a) }

var fileDescriptor_1fa39c494238874a = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xce, 0x3a, 0x5f, 0xce, 0xb4, 0x6f, 0xde, 0xb2, 0x87, 0xc8, 0x84, 0x82, 0xc2, 0x2a, 0xaa,
	0x72, 0x40, 0x11, 0x0a, 0x17, 0x90, 0xb8, 0x60, 0x20, 0xa2, 0x2a, 0x52, 0xa3, 0x8d, 0x8a, 0xb8,
	0x20, 0xe4, 0xb4, 0xdb, 0xc4, 0x92, 0xed, 0x35, 0xde, 0x31, 0xc4, 0xff, 0x80, 0x23, 0x67, 0x8e,
	0xfc, 0x2c, 0x7e, 0x0d, 0xda, 0x75, 0x12, 0x1c, 0x1a, 0x57, 0xb9, 0x70, 0xe2, 0xb6, 0x33, 0xfb,
	0x3c, 0xcf, 0x7c, 0xec, 0x78, 0x0c, 0x77, 0xaf, 0x53, 0x4c, 0x13, 0x31, 0xf3, 0x10, 0x03, 0x31,
	0x4f, 0x64, 0x1a, 0x5d, 0xa9, 0x61, 0x9c, 0x48, 0x94, 0xb4, 0xb3, 0xe3, 0x8a, 0x4f, 0x5e, 0xb2,
	0x3e, 0x1c, 0x4d, 0x17, 0x7e, 0x3c, 0x8d, 0xbd, 0x2f, 0x11, 0x17, 0x9f, 0x52, 0xa1, 0x90, 0x1e,
	0x41, 0xf5, 0x74, 0x3c, 0x76, 0x48, 0x8f, 0x0c, 0x5a, 0x5c, 0x1f, 0x59, 0x1f, 0xda, 0x05, 0x54,
	0x1c, 0x64, 0x94, 0x42, 0xed, 0xe2, 0xe2, 0xf4, 0xd5, 0x0a, 0x64, 0xce, 0x6c, 0x09, 0x1d, 0x8d,
	0x7a, 0x71, 0x89, 0xbe, 0x8c, 0xa6, 0xe8, 0xa1, 0x58, 0x2b, 0xee, 0x40, 0xd3, 0x0e, 0x34, 0x70,
	0x91, 0xa4, 0x0a, 0x1d, 0xab, 0x47, 0x06, 0x16, 0x5f, 0x59, 0xda, 0x9f, 0x48, 0x4d, 0x76, 0xaa,
	0xb9, 0x3f, 0xb7, 0x68, 0x17, 0x6c, 0xb5, 0x90, 0x12, 0xfd, 0x68, 0xee, 0xd4, 0x7a, 0x64, 0x60,
	0xf3, 0x8d, 0xcd, 0xee, 0x41, 0xcb, 0xf5, 0x94, 0xc8, 0x53, 0x6b, 0x83, 0x75, 0x7e, 0x66, 0x42,
	0xd9, 0xdc, 0x3a, 0x3f, 0x63, 0x4d, 0xa8, 0xbf, 0x0e, 0x63, 0xcc, 0xd8, 0x09, 0xd8, 0xef, 0xc4,
	0x25, 0xca, 0x64, 0x74, 0x45, 0x0f, 0x81, 0x2c, 0x0d, 0x86, 0x70, 0xb2, 0xd4, 0x56, 0x66, 0xd2,
	0x20, 0x9c, 0x64, 0xec, 0xa7, 0x05, 0x35, 0x5d, 0x08, 0x7d, 0x0e, 0x76, 0x2c, 0x95, 0xaf, 0xcb,
	0x31, 0xd8, 0x83, 0x51, 0x6f, 0xb8, 0xbb, 0x8f, 0xc3, 0xb5, 0x30, 0xdf, 0x30, 0xe8, 0x1b, 0x68,
	0x87, 0xf2, 0xb3, 0x08, 0x45, 0x84, 0xf9, 0xad, 0x89, 0xb0, 0x8f, 0xc6, 0x1f, 0x3c, 0xad, 0x64,
	0x9a, 0xe0, 0xcb, 0x68, 0xa5, 0x54, 0xdd, 0x57, 0x69, 0x9b, 0x47, 0x1f, 0xc1, 0x9d, 0x44, 0x84,
	0x9e, 0x1f, 0xf9, 0xd1, 0xfc, 0xad, 0x7f, 0x2d, 0xd0, 0x0f, 0x85, 0xe9, 0x26, 0xe1, 0x37, 0x2f,
	0xd6, 0x83, 0x50, 0xdf, 0x0c, 0x82, 0x7e, 0xc8, 0x45, 0x1a, 0x04, 0x4e, 0xc3, 0x50, 0xcc, 0x99,
	0x3a, 0xd0, 0xd4, 0x09, 0x88, 0x24, 0x73, 0x9a, 0xc6, 0xbd, 0x36, 0x37, 0xcf, 0x6e, 0x17, 0x86,
	0xe4, 0xab, 0x05, 0x0d, 0x37, 0x0d, 0x02, 0x81, 0xff, 0x7a, 0x7b, 0xd9, 0x77, 0x02, 0x87, 0x6e,
	0x41, 0x9b, 0x8e, 0xa0, 0xae, 0x16, 0x7e, 0xac, 0x1c, 0xd2, 0xab, 0x0e, 0x0e, 0x46, 0xc7, 0x65,
	0xf1, 0xf5, 0x70, 0xf2, 0x1c, 0x4a, 0x9f, 0x42, 0x73, 0x66, 0xda, 0xa9, 0x1c, 0xcb, 0xb0, 0x1e,
	0x94, 0xb1, 0xf2, 0xae, 0xf3, 0x35, 0x9c, 0x1e, 0x43, 0x4b, 0xa7, 0xa1, 0xd0, 0x0b, 0x63, 0x53,
	0x31, 0xe1, 0xbf, 0x1d, 0xa3, 0x6f, 0x55, 0xf8, 0xaf, 0x98, 0x9c, 0xa2, 0xef, 0xe1, 0xff, 0xb9,
	0xc0, 0xad, 0x84, 0xef, 0x97, 0xc5, 0x32, 0x1f, 0x5c, 0xb7, 0x5f, 0x9a, 0x4a, 0xc1, 0xc1, 0x2a,
	0xf4, 0x03, 0x50, 0x85, 0x89, 0xf0, 0xc2, 0xbf, 0x20, 0xfe, 0x98, 0xd0, 0x8f, 0xd0, 0x52, 0x7a,
	0x73, 0x99, 0x6f, 0x7a, 0x70, 0x5b, 0x53, 0x8b, 0x6b, 0xb0, 0x7b, 0xb2, 0x07, 0x32, 0x0e, 0x32,
	0x56, 0xa1, 0x02, 0xda, 0x4a, 0x60, 0x61, 0xef, 0xd1, 0xe1, 0x6d, 0xdc, 0x9b, 0x0b, 0xb2, 0xfb,
	0xb0, 0xbc, 0x98, 0xd5, 0x5a, 0x63, 0x15, 0xf7, 0x19, 0x94, 0x6c, 0x71, 0xd7, 0x19, 0x1b, 0xff,
	0xd6, 0x7b, 0x4d, 0xf4, 0xde, 0x9f, 0x90, 0x1f, 0x96, 0x35, 0x76, 0x67, 0x0d, 0xf3, 0x17, 0x78,
	0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x54, 0xcd, 0xd7, 0x22, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BattlegroundsClient is the client API for Battlegrounds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BattlegroundsClient interface {
	GetBattleground(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Battleground, error)
	StreamBattleground(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Battlegrounds_StreamBattlegroundClient, error)
	SpawnShip(ctx context.Context, in *ShipSpawnRequest, opts ...grpc.CallOption) (*ShipSpawnReply, error)
	SetActionState(ctx context.Context, in *ShipActionStateRequest, opts ...grpc.CallOption) (*BaseReply, error)
}

type battlegroundsClient struct {
	cc *grpc.ClientConn
}

func NewBattlegroundsClient(cc *grpc.ClientConn) BattlegroundsClient {
	return &battlegroundsClient{cc}
}

func (c *battlegroundsClient) GetBattleground(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Battleground, error) {
	out := new(Battleground)
	err := c.cc.Invoke(ctx, "/futurebattlegroundsRPC.Battlegrounds/getBattleground", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battlegroundsClient) StreamBattleground(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Battlegrounds_StreamBattlegroundClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Battlegrounds_serviceDesc.Streams[0], "/futurebattlegroundsRPC.Battlegrounds/streamBattleground", opts...)
	if err != nil {
		return nil, err
	}
	x := &battlegroundsStreamBattlegroundClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Battlegrounds_StreamBattlegroundClient interface {
	Recv() (*Battleground, error)
	grpc.ClientStream
}

type battlegroundsStreamBattlegroundClient struct {
	grpc.ClientStream
}

func (x *battlegroundsStreamBattlegroundClient) Recv() (*Battleground, error) {
	m := new(Battleground)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *battlegroundsClient) SpawnShip(ctx context.Context, in *ShipSpawnRequest, opts ...grpc.CallOption) (*ShipSpawnReply, error) {
	out := new(ShipSpawnReply)
	err := c.cc.Invoke(ctx, "/futurebattlegroundsRPC.Battlegrounds/spawnShip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battlegroundsClient) SetActionState(ctx context.Context, in *ShipActionStateRequest, opts ...grpc.CallOption) (*BaseReply, error) {
	out := new(BaseReply)
	err := c.cc.Invoke(ctx, "/futurebattlegroundsRPC.Battlegrounds/setActionState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BattlegroundsServer is the server API for Battlegrounds service.
type BattlegroundsServer interface {
	GetBattleground(context.Context, *Empty) (*Battleground, error)
	StreamBattleground(*Empty, Battlegrounds_StreamBattlegroundServer) error
	SpawnShip(context.Context, *ShipSpawnRequest) (*ShipSpawnReply, error)
	SetActionState(context.Context, *ShipActionStateRequest) (*BaseReply, error)
}

// UnimplementedBattlegroundsServer can be embedded to have forward compatible implementations.
type UnimplementedBattlegroundsServer struct {
}

func (*UnimplementedBattlegroundsServer) GetBattleground(ctx context.Context, req *Empty) (*Battleground, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBattleground not implemented")
}
func (*UnimplementedBattlegroundsServer) StreamBattleground(req *Empty, srv Battlegrounds_StreamBattlegroundServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBattleground not implemented")
}
func (*UnimplementedBattlegroundsServer) SpawnShip(ctx context.Context, req *ShipSpawnRequest) (*ShipSpawnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpawnShip not implemented")
}
func (*UnimplementedBattlegroundsServer) SetActionState(ctx context.Context, req *ShipActionStateRequest) (*BaseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActionState not implemented")
}

func RegisterBattlegroundsServer(s *grpc.Server, srv BattlegroundsServer) {
	s.RegisterService(&_Battlegrounds_serviceDesc, srv)
}

func _Battlegrounds_GetBattleground_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattlegroundsServer).GetBattleground(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/futurebattlegroundsRPC.Battlegrounds/GetBattleground",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattlegroundsServer).GetBattleground(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battlegrounds_StreamBattleground_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BattlegroundsServer).StreamBattleground(m, &battlegroundsStreamBattlegroundServer{stream})
}

type Battlegrounds_StreamBattlegroundServer interface {
	Send(*Battleground) error
	grpc.ServerStream
}

type battlegroundsStreamBattlegroundServer struct {
	grpc.ServerStream
}

func (x *battlegroundsStreamBattlegroundServer) Send(m *Battleground) error {
	return x.ServerStream.SendMsg(m)
}

func _Battlegrounds_SpawnShip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipSpawnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattlegroundsServer).SpawnShip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/futurebattlegroundsRPC.Battlegrounds/SpawnShip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattlegroundsServer).SpawnShip(ctx, req.(*ShipSpawnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battlegrounds_SetActionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShipActionStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattlegroundsServer).SetActionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/futurebattlegroundsRPC.Battlegrounds/SetActionState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattlegroundsServer).SetActionState(ctx, req.(*ShipActionStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Battlegrounds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "futurebattlegroundsRPC.Battlegrounds",
	HandlerType: (*BattlegroundsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBattleground",
			Handler:    _Battlegrounds_GetBattleground_Handler,
		},
		{
			MethodName: "spawnShip",
			Handler:    _Battlegrounds_SpawnShip_Handler,
		},
		{
			MethodName: "setActionState",
			Handler:    _Battlegrounds_SetActionState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "streamBattleground",
			Handler:       _Battlegrounds_StreamBattleground_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "futurebattlegrounds.proto",
}
