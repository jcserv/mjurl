package url

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jcserv/mjurl/model"
)

type PSQLStore struct {
	dbpool *pgxpool.Pool
}

func NewPSQLStore(dbpool *pgxpool.Pool) *PSQLStore {
	return &PSQLStore{
		dbpool,
	}
}

// CreateURL inserts a new url into the PSQL database
func (s *PSQLStore) CreateURL(ctx context.Context, url *model.URL) error {
	conn, err := s.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	statement, err := conn.Conn().Prepare(ctx, "createURL", "INSERT INTO url (short_url, long_url) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = conn.Conn().Exec(ctx, statement.Name, url.Short, url.Long)
	return nil, err
}

// QueryURL queries the database for a URL with the given ID.
func (s *PSQLStore) QueryURL(ctx context.Context, id model.URLID) (*model.URL, error) {
	conn, err := s.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Conn().Prepare(ctx, "queryURL", "SELECT id, short_url, long_url FROM url WHERE id = $1")
	if err != nil {
		return nil, err
	}

	return rowToURL(conn.QueryRow(ctx, "queryURL", id))

}

// QueryURLByShort queries the database for a URL with the given short URL.
func (s *PSQLStore) QueryURLByShort(ctx context.Context, short model.ShortURL) (*model.URL, error) {
	conn, err := s.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Conn().Prepare(ctx, "queryURLByShort", "SELECT id, short_url, long_url FROM url WHERE short_url = $1")
	if err != nil {
		return nil, err
	}
	return rowToURL(conn.QueryRow(ctx, "queryURLByShort", short))
}

// QueryURLByLong queries the database for a URL with the given long URL.
func (s *PSQLStore) QueryURLByLong(ctx context.Context, long model.LongURL) (*model.URL, error) {
	conn, err := s.dbpool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	_, err = conn.Conn().Prepare(ctx, "queryURLByLong", "SELECT id, short_url, long_url FROM url WHERE long_url = $1")
	if err != nil {
		return nil, err
	}

	return rowToURL(conn.QueryRow(ctx, "queryURLByLong", long))
}

func rowToURL(row pgx.Row) (*model.URL, error) {
	var url model.URL
	err := row.Scan(&url.ID, &url.Short, &url.Long)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
