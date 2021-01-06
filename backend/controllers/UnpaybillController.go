package controllers

import (
	"context"
	"strconv"
	"fmt"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/unpaybill"
	"github.com/team10/app/ent/treatment"
	"github.com/gin-gonic/gin"
)


// UnpaybillController defines the struct for the unpaybill controller
type UnpaybillController struct {
	client *ent.Client
	router gin.IRouter
}

// Unpaybill defines the struct for the Unpaybill entity
type Unpaybill struct {
	Status string
	Treatment int
}
// CreateUnpaybill handles POST requests for adding unpaybill entities
// @Summary Create unpaybill
// @Description Create unpaybill
// @ID create-unpaybill
// @Accept   json
// @Produce  json
// @Param unpaybill body ent.Unpaybill true "Unpaybill entity"
// @Success 200 {object} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills [post]
func (ctl *UnpaybillController) CreateUnpaybill(c *gin.Context) {
	obj := Unpaybill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "unpaybill binding failed",
		})
		return
	}
	t, err := ctl.client.Treatment.
		Query().
		Where(treatment.IDEQ(int(obj.Treatment))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}
	u, err := ctl.client.Unpaybill.
		Create().
		SetStatus("Unpaid").
		SetTreatment(t).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetUnpaybill handles GET requests to retrieve a unpaybill entity
// @Summary Get a unpaybill entity by ID
// @Description get unpaybill by ID
// @ID get-unpaybill
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Success 200 {object} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [get]
func (ctl *UnpaybillController) GetUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Unpaybill.
		Query().
		Where(unpaybill.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUnpaybill handles request to get a list of unpaybill entities
// @Summary List unpaybill entities
// @Description list unpaybill entities
// @ID list-unpaybill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills [get]
func (ctl *UnpaybillController) ListUnpaybill(c *gin.Context) {
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

	unpaybills, err := ctl.client.Unpaybill.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, unpaybills)
}

// DeleteUnpaybill handles DELETE requests to delete a unpaybill entity
// @Summary Delete a unpaybill entity by ID
// @Description get unpaybill by ID
// @ID delete-unpaybill
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [delete]
func (ctl *UnpaybillController) DeleteUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Unpaybill.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateUnpaybill handles PUT requests to update a unpaybill entity
// @Summary Update a unpaybill entity by ID
// @Description update unpaybill by ID
// @ID update-unpaybill
// @Accept   json
// @Produce  json
// @Param id path int true "Unpaybill ID"
// @Param unpaybill body ent.Unpaybill true "Unpaybill entity"
// @Success 200 {object} ent.Unpaybill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /unpaybills/{id} [put]
func (ctl *UnpaybillController) UpdateUnpaybill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Unpaybill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "unpaybill binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Unpaybill.
		UpdateOneID(int(id)).
		SetStatus("Paid").
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewUnpaybillController creates and registers handles for the unpaybill controller
func NewUnpaybillController(router gin.IRouter, client *ent.Client) *UnpaybillController {
	uc := &UnpaybillController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *UnpaybillController) register() {
	unpaybills := ctl.router.Group("/unpaybills")

	unpaybills.GET("", ctl.ListUnpaybill)

	// CRUD
	unpaybills.POST("", ctl.CreateUnpaybill)
	unpaybills.GET(":id", ctl.GetUnpaybill)
	unpaybills.PUT(":id", ctl.UpdateUnpaybill)
	unpaybills.DELETE(":id", ctl.DeleteUnpaybill)
}