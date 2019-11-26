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

func pv(v v.T) *v.T {
	return &v
}

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

	var ship *pb.Ship
	for i := 0; i < len(r.Ships); i++ {
		if r.Ships[i].UUID == uuid {
			ship = r.Ships[i]
		}
	}

	if ship != nil {
		var rotation v.T
		var position v.T
		rotation = v.T{ship.RotationVector.X, ship.RotationVector.Y}
		position = v.T{ship.Position.X, ship.Position.Y}

		log.Printf("ship: %s %f %f %f", ship.Position, (rotation.Angle()*180/math.Pi)+90, ship.Hull, ship.Battery)
		_ = position

		var closestShip *pb.Ship
		for i := 0; i < len(r.Ships); i++ {
			if r.Ships[i] != ship {
				if closestShip == nil {
					closestShip = r.Ships[i]
				} else {
					len := (pv(v.T{r.Ships[i].Position.X, r.Ships[i].Position.Y})).Sub(&position).Length()
					if len < (pv(v.T{closestShip.Position.X, closestShip.Position.Y})).Sub(&position).Length() {
						closestShip = r.Ships[i]
					}
				}
			}
		}
		if closestShip != nil {
			log.Printf("Closest: %s", closestShip.Position)
			closestShipPosRel := position.Scale(1).Sub(pv(v.T{closestShip.Position.X, closestShip.Position.Y}))
			var relRot int
			relRot = 0
			if v.IsLeftWinding(&rotation, closestShipPosRel) {
				log.Printf("turn left!")
				relRot = -1
			} else {
				log.Printf("turn right!")
				relRot = 1
			}
			_ = relRot

			x, err := c.SetActionState(ctx, &pb.ShipActionStateRequest{UUID: uuid, Thrust: 0, Rotate: float32(relRot), Shooting: false})
			if err != nil {
				log.Fatalf("fail: %v", err)
			}
			log.Printf("Command success? %t", x.OK)
		}
	}
}
