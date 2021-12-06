package dm

import (
	"context"

	"github.com/nnnewb/dt/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type DMService struct {
	pb.UnimplementedDMServer
	DB *gorm.DB
}

func (dms *DMService) CreateGlobalTransaction(_ context.Context, _ *pb.CreateGlobalTransactionReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (dms *DMService) CreateSubTransaction(_ context.Context, _ *pb.CreateSubTransactionReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (dms *DMService) CommitGlobalTransaction(_ context.Context, _ *pb.CommitGlobalTransactionReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (dms *DMService) RollbackGlobalTransaction(_ context.Context, _ *pb.RollbackGlobalTransactionReq) (*emptypb.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (dms *DMService) mustEmbedUnimplementedDMServer() {
	panic("not implemented") // TODO: Implement
}
