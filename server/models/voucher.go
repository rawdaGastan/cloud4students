// Package models for database models
package models

// Voucher struct holds data of vouchers
type Voucher struct {
	ID      int    `json:"id" gorm:"unique" binding:"required"`
	Voucher string `json:"voucher" gorm:"unique" binding:"required"`
	VMs     int    `json:"vms" binding:"required"`
	K8s     int    `json:"k8s" binding:"required"`
}