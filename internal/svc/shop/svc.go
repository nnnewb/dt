package shop

import (
	"context"

	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct{}

func (o *OrderService) ListOrders(_ context.Context, _ *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) GetOrder(_ context.Context, _ *pb.GetOrderRequest) (*pb.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) CreateOrder(_ context.Context, _ *pb.CreateOrderRequest) (*pb.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) UpdateOrder(_ context.Context, _ *pb.UpdateOrderRequest) (*pb.Order, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) DeleteOrder(_ context.Context, _ *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (o *OrderService) mustEmbedUnimplementedOrderServiceServer() {
	panic("not implemented") // TODO: Implement
}
