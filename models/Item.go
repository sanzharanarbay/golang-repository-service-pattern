package models

import "time"

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" binding:"required`
	Cost        float64   `json:"cost" binding:"required`
	Description string    `json:"description" binding:"required`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
