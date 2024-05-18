package usecase

import (
	"context"
	applicationsvc "github.com/augustomcosta/client-api/internal/client/application/service"
	"github.com/augustomcosta/client-api/internal/client/domain/entity"
	"github.com/augustomcosta/client-api/internal/client/domain/repository"
	"github.com/augustomcosta/client-api/internal/client/domain/vo"
	"time"
)

type CreateClientInputDTO struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	Facebook  string `json:"facebook"`
}

type CreateClientOutputDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateClientUseCase struct {
	repository repository.ClientRepository
}

func NewCreateClientUseCase(repository repository.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{
		repository: repository,
	}
}

func (uc *CreateClientUseCase) Execute(ctx context.Context, input CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	createdSocialInfo, err := vo.NewSocialInfo(input.Instagram, input.Linkedin, input.Facebook)
	if err != nil {
		return nil, err
	}

	createdClient, err := entity.NewClient(input.Name, input.Age, *createdSocialInfo)
	if err != nil {
		return nil, err
	}

	err = applicationsvc.VerifyJusbrasilHistory(createdClient)
	if err != nil {
		return nil, err
	}

	err = uc.repository.Save(ctx, createdClient)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        createdClient.ID.String(),
		CreatedAt: time.Now(),
	}, nil

}
