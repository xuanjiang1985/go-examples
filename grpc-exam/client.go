package main

import (
    "log"
    "os"

    pb "grpc-exam/hello"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

const (
    address     = "localhost:50051"
	defaultName = "world"
	dage = 32
)

func main() {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)

    // Contact the server and print out its response.
	name := defaultName
	age := dage
    if len(os.Args) > 1 {
        name = os.Args[1]
    }
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name, Age: int64(age),})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s, %d", r.Message, r.Fuck)
}