package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
)

// PaytypeController defines the struct for the paytype controller
type PaytypeController struct {
	client *ent.Client
	router gin.IRouter
}

// ListPaytype handles request to get a list of paytype entities
// @Summary List paytype entities
// @Description list paytype entities
// @ID list-paytype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Paytype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /paytype [get]
func (ctl *PaytypeController) ListPaytype(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	paytypes, err := ctl.client.Paytype.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, paytypes)
}

// NewPaytypeController creates and registers handles for the paytype controller
func NewPaytypeController(router gin.IRouter, client *ent.Client) *PaytypeController {
	uc := &PaytypeController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *PaytypeController) register() {
	paytypes := ctl.router.Group("/paytypes")

	paytypes.GET("", ctl.ListPaytype)

}
