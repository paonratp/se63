package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/medicalrecordstaff"
)

// MedicalrecordstaffController defines the struct for the medicalrecordstaff controller
type MedicalrecordstaffController struct {
	client *ent.Client
	router gin.IRouter
}

// GetMedicalrecordstaff handles GET requests to retrieve a medicalrecordstaff entity
// @Summary Get a medicalrecordstaff entity by ID
// @Description get medicalrecordstaff by ID
// @ID get-medicalrecordstaff
// @Produce  json
// @Param id path int true "Medicalrecordstaff ID"
// @Success 200 {object} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Medicalrecordstaff/{id} [get]
func (ctl *MedicalrecordstaffController) GetMedicalrecordstaff(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Medicalrecordstaff.
		Query().
		Where(medicalrecordstaff.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMedicalrecordstaff handles request to get a list of medicalrecordstaff entities
// @Summary List medicalrecordstaff entities
// @Description list medicalrecordstaff entities
// @ID list-medicalrecordstaff
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicalrecordstaff
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalrecordstaff [get]
func (ctl *MedicalrecordstaffController) ListMedicalrecordstaff(c *gin.Context) {
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

	medicalrecordstaff, err := ctl.client.Medicalrecordstaff.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, medicalrecordstaff)
}

// NewMedicalrecordstaffController creates and registers handles for the medicalrecordstaff controller
func NewMedicalrecordstaffController(router gin.IRouter, client *ent.Client) *MedicalrecordstaffController {
	mc := &MedicalrecordstaffController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMedicalrecordstaffController registers routes to the main engine
func (ctl *MedicalrecordstaffController) register() {
	medicalrecordstaff := ctl.router.Group("/medicalrecordstaff")

	medicalrecordstaff.GET("", ctl.ListMedicalrecordstaff)

	// CRUD
	medicalrecordstaff.GET(":id", ctl.GetMedicalrecordstaff)
}
