package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
)

type PersonController struct {
	DB *sql.DB
}

func NewPersonController(db_arg *sql.DB) PersonController {
	return PersonController{DB: db_arg}
}

///////////////////////////////////////////
////// later need to use TOKEN with////////
///////////////////////////////////////////

func (pc *PersonController) PersonSubjectInfo(ctx *gin.Context) {
	qry := `select p.id id_people, 
				   p."name" name_people, 
				   s.subject name_subject 
			from person p
				 left join subject s on s.id_person = p.id
			order by p."name" asc`

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
		err = rows.Scan(&subject_info.Person.ID, &subject_info.Person.Name, &subject_info.Subject.SubjectName)
		if err != nil {
			log.Fatal(err)
		}
		rows_subject_info = append(rows_subject_info, subject_info)
	}

	// return &models.Person{}, nil
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": rows_subject_info, "status_db": get_status})
	// make sure call return to stop here
	return
}
