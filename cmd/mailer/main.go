package main

import (
	"database/sql"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"kannon.gyozatech.dev/generated/pb"
)

func main() {
	logrus.SetLevel(log.DebugLevel)
	runGrpcServer()
}

func runGrpcServer() error {
	godotenv.Load()

	dbi, err := sql.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		panic(err)
	}

	mailerService, err := newMailerService(dbi)
	if err != nil {
		return err
	}

	log.Info("😃 Open TCP Connection")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	defer lis.Close()

	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterMailerServer(s, mailerService)

	log.Printf("🚀 starting gRPC... Listening on %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
