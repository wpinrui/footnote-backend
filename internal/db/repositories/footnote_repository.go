package repositories

import (
	"database/sql"
	"footnote-backend/internal/api/models"
)

type FootnoteRepository struct {
	DB *sql.DB
}

func NewFootnoteRepository(db *sql.DB) *FootnoteRepository {
	return &FootnoteRepository{DB: db}
}

func (fr *FootnoteRepository) Create(f *models.Footnote) (int, error) {
	var id int
	err := fr.DB.QueryRow(`
		INSERT INTO footnotes (uid, content, day)
		VALUES ($1, $2, $3)
		RETURNING id
	`, f.UserId, f.Content, f.Day).Scan(&id)
	return id, err
}

func (fr *FootnoteRepository) ListByUser(userId int) ([]*models.Footnote, error) {
	rows, err := fr.DB.Query(`
		SELECT id, uid, content, day
		FROM footnotes
		WHERE uid = $1
		ORDER BY id DESC
	`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var footnotes []*models.Footnote = make([]*models.Footnote, 0) // ! Prevents null return
	for rows.Next() {
		var f models.Footnote
		if err := rows.Scan(&f.Id, &f.UserId, &f.Content); err != nil {
			return nil, err
		}
		footnotes = append(footnotes, &f)
	}
	return footnotes, nil
}

func (fr *FootnoteRepository) GetByID(id, userId int) (*models.Footnote, error) {
	var f models.Footnote
	err := fr.DB.QueryRow(`
		SELECT id, uid, content, day
		FROM footnotes
		WHERE id = $1 AND uid = $2
	`, id, userId).Scan(&f.Id, &f.UserId, &f.Content)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &f, err
}

func (fr *FootnoteRepository) Update(id, userId int, content string) error {
	_, err := fr.DB.Exec(`
		UPDATE footnotes
		SET content = $1
		WHERE id = $2 AND uid = $3
	`, content, id, userId)
	return err
}

func (fr *FootnoteRepository) Delete(id, userId int) error {
	_, err := fr.DB.Exec(`
		DELETE FROM footnotes
		WHERE id = $1 AND uid = $2
	`, id, userId)
	return err
}

func (fr *FootnoteRepository) Search(userId int, query string) ([]*models.Footnote, error) {
	rows, err := fr.DB.Query(`
		SELECT id, uid, content, day
		FROM footnotes
		WHERE uid = $1 AND content ILIKE '%' || $2 || '%'
		ORDER BY id DESC
	`, userId, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Footnote
	for rows.Next() {
		var f models.Footnote
		if err := rows.Scan(&f.Id, &f.UserId, &f.Content); err != nil {
			return nil, err
		}
		results = append(results, &f)
	}
	return results, nil
}
