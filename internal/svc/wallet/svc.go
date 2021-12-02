package wallet

import (
	"context"
	"log"
	"path"
	"strconv"

	"github.com/nnnewb/dt/pkg/models"
	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type WalletService struct {
	pb.UnimplementedWalletServiceServer
	DB *gorm.DB
}

func (w *WalletService) ListWallets(ctx context.Context, req *pb.ListWalletsRequest) (*pb.ListWalletsResponse, error) {
	session := w.DB.Session(&gorm.Session{Context: ctx})
	wallets := make([]models.Wallet, 0)
	offset, err := strconv.ParseUint(req.GetPageToken(), 10, 64)
	if err != nil {
		log.Printf("invalid page token %s", req.GetPageToken())
	}

	result := session.Offset(int(offset)).Limit(int(req.GetPageSize())).Find(&wallets)
	if result.Error != nil {
		log.Printf("find wallets failed, error %v", result.Error)
		return nil, status.Errorf(codes.Internal, "find wallets failed")
	}

	offset += uint64(result.RowsAffected)
	resp := &pb.ListWalletsResponse{
		Wallets:       []*pb.Wallet{},
		NextPageToken: strconv.FormatUint(offset, 10),
	}

	for _, wallet := range wallets {
		resp.Wallets = append(resp.Wallets, &pb.Wallet{
			Name:       wallet.Name,
			Balance:    uint64(wallet.Balance),
			CreateTime: &timestamppb.Timestamp{Seconds: wallet.CreatedAt.Unix()},
			UpdateTime: &timestamppb.Timestamp{Seconds: wallet.UpdatedAt.Unix()},
		})
	}

	return resp, nil
}

func (w *WalletService) GetWallet(ctx context.Context, req *pb.GetWalletRequest) (*pb.Wallet, error) {
	id := path.Base(req.GetName())
	session := w.DB.Session(&gorm.Session{Context: ctx})
	wallet := &models.Wallet{}
	result := session.Take(wallet, &models.Wallet{Name: id})
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "wallet not found")
		}
		log.Printf("find wallet failed, error %v", result.Error)
		return nil, status.Errorf(codes.Internal, "find wallet failed")
	}

	return &pb.Wallet{
		Name:    wallet.Name,
		Balance: uint64(wallet.Balance),
		CreateTime: &timestamppb.Timestamp{
			Seconds: wallet.CreatedAt.Unix(),
		},
		UpdateTime: &timestamppb.Timestamp{
			Seconds: wallet.UpdatedAt.Unix(),
		},
	}, nil
}

func (w *WalletService) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.Wallet, error) {
	session := w.DB.Session(&gorm.Session{Context: ctx})
	err := session.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&models.Wallet{
			Name:    req.GetWallet().GetName(),
			Balance: int64(req.GetWallet().Balance),
		})
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		log.Printf("transaction failed, error %v", err)
		return nil, status.Error(codes.Internal, "create wallet failed")
	}

	return req.GetWallet(), nil
}

func (w *WalletService) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.Wallet, error) {
	id := path.Base(req.GetWallet().GetName())
	m := &models.Wallet{}
	session := w.DB.Session(&gorm.Session{Context: ctx})
	err := session.Transaction(func(tx *gorm.DB) error {
		result := session.
			Where(&models.Wallet{Name: id}).
			Updates(&models.Wallet{Balance: int64(req.GetWallet().GetBalance())})
		if result.Error != nil {
			return result.Error
		}

		result = session.Take(m, &models.Wallet{Name: id})
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "update wallet failed")
	}

	return &pb.Wallet{
		Name:    m.Name,
		Balance: uint64(m.Balance),
		CreateTime: &timestamppb.Timestamp{
			Seconds: m.CreatedAt.Unix(),
		},
		UpdateTime: &timestamppb.Timestamp{
			Seconds: m.UpdatedAt.Unix(),
		},
	}, nil
}

func (w *WalletService) DeleteWallet(ctx context.Context, req *pb.DeleteWalletRequest) (*emptypb.Empty, error) {
	id := path.Base(req.GetName())
	session := w.DB.Session(&gorm.Session{Context: ctx})

	result := session.Delete(&models.Wallet{Name: id})
	if result.Error != nil {
		log.Printf("delete wallet failed, error %v", result.Error)
		return nil, status.Error(codes.Internal, "delete wallet failed")
	}
	return &emptypb.Empty{}, nil
}
