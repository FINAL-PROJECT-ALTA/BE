package midtrans

import (
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateTransaction(core coreapi.Client, req *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {

	apiRes, err := core.ChargeTransaction(req)
	if err != nil {
		log.Warn("Payment Error")
	}
	return apiRes, nil
}
