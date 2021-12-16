package client

import (
	"context"
	"net/http"

	"github.com/nnnewb/dt/internal/svc/tm"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type TMClient struct {
	BaseClient
}

func NewTMClient(baseUrl string) *TMClient {
	return &TMClient{
		BaseClient{
			BaseUrl: baseUrl,
			HTTPClient: &http.Client{
				Transport: otelhttp.NewTransport(http.DefaultTransport),
			},
		},
	}
}

func (d *TMClient) CreateGlobalTx(ctx context.Context, payload *tm.CreateGlobalTxReq) (resp *tm.CreateGlobalTxResp, err error) {
	resp = new(tm.CreateGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/create_global_tx", payload, resp)
	return
}

func (d *TMClient) RegisterLocalTx(ctx context.Context, payload *tm.RegisterLocalTxReq) (resp *tm.RegisterLocalTxResp, err error) {
	resp = new(tm.RegisterLocalTxResp)
	err = d.post(ctx, "/v1alpha1/register_local_tx", payload, resp)
	return
}

func (d *TMClient) CommitGlobalTx(ctx context.Context, payload *tm.CommitGlobalTxReq) (resp *tm.CommitGlobalTxResp, err error) {
	resp = new(tm.CommitGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/commit_global_tx", payload, resp)
	return
}

func (d *TMClient) RollbackGlobalTx(ctx context.Context, payload *tm.RollbackGlobalTxReq) (resp *tm.RollbackGlobalTxResp, err error) {
	resp = new(tm.RollbackGlobalTxResp)
	err = d.post(ctx, "/v1alpha1/rollback_global_tx", payload, resp)
	return
}
