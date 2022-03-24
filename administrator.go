package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAdministrators(w http.ResponseWriter, r *http.Request) {

	administrators, err := GetAdminsFromDatabase()
	if err != nil {
		log.Println(err.Error())
		SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, http.StatusOK, administrators, false)

}

func PostAdministrator(w http.ResponseWriter, r *http.Request) {

	administratorName, ownerId, criticalityId, err := TransformAdministratorData(r)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, err.Error())
	}

	err = QueryDatabase(`INSERT INTO admins.admins (adminame, ownerId, criticalityId) VALUES(?, ?, ?);`, administratorName, ownerId, criticalityId)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, http.StatusOK, "Administrator created successfully!")

}

func UpdateAdministrator(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		SendResponse(w, http.StatusBadRequest, "Administrator's ID invalid")
	}

	administratorName, ownerId, criticalityId, err := TransformAdministratorData(r)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, err.Error())
	}

	err = QueryDatabase(`UPDATE admins.admins SET adminame=?, ownerId=?, criticalityId=? WHERE id=?;`, administratorName, ownerId, criticalityId, id)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendResponse(w, http.StatusOK, "Administrator updated successfully!")

}

func TransformAdministratorData(r *http.Request) (string, int, int, error) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", -1, -1, err
	}

	var adminBody Administrator

	json.Unmarshal(reqBody, &adminBody)

	administratorName := adminBody.Name
	ownerId := GetOwnerId(adminBody.Owner)
	criticalityId := GetCriticalityId(adminBody.Citicality)

	if administratorName == "" {
		return "", -1, -1, errors.New("administrator's name is missing")
	}

	if ownerId == -1 {
		return "", -1, -1, errors.New("the owner doesn't exists")
	}

	if criticalityId == -1 {
		return "", -1, -1, errors.New("the criticality doesn't exists")
	}

	return administratorName, ownerId, criticalityId, nil

}
