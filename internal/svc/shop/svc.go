package shop

import (
	"context"
	"log"
	"strconv"

	"github.com/nnnewb/dt/pkg/models"
	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	DB *gorm.DB
}

func (o *OrderService) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders := make([]models.Order, 0)
	offset, err := strconv.ParseUint(req.GetPageToken(), 10, 64)
	if err != nil {
		log.Printf("invalid page token %s", req.GetPageToken())
	}

	result := o.DB.Offset(int(offset)).Limit(int(req.GetPageSize())).Find(&orders)
	if result.Error != nil {
		log.Printf("find orders failed, error %v", result.Error)
		return nil, status.Errorf(codes.Internal, "find orders failed")
	}

	offset += uint64(result.RowsAffected)
	resp := &pb.ListOrdersResponse{
		Orders:        []*pb.Order{},
		NextPageToken: strconv.FormatUint(offset, 10),
	}

	for _, order := range orders {
		resp.Orders = append(resp.Orders, &pb.Order{
			Name: order.Name,
			CreateTime: &timestamppb.Timestamp{
				Seconds: order.CreatedAt.Unix(),
			},
			UpdateTime: &timestamppb.Timestamp{
				Seconds: order.UpdatedAt.Unix(),
			},
		})
	}

	return resp, nil
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
