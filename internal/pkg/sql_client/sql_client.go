package sql_client

import "gorm.io/gorm"

type ListProject interface {
	List(owner string) []string
}

type StoreProject interface {
	Store(name, owner string) error
}

type projectModel struct {
	gorm.Model
	Name  string
	Owner string
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *projectRepository {
	db.AutoMigrate(&projectModel{})
	return &projectRepository{db}
}

func (p *projectRepository) Store(name, owner string) error {
	return p.db.Create(&projectModel{Name: name, Owner: owner}).Error
}

func (p *projectRepository) List(owner string) []string {
	var projects []projectModel
	res := p.db.Find(&projects, "owner = ?", owner)
	if res == nil {
		return []string{}
	}

	if res.RowsAffected == 0 {
		return []string{}
	}

	allProjects := []string{}

	for _, p := range projects {
		allProjects = append(allProjects, p.Name)
	}

	return allProjects
}
