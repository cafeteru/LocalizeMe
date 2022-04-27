package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
)

type BaseStringService interface {
	Create(request domain.BaseString) (domain.BaseString, error)
}
