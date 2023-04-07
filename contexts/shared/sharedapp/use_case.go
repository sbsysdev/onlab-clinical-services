package sharedapp

import "context"

type CommandUseCase interface {
	Command(interface{}) error
}

type QueryUseCase interface {
	Query(interface{}) (interface{}, error)
}

type EventHandler interface {
	Handle(context.Context)
}
