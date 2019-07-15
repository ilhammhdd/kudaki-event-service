// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/kudaki_event.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	kudaki_event "github.com/ilhammhdd/kudaki-event-service/entities/aggregates/kudaki_event"
	user "github.com/ilhammhdd/kudaki-event-service/entities/aggregates/user"
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

type EventServiceCommandTopic int32

const (
	EventServiceCommandTopic_ADD_KUDAKI_EVENT       EventServiceCommandTopic = 0
	EventServiceCommandTopic_UPDATE_KUDAKI_EVENT    EventServiceCommandTopic = 1
	EventServiceCommandTopic_DELETE_KUDAKI_EVENT    EventServiceCommandTopic = 2
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENT  EventServiceCommandTopic = 3
	EventServiceCommandTopic_RETRIEVE_KUDAKI_EVENTS EventServiceCommandTopic = 4
)

var EventServiceCommandTopic_name = map[int32]string{
	0: "ADD_KUDAKI_EVENT",
	1: "UPDATE_KUDAKI_EVENT",
	2: "DELETE_KUDAKI_EVENT",
	3: "RETRIEVE_KUDAKI_EVENT",
	4: "RETRIEVE_KUDAKI_EVENTS",
}

var EventServiceCommandTopic_value = map[string]int32{
	"ADD_KUDAKI_EVENT":       0,
	"UPDATE_KUDAKI_EVENT":    1,
	"DELETE_KUDAKI_EVENT":    2,
	"RETRIEVE_KUDAKI_EVENT":  3,
	"RETRIEVE_KUDAKI_EVENTS": 4,
}

func (x EventServiceCommandTopic) String() string {
	return proto.EnumName(EventServiceCommandTopic_name, int32(x))
}

