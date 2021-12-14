package dmcli

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nnnewb/dt/internal/svc/dm"
)

func CreateGlobalTransaction(ctx context.Context, payload *dm.CreateGlobalTxReq) error {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "dm:5000/v1alpha1/create-global-transaction", bytes.NewReader(marshaled))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func RegisterLocalTransaction(ctx context.Context, gid string, payload *dm.RegisterLocalTxReq) error {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "dm:5000/v1alpha1/register-local-transaction", bytes.NewReader(marshaled))
	if err != nil {
		return err
	}

	req.Header.Add("X-DM-GID", gid)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil

}
