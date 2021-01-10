package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/symptomseverity"
)

// HistorytakingController defines the struct for the historytaking controller
type HistorytakingController struct {
	client *ent.Client
	router gin.IRouter
}

// Historytaking defines the struct for the Historytaking entity
type Historytaking struct {
	Nurse           int
	Department      int
	Symptomseverity int
	Patientrecord   int
	Hight           string
	Weight          string
	Temp            string
	Pulse           string
	Respiration     string
	Bp              string
	Oxygen          string
	Symptom         string
	Datetime        string
}

// CreateHistorytaking handles POST requests for adding historytaking entities
// @Summary Create historytaking
// @Description Create historytaking
// @ID create-historytaking
// @Accept   json
// @Produce  json
// @Param historytaking body Historytaking true "Historytaking entity"
// @Success 200 {object} ent.Historytaking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking [post]
func (ctl *HistorytakingController) CreateHistorytaking(c *gin.Context) {
	obj := Historytaking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Historytaking binding failed",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	d, err := ctl.client.Department.
		Query().
		Where(department.IDEQ(int(obj.Department))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	ss, err := ctl.client.Symptomseverity.
		Query().
		Where(symptomseverity.IDEQ(int(obj.Symptomseverity))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	pt, err := ctl.client.Patientrecord.
		Query().
		Where(patientrecord.IDEQ(int(obj.Patientrecord))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	times := time.Now().Local()

	var h float32
	if hights, err := strconv.ParseFloat(obj.Hight, 64); err == nil {
		h = float32(hights)
	}

	var w float32
	if weights, err := strconv.ParseFloat(obj.Weight, 64); err == nil {
		w = float32(weights)
	}
	var t float32
	if temps, err := strconv.ParseFloat(obj.Temp, 64); err == nil {
		t = float32(temps)
	}

	pulses, err := strconv.Atoi(obj.Pulse)
	respirations, err := strconv.Atoi(obj.Respiration)
	bps, err := strconv.Atoi(obj.Bp)

	ht, err := ctl.client.Historytaking.
		Create().
		SetNurse(n).
		SetDepartment(d).
		SetSymptomseverity(ss).
		SetPatientrecord(pt).
		SetHight(h).
		SetWeight(w).
		SetTemp(t).
		SetPulse(pulses).
		SetRespiration(respirations).
		SetBp(bps).
		SetOxygen(obj.Oxygen + "%").
		SetSymptom(obj.Symptom).
		SetDatetime(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ht)
}

// GetHistorytaking handles GET requests to retrieve a historytaking entity
// @Summary Get a historytaking entity by ID
// @Description get historytaking by ID
// @ID get-historytaking
// @Produce  json
// @Param id path int true "Historytaking ID"
// @Success 200 {object} ent.Historytaking
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking/{id} [get]
func (ctl *HistorytakingController) GetHistorytaking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pr, err := ctl.client.Historytaking.
		Query().
		Where(historytaking.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pr)
}

// ListHistorytaking handles request to get a list of historytaking entities
// @Summary List historytaking entities
// @Description list historytaking entities
// @ID list-historytaking
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Historytaking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytaking [get]
func (ctl *HistorytakingController) ListHistorytaking(c *gin.Context) {
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

	historytaking, err := ctl.client.Historytaking.
		Query().
		WithNurse().
		WithDepartment().
		WithSymptomseverity().
		WithPatientrecord().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, historytaking)
}

// DeleteHistorytaking handles DELETE requests to delete a historytaking entity
// @Summary Delete a historytaking entity by ID
// @Description get historytaking by ID
// @ID delete-historytaking
// @Produce  json
// @Param id path int true "Historytaking ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /historytakings/{id} [delete]
func (ctl *HistorytakingController) DeleteHistorytaking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Historytaking.
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

// NewHistorytakingController creates and registers handles for the historytaking controller
func NewHistorytakingController(router gin.IRouter, client *ent.Client) *HistorytakingController {
	uc := &HistorytakingController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *HistorytakingController) register() {
	historytakings := ctl.router.Group("/historytakings")

	historytakings.GET("", ctl.ListHistorytaking)
	historytakings.POST("", ctl.CreateHistorytaking)
	historytakings.GET(":id", ctl.GetHistorytaking)
	historytakings.DELETE(":id", ctl.DeleteHistorytaking)
}
