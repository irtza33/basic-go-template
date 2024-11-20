package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type sqlClient struct {
	conn *pgx.Conn
}

func newSqlClient(ctx context.Context, username, password, host , db string) (sqlClient, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, host, db)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return sqlClient{}, err
	}

	return sqlClient{conn: conn}, nil
}