package model

import "time"

type AppVersion struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	VersionNumber int       `json:"version_number"`
	Version       string    `json:"version"`
}
