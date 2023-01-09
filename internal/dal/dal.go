package dal

import (
	"fgzs-single/internal/errorx"
	"gorm.io/gorm"
)

type Sort struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}

func SqlErrCheck(err error) error {
	if err != nil && err != gorm.ErrRecordNotFound {
		return errorx.DataSqlErr.WithDetail(err)
	}
	return nil
}
