// Package models for database models
package models

// Quota struct holds available vms && k8s for each user
type Quota struct {
	UserID string `json:"user_id"`
	Vms    int    `json:"vms"`
	K8s    int    `json:"k8s"`
}
