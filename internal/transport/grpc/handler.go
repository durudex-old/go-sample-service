package grpc

import (
	"github.com/durudex/go-sample-service/internal/service"
	v1 "github.com/durudex/go-sample-service/internal/transport/grpc/v1"

	"google.golang.org/grpc"
)

// gRPC server handler structure.
type Handler struct{ service *service.Service }

// Creating a new gRPC handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// Registering gRPC version handlers.
func (h *Handler) RegisterHandlers(srv *grpc.Server) {
	v1.NewHandler(h.service).RegisterHandlers(srv)
}