func (EventServiceCommandTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

type EventServiceEventTopic int32

const (
	EventServiceEventTopic_KUDAKI_EVENT_ADDED      EventServiceEventTopic = 0
	EventServiceEventTopic_KUDAKI_EVENT_UPDATED    EventServiceEventTopic = 1
	EventServiceEventTopic_KUDAKI_EVENT_DELETED    EventServiceEventTopic = 2
	EventServiceEventTopic_KUDAKI_EVENT_RETRIEVED  EventServiceEventTopic = 3
	EventServiceEventTopic_KUDAKI_EVENTS_RETRIEVED EventServiceEventTopic = 4
)

var EventServiceEventTopic_name = map[int32]string{
	0: "KUDAKI_EVENT_ADDED",
	1: "KUDAKI_EVENT_UPDATED",
	2: "KUDAKI_EVENT_DELETED",
	3: "KUDAKI_EVENT_RETRIEVED",
	4: "KUDAKI_EVENTS_RETRIEVED",
}

var EventServiceEventTopic_value = map[string]int32{
	"KUDAKI_EVENT_ADDED":      0,
	"KUDAKI_EVENT_UPDATED":    1,
	"KUDAKI_EVENT_DELETED":    2,
	"KUDAKI_EVENT_RETRIEVED":  3,
	"KUDAKI_EVENTS_RETRIEVED": 4,
}

func (x EventServiceEventTopic) String() string {
	return proto.EnumName(EventServiceEventTopic_name, int32(x))
}

func (EventServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

type EventPaymentServiceEventTopic int32

const (
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_REQUESTED EventPaymentServiceEventTopic = 0
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_IDENTIFY  EventPaymentServiceEventTopic = 1
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_REDIRECT  EventPaymentServiceEventTopic = 2
	EventPaymentServiceEventTopic_EVENT_DOKU_PAYMENT_NOTIFY    EventPaymentServiceEventTopic = 3
)

var EventPaymentServiceEventTopic_name = map[int32]string{
	0: "EVENT_DOKU_PAYMENT_REQUESTED",
	1: "EVENT_DOKU_PAYMENT_IDENTIFY",
	2: "EVENT_DOKU_PAYMENT_REDIRECT",
	3: "EVENT_DOKU_PAYMENT_NOTIFY",
}

var EventPaymentServiceEventTopic_value = map[string]int32{
	"EVENT_DOKU_PAYMENT_REQUESTED": 0,
	"EVENT_DOKU_PAYMENT_IDENTIFY":  1,
	"EVENT_DOKU_PAYMENT_REDIRECT":  2,
	"EVENT_DOKU_PAYMENT_NOTIFY":    3,
}

func (x EventPaymentServiceEventTopic) String() string {
	return proto.EnumName(EventPaymentServiceEventTopic_name, int32(x))
}

func (EventPaymentServiceEventTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

type AddKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Venue                string   `protobuf:"bytes,2,opt,name=venue,proto3" json:"venue,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DurationFrom         int64    `protobuf:"varint,4,opt,name=duration_from,json=durationFrom,proto3" json:"duration_from,omitempty"`
	DurationTo           int64    `protobuf:"varint,5,opt,name=duration_to,json=durationTo,proto3" json:"duration_to,omitempty"`
	AdDurationFrom       int64    `protobuf:"varint,7,opt,name=ad_duration_from,json=adDurationFrom,proto3" json:"ad_duration_from,omitempty"`
	AdDurationTo         int64    `protobuf:"varint,8,opt,name=ad_duration_to,json=adDurationTo,proto3" json:"ad_duration_to,omitempty"`
	KudakiToken          string   `protobuf:"bytes,9,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddKudakiEvent) Reset()         { *m = AddKudakiEvent{} }
func (m *AddKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*AddKudakiEvent) ProtoMessage()    {}
func (*AddKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{0}
}

func (m *AddKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKudakiEvent.Unmarshal(m, b)
}
func (m *AddKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKudakiEvent.Marshal(b, m, deterministic)
}
func (m *AddKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKudakiEvent.Merge(m, src)
}
func (m *AddKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_AddKudakiEvent.Size(m)
}
func (m *AddKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AddKudakiEvent proto.InternalMessageInfo

func (m *AddKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddKudakiEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddKudakiEvent) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *AddKudakiEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddKudakiEvent) GetDurationFrom() int64 {
	if m != nil {
		return m.DurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetDurationTo() int64 {
	if m != nil {
		return m.DurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetAdDurationFrom() int64 {
	if m != nil {
		return m.AdDurationFrom
	}
	return 0
}

func (m *AddKudakiEvent) GetAdDurationTo() int64 {
	if m != nil {
		return m.AdDurationTo
	}
	return 0
}

func (m *AddKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type DeleteKudakiEvent struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	EventUuid            string   `protobuf:"bytes,2,opt,name=event_uuid,json=eventUuid,proto3" json:"event_uuid,omitempty"`
	KudakiToken          string   `protobuf:"bytes,3,opt,name=kudaki_token,json=kudakiToken,proto3" json:"kudaki_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKudakiEvent) Reset()         { *m = DeleteKudakiEvent{} }
func (m *DeleteKudakiEvent) String() string { return proto.CompactTextString(m) }
func (*DeleteKudakiEvent) ProtoMessage()    {}
func (*DeleteKudakiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{1}
}

func (m *DeleteKudakiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKudakiEvent.Unmarshal(m, b)
}
func (m *DeleteKudakiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKudakiEvent.Marshal(b, m, deterministic)
}
func (m *DeleteKudakiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKudakiEvent.Merge(m, src)
}
func (m *DeleteKudakiEvent) XXX_Size() int {
	return xxx_messageInfo_DeleteKudakiEvent.Size(m)
}
func (m *DeleteKudakiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKudakiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKudakiEvent proto.InternalMessageInfo

func (m *DeleteKudakiEvent) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetEventUuid() string {
	if m != nil {
		return m.EventUuid
	}
	return ""
}

func (m *DeleteKudakiEvent) GetKudakiToken() string {
	if m != nil {
		return m.KudakiToken
	}
	return ""
}

type KudakiEventAdded struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.Profile             `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	AddKudakiEvent       *AddKudakiEvent           `protobuf:"bytes,5,opt,name=add_kudaki_event,json=addKudakiEvent,proto3" json:"add_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventAdded) Reset()         { *m = KudakiEventAdded{} }
func (m *KudakiEventAdded) String() string { return proto.CompactTextString(m) }
func (*KudakiEventAdded) ProtoMessage()    {}
func (*KudakiEventAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{2}
}

func (m *KudakiEventAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventAdded.Unmarshal(m, b)
}
func (m *KudakiEventAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventAdded.Marshal(b, m, deterministic)
}
func (m *KudakiEventAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventAdded.Merge(m, src)
}
func (m *KudakiEventAdded) XXX_Size() int {
	return xxx_messageInfo_KudakiEventAdded.Size(m)
}
func (m *KudakiEventAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventAdded.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventAdded proto.InternalMessageInfo

func (m *KudakiEventAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventAdded) GetOrganizer() *user.Profile {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventAdded) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventAdded) GetAddKudakiEvent() *AddKudakiEvent {
	if m != nil {
		return m.AddKudakiEvent
	}
	return nil
}

type KudakiEventDeleted struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.User                `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	KudakiEvent          *kudaki_event.KudakiEvent `protobuf:"bytes,4,opt,name=kudaki_event,json=kudakiEvent,proto3" json:"kudaki_event,omitempty"`
	DeleteKudakiEvent    *DeleteKudakiEvent        `protobuf:"bytes,5,opt,name=delete_kudaki_event,json=deleteKudakiEvent,proto3" json:"delete_kudaki_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDeleted) Reset()         { *m = KudakiEventDeleted{} }
func (m *KudakiEventDeleted) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDeleted) ProtoMessage()    {}
func (*KudakiEventDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{3}
}

func (m *KudakiEventDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDeleted.Unmarshal(m, b)
}
func (m *KudakiEventDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDeleted.Marshal(b, m, deterministic)
}
func (m *KudakiEventDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDeleted.Merge(m, src)
}
func (m *KudakiEventDeleted) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDeleted.Size(m)
}
func (m *KudakiEventDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDeleted proto.InternalMessageInfo

func (m *KudakiEventDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDeleted) GetOrganizer() *user.User {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDeleted) GetKudakiEvent() *kudaki_event.KudakiEvent {
	if m != nil {
		return m.KudakiEvent
	}
	return nil
}

func (m *KudakiEventDeleted) GetDeleteKudakiEvent() *DeleteKudakiEvent {
	if m != nil {
		return m.DeleteKudakiEvent
	}
	return nil
}

type KudakiEventDokuPaymentRequested struct {
	Uid                  string                    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Organizer            *user.Profile             `protobuf:"bytes,2,opt,name=organizer,proto3" json:"organizer,omitempty"`
	EventStatus          *Status                   `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	DokuInvoice          *kudaki_event.DokuInvoice `protobuf:"bytes,4,opt,name=doku_invoice,json=dokuInvoice,proto3" json:"doku_invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *KudakiEventDokuPaymentRequested) Reset()         { *m = KudakiEventDokuPaymentRequested{} }
func (m *KudakiEventDokuPaymentRequested) String() string { return proto.CompactTextString(m) }
func (*KudakiEventDokuPaymentRequested) ProtoMessage()    {}
func (*KudakiEventDokuPaymentRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa36ed941739bd5b, []int{4}
}

func (m *KudakiEventDokuPaymentRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KudakiEventDokuPaymentRequested.Unmarshal(m, b)
}
func (m *KudakiEventDokuPaymentRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KudakiEventDokuPaymentRequested.Marshal(b, m, deterministic)
}
func (m *KudakiEventDokuPaymentRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KudakiEventDokuPaymentRequested.Merge(m, src)
}
func (m *KudakiEventDokuPaymentRequested) XXX_Size() int {
	return xxx_messageInfo_KudakiEventDokuPaymentRequested.Size(m)
}
func (m *KudakiEventDokuPaymentRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_KudakiEventDokuPaymentRequested.DiscardUnknown(m)
}

var xxx_messageInfo_KudakiEventDokuPaymentRequested proto.InternalMessageInfo

func (m *KudakiEventDokuPaymentRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KudakiEventDokuPaymentRequested) GetOrganizer() *user.Profile {
	if m != nil {
		return m.Organizer
	}
	return nil
}

func (m *KudakiEventDokuPaymentRequested) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *KudakiEventDokuPaymentRequested) GetDokuInvoice() *kudaki_event.DokuInvoice {
	if m != nil {
		return m.DokuInvoice
	}
	return nil
}

func init() {
	proto.RegisterEnum("aggregates.event.EventServiceCommandTopic", EventServiceCommandTopic_name, EventServiceCommandTopic_value)
	proto.RegisterEnum("aggregates.event.EventServiceEventTopic", EventServiceEventTopic_name, EventServiceEventTopic_value)
	proto.RegisterEnum("aggregates.event.EventPaymentServiceEventTopic", EventPaymentServiceEventTopic_name, EventPaymentServiceEventTopic_value)
	proto.RegisterType((*AddKudakiEvent)(nil), "aggregates.event.AddKudakiEvent")
	proto.RegisterType((*DeleteKudakiEvent)(nil), "aggregates.event.DeleteKudakiEvent")
	proto.RegisterType((*KudakiEventAdded)(nil), "aggregates.event.KudakiEventAdded")
	proto.RegisterType((*KudakiEventDeleted)(nil), "aggregates.event.KudakiEventDeleted")
	proto.RegisterType((*KudakiEventDokuPaymentRequested)(nil), "aggregates.event.KudakiEventDokuPaymentRequested")
}

func init() { proto.RegisterFile("events/kudaki_event.proto", fileDescriptor_fa36ed941739bd5b) }

var fileDescriptor_fa36ed941739bd5b = []byte{
	// 750 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcf, 0x6f, 0x12, 0x41,
	0x18, 0xed, 0x02, 0xad, 0xf2, 0xd1, 0x36, 0xdb, 0xe9, 0xaf, 0x2d, 0x95, 0x14, 0x69, 0x0f, 0x0d,
	0x49, 0xc1, 0xb4, 0x89, 0x89, 0xf1, 0x84, 0x9d, 0xad, 0x41, 0xb4, 0xc5, 0x65, 0x69, 0x52, 0x2f,
	0x9b, 0x2d, 0x33, 0xa5, 0x1b, 0xba, 0x3b, 0xb8, 0x3f, 0x48, 0xf4, 0x3f, 0xf1, 0xe0, 0x45, 0x0f,
	0x1e, 0xfc, 0xbf, 0xfc, 0x3b, 0xcc, 0xce, 0x40, 0x99, 0x65, 0xd1, 0x1e, 0xf5, 0x42, 0x66, 0xde,
	0x7b, 0x7c, 0xdf, 0x9b, 0x37, 0x1f, 0x0c, 0xec, 0xd0, 0x11, 0xf5, 0xc2, 0xa0, 0x3e, 0x88, 0x88,
	0x3d, 0x70, 0x2c, 0xbe, 0xab, 0x0d, 0x7d, 0x16, 0x32, 0xa4, 0xda, 0xfd, 0xbe, 0x4f, 0xfb, 0x76,
	0x48, 0x83, 0x1a, 0xc7, 0x8b, 0xc5, 0x29, 0x52, 0x8f, 0x02, 0xea, 0xf3, 0x0f, 0xa1, 0x2e, 0x56,
	0x25, 0x4e, 0x2e, 0x56, 0x27, 0x6c, 0x10, 0x59, 0x8e, 0x37, 0x62, 0x4e, 0x8f, 0x8e, 0xb5, 0xa5,
	0xd9, 0x3a, 0x43, 0x9f, 0xdd, 0x38, 0x77, 0xf4, 0xa1, 0x52, 0x69, 0x93, 0xc5, 0xf5, 0xb1, 0xff,
	0x20, 0xb4, 0xc3, 0x28, 0x10, 0x60, 0xe5, 0x47, 0x06, 0x56, 0x1b, 0x84, 0xb4, 0xb8, 0x5c, 0x8f,
	0x05, 0x48, 0x85, 0x6c, 0xe4, 0x10, 0x4d, 0x29, 0x2b, 0x87, 0x79, 0x23, 0x5e, 0x22, 0x04, 0x39,
	0xcf, 0x76, 0xa9, 0xb6, 0xc4, 0x21, 0xbe, 0x46, 0x1b, 0xb0, 0x38, 0xa2, 0x5e, 0x44, 0xb5, 0x0c,
	0x07, 0xc5, 0x06, 0x95, 0xa1, 0x40, 0x68, 0xd0, 0xf3, 0x9d, 0x61, 0xe8, 0x30, 0x4f, 0xcb, 0x72,
	0x4e, 0x86, 0xd0, 0x3e, 0xac, 0x90, 0xc8, 0xb7, 0xe3, 0xb5, 0x75, 0xe3, 0x33, 0x57, 0xcb, 0x95,
	0x95, 0xc3, 0xac, 0xb1, 0x3c, 0x01, 0xcf, 0x7c, 0xe6, 0xa2, 0x3d, 0x28, 0xdc, 0x8b, 0x42, 0xa6,
	0x2d, 0x72, 0x09, 0x4c, 0x20, 0x93, 0xa1, 0x43, 0x50, 0x6d, 0x62, 0x25, 0x0b, 0x3d, 0xe2, 0xaa,
	0x55, 0x9b, 0x60, 0xb9, 0xd4, 0x01, 0xac, 0xca, 0xca, 0x90, 0x69, 0x8f, 0x45, 0xc3, 0xa9, 0xce,
	0x64, 0xe8, 0x29, 0x2c, 0x8f, 0x13, 0x0b, 0xd9, 0x80, 0x7a, 0x5a, 0x5e, 0x18, 0x17, 0x98, 0x19,
	0x43, 0x95, 0x3e, 0xac, 0x61, 0x7a, 0x47, 0x43, 0xfa, 0xf7, 0xac, 0x4a, 0x00, 0x3c, 0x67, 0x2b,
	0x8a, 0x09, 0x11, 0x4e, 0x9e, 0x23, 0xdd, 0x98, 0x9e, 0x6d, 0x94, 0x4d, 0x37, 0xfa, 0x9e, 0x01,
	0x55, 0xea, 0xd1, 0x20, 0x84, 0x92, 0x39, 0x8d, 0x9e, 0x43, 0x9e, 0xf9, 0x7d, 0xdb, 0x73, 0x3e,
	0x53, 0x9f, 0xf7, 0x29, 0x1c, 0x6b, 0x35, 0x69, 0x0e, 0xf9, 0xc0, 0xb5, 0xc5, 0xb4, 0x18, 0x53,
	0x29, 0x7a, 0x06, 0xcb, 0xc2, 0xa0, 0x98, 0x03, 0xee, 0xa0, 0x70, 0xbc, 0x22, 0xe6, 0xb6, 0xd6,
	0xe1, 0xa0, 0x51, 0xe0, 0x3b, 0xb1, 0x41, 0xaf, 0xef, 0x3d, 0x73, 0x94, 0xdf, 0x58, 0xe1, 0xf8,
	0x40, 0x6e, 0x96, 0x18, 0x37, 0xc9, 0xfc, 0xe4, 0x64, 0x22, 0xad, 0x37, 0xf1, 0xad, 0x11, 0x2b,
	0x51, 0x6c, 0x91, 0x17, 0x2b, 0xd7, 0x66, 0x7f, 0x41, 0xb5, 0xe4, 0x54, 0xc6, 0xf7, 0x2a, 0xef,
	0x2b, 0x3f, 0x33, 0x80, 0xa4, 0xbd, 0xb8, 0x9a, 0x79, 0x39, 0x9d, 0xa4, 0x73, 0xda, 0x4c, 0xe5,
	0xd4, 0x0d, 0xa8, 0xff, 0x9f, 0x84, 0xd4, 0x81, 0x75, 0xc2, 0x0f, 0x33, 0x2f, 0xa7, 0xfd, 0x74,
	0x4e, 0xa9, 0xa1, 0x34, 0xd6, 0xc8, 0x2c, 0x54, 0xf9, 0xa5, 0xc0, 0x9e, 0x9c, 0x16, 0x1b, 0x44,
	0x6d, 0xfb, 0x93, 0x1b, 0xab, 0xe9, 0xc7, 0x88, 0x06, 0xe1, 0xbf, 0x1f, 0x31, 0xf9, 0xcf, 0xef,
	0xc1, 0xf4, 0xe2, 0x03, 0x34, 0x85, 0xd6, 0x28, 0x90, 0xe9, 0xa6, 0xfa, 0x45, 0x01, 0x8d, 0x1f,
	0xb1, 0x43, 0xfd, 0x91, 0xd3, 0xa3, 0xa7, 0xcc, 0x75, 0x6d, 0x8f, 0x98, 0x6c, 0xe8, 0xf4, 0xd0,
	0x06, 0xa8, 0x0d, 0x8c, 0xad, 0x56, 0x17, 0x37, 0x5a, 0x4d, 0x4b, 0xbf, 0xd4, 0xcf, 0x4d, 0x75,
	0x01, 0x6d, 0xc3, 0x7a, 0xb7, 0x8d, 0x1b, 0xa6, 0x9e, 0x24, 0x94, 0x98, 0xc0, 0xfa, 0x5b, 0x7d,
	0x96, 0xc8, 0xa0, 0x1d, 0xd8, 0x34, 0x74, 0xd3, 0x68, 0xea, 0x97, 0x33, 0x54, 0x16, 0x15, 0x61,
	0x6b, 0x2e, 0xd5, 0x51, 0x73, 0xd5, 0xaf, 0x0a, 0x6c, 0xc9, 0xde, 0xf8, 0x5a, 0x38, 0xdb, 0x02,
	0x24, 0xab, 0xad, 0x06, 0xc6, 0x3a, 0x56, 0x17, 0x90, 0x06, 0x1b, 0x09, 0x5c, 0x18, 0xc5, 0xaa,
	0x92, 0x62, 0x84, 0x53, 0xac, 0x66, 0x62, 0x0b, 0x09, 0x66, 0xe2, 0x07, 0xab, 0x59, 0xb4, 0x0b,
	0xdb, 0x09, 0x57, 0x12, 0x99, 0xab, 0x7e, 0x53, 0xa0, 0xc4, 0x3d, 0x8d, 0x47, 0x23, 0x6d, 0xb3,
	0x0c, 0x4f, 0xc6, 0xdd, 0x2e, 0x5a, 0x5d, 0xab, 0xdd, 0xb8, 0x7a, 0x27, 0x1a, 0xbc, 0xef, 0xea,
	0x1d, 0x93, 0x1b, 0xde, 0x83, 0xdd, 0x39, 0x8a, 0x26, 0xd6, 0xcf, 0xcd, 0xe6, 0xd9, 0x95, 0xaa,
	0xfc, 0x41, 0x60, 0xe8, 0xb8, 0x69, 0xe8, 0xa7, 0x71, 0xb8, 0x25, 0xd8, 0x99, 0x23, 0x38, 0xbf,
	0xe0, 0xdf, 0xcf, 0xbe, 0x7a, 0xf9, 0xe1, 0x45, 0xdf, 0x09, 0x6f, 0xa3, 0xeb, 0x5a, 0x8f, 0xb9,
	0x75, 0xe7, 0xee, 0xd6, 0x76, 0xdd, 0x5b, 0x42, 0xc6, 0x0f, 0xde, 0x11, 0x1f, 0x8f, 0xa3, 0x40,
	0x38, 0x8f, 0x5f, 0xcc, 0x9e, 0x1d, 0xd0, 0xa0, 0x2e, 0x9e, 0xbe, 0xeb, 0x25, 0xfe, 0xe8, 0x9d,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x55, 0x47, 0xc2, 0xcb, 0x07, 0x00, 0x00,
}
