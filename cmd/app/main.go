package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nnnewb/dt/internal/dmcli"
	"github.com/nnnewb/dt/internal/svc/bank"
)

type GeneralResp struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type TransferReq struct {
	ID     int64 `json:"id,omitempty"`
	Bank   int32 `json:"bank,omitempty"`
	Amount int64 `json:"amount,omitempty"`
	ToID   int64 `json:"to_id,omitempty"`
	ToBank int32 `json:"to_bank,omitempty"`
}

func main() {
	r := gin.Default()
	r.POST("/v1alpha1/transfer", func(c *gin.Context) {
		req := &TransferReq{}
		c.BindJSON(req)
		dmcli.GlobalTx(c, func(gid string) error {
			transInReq := &bank.TransInReq{ID: req.ToID, Amount: req.Amount}
			if err := TransIn(c, gid, transInReq); err != nil {
				return err
			}

			transOutReq := &bank.TransOutReq{ID: req.ID, Amount: req.Amount}
			if err := TransOut(c, gid, transOutReq); err != nil {
				return err
			}

			return nil
		})
	})
	r.Run()
}

func TransIn(c context.Context, gid string, payload *bank.TransInReq) error {
	transInPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(c, "POST", "bank2:5000/v1alpha1/TransIn", bytes.NewReader(transInPayload))
	if err != nil {
		return err
	}

	req.Header.Add("X-DM-GID", gid)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	generalResp := &GeneralResp{}
	err = json.Unmarshal(data, generalResp)
	if err != nil {
		return err
	}

	if generalResp.Code != 0 {
		return fmt.Errorf("error code %d, %s", generalResp.Code, generalResp.Message)
	}

	return nil
}

func TransOut(c context.Context, gid string, payload *bank.TransOutReq) error {
	transOutPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(c, "POST", "bank2:5000/v1alpha1/TransOut", bytes.NewReader(transOutPayload))
	if err != nil {
		return err
	}

	req.Header.Add("X-DM-GID", gid)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	generalResp := &GeneralResp{}
	err = json.Unmarshal(data, generalResp)
	if err != nil {
		return err
	}

	if generalResp.Code != 0 {
		return fmt.Errorf("error code %d, %s", generalResp.Code, generalResp.Message)
	}

	return nil
}
