package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nnnewb/dt/internal/client"
	"github.com/nnnewb/dt/internal/dmcli"
	"github.com/nnnewb/dt/internal/svc/bank"
)

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
			cli1 := client.NewBankClient("http://bank1/")
			_, err := cli1.TransIn(c, &bank.TransInReq{ID: req.ToID, Amount: req.Amount})
			if err != nil {
				return err
			}

			cli2 := client.NewBankClient("http://bank2/")
			_, err = cli2.TransIn(c, &bank.TransInReq{ID: req.ToID, Amount: req.Amount})
			if err != nil {
				return err
			}

			return nil
		})

		c.JSONP(200, map[string]interface{}{
			"code":    0,
			"message": "ok",
		})
	})
	r.Run(":5000")
}
