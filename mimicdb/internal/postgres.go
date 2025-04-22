package internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func GetPGConnection(dns string) (*pgconn.PgConn, error) {
	c, err := GetPGConfig(false)
	if err != nil {
		return nil, err
	}

	con, err := pgx.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return con.PgConn(), nil
}

func GetPGConfig(ssl bool) (*pgx.ConnConfig, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s")
	if ssl {
		connStr += "?sslmode=required"
	}

	c, err := pgx.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	if ssl {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(""))
		c.TLSConfig = &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            caCertPool,
		}
	}

	return c, nil
}
