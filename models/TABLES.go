package models

import (
	"database.dictionary/dao"
)

type Tables struct {
	TableName    string `gorm:"column:TABLE_NAME"`
	TableComment string `gorm:"column:TABLE_COMMENT"`
}

func GetAllByTableSchema(tableSchema string) (tableList []*Tables, err error) {
	if err = dao.DB.Where("table_schema=?", tableSchema).Find(&tableList).Error; err != nil {
		return nil, err
	}
	return
}
