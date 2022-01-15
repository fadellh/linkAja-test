package account

import (
	"link-test/api/account/request"
	"link-test/api/account/response"
	"link-test/api/common"
	"link-test/business/account"

	"github.com/labstack/echo"
)

type Controller struct {
	service account.Service
}

func NewController(service account.Service) *Controller {
	return &Controller{
		service,
	}
}

func (ctr *Controller) FindBalanceByAccNo(c echo.Context) error {

	accNo := c.Param("number")

	accounts, err := ctr.service.FindBalanceByAccNo(accNo)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAccountResponse(*accounts)

	return c.JSON(common.NewSuccessResponse(response))
}

func (ctr *Controller) TransBalance(c echo.Context) error {

	accNo := c.Param("from_number")
	tr := new(request.TransferRequest)
	if err := c.Bind(tr); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	at := account.TransferRequest{
		FromAccNo: accNo,
		ToAccNo:   tr.ToAccNumber,
		Amount:    tr.Amount,
	}
	err := ctr.service.TransBalance(at)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseCreated())
}
