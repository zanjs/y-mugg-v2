package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v2/app/models"
	"github.com/zanjs/y-mugg-v2/app/services"
)

// SaleController is 销量记录
type SaleController struct {
	Controller
}

// GetAll is get all Sales
func (ctl SaleController) GetAll(c echo.Context) error {

	var (
		datas       []models.Sale
		queryparams models.QueryParams
		page        models.PageModel
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	datas, page, err = services.SaleServices{}.GetAll(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	dataAll := echo.Map{
		"data": datas,
		"page": page,
	}

	return ctl.ResponseSuccess(c, dataAll)
}

// GetAllWhereTime is get all Sales
func (ctl SaleController) GetAllWhereTime(c echo.Context) error {

	var (
		sales       []models.Sale
		queryparams models.QueryParams
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	sales, err = services.SaleServices{}.WhereTime(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	return ctl.ResponseSuccess(c, sales)
}
