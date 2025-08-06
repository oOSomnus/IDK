package repo

import (
	"I_Dev_Kit/internal/entity"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *entity.Project) error
	FindById(id int64) (*entity.Project, error)
	FindAll() ([]entity.Project, error)
	Update(project *entity.Project) error
	Delete(id int64) error
	GetByPage(page int) ([]entity.Project, error)
}

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) ProjectRepository {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(project *entity.Project) error {
	result := r.db.Create(project)
	return result.Error
}

func (r *projectRepo) FindById(id int64) (*entity.Project, error) {
	var project entity.Project
	result := r.db.Preload("Features").Preload("Features.Decisions").Preload("Features.Diagrams").Preload("Features.Todos").First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (r *projectRepo) FindAll() ([]entity.Project, error) {
	var projects []entity.Project
	result := r.db.Find(&projects)
	return projects, result.Error
}

func (r *projectRepo) Update(project *entity.Project) error {
	result := r.db.Save(project)
	return result.Error
}

func (r *projectRepo) Delete(id int64) error {
	return r.db.Unscoped().Delete(&entity.Project{ID: id}).Error
}

func (r *projectRepo) GetByPage(page int) ([]entity.Project, error) {
	var projects []entity.Project
	result := r.db.Order("updated_at DESC").Limit(10).Offset((page - 1) * 10).Find(&projects)
	return projects, result.Error
}
