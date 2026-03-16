package service

import (
	"github.com/InkSha/Resummer/internal/service/model"
)

func ListModels() []any {
	return []any{
		&model.User{},
	}
}
