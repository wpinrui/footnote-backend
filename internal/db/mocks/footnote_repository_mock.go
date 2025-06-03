package mocks

import (
	"footnote-backend/internal/api/models"
	"footnote-backend/internal/db/repositories"

	"github.com/stretchr/testify/mock"
)

type FootnoteRepositoryMock struct {
	mock.Mock
}

func (m *FootnoteRepositoryMock) Create(f *models.Footnote) (int, error) {
	args := m.Called(f)
	return args.Int(0), args.Error(1)
}

func (m *FootnoteRepositoryMock) ListByUser(userId int) ([]*models.Footnote, error) {
	args := m.Called(userId)
	return args.Get(0).([]*models.Footnote), args.Error(1)
}

func (m *FootnoteRepositoryMock) GetByID(id, userId int) (*models.Footnote, error) {
	args := m.Called(id, userId)
	f := args.Get(0)
	if f == nil {
		return nil, args.Error(1)
	}
	return f.(*models.Footnote), args.Error(1)
}

func (m *FootnoteRepositoryMock) Update(id, userId int, content string) error {
	args := m.Called(id, userId, content)
	return args.Error(0)
}

func (m *FootnoteRepositoryMock) Delete(id, userId int) error {
	args := m.Called(id, userId)
	return args.Error(0)
}

func (m *FootnoteRepositoryMock) Search(userId int, query string) ([]*models.Footnote, error) {
	args := m.Called(userId, query)
	return args.Get(0).([]*models.Footnote), args.Error(1)
}

var _ repositories.FootnoteRepositoryInterface = (*FootnoteRepositoryMock)(nil)
