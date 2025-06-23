package services

import (
	"log"
	"net/http"
	"net/url"

	"github.com/herdiansc/orderfaz/auth/models"
)

// LogisticFinder defines logistic function
type LogisticFinder interface {
	FindByUUID(uuid string) (models.Logistic, error)
	List(filter map[string]interface{}) ([]models.Logistic, error)
}

// LogisticServices defines logistic services struct
type LogisticServices struct {
	repo LogisticFinder
}

// NewLogisticServices inits LogisticServices
func NewLogisticServices(ac LogisticFinder) LogisticServices {
	return LogisticServices{
		repo: ac,
	}
}

// GetLogisticByUUID gets logistic data by uuid
func (svc LogisticServices) GetLogisticByUUID(uuid string) (int, models.Response) {
	data, err := svc.repo.FindByUUID(uuid)
	if err != nil {
		log.Printf("Failed to get data: %+v\n", err.Error())
		return http.StatusNotFound, models.Response{Message: "Internal server error", Data: err.Error()}
	}

	return http.StatusOK, models.Response{Message: "ok", Data: data}
}

// ListLogistics lists all logistic by query string
func (svc LogisticServices) ListLogistics(q url.Values) (int, models.Response) {
	filter := make(map[string]interface{})
	for k, v := range q {
		filter[k] = v[0]
	}
	data, _ := svc.repo.List(filter)
	if len(data) == 0 {
		log.Printf("Failed to get data")
		return http.StatusNotFound, models.Response{Message: "Not found", Data: nil}
	}

	return http.StatusOK, models.Response{Message: "ok", Data: data}
}
