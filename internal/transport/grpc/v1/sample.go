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

package v1

import (
	"context"

	"github.com/durudex/go-sample-service/internal/service"
	v1 "github.com/durudex/go-sample-service/pkg/pb/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sample gRPC server handler.
type SampleHandler struct {
	service service.Sample
	v1.UnimplementedSampleServiceServer
}

// Creating a new sample gRPC handler.
func NewSampleHandler(service service.Sample) *SampleHandler {
	return &SampleHandler{service: service}
}

// Creating a new sample element.
func (h *SampleHandler) CreateElement(ctx context.Context, input *v1.CreateElementRequest) (*v1.CreateElementResponse, error) {
	id, err := h.service.Create(ctx, input.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.CreateElementResponse{Id: int64(id)}, nil
}

// Deleting a sample element.
func (h *SampleHandler) DeleteElement(ctx context.Context, input *v1.DeleteElementRequest) (*v1.DeleteElementResponse, error) {
	err := h.service.Delete(ctx, int(input.Id))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.DeleteElementResponse{}, nil
}
