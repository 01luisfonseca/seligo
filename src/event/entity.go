package event

import "github.com/01luisfonseca/seligo/src/common"

type EventInputDTO struct {
	Payload string `json:"payload"`
}

type EventRegistryDTO struct {
	common.CommonDTO
	EventInputDTO
}

type EventInterface interface {
	GetEvents() []EventRegistryDTO
	GetEvent(id string) EventRegistryDTO
	CreateEvent(event EventInputDTO) EventRegistryDTO
	UpdateEvent(id string, event EventInputDTO) EventRegistryDTO
	DeleteEvent(id string) EventRegistryDTO
}
