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

package grpc

import (
	"net"

	"github.com/durudex/go-sample-service/internal/config"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// gRPC server structure.
type Server struct {
	server *grpc.Server
	config config.ServerConfig
}

// Creating a new gRPC server.
func NewServer(cfg config.ServerConfig) *Server {
	return &Server{server: grpc.NewServer(getOptions(cfg.TLS)), config: cfg}
}

// Running gRPC server.
func (s *Server) Run() {
	log.Info().Msg("Running gRPC server...")

	address := s.config.Host + ":" + s.config.Port

	// Creating a new TCP listener.
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating tcp listener")
	}

	// Running gRPC server.
	if err := s.server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("error running gRPC server")
	}
}

// Stoping gRPC server.
func (s *Server) Stop() {
	log.Info().Msg("Stopping gRPC server...")

	s.server.Stop()
}
