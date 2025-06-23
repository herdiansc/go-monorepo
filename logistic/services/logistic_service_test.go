package services

import (
	"errors"
	"log"
	"net/url"
	"testing"

	"github.com/herdiansc/orderfaz/auth/models"
)

type mockLogisticFinder struct {
	a  models.Logistic
	e  error
	l  []models.Logistic
	le error
}

func (m mockLogisticFinder) FindByUUID(uuid string) (models.Logistic, error) {
	return m.a, m.e
}

func (m mockLogisticFinder) List(filter map[string]interface{}) ([]models.Logistic, error) {
	return m.l, m.le
}

var (
	mockSuccessLogisticFinder = mockLogisticFinder{
		a: models.Logistic{},
		e: nil,
	}
	mockFailedLogisticFinder = mockLogisticFinder{
		a: models.Logistic{},
		e: errors.New("error"),
	}
)

func TestLogisticServices_GetLogisticByUUID(t *testing.T) {
	cases := []struct {
		name string
		repo mockLogisticFinder
		want int
	}{
		{
			name: "Positive",
			repo: mockSuccessLogisticFinder,
			want: 200,
		},
		{
			name: "Negative: Failed find data",
			repo: mockFailedLogisticFinder,
			want: 404,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewLogisticServices(tt.repo)
			code, _ := svc.GetLogisticByUUID("123")
			if code != tt.want {
				t.Errorf("Expected resp to be %q but it was %q", tt.want, code)
			}
		})
	}
}

var (
	mockSuccessListLogisticFinder = mockLogisticFinder{
		l:  []models.Logistic{{LogisticName: "JNE"}},
		le: nil,
	}
	mockFailedListLogisticFinder = mockLogisticFinder{
		l:  []models.Logistic{},
		le: errors.New("error"),
	}
)

func TestLogisticServices_GetLogistic(t *testing.T) {
	cases := []struct {
		name string
		repo mockLogisticFinder
		want int
	}{
		{
			name: "Positive",
			repo: mockSuccessListLogisticFinder,
			want: 200,
		},
		{
			name: "Negative: Failed find data",
			repo: mockFailedListLogisticFinder,
			want: 404,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewLogisticServices(tt.repo)
			code, _ := svc.ListLogistics(url.Values{})
			log.Printf("code: %+v, want: %+v", code, tt.want)
			if code != tt.want {
				t.Errorf("Expected resp to be %q but it was %q", tt.want, code)
			}
		})
	}
}
