package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/herdiansc/orderfaz/auth/respositories"
	"github.com/herdiansc/orderfaz/auth/services"
	"gorm.io/gorm"
)

type LogisticHandler struct {
	db *gorm.DB
}

func NewLogisticHandler(db *gorm.DB) LogisticHandler {
	return LogisticHandler{
		db: db,
	}
}

// List gets list
// @Summary		Add a new auth to database
// @Description	Add a new auth to database
// @Accept		json
// @Produce		json
// @Param Authorization header string true "With the bearer started"
// @Param       origin_name  query       string  false  "origin"
// @Param       destination_name  query       string  false  "destination"
// @Success		200		{object}	models.Response			"ok"
// @Failure		400		{object}    models.Response	"bad request"
// @Failure		500		{object}    models.Response	"internal server error"
// @Router		/logistics [get]
func (h LogisticHandler) List(w http.ResponseWriter, r *http.Request) {
	lf := respositories.NewLogisticRepository(h.db)
	svc := services.NewLogisticServices(lf)
	code, res := svc.ListLogistics(r.URL.Query())
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

// GetLogisticByUUID gets logistic
// @Summary		Add a new auth to database
// @Description	Add a new auth to database
// @Accept		json
// @Produce		json
// @Param Authorization header string true "With the bearer started"
// @Param       uuid    path        string  true  "UUID"
// @Success		200		{object}	models.Response			"ok"
// @Failure		400		{object}    models.Response	"bad request"
// @Failure		500		{object}    models.Response	"internal server error"
// @Router		/logistics/{uuid} [get]
func (h LogisticHandler) GetLogisticByUUID(w http.ResponseWriter, r *http.Request) {
	lf := respositories.NewLogisticRepository(h.db)
	svc := services.NewLogisticServices(lf)

	code, res := svc.GetLogisticByUUID(r.PathValue("uuid"))
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
