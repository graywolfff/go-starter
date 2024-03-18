package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryMysql struct {
	client *sql.DB
}

func NewCustomerRepositoryMysql() CustomerRepositoryMysql {
	client, err := sql.Open("mysql", "root:123456@tcp(db:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryMysql{client: client}
}

func (r *CustomerRepositoryMysql) FindAll() ([]*Customer, error) {

	findAllSql := "select id, name, city, zipcode, date_of_birth, status from customer"
	rows, err := r.client.Query(findAllSql)
	if err != nil {
		log.Fatal("db error")
	}
	customers := make([]*Customer, 0)
	var dob string
	for rows.Next() {
		var c Customer
		rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &dob, &c.Status)
		c.DateOfBirth, _ = time.Parse("2006-01-02", dob)
		customers = append(customers, &c)
	}
	return customers, nil
}
