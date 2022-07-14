package controllers

import (
	"net/http/httptest"
	"testing"
)

func TestCreateLog(t *testing.T) {
	req, err := http.NewRequest("GET", "/logs", nil)

}
