package common

type Id string
type Status bool
type Date string
type UpdatedBy string
type UpdatedAt Date
type CreatedBy string
type CreatedAt Date
type DeletedBy string
type DeletedAt Date
type MoneyValue float64

type CommonDTO struct {
	Id        Id        `json:"id"`
	Status    Status    `json:"status"`
	UpdatedAt UpdatedAt `json:"updated_at"`
	UpdatedBy UpdatedBy `json:"updated_by"`
	CreatedAt CreatedAt `json:"created_at"`
	CreatedBy CreatedBy `json:"created_by"`
	DeletedAt DeletedAt `json:"deleted_at,omitempty"`
	DeletedBy DeletedBy `json:"deleted_by,omitempty"`
}
