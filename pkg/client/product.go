package client

import (
	"context"
	"fmt"

	"github.com/ab3llo/go-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func NewProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(id string) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: id,
	}
	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DeacreaseStock(id string, orderId string, quantity int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:       id,
		OrderId:  orderId,
		Quantity: quantity,
	}
	return c.Client.DecreaseStock(context.Background(), req)
}
