package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (db *sql.DB, e error) {
	user := "root"
	password := "GNKrjohAt0Ff4Dd0"
	host := "tcp(34.148.102.208:3306)"
	database := "admins"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", user, password, host, database))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetAdminsFromDatabase(parameters ...string) ([]Administrator, error) {
	db, err := GetDatabase()
	admins := []Administrator{}

	if err != nil {
		fmt.Printf("error with the database connection: %v", err)
		return []Administrator{}, err
	}

	defer db.Close()

	results, err := db.Query(`SELECT a.id, a.adminame AS name, o.name AS owner, c.criticality
								FROM admins a
								INNER JOIN owners o ON a.ownerId = o.ownerId
								INNER JOIN criticality c ON a.criticalityId = c.criticalityId`)
	if err != nil {
		log.Println(err.Error())
	}

	defer results.Close()

	var admin Administrator

	for results.Next() {
		err := results.Scan(&admin.ID, &admin.Name, &admin.Owner, &admin.Citicality)
		if err != nil {
			return nil, err
		}
		admins = append(admins, admin)
	}

	return admins, nil
}

func GetOwnerId(ownerName string) int {
	id := -1

	db, err := GetDatabase()
	if err != nil {
		fmt.Printf("error with the database connection: %v", err)
		return id
	}

	defer db.Close()

	result, err := db.Query(`select ownerId from owners o where UPPER(name) like UPPER(?)`, ownerName)
	if err != nil {
		log.Println(err.Error())
	}
	result.Next()

	err = result.Scan(&id)
	if err != nil {
		return id
	}
	return id
}

func QueryDatabase(query string, administratorName string, ownerId int, criticalityId int, updateParams ...int) error {
	db, err := GetDatabase()
	if err != nil {
		return errors.New("error with the database connection -> " + err.Error())
	}

	defer db.Close()

	querySentence, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer querySentence.Close()

	if len(updateParams) > 0 {
		_, err = querySentence.Exec(administratorName, ownerId, criticalityId, updateParams[0])
		if err != nil {
			return err
		}

		return nil
	}

	_, err = querySentence.Exec(administratorName, ownerId, criticalityId)
	if err != nil {
		return err
	}

	return nil

}
