package main

import (
	"context"
	"io"
	"log"
	"math"

	pb "github.com/jorisvddonk/future-battlegrounds-go-bot/futurebattlegrounds"
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

	// Establish a context (TODO: do this properly..)
	ctx := context.Background()

	// Spawn a ship
	x, err := c.SpawnShip(ctx, &pb.ShipSpawnRequest{IFF: "go-bot"})
	if err != nil {
		log.Fatalf("could not spawn ship: %v", err)
	}

	var uuid string
	uuid = x.UUID
	log.Printf("spawned ship: %s", uuid)

	stream, err := c.StreamBattleground(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not stream battleground: %v", err)
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("receiving battleground err: %v", err)
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
				closestShipPosRel := position.Scale(1).Sub(pv(v.T{closestShip.Position.X, closestShip.Position.Y}))
				var relRot int
				var thrust float32
				var shooting bool
				relRot = 0
				thrust = 0
				shooting = false

				if v.IsLeftWinding(&rotation, closestShipPosRel) {
					relRot = -1
				} else {
					relRot = 1
				}

				ang := v.Angle(closestShipPosRel, rotation.Scale(-1))
				if ang < 0.1 {
					shooting = true
				}

				distance := closestShipPosRel.Length()
				if ang < 0.3 {
					if distance > 200 {
						thrust = 1
					} else if distance < 100 {
						thrust = -1
					} else {
						thrust = 0
					}
				}

				x, err := c.SetActionState(ctx, &pb.ShipActionStateRequest{UUID: uuid, Thrust: thrust, Rotate: float32(relRot), Shooting: shooting})
				if err != nil {
					log.Fatalf("fail: %v", err)
				}
				log.Printf("Command success? %t", x.OK)
			}
		}
	}
}
