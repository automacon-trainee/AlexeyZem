package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"project/internal/library/models"
	"project/internal/myerror"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresDB {
	return &PostgresDB{
		db: db,
	}
}

func (p *PostgresDB) Create(ctx context.Context, book models.Book) error {
	str := `
INSERT INTO book (title, author, count)
VALUES ($1, $2, $3)
`

	_, err := p.db.ExecContext(ctx, str, book.Title, book.Author, book.Count)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) GetAll(ctx context.Context) ([]models.Book, error) {
	str := `SELECT title, author, count FROM book`
	rows, err := p.db.QueryContext(ctx, str)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Count); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (p *PostgresDB) TakeBook(ctx context.Context, id int) (models.Book, error) {
	str := `SELECT title, author, count FROM book WHERE id = $1`
	row := p.db.QueryRowContext(ctx, str, id)
	var book models.Book
	if err := row.Scan(&book.Title, &book.Author, &book.Count); err != nil {
		return models.Book{}, err
	}

	if book.Count == 0 {
		return models.Book{}, myerror.ErrNotBook
	}

	str = `UPDATE book SET count = count - 1 WHERE id = $1`
	_, err := p.db.ExecContext(ctx, str, id)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (p *PostgresDB) ReturnBook(ctx context.Context, id int) error {
	str := `UPDATE book SET count = count + 1 WHERE id = $1`
	res, err := p.db.ExecContext(ctx, str, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return myerror.ErrNotBook
	}

	return nil
}

func (p *PostgresDB) GetByID(ctx context.Context, id int) (models.Book, error) {
	str := `SELECT title, author, count FROM book WHERE id = $1`
	row := p.db.QueryRowContext(ctx, str, id)
	var book models.Book
	if err := row.Scan(&book.Title, &book.Author, &book.Count); err != nil {
		return models.Book{}, err
	}

	if book.Count == 0 {
		return models.Book{}, myerror.ErrNotBook
	}

	return book, nil
}

func (p *PostgresDB) CreateNewBookTable(ctx context.Context) error {
	newTableString := `CREATE TABLE IF NOT EXISTS book (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		count integer NOT NULL
	);`

	_, err := p.db.ExecContext(ctx, newTableString)
	if err != nil {
		return err
	}

	return nil
}
