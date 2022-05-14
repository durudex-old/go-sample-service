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
	"fmt"

	"github.com/durudex/go-sample-service/pkg/database/postgres"
)

// Sample table name.
const SampleTable string = "sample"

// Sample repository interface.
type Sample interface {
	Create(ctx context.Context, text string) (int, error)
	Delete(ctx context.Context, id int) error
}

// Sample repository structure.
type SampleRepository struct{ psql postgres.Postgres }

// Creating a new sample repository.
func NewSampleRepository(psql postgres.Postgres) *SampleRepository {
	return &SampleRepository{psql: psql}
}

// Create sample element in postgres database.
func (r *SampleRepository) Create(ctx context.Context, text string) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO "%s" (text) VALUES ($1) RETURNING "id"`, SampleTable)

	// Execute query.
	err := r.psql.QueryRow(ctx, query, text).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Delete sample element in postgres database.
func (r *SampleRepository) Delete(ctx context.Context, id int) error {
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE "id"=$1`, SampleTable)

	// Execute query.
	_, err := r.psql.Exec(ctx, query, id)

	return err
}
