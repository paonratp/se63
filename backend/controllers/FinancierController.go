package controllers

import (
	"context"
	"strconv"

	"github.com/team10/app/ent"
	"github.com/team10/app/ent/financier"
	"github.com/gin-gonic/gin"
)

// FinancierController defines the struct for the financier controller
type FinancierController struct {
	client *ent.Client
	router gin.IRouter
}

// GetFinancier handles GET requests to retrieve a financier entity
// @Summary Get a financier entity by ID
// @Description get financier by ID
// @ID get-financier
// @Produce  json
// @Param id path int true "Financier ID"
// @Success 200 {object} ent.Financier
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /financier/{id} [get]
func (ctl *FinancierController) GetFinancier(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Financier.
		Query().
		WithUser().
		Where(financier.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListFinancier handles request to get a list of financier entities
// @Summary List financier entities
// @Description list financier entities
// @ID list-financier
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Financier
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /financier [get]
func (ctl *FinancierController) ListFinancier(c *gin.Context) {
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

	financiers, err := ctl.client.Financier.
		Query().
		WithUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, financiers)
}
// NewFinancierController creates and registers handles for the financier controller
func NewFinancierController(router gin.IRouter, client *ent.Client) *FinancierController {
	uc := &FinancierController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *FinancierController) register() {
	financiers := ctl.router.Group("/financiers")

	financiers.GET("", ctl.ListFinancier)
	financiers.GET(":id", ctl.GetFinancier)

}