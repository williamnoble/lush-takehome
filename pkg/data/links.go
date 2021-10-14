package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Link struct {
	Id       int64  `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

type Models struct {
	Links LinkModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Links: LinkModel{DB: db},
	}
}

type LinkModel struct {
	DB *sql.DB
}

// GetLink retrieves a Long URL given a short URL.
func (l LinkModel) GetLink(url string) (Link, error) {
	query := `SELECT short_url, long_url FROM links WHERE short_url=$1`
	var link Link
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := l.DB.QueryRowContext(ctx, query, url).Scan(&link.ShortURL, &link.LongURL)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Link{}, ErrRecordNotFound
		default:
			return Link{}, err
		}
	}
	return link, nil
}

// CreateLink saves a Link (with Long & Short URL) to the underlying database.
func (l LinkModel) CreateLink(link Link) (*Link, error) {
	query := "INSERT INTO links (short_url, long_url) VALUES ($1, $2) RETURNING (id)"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := l.DB.QueryRowContext(ctx, query, link.ShortURL, link.LongURL).Scan(&link.Id)
	return &link, err
}
