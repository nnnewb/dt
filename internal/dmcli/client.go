package dmcli

import (
	"context"
	"database/sql"

	"github.com/nnnewb/dt/internal/svc/dm"
)

func GlobalTx(ctx context.Context, fn func(gid string) error) error {
	var gid string
	// TODO 生成GID
	CreateGlobalTransaction(ctx, &dm.CreateGlobalTxReq{GID: gid})

	// 执行业务代码
	// TODO 根据结果决定提交还是回滚
	if err := fn(gid); err != nil {
		return err
	}

	return nil
}

func LocalTx(ctx context.Context, db *sql.DB, fn func(db *sql.DB) error) error {
	return nil
}
