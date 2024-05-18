package entity

import (
	"github.com/augustomcosta/client-api/internal/client/domain/vo"
	"github.com/google/uuid"
)

type Client struct {
	ID         uuid.UUID     `json:"id"`
	Name       string        `json:"name"`
	Age        int           `json:"age"`
	SocialInfo vo.SocialInfo `json:"social_info"`
}

func NewClient(name string, age int, socialInfo vo.SocialInfo) (*Client, error) {
	return &Client{
		ID:         uuid.New(),
		Name:       name,
		Age:        age,
		SocialInfo: socialInfo,
	}, nil
}
