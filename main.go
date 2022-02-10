package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/DymaSV/shippy-consignment-server/proto/consignment"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/v2/client"
)

const (
	address  = "localhost:50051"
	filename = "consignment.json"
)

func main() {
	service := micro.NewService(micro.Name("shippy.cli.consignment"))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", client.NewClient())

	file := filename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Success)

	getAll, err := client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	if getAll.Success {
		for _, v := range getAll.Consignments {
			log.Println(v)
		}
	}
}

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Cannot read file %v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		fmt.Printf("Cannot read json %v", err)
		return nil, err
	}
	return consignment, err
}
