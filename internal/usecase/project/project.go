package project

import "I_Dev_Kit/internal/repo"

type UseCase struct {
	repo repo.ProjectRepository
}

func New(r repo.ProjectRepository) *UseCase {
	return &UseCase{repo: r}
}
