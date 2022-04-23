package services

import (
	"context"
	"net/http"

	"github.com/ab3llo/go-order-svc/pkg/client"
	"github.com/ab3llo/go-order-svc/pkg/db"
	"github.com/ab3llo/go-order-svc/pkg/models"
	"github.com/ab3llo/go-order-svc/pkg/pb"
	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	DbConnection db.DatabaseConnection
	ProductSvc   client.ProductServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Not enough in stock"}, nil
	}

	order := models.Order{
		Id:        uuid.New().String(),
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	s.DbConnection.DB.Create(&order)

	res, err := s.ProductSvc.DeacreaseStock(req.ProductId, order.Id, req.Quantity)

	if err != nil {

	} else if res.Status == http.StatusConflict {
		s.DbConnection.DB.Delete(&models.Order{}, order.Id)
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
