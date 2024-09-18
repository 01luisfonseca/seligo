package participant

import "github.com/01luisfonseca/seligo/entities/common"

type ParticipantInputDTO struct {
	ProjectId common.Id `json:"project_id"`
	UserId    common.Id `json:"user_id"`
	StartDate string    `json:"start_date,omitempty"`
}

type ParticipantRegistryDTO struct {
	common.CommonDTO
	ParticipantInputDTO
}

type ParticipantInterface interface {
	GetParticipants() []ParticipantRegistryDTO
	GetParticipant(id string) ParticipantRegistryDTO
	CreateParticipant(participant ParticipantInputDTO) ParticipantRegistryDTO
	UpdateParticipant(id string, participant ParticipantInputDTO) ParticipantRegistryDTO
	DeleteParticipant(id string) ParticipantRegistryDTO
}
