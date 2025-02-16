package core

import (
	"time"
)

type SystemRegistration struct {
	Address            string            `json:"address"`
	AuthenticationInfo string            `json:"authenticationInfo"`
	Metadata           map[string]string `json:"metadata"`
	Port               int               `json:"port"`
	SystemName         string            `json:"systemName"`
}

type SystemsResponse struct {
	Systems []System `json:"data"`
	Count   int      `json:"count"`
}

type System struct {
	ID                 int               `json:"id"`
	SystemName         string            `json:"systemName"`
	Address            string            `json:"address"`
	Port               int               `json:"port"`
	AuthenticationInfo string            `json:"authenticationInfo,omitempty"`
	CreatedAt          time.Time         `json:"createdAt,omitempty"`
	UpdatedAt          time.Time         `json:"updatedAt,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}
