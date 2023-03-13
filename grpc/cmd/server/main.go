package main

import (
	"database/sql"
	"net"

	//"github.com/kotayz/grpc/internal/database"
	//"github.com/kotayz/grpc/internal/pb"
	//"github.com/kotayz/grpc/internal/service"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlie3", "file:./db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	refletion.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
