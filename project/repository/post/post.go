package post

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alochym01/web-w-golang/project/models"
	"github.com/alochym01/web-w-golang/project/repository"
)

// sqlsqlPostRepo ...
type sqlPostRepo struct {
	Conn *sql.DB
}

// NewSQLPostRepo ...
func NewSQLPostRepo(Conn *sql.DB) repository.PostRepo {
	return &sqlPostRepo{
		Conn: Conn,
	}
}

func (m *sqlPostRepo) Fetch(ctx context.Context, numberRecord int64) ([]*models.Post, error) {
	// p := make([]*models.Post, 0)
	// gorm
	fmt.Println("Get all posts")
	query := "Select id, title, content From posts limit ?"
	fmt.Println("Get all posts")
	return m.fetch(ctx, query, numberRecord)
	// return p, nil
}

func (m *sqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	p := &models.Post{}
	query := "Select id, title, content From posts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(rows) > 0 {
		p = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	fmt.Println("Get a posts by id")
	return p, nil
}

func (m *sqlPostRepo) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "Insert posts SET title=?, content=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *sqlPostRepo) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	fmt.Println("Update a posts")
	return p, nil
}

func (m *sqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	fmt.Println("Delete post by id")
	return true, nil
}

func (m *sqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
