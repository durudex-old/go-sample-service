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

package service

import (
	"context"

	"github.com/durudex/go-sample-service/internal/repository/postgres"
)

// Sample service structure.
type Sample interface {
	Create(ctx context.Context, text string) (int, error)
	Delete(ctx context.Context, id int) error
}

// Sample service structure.
type SampleService struct{ repos postgres.Sample }

// Creating a new sample service.
func NewSampleService(repos postgres.Sample) *SampleService {
	return &SampleService{repos: repos}
}

// Creating a new sample element.
func (s *SampleService) Create(ctx context.Context, text string) (int, error) {
	return s.repos.Create(ctx, text)
}

// Deleting a sample element.
func (s *SampleService) Delete(ctx context.Context, id int) error {
	return s.repos.Delete(ctx, id)
}
