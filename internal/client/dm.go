package client

import (
	"context"
	"net/http"

	"github.com/nnnewb/dt/internal/svc/dm"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type DMClient struct {
	BaseClient
}

func NewDMClient(baseUrl string) *DMClient {
	return &DMClient{
		BaseClient{
			BaseUrl: baseUrl,
			HTTPClient: &http.Client{
				Transport: otelhttp.NewTransport(http.DefaultTransport),
			},
		},
	}
}

func (d *DMClient) CreateGlobalTx(ctx context.Context, payload *dm.CreateGlobalTxReq) (resp *dm.CreateGlobalTxResp, err error) {
	resp = new(dm.CreateGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/create_global_tx", payload, resp)
	return
}

func (d *DMClient) RegisterLocalTx(ctx context.Context, payload *dm.RegisterLocalTxReq) (resp *dm.RegisterLocalTxResp, err error) {
	resp = new(dm.RegisterLocalTxResp)
	err = d.post(ctx, "/v1alpha1/register_local_tx", payload, resp)
	return
}

func (d *DMClient) CommitGlobalTx(ctx context.Context, payload *dm.CommitGlobalTxReq) (resp *dm.CommitGlobalTxResp, err error) {
	resp = new(dm.CommitGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/commit_global_tx", payload, resp)
	return
}

func (d *DMClient) RollbackGlobalTx(ctx context.Context, payload *dm.RollbackGlobalTxReq) (resp *dm.RollbackGlobalTxResp, err error) {
	resp = new(dm.RollbackGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/rollback_global_tx", payload, resp)
	return
}
