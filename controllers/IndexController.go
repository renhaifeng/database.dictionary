package controllers

import (
	"database.dictionary/models"
	"database.dictionary/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	*models.Tables
	ColumnList []*models.Columns
}

func IndexHandler(c *gin.Context) {

	tableSchema := setting.Conf.WatchDb

	tableList, _ := models.GetAllByTableSchema(tableSchema)

	var rsp []Item

	for _, table := range tableList {

		columnList, _ := models.GetAllByTableName(table.TableName, tableSchema)

		rsp = append(rsp, Item{table, columnList})

	}

	c.HTML(http.StatusOK, "index.html", rsp)
}
