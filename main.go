package main

import (
	"context"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/milkywoosh/go_postgre/orm"
)

// _ di import fungsi untuk "init" package dibelakang

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// coba batasi maximum koneksi ke datatabase agar tidak berbahaya di postgres berapa batasnya ????
	// minimal dan maximal harus di set juga

	// db.SetConnMaxIdleTime()
	// db.SetMaxOpenConns()

	defer db.Close()
	// ingat variable "err" harus reassign, dan tidak boleh buat baru => err := "xxxx"
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("db is now connected")

	// create new constructor
	r := orm.New(db)

	// fmt.Println(time.Now().Format(time.RFC3339))
	// return
	cek_val, err := r.TestInsertSchoolsExecQuery(context.TODO(), "sman z", "jl. okay", "zschool@school.org")
	if err != nil {
		panic(err)
	}
	fmt.Println(cek_val)
	/*

		arrayPeople, err := r.FindAllPeople(context.TODO())
		if err != nil {
			panic(err)
		}

		for _, val := range arrayPeople {
			fmt.Println(val)
		}

		getBenData, err := r.GetPeopleSchoolByJoin(context.TODO(), 1)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(getBenData)
		fmt.Println("email sekolah: ", getBenData.School.Email)
		fmt.Println("alamat sekolah: ", getBenData.School.Address)
	*/

}
