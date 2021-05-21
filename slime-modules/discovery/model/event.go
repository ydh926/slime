package model

import "github.com/gogo/protobuf/proto"

type EventType int

const (
	Add = iota
	Delete
	Update
)

type Meta struct {
	Name      string
	Namespace string
	TypeUrl   string
}

type Event struct {
	EventType
	Meta
	Version int64
	Message proto.Message
}
