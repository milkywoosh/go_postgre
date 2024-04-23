package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
)

type PersonController struct {
	DB *sql.DB
}

// note : return POINTER atau VALUE, tidak maasalah
//
//	itu tergantung penggunaan
//	return POINTER enable for changing the VALUE attached to address
func NewPersonController(db_arg *sql.DB) PersonController {
	return PersonController{DB: db_arg}
}

///////////////////////////////////////////
////// later need to use TOKEN with////////
///////////////////////////////////////////

// ctx disini penyalur request dari FRONT END
func (pc *PersonController) PersonSubjectInfo(ctx *gin.Context) {
	qry := `select p.id_person, 
				   p.name_person,
				   s.subject_name
			from person p
				 left join subject s on s.id_person = p.id_person
			order by p.id_person asc`

	// defer pc.DB.Close()

	// perhatikan saat close DB, perlu close db conn setelah call each function ?????

	var subject_info models.SubjectInfo
	var rows_subject_info []models.SubjectInfo
	rows, err := pc.DB.QueryContext(ctx, qry)
	get_status := pc.DB.Stats()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "No data is found",
			"error":   err.Error(),
		})
		// make sure call return to stop here
		return
	}

	for rows.Next() {
		err = rows.Scan(&subject_info.IdSubject,
			&subject_info.NamePerson,
			&subject_info.SubjectName)
		if err != nil {
			log.Fatal(err)
		}
		rows_subject_info = append(rows_subject_info, subject_info)
	}

	// return &models.Person{}, nil
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": rows_subject_info, "status_db": get_status})
	// make sure call return to stop here

}

func (pc *PersonController) CreateNewPerson(ctx *gin.Context) {
	// check what inside context
	var Person *models.Person
	var now_time time.Time = time.Now()
	// ShouldBinJSON harus represent body request
	if err := ctx.ShouldBindJSON(&Person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "err request", "error info": err})
		return
	}
	Person.CreatedAt = &now_time

	log.Println("cek: ===> ", Person)
	// note: LOCALTIMESTAMP is without TIMEZONE
	// CURRENT_TIMESTAMP is WITH TIMEZONE
	Person.CreatedAt = &now_time // byPass context yg kirim value dari request Body !
	qry_insert := `insert into person (name_person, id_school) values($1, $2)`
	result, err := pc.DB.ExecContext(ctx, qry_insert, Person.NamePerson, Person.SchoolID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to create new person data", "error info": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "success", "result": result})

}

// BELUM SELESAI SAMPE INSERT KE TABLE
func (pc *PersonController) UploadMultiplePerson(ctx *gin.Context) {
	var Person models.Person
	var Persons []models.Person
	//
	// what key?
	file, err := ctx.FormFile("excel_file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file_data, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	xl_file, err := excelize.OpenReader(file_data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	now := time.Now()
	rows := xl_file.GetRows("Sheet1")
	for i, row := range rows {
		if i == 0 {
			continue
		}

		Person.CreatedAt = &now
		Person.NamePerson = &row[0]
		school_id, _ := strconv.Atoi(row[1])
		Person.SchoolID = &school_id
		created_by, _ := strconv.Atoi(row[2])
		Person.CreatedBy = &created_by
		// log.Printf("%+v", Person)
		// log.Printf("%+v", row[1])
		Persons = append(Persons, Person)
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": len(Persons)})
	log.Println(time.Since(now))
}

// api/person/plsql-one
func (pc *PersonController) PlSqlCallDefinedFuncOne(ctx *gin.Context) {
	// what to do ?

	id_param := ctx.Param("id")
	var Teacher models.Teacher
	var Teachers []models.Teacher
	// if err := ctx.ShouldBindJSON(&Teacher); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "err request", "error info": err})
	// 	return
	// }

	qry := `SELECT * FROM personbyid($1)`
	rows, err := pc.DB.QueryContext(ctx, qry, id_param) // harus "1", kenapa id_param gagal??=> karen salah penempatan param di POSTMAN !!
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "No data is found",
			"error":   err.Error(),
		})
		return
	}
	for rows.Next() {
		err = rows.Scan(&Teacher.IdTeacher, &Teacher.NameTeacher, &Teacher.Email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(&Teacher)
		Teachers = append(Teachers, Teacher)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": Teachers})
}

// call PROCEDURE
func (pc *PersonController) PlSqlCallDefinedProcOne(ctx *gin.Context) {

}
