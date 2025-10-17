package main

import (
	"context"
	"fmt"
	pb "go_grpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	fmt.Println("Creating a new person...")
	createReq := &pb.CreatePersonRequest{
		Name:        "Hussain Al-Shammari",
		Email:       "hussain@gmail.com",
		PhoneNumber: "123-456-7890",
	}
	createRes, err := client.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Error during Create: %v", err)
	}
	fmt.Printf("Person created: %+v\n", createRes)

	fmt.Println("Reading the person by ID...")
	readReq := &pb.SinglePersonRequest{
		Id: createRes.GetId(),
	}
	readRes, err := client.Read(ctx, readReq)
	if err != nil {
		log.Fatalf("Error during Read: %v", err)
	}
	fmt.Printf("Person details: %+v\n", readRes)

	fmt.Println("Updating the person's details...")
	updateReq := &pb.UpdatePersonRequest{
		Id:          createRes.GetId(),
		Name:        "Not Hussain",
		Email:       "newemail@gmail.com",
		PhoneNumber: "987-654-3210",
	}
	updateRes, err := client.Update(ctx, updateReq)
	if err != nil {
		log.Fatalf("Error during Update: %v", err)
	}
	fmt.Printf("Update response: %s\n", updateRes.GetResponse())

	fmt.Println("Deleting the person by ID...")
	deleteReq := &pb.SinglePersonRequest{
		Id: createRes.GetId(),
	}
	deleteRes, err := client.Delete(ctx, deleteReq)
	if err != nil {
		log.Fatalf("Error during Delete: %v", err)
	}
	fmt.Printf("Delete response: %s\n", deleteRes.GetResponse())
}