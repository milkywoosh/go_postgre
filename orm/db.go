package orm

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

// kalo gak mau edit gak perlu pake *DB
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

	// ini apa??
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
		ID).Scan(&people.ID, &people.Name, &people.SchoolID, &school.ID, &school.NameSchool, &school.Address, &school.CreatedAt, &school.Email)

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
