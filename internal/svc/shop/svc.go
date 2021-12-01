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
	session := o.DB.Session(&gorm.Session{Context: ctx})
	orders := make([]models.Order, 0)
	offset, err := strconv.ParseUint(req.GetPageToken(), 10, 64)
	if err != nil {
		log.Printf("invalid page token %s", req.GetPageToken())
	}

	result := session.Offset(int(offset)).Limit(int(req.GetPageSize())).Find(&orders)
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
			Name:       order.Name,
			CreateTime: &timestamppb.Timestamp{Seconds: order.CreatedAt.Unix()},
			UpdateTime: &timestamppb.Timestamp{Seconds: order.UpdatedAt.Unix()},
		})
	}

	return resp, nil
}

func (o *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	session := o.DB.Session(&gorm.Session{Context: ctx})
	order := &models.Order{}
	result := session.Where("name=?", req.GetName()).Take(order)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		log.Printf("find order failed, error %v", result.Error)
		return nil, status.Errorf(codes.Internal, "find order failed")
	}

	return &pb.Order{
		Name: order.Name,
		CreateTime: &timestamppb.Timestamp{
			Seconds: order.CreatedAt.Unix(),
		},
		UpdateTime: &timestamppb.Timestamp{
			Seconds: order.UpdatedAt.Unix(),
		},
	}, nil
}

func (o *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	session := o.DB.Session(&gorm.Session{Context: ctx})
	err := session.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&models.Order{Name: req.GetOrder().GetName()})
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		log.Printf("transaction failed, error %v", err)
		return nil, status.Error(codes.Internal, "create order failed")
	}

	return &pb.Order{
		Name: req.GetOrder().GetName(),
	}, nil
}

func (o *OrderService) UpdateOrder(_ context.Context, req *pb.UpdateOrderRequest) (*pb.Order, error) {
	order := &models.Order{}
	result := o.DB.Where("name=?", req.GetOrder().GetName()).Take(order)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		log.Printf("find order failed, error %v", result.Error)
		return nil, status.Errorf(codes.Internal, "find order failed")
	}

	return &pb.Order{
		Name: order.Name,
		CreateTime: &timestamppb.Timestamp{
			Seconds: order.CreatedAt.Unix(),
		},
		UpdateTime: &timestamppb.Timestamp{
			Seconds: order.UpdatedAt.Unix(),
		},
	}, nil
}

func (o *OrderService) DeleteOrder(_ context.Context, req *pb.DeleteOrderRequest) (*emptypb.Empty, error) {
	result := o.DB.Delete(&models.Order{Name: req.GetName()})
	if result.Error != nil {
		log.Printf("delete order failed, error %v", result.Error)
		return nil, status.Error(codes.Internal, "delete order failed")
	}
	return &emptypb.Empty{}, nil
}
