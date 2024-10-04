package rule

import "github.com/01luisfonseca/seligo/src/common"

type RuleInputDTO struct {
	PolicyId   common.Id `json:"policy_id"`
	Conditions string    `json:"conditions"`
	Events     string    `json:"events"`
}

type RuleRegistryDTO struct {
	common.CommonDTO
	RuleInputDTO
}

type RuleInterface interface {
	GetRules() []RuleRegistryDTO
	GetRule(id string) RuleRegistryDTO
	CreateRule(rule RuleInputDTO) RuleRegistryDTO
	UpdateRule(id string, rule RuleInputDTO) RuleRegistryDTO
	DeleteRule(id string) RuleRegistryDTO
}
