package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/expertmaksud/go-clean-archi/helpers"
	"github.com/expertmaksud/go-clean-archi/services/interfaces"
)

//ProducrController ..
type ProducrController struct {
	service interfaces.IProductService
}

//NewProductController facorty...
func NewProductController(productService interfaces.IProductService) *ProducrController {
	return &ProducrController{service: productService}
}

//GetByTenetID return all products by tenent id
func (controller *ProducrController) GetByTenetID(res http.ResponseWriter, req *http.Request) {

	httpHelper := &helpers.HTTPHelper{}

	tenantID, _ := strconv.Atoi(chi.URLParam(req, "tenantId"))

	products, err := controller.service.GetAllProductsByTenant(tenantID)

	if err != nil {
		httpHelper.RespondwithJSON(res, http.StatusInternalServerError, "Server Error")
	}

	httpHelper.RespondwithJSON(res, http.StatusOK, products)

}
