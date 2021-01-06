package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team10/app/controllers"
	_ "github.com/team10/app/docs"
	"github.com/team10/app/ent"
	"github.com/team10/app/ent/user"
)

// struct By team 10
//-------------------------------------------------------------------




// Struct By Historytaking System

//*******************************************************************

// Nurses defines the struct for the Nurses
type Nurses struct {
	Nurse []Nurse
}

// Nurse defines the struct for the Nurse
type Nurse struct {
	Name           string
	Nursinglicense string
	Position       string
	User           int
}

// Symptomseveritys defines the struct for the Symptomseveritys
type Symptomseveritys struct {
	Symptomseverity []Symptomseverity
}

// Symptomseverity defines the struct for the Symptomseverity
type Symptomseverity struct {
	Symptomseverity string
}

// Departments defines the struct for the Departments
type Departments struct {
	Department []Department
}

// Department defines the struct for the Department
type Department struct {
	Department string
}

// Patientrecords defines the struct for the Patientrecords
type Patientrecords struct {
	Patientrecord []Patientrecord
}

// Patientrecord defines the struct for the Patientrecord
type Patientrecord struct {
	Name string
}

//*******************************************************************

//-------------------------------------------------------------------

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {

	// Set router By Team10
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")

	// Controller By Team 10 System
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	controllers.NewUserController(v1, client)
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	// Controller By Patientrights System
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	// Controller By Patientrecord System
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	controllers.NewPatientrecordController(v1, client)
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	//Controller By Bill System
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	//Controller By Historytaking System
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	controllers.NewHistorytakingController(v1, client)
	controllers.NewNurseController(v1, client)
	controllers.NewSymptomseverityController(v1, client)
	controllers.NewDepartmentController(v1, client)
	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	// Set Postman By Team10
	//-------------------------------------------------------------------
	// Set Postman By Team10 System
	//*******************************************************************
	User := []string{"Khatadet_khianchainat","nara_haru","morani_rode","faratell_yova","pulla_visan","omaha_yad",}
	for _, r := range User {
		client.User.
		Create().
		SetEmail(r+"@gmail.com").
		SetPassword("123456").
		Save(context.Background())
	}
	//*******************************************************************

	// Set Postman By Patientrights System
	//*******************************************************************

	
	// Set Postman By Historytaking System
	//*******************************************************************

	//Set nurse data
	nurses := Nurses{
		Nurse: []Nurse{
			Nurse{"Paonrat Panjainam", "Nurse123456", "พยาบาลวิชาชีพ", 2},
			Nurse{"Name Surname", "Nurse001122", "พยาบาลวิชาชีพ", 3},
		},
	}

	for _, n := range nurses.Nurse {
		u, err := client.User.
			Query().
			Where(user.IDEQ(int(n.User))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Nurse.
			Create().
			SetUser(u).
			SetName(n.Name).
			SetNursinglicense(n.Nursinglicense).
			SetPosition(n.Position).
			Save(context.Background())
	}

	//Set Symptomseverity data
	symptomseveritys := Symptomseveritys{
		Symptomseverity: []Symptomseverity{
			Symptomseverity{"ฉุกเฉิน"},
			Symptomseverity{"ฉุกเฉินรอได้"},
			Symptomseverity{"ปานกลาง"},
		},
	}

	for _, ss := range symptomseveritys.Symptomseverity {
		client.Symptomseverity.
			Create().
			SetSymptomseverity(ss.Symptomseverity).
			Save(context.Background())
	}

	//Set Department data
	departments := Departments{
		Department: []Department{
			Department{"ตาคอหู"},
			Department{"กระดูก"},
			Department{"อายุรกรรม"},
		},
	}

	for _, d := range departments.Department {
		client.Department.
			Create().
			SetDepartment(d.Department).
			Save(context.Background())
	}
	//*******************************************************************
	
	//-------------------------------------------------------------------

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
