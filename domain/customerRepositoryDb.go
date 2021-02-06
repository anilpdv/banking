package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/anilpdv/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

// CustomerRespotoryDb : struct
type CustomerRepositoryDb struct {
	db *sql.DB
}

// FindAll : reciever
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, error) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.db.Query(findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.db.Query(findAllSQL, status)
	}

	if err != nil {
		log.Println("Got Error when retrieving data", err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

		if err != nil {
			if err != sql.ErrNoRows {
				return nil, errors.New("Customers Not Found")
			} else {
				log.Println("Got Error when retrieving data", err.Error())
				return nil, err
			}
		}

		customers = append(customers, c)
	}

	return customers, nil
}

// ById : receiver
func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.db.QueryRow(customerSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateofBirth, &c.City, &c.Status, &c.Zipcode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			log.Println("Got Error while scanning customer", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Database error")
		}
	}

	return &c, nil
}

// NewCustomerRepositoryDb : Adaptor
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:13551a0396@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}
}
