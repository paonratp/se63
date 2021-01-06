package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/user"
)

// NurseController defines the struct for the nurse controller
type NurseController struct {
	client *ent.Client
	router gin.IRouter
}

type Nurse struct {
	Name			string
	Nursinglicense	string
	Position		string
	User			int
}

// CreateNurse handles POST requests for adding nurse entities
// @Summary Create nurse
// @Description Create nurse
// @ID create-nurse
// @Accept   json
// @Produce  json
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurse [post]
func (ctl *NurseController) CreateNurse(c *gin.Context) {
	obj := Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	n, err := ctl.client.Nurse.
		Create().
		SetUser(u).
		SetName(obj.Name).
		SetNursinglicense(obj.Nursinglicense).
		SetPosition(obj.Position).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, n)
}

// GetNurse handles GET requests to retrieve a nurse entity
// @Summary Get a nurse entity by ID
// @Description get nurse by ID
// @ID get-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurse/{id} [get]
func (ctl *NurseController) GetNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListNurse handles request to get a list of nurse entities
// @Summary List nurse entities
// @Description list nurse entities
// @ID list-nurse
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurse [get]
func (ctl *NurseController) ListNurse(c *gin.Context) {
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

	nurse, err := ctl.client.Nurse.
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

	c.JSON(200, nurse)
}

// DeleteNurse handles DELETE requests to delete a nurse entity
// @Summary Delete a nurse entity by ID
// @Description get nurse by ID
// @ID delete-nurse
// @Produce  json
// @Param id path int true "Nurse ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurse/{id} [delete]
func (ctl *NurseController) DeleteNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Nurse.
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

// UpdateNurse handles PUT requests to update a nurse entity
// @Summary Update a nurse entity by ID
// @Description update nurse by ID
// @ID update-nurse
// @Accept   json
// @Produce  json
// @Param id path int true "Nurse ID"
// @Param nurse body ent.Nurse true "Nurse entity"
// @Success 200 {object} ent.Nurse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /nurse/{id} [put]
func (ctl *NurseController) UpdateNurse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Nurse{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "nurse binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Nurse.
		UpdateOneID(int(id)).
		SetName(obj.Name).
		SetNursinglicense(obj.Nursinglicense).
		SetPosition(obj.Position).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewNurseController creates and registers handles for the nurse controller
func NewNurseController(router gin.IRouter, client *ent.Client) *NurseController {
	uc := &NurseController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *NurseController) register() {
	nurse := ctl.router.Group("/nurse")

	nurse.GET("", ctl.ListNurse)

	// CRUD
	nurse.POST("", ctl.CreateNurse)
	nurse.GET(":id", ctl.GetNurse)
	nurse.PUT(":id", ctl.UpdateNurse)
	nurse.DELETE(":id", ctl.DeleteNurse)
}
