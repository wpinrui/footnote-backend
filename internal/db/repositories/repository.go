package repositories

import (
	"database/sql"
	"footnote-backend/internal/api/models"
)

type Repositories struct {
	UserRepository     *UserRepository
	FootnoteRepository *FootnoteRepository
}

type FootnoteRepositoryInterface interface {
	Create(*models.Footnote) (int, error)
	ListByUser(userId int) ([]*models.Footnote, error)
	GetByID(id, userId int) (*models.Footnote, error)
	Update(id, userId int, content string) error
	Delete(id, userId int) error
	Search(userId int, query string) ([]*models.Footnote, error)
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository:     NewUserRepository(db),
		FootnoteRepository: NewFootnoteRepository(db),
	}
}
