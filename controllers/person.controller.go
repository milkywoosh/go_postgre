package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/360EntSecGroup-Skylar/excelize"
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

func (pc PersonController) AllPerson(ctx *gin.Context) {
	qry := `select * from person`

	rows, err := pc.DB.QueryContext(ctx, qry)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	var Person models.Person
	var Persons []models.Person
	for rows.Next() {
		err = rows.Scan(
			&Person.ID,
			&Person.FirstName,
			&Person.LastName,
			&Person.Address,
			&Person.PhoneNumber,
			&Person.Job,
			&Person.InstagramName,
			&Person.FBName,
			&Person.PostalCode,
			&Person.CreatedAt,
			&Person.IDUser,
		)

		if err != nil {
			// log.Fatal(err.Error())
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"status": "fail",
				"error":  err.Error(),
			})
			return
		}
		Persons = append(Persons, Person)
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"body":    Persons,
		"message": "success",
	})

}

func (pc PersonController) InsertNewPerson(ctx *gin.Context) {
	// fail handling
	// fail := func(err error) (int64, error) {
	// 	return 0, fmt.Errorf("CreateOrder: %v", err)
	// }
	var err error
	var Person *models.Person
	// shouldBindJSON : bind body request to Person model
	if err = ctx.ShouldBindJSON(&Person); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail1",
			"error":  err.Error(),
		})
		return
	}

	var tx *sql.Tx
	// apa maksud Isolation: sql.LevelDefault ????
	tx, err = pc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if err != nil {
		// log.Fatal(err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// testing salah insert value type [first_name, 100], expected error
	insertQry := `insert into person(first_name, last_name, address, phone_number, job, instagram_uname, facebook_uname, postal_code, created_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	// person ini get value from BODY REQUEST
	result, err := tx.ExecContext(ctx, insertQry, &Person.FirstName, &Person.LastName, &Person.Address, &Person.PhoneNumber, &Person.Job, &Person.InstagramName, &Person.FBName, &Person.PostalCode, time.Now())
	if err != nil {

		// if something wrong do rollback
		tx.Rollback()

		ctx.JSON(http.StatusFailedDependency, gin.H{
			"status": "fail2",
			"error":  err.Error(),
		})
		return
	}
	fmt.Println("ape tuu result: ", result)

	if err = tx.Commit(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"status": "fail3",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"status": "success insert new person",
	})
	return

}

/**
func (pc PersonController) InsertNewUsers(ctx *gin.Context) {
	pc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelLinearizable})
}

func (pc PersonController) InsertNewPerson2(ctx *gin.Context) {
	pc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}
func (pc PersonController) InsertNewPerson3(ctx *gin.Context) {
	// pc.DB.BeginTx(ctx, &sql.TxOptions{Isolation: sql})
}
*/

// ctx disini penyalur request dari FRONT END
// func (pc *PersonController) PersonSubjectInfo(ctx *gin.Context) {
// 	qry := `select p.id id_person,
// 				   p."name" name_person,
// 				   s.subject name_subject
// 			from person p
// 				 left join subject s on s.id_person = p.id
// 			order by p."name" asc`

// 	// defer pc.DB.Close()
// 	// perhatikan saat close DB, perlu close db conn setelah call each function ?????

// 	var subject_info models.SubjectInfo
// 	var rows_subject_info []models.SubjectInfo
// 	rows, err := pc.DB.QueryContext(ctx, qry)
// 	get_status := pc.DB.Stats()
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"status":  "fail",
// 			"message": "No data is found",
// 			"error":   err.Error(),
// 		})
// 		// make sure call return to stop here
// 		return
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&subject_info.Person.ID, &subject_info.Person.NamePerson, &subject_info.Subject.SubjectName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		rows_subject_info = append(rows_subject_info, subject_info)
// 	}

// 	// return &models.Person{}, nil
// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": rows_subject_info, "status_db": get_status})
// 	// make sure call return to stop here
// }

// func (pc *PersonController) CreateNewPerson(ctx *gin.Context) {
// 	// check what inside context
// 	var Person *models.Person
// 	if err := ctx.ShouldBindJSON(&Person); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "err request", "error info": err})
// 		return
// 	}

// 	log.Println("cek: ===> ", Person)
// 	// note: LOCALTIMESTAMP is without TIMEZONE
// 	// CURRENT_TIMESTAMP is WITH TIMEZONE
// 	Person.CreatedAt = time.Now() // byPass context yg kirim value dari request Body !
// 	qry_insert := `insert into person (name_person, school_id, created_by, created_at) values($1, $2, $3, $4)`
// 	result, err := pc.DB.ExecContext(ctx, qry_insert, Person.NamePerson, Person.SchoolID, Person.CreatedBy, Person.CreatedAt)

// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to create new person data", "error info": err})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "success", "result": result})

// }

// BELUM SELESAI SAMPE INSERT KE TABLE
// func (pc *PersonController) UploadMultiplePerson(ctx *gin.Context) {
// 	var Person models.Person
// 	var Persons []models.Person
// 	//
// 	// what key?
// 	file, err := ctx.FormFile("excel_file")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	file_data, err := file.Open()
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		return
// 	}

// 	xl_file, err := excelize.OpenReader(file_data)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		return
// 	}
// 	now := time.Now()
// 	rows := xl_file.GetRows("Sheet1")
// 	for i, row := range rows {
// 		if i == 0 {
// 			continue
// 		}

// 		Person.CreatedAt = now
// 		Person.NamePerson = row[0]
// 		school_id, _ := strconv.Atoi(row[1])
// 		Person.SchoolID = school_id
// 		created_by, _ := strconv.Atoi(row[2])
// 		Person.CreatedBy = created_by
// 		// log.Printf("%+v", Person)
// 		// log.Printf("%+v", row[1])
// 		Persons = append(Persons, Person)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": len(Persons)})
// 	log.Println(time.Since(now))
// }

// api/person/plsql-one
// func (pc *PersonController) PlSqlCallDefinedFuncOne(ctx *gin.Context) {
// 	// what to do ?

// 	id_param := ctx.Param("id")
// 	var Teacher models.Teacher
// 	var Teachers []models.Teacher
// 	// if err := ctx.ShouldBindJSON(&Teacher); err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": "err request", "error info": err})
// 	// 	return
// 	// }

// 	qry := `SELECT * FROM personbyid($1)`
// 	rows, err := pc.DB.QueryContext(ctx, qry, id_param) // harus "1", kenapa id_param gagal??=> karen salah penempatan param di POSTMAN !!
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"status":  "fail",
// 			"message": "No data is found",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&Teacher.IdTeacher, &Teacher.NameTeacher, &Teacher.Email)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(&Teacher)
// 		Teachers = append(Teachers, Teacher)
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": Teachers})
// }
