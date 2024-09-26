package repository

import (
	"context"
	"database/sql"

	"project/internal/myerror"
	"project/internal/profile/models"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

func (p *PostgresDB) Create(ctx context.Context, profile models.Profile) error {
	str := `
INSERT INTO profile (name, lastname)
VALUES ($1, $2)
`

	_, err := p.db.ExecContext(ctx, str, profile.Name, profile.Lastname)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) GetProfile(ctx context.Context, id int) (*models.Profile, []int, error) {
	str := `SELECT name, lastname FROM profile WHERE id = $1`

	row := p.db.QueryRowContext(ctx, str, id)
	if row == nil {
		return nil, nil, myerror.ErrUserNotFound
	}

	var profile models.Profile
	err := row.Scan(&profile.Name, &profile.Lastname)
	if err != nil {
		return nil, nil, err
	}

	str = `SELECT book_id FROM profile_book WHERE profile_id = $1`

	rows, err := p.db.QueryContext(ctx, str, id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var booksID []int
	for rows.Next() {
		var bookId int
		if err := rows.Scan(&bookId); err != nil {
			return nil, nil, err
		}

		booksID = append(booksID, bookId)

	}

	return &profile, booksID, nil
}

func (p *PostgresDB) TakeBook(ctx context.Context, profileID, bookID int) error {
	str := `INSERT INTO profile_book (profile_id, book_id)
VALUES ($1, $2)
`
	_, err := p.db.ExecContext(ctx, str, profileID, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) ReturnBook(ctx context.Context, profileID, bookID int) error {
	str := `DELETE FROM profile_book WHERE profile_id = $1 AND book_id = $2`

	_, err := p.db.ExecContext(ctx, str, profileID, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) CreateTableProfile(ctx context.Context) error {
	str := `CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    name TEXT,
    lastname TEXT,
)`
	_, err := p.db.ExecContext(ctx, str)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) CreateProfileBookTable(ctx context.Context) error {
	str := `CREATE TABLE IF NOT EXISTS profile_book (
    profile_id INTEGER,
    book_id INTEGER
)`
	_, err := p.db.ExecContext(ctx, str)
	if err != nil {
		return err
	}

	return nil
}
