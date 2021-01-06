package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/symptomseverity"
)

// SymptomseverityController defines the struct for the symptomseverity controller
type SymptomseverityController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSymptomseverity handles POST requests for adding symptomseverity entities
// @Summary Create symptomseverity
// @Description Create symptomseverity
// @ID create-symptomseverity
// @Accept   json
// @Produce  json
// @Param symptomseverity body ent.Symptomseverity true "Symptomseverity entity"
// @Success 200 {object} ent.Symptomseverity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomseverity [post]
func (ctl *SymptomseverityController) CreateSymptomseverity(c *gin.Context) {
	obj := ent.Symptomseverity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "symptomseverity binding failed",
		})
		return
	}

	u, err := ctl.client.Symptomseverity.
		Create().
		SetSymptomseverity(obj.Symptomseverity).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetSymptomseverity handles GET requests to retrieve a symptomseverity entity
// @Summary Get a symptomseverity entity by ID
// @Description get symptomseverity by ID
// @ID get-symptomseverity
// @Produce  json
// @Param id path int true "Symptomseverity ID"
// @Success 200 {object} ent.Symptomseverity
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomseverity/{id} [get]
func (ctl *SymptomseverityController) GetSymptomseverity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Symptomseverity.
		Query().
		Where(symptomseverity.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSymptomseverity handles request to get a list of symptomseverity entities
// @Summary List symptomseverity entities
// @Description list symptomseverity entities
// @ID list-symptomseverity
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Symptomseverity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomseverity [get]
func (ctl *SymptomseverityController) ListSymptomseverity(c *gin.Context) {
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

	symptomseverity, err := ctl.client.Symptomseverity.
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

	c.JSON(200, symptomseverity)
}

// DeleteSymptomseverity handles DELETE requests to delete a symptomseverity entity
// @Summary Delete a symptomseverity entity by ID
// @Description get symptomseverity by ID
// @ID delete-symptomseverity
// @Produce  json
// @Param id path int true "Symptomseverity ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomseverity/{id} [delete]
func (ctl *SymptomseverityController) DeleteSymptomseverity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Symptomseverity.
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

// UpdateSymptomseverity handles PUT requests to update a symptomseverity entity
// @Summary Update a symptomseverity entity by ID
// @Description update symptomseverity by ID
// @ID update-symptomseverity
// @Accept   json
// @Produce  json
// @Param id path int true "Symptomseverity ID"
// @Param symptomseverity body ent.Symptomseverity true "Symptomseverity entity"
// @Success 200 {object} ent.Symptomseverity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /symptomseverity/{id} [put]
func (ctl *SymptomseverityController) UpdateSymptomseverity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Symptomseverity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "symptomseverity binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Symptomseverity.
		UpdateOneID(int(id)).
		SetSymptomseverity(obj.Symptomseverity).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewSymptomseverityController creates and registers handles for the symptomseverity controller
func NewSymptomseverityController(router gin.IRouter, client *ent.Client) *SymptomseverityController {
	uc := &SymptomseverityController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *SymptomseverityController) register() {
	symptomseverity := ctl.router.Group("/symptomseverity")

	symptomseverity.GET("", ctl.ListSymptomseverity)

	// CRUD
	symptomseverity.POST("", ctl.CreateSymptomseverity)
	symptomseverity.GET(":id", ctl.GetSymptomseverity)
	symptomseverity.PUT(":id", ctl.UpdateSymptomseverity)
	symptomseverity.DELETE(":id", ctl.DeleteSymptomseverity)
}
