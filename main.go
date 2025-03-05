package main

import (
	"context"
	"net"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	db "UsersManagement/db/sqlc"
	"UsersManagement/gapi"
	pbv "UsersManagement/pb/session"
	pbu "UsersManagement/pb/users"
	"UsersManagement/util"
)

func main() {
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Error().Err(err).Msg("app.env is not found")
		os.Exit(1)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to db")
		os.Exit(1)
	}

	initMigration(config.MigrationURL, config.DBSource)

	defer conn.Close()

	store := db.NewStore(conn)

	server, err := gapi.ServerSetup(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create server")
		os.Exit(1)
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	pbv.RegisterSessionServer(grpcServer, server)
	pbu.RegisterUserServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.UsersManagementPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
		os.Exit(1)
	}

	log.Info().Msgf("Connect to grpc server at %s", listener.Addr().String())

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal().Err(err).Msg("Failed to connecting server")
		os.Exit(1)
	}
}

func initMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration database source")
		os.Exit(1)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("Fail to Up migrate database")
		os.Exit(1)
	}

	log.Info().Msg("Successfully created migration database")
}
