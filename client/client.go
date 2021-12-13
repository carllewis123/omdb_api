package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "omdb_api/models"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50001"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFilmManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/*uncomment one of the client requests below*/

	// Get List Film
	// r, err := c.SearchByParamGrpc(ctx, &pb.RequestParameter{ApiKey: "faf7e5bb", Page: "2", Title: "Batman"})
	// if err != nil {
	// 	log.Fatalf("could not create user: %v", err)
	// }
	// fmt.Printf("%+v", r)

	//Get List Detail Film
	r, err := c.GetDetailFilmGrpc(ctx, &pb.RequestParameter{ApiKey: "faf7e5bb", ImdbID: "tt4853102"})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	fmt.Printf("%+v", r)
}
