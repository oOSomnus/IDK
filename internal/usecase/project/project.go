package project

import (
	"I_Dev_Kit/internal/entity"
	"I_Dev_Kit/internal/repo"
	"fmt"
)

type UseCase struct {
	repo repo.ProjectRepository
}

func New(r repo.ProjectRepository) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) GetProjectsByPage(page int) ([]entity.Project, error) {
	projects, err := u.repo.GetByPage(page)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects by page: %w", err)
	}
	return projects, nil
}
