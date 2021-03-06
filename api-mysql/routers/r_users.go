package routers

import (
	"github.com/dimaskiddo/simple-go/api-mysql/models"
)

type ResponseGetUser struct {
	Status  bool          `json:"status"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}
