package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Artpou/wiki_golang/handler/respond"
	"github.com/Artpou/wiki_golang/models"
	"github.com/Artpou/wiki_golang/views"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	db.Find(&users)
	respond.RespondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(w, r) {
		return
	}
	rawUser := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rawUser); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	if rawUser.Username == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Username"))
		return
	}
	if rawUser.Password == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Password"))
		return
	}
	user := models.NewUser(rawUser.Username, rawUser.Password)
	if err := db.Save(&user).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusCreated, user)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	respond.RespondJSON(w, http.StatusOK, user)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	oldUser := models.User{}
	newUser := models.User{}
	if err := db.First(&oldUser, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if newUser.Password == "" {
		respond.RespondError(w, http.StatusBadRequest, views.FieldRequiered("Password"))
		return
	}

	updatedUser := models.UpdateUser(oldUser, newUser.Password)

	if err := db.Save(&updatedUser).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusOK, updatedUser)
}

func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	if !IsAdmin(w, r) {
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	uid64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		respond.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	uid := uint(uid64)
	user := models.User{}
	if err := db.First(&user, models.User{ID: uid}).Error; err != nil {
		respond.RespondError(w, http.StatusNotFound, views.FieldNotFound("User"))
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		respond.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond.RespondJSON(w, http.StatusNoContent, nil)
}

func GetSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.ShowUser(*user)))*/
}

func UpdateSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.UpdateUser(*user)))*/
}

func DeleteSelf(w http.ResponseWriter, r *http.Request) {
	/*w.WriteHeader(http.StatusCreated)
	user := models.NewUser("test", "1234")
	w.Write([]byte(views.DeleteUser(*user)))*/
}
