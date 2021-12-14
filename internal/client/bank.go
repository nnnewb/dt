package client

import (
	"context"
	"net/http"

	"github.com/nnnewb/dt/internal/svc/bank"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type BankClient struct {
	BaseClient
}

func NewBankClient(baseUrl string) *BankClient {
	return &BankClient{
		BaseClient{
			BaseUrl: baseUrl,
			HTTPClient: &http.Client{
				Transport: otelhttp.NewTransport(http.DefaultTransport),
			},
		},
	}
}

func (b *BankClient) TransIn(ctx context.Context, payload *bank.TransInReq) (resp *bank.TransInResp, err error) {
	resp = new(bank.TransInResp)
	err = b.post(ctx, "/v1alpha1/trans_in", payload, resp)
	return
}

func (b *BankClient) TransOut(ctx context.Context, payload *bank.TransOutReq) (resp *bank.TransOutResp, err error) {
	resp = new(bank.TransOutResp)
	err = b.post(ctx, "/v1alpha1/trans_out", payload, resp)
	return
}
