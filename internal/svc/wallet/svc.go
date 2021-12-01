package wallet

import (
	"context"

	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WalletService struct {
	pb.UnimplementedWalletServiceServer
}

func (w *WalletService) ListWallets(_ context.Context, _ *pb.ListWalletsRequest) (*pb.ListWalletsResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WalletService) GetWallet(_ context.Context, _ *pb.GetWalletRequest) (*pb.Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WalletService) CreateWallet(_ context.Context, _ *pb.CreateWalletRequest) (*pb.Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WalletService) UpdateWallet(_ context.Context, _ *pb.UpdateWalletRequest) (*pb.Wallet, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WalletService) DeleteWallet(_ context.Context, _ *pb.DeleteWalletRequest) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}
