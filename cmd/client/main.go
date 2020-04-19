package main

import (
	"Test_gRPC/pkg/api"
	"context"
	"flag"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewAdderClient(conn)
	res, err := client.Add(
		context.Background(),
		&api.AddRequest{X: int32(x), Y: int32(y)},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
