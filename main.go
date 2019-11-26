package main

import (
	"context"
	"log"
	"math"
	"time"

	pb "./futurebattlegrounds"
	v "github.com/ungerik/go3d/float64/vec2"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// North := v.T{0, -1}
	// East := v.T{1, 0}

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

	var uuid string
	uuid = x.UUID
	log.Printf("spawned ship: %s", uuid)

	r, err := c.GetBattleground(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get battleground: %v", err)
	}

	log.Printf("battleground: %d ships", len(r.Ships))
	for i := 0; i < len(r.Ships); i++ {
		if r.Ships[i].UUID == uuid {
			var ship pb.Ship
			var rotation v.T
			var position v.T

			ship = *r.Ships[i]
			rotation = v.T{ship.RotationVector.X, ship.RotationVector.Y}
			position = v.T{ship.Position.X, ship.Position.Y}

			log.Printf("ship[%d]: %s %f %f %f", i, ship.Position, (rotation.Angle()*180/math.Pi)+90, ship.Hull, ship.Battery)
			_ = position
		}
	}
}
