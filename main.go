package main

import (
	"context"
	"log"
	"time"

	pb "./futurebattlegrounds"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	log.Printf("Running...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBattlegroundsClient(conn)

	// Contact the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Spawn a ship
	x, err := c.SpawnShip(ctx, &pb.ShipSpawnRequest{IFF: "go-bot"})
	if err != nil {
		log.Fatalf("could not spawn ship: %v", err)
	}
	log.Printf("spawned ship: %s", x)

	r, err := c.GetBattleground(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get battleground: %v", err)
	}
	log.Printf("battleground: %s", r)
}
