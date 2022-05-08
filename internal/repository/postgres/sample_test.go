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

package postgres_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/pashagolub/pgxmock"

	"github.com/durudex/go-sample-service/internal/repository/postgres"
)

// Testing creating sample element.
func TestSampleRepository_Create(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ text string }

	// Test bahavior.
	type mockBehavior func(args args, id int)

	// Creating a new repository.
	repos := postgres.NewSampleRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         int
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{text: "test message"},
			want: 1,
			mockBehavior: func(args args, id int) {
				query := fmt.Sprintf(`INSERT INTO "%s"`, postgres.UserTable)
				mock.ExpectQuery(query).
					WithArgs(args.text).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(id))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Create sample element in postgres database.
			got, err := repos.Create(context.Background(), tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating sample element: %s", err.Error())
			}

			// Check for similarity of id.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error id are not similar")
			}
		})
	}
}

// Testing deleting sample element.
func TestSampleRepository_Delete(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ id int }

	// Test bahavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewSampleRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{id: 1},
			mockBehavior: func(args args) {
				query := fmt.Sprintf(`DELETE FROM "%s"`, postgres.UserTable)
				mock.ExpectExec(query).
					WithArgs(args.id).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Delete sample element in postgres database.
			err := repos.Delete(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error deleting sample element: %s", err.Error())
			}
		})
	}
}
