package model

import (
	"istio.io/istio-mcp/pkg/model"
)

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
	model.Config
	Version int64
}
