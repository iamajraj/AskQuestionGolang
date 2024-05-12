package controllers

import (
	"ask-anon-ques/db"
	"ask-anon-ques/models"
	"ask-anon-ques/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

func init(){
	validate = validator.New()
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Model(&models.User{}).Preload("Questions").Find(&users)

	utils.SendJSON(users, &w);
}

func CreateUser(w http.ResponseWriter, r*http.Request){
	data, err := io.ReadAll(r.Body)

	if err != nil {
		utils.SendJSON(utils.JSONT{
			"error":"Failed to read body",
		}, &w)
		return;
	}

	var user models.User
	err = json.Unmarshal(data, &user);

	if err != nil {
		utils.SendJSON(utils.JSONT{
			"error":"Failed to parse data",
		}, &w)
		return;
	}

	err = validate.Struct(&user)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		utils.SendJSON(utils.JSONT{
			"validation_error": validationErrors[0].Tag() + ": " + validationErrors[0].Field(),
		}, &w)
		return;
	}

	result := db.DB.Create(&user)
	
	if result.RowsAffected == 1 {
		utils.SendJSON(utils.JSONT{
			"success":true,
			"user":user,
		}, &w)
	}else{
		utils.SendJSON(utils.JSONT{
			"success":false,
			"message":"Failed to create user",
		}, &w)
	}
}

func CreateQuestion(w http.ResponseWriter, r *http.Request){
	data, err := io.ReadAll(r.Body)

	if err != nil {
		utils.SendJSON(utils.JSONT{
			"error":"Failed to read body",
		}, &w)
		return;
	}

	var question models.Question
	err = json.Unmarshal(data, &question);

	if err != nil {
		utils.SendJSON(utils.JSONT{
			"error":"Failed to parse data",
		}, &w)
		return;
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	
	if err != nil {
		utils.SendJSON(utils.JSONT{
			"error":"Failed to parse id, Invalid id",
		}, &w)
		return;
	}

	question.UserID = uint(id)

	err = validate.Struct(&question)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		utils.SendJSON(utils.JSONT{
			"validation_error": validationErrors[0].Tag() + ": " + validationErrors[0].Field(),
		}, &w)
		return;
	}

	result:= db.DB.Create(&question)

	if result.RowsAffected == 1 {
		utils.SendJSON(utils.JSONT{
			"success":true,
			"question":question,
		}, &w)
	}else{
		utils.SendJSON(utils.JSONT{
			"success":false,
			"message":"Failed to create question",
		}, &w)
	}
}