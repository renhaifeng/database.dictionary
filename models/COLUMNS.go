package models

import (
	"database.dictionary/dao"
)

type Columns struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	Extra         string `gorm:"column:EXTRA"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}

func GetAllByTableName(tableName string, tableSchema string) (columnList []*Columns, err error) {
	if err = dao.DB.Where("table_name=?", tableName).Where("table_schema=?", tableSchema).Find(&columnList).Error; err != nil {
		return nil, err
	}
	return
}
