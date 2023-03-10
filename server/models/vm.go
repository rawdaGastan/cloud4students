// Package models for database models
package models

// VM struct for vms data
type VM struct {
	ID                int    `json:"id" gorm:"primaryKey"`
	UserID            string `json:"user_id"`
	Name              string `json:"name"`
	IP                string `json:"ip"`
	Resources         string `json:"resources"`
	SRU               uint64 `json:"sru"`
	CRU               uint64 `json:"cru"`
	MRU               uint64 `json:"mru"`
	ContractID        uint64 `json:"contractID"`
	NetworkContractID uint64 `json:"networkContractID"`
}
