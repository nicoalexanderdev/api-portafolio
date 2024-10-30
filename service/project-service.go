package service

import "github.com/nicoalexanderdev/api-portafolio/entity"

type ProjectService interface {
	Save(entity.Project) entity.Project
	FindAll() []entity.Project
}

type projectServiceImpl struct {
	projects []entity.Project
}

func New() ProjectService {
	return &projectServiceImpl{}
}

func (service *projectServiceImpl) Save(project entity.Project) entity.Project {
	service.projects = append(service.projects, project)
	return project
}

func (service *projectServiceImpl) FindAll() []entity.Project {
	return service.projects
}
