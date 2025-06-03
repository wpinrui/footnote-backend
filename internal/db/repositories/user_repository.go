package repositories

import (
	"database/sql"
	"footnote-backend/internal/api/models"
)

type UserRepositoryInterface interface {
	Create(*models.User) (int, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) (int, error) {
	query := `
		INSERT INTO users (email, hashed_password)
		VALUES ($1, $2)
		RETURNING id
	`
	var id int
	err := ur.db.QueryRow(query, user.Email, user.HashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, hashed_password, is_email_verified, date_created, date_updated
		FROM users
		WHERE email = $1
	`
	user := &models.User{}
	err := ur.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.Email,
		&user.HashedPassword,
		&user.IsEmailVerified,
		&user.DateCreated,
		&user.DateUpdated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err // Other error
	}
	return user, nil
}
