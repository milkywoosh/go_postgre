package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
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

/*
	postId := ctx.Param("postId")

	var post models.Post
	result := pc.DB.First(&post, "id = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": post})

*/

func (pc *PersonController) PersonSubjectInfo(ctx *gin.Context) {
	qry := `select p.id id_people, 
				   p."name" name_people, 
				   s.subject name_subject 
			from person p
				 left join subject s on s.id_person = p.id
			order by p."name" asc`

	defer pc.DB.Close()
	// perhatikan saat close DB, perlu close db conn setelah call each function ?????

	_, err := pc.DB.QueryContext(ctx, qry)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "No data is found",
			"error":   err.Error(),
		})
		// make sure call return to stop here
		return
	}

	// return &models.Person{}, nil
	ctx.JSON(http.StatusOK, gin.H{"message": "implemented success"})
	// make sure call return to stop here
	return
}
