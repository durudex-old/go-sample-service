/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"

	"github.com/durudex/go-sample-service/internal/config"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

// Postgres driver interface.
type Postgres interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// Postgres repository structure.
type PostgresRepository struct{ Sample }

// Creating a new postgres repository.
func NewPostgresRepository(cfg config.PostgresConfig) *PostgresRepository {
	log.Debug().Msg("Creating a new postgres repository")

	// Creating a new postgres client.
	client, err := NewClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create postgres client")
	}

	return &PostgresRepository{Sample: NewSampleRepository(client)}
}

// Creating a new postgres client.
func NewClient(cfg config.PostgresConfig) (Postgres, error) {
	log.Debug().Msg("Creating a new postgres client")

	// Parsing database url.
	config, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, err
	}

	// Create a new pool connections by config.
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Ping a database connection.
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
