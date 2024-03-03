package modules

import (
	"context"
	"database/sql"
	"log"

	"github.com/milkywoosh/go_postgre/models"
)

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

// kalo gak mau edit gak perlu pake *DB ??
func (d DB) FindAllPeople(ctx context.Context) ([]models.People, error) {

	queryString := `select * from people`
	rows, err := d.db.QueryContext(ctx, queryString)
	if err != nil {
		// panic(err)
		return nil, err
	}

	defer rows.Close() // untuk defer Close() rows di line selanjutnya

	var peoples []models.People // initiate new array

	for rows.Next() {
		var people models.People // initiate new struct before pushed to array

		err := rows.Scan(&people.ID, &people.Name, &people.SchoolID)
		if err != nil {
			log.Fatal(err)
		}
		peoples = append(peoples, people) // push to array
	}
	// close looping from rows
	rerr := rows.Close()

	// ini apa?? ===> RETURN ERROR format for checking
	if rerr != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return peoples, nil

}

func (d DB) GetPeopleSchoolByJoin(ctx context.Context, ID uint) (*models.People, error) {

	var people models.People
	var school models.School

	//  NOTEE!!!!! prepared statment harus pake $1, $2, $3 param DST
	err := d.db.QueryRowContext(ctx, `select p.*, s.* from people p 
										inner join schools s on s.id = p.school_id where p.id = $1`,
		ID).Scan(&people.ID, &people.Name, &people.SchoolID, &school.ID, &school.NameSchool, &school.Address, &school.CreatedAt, &school.EmailSchool)

	if err != nil {
		return nil, err
	}

	people.School = &school
	return &people, nil
}

func (d DB) TestInsertExecQuery(ctx context.Context, any_query_str string, arg1 string, arg2 string) (int64, error) {
	// d.db.Conn()
	result, err := d.db.ExecContext(ctx, any_query_str)
	if err != nil {
		// panic(err.Error("err exec context"))
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		// panic(err.Error("err rows affected"))
		return -1, err
	}

	return rows, nil

}

// note !!!
// try multiple insert golang with postgres
// medium :=> https://medium.com/@amoghagarwal/insert-optimisations-in-golang-26884b183b35
//
// https://stackoverflow.com/questions/71684703/bulk-insert-with-golang-and-gorm-deadlock-in-concurrency-goroutines
func (d DB) TestInsertSchoolsExecQuery(ctx context.Context, name_school string, address string, email string) (int64, error) {
	// d.db.Conn()
	result, err := d.db.ExecContext(ctx, `insert into schools (name_school, address, created_at, email) values($1, $2, current_date, $3)`, name_school, address, email)
	if err != nil {
		// panic(err.Error("err exec context"))
		return -1, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		// panic(err.Error("err rows affected"))
		return -1, err
	}

	return rows, nil

}

// method struct
func (d DB) UseJoinSQL(ctx context.Context, id_people uint) (*models.People, error) {

	query_str := `
	select p.id id_people, p."name" name_people, s.subject name_subject  from people p
		left join subject s on s.id_people = p.id
		where p.id = $1
		order by p."name" asc`
	//  NOTEE!!!!! prepared statment harus pake $1, $2, $3 param DST
	rows, err := d.db.QueryContext(ctx, query_str, id_people)

	if err != nil {
		return nil, err
	}
	// NOTE PENGGUNAAN "&"
	// [var people models.People RETURN harus '&people'] == SETARA == [people := &models.People{}  RETURN langsung 'people' ]
	var people models.People
	var subject models.Subject

	for rows.Next() {

		err = rows.Scan(&people.ID, &people.Name, &subject.SubjectName)
		if err != nil {
			log.Fatal(err)
		}
		people.Subjects = append(people.Subjects, subject)

	}

	return &people, nil

	/* convert to JSON for RESPONSE API
	var jsonData, err := json.Marshal(data of any type, like struct, map)
	if err != nil {
		panic(err.Error())
	}

	var jsonString := string(jsonData)
	fmt.Println(jsonString)
	// test := reflect.TypeOf(data.Subjects)
	// fmt.Println(test)
	*/

	/* hasil to MARSHAL
	WITH OMITEMPTY
	{
		"id": 1,
		"name": "ben",
		"Subjects": [
			{ "subject": "MATH" },
			{ "subject": "PHYSIC" },
			{ "subject": "LOGIC" },
			{ "subject": "HISTORY" },
			{ "subject": "MAGIC" }
		]
	}


	=====================================================================
	{
	  "id": 1,
	  "name": "ben",
	  "school_id": 0,
	  "Subjects": [
	    { "id_subject": 0, "subject": "MATH", "id_people": 0 },
	    { "id_subject": 0, "subject": "PHYSIC", "id_people": 0 },
	    { "id_subject": 0, "subject": "LOGIC", "id_people": 0 },
	    { "id_subject": 0, "subject": "HISTORY", "id_people": 0 },
	    { "id_subject": 0, "subject": "MAGIC", "id_people": 0 }
	  ]
	}


	*/
} // end function

/*
SELECT p.name as p_name, t.name_teacher as name_teacher, t.email, s.subject

	FROM people p
	inner join teacher t ON p.id = t.id_people
	inner join subject s ON t.id_subject = s.id_subject;
*/
func (d DB) UseTripleJoin(ctx context.Context) ([]models.CompleteData, error) {
	query := `SELECT p.name as p_name, t.name_teacher as name_teacher, t.email, s.subject
	FROM people p
	inner join teacher t ON p.id = t.id_people
	inner join subject s ON t.id_subject = s.id_subject;`

	var CompleteDataVals []models.CompleteData
	var CompleteData models.CompleteData

	// var people models.People
	// var teacher models.Teacher
	// var subject models.Subject
	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err = rows.Scan(&CompleteData.Name, &CompleteData.NameTeacher, &CompleteData.EmailTeacher, &CompleteData.SubjectName)
		if err != nil {
			log.Fatal(err)
		}
		CompleteDataVals = append(CompleteDataVals, CompleteData)
	}

	return CompleteDataVals, nil
}
