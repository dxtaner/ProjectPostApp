package handlers

import (
	"PostApp/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"PostApp/config"
	"PostApp/models"

	"github.com/gorilla/mux"
)

func init() {
	// eğer başlangıçta veri eklemek istiyorsak init() çalışır ve verileri ekler
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var posts []models.Post
	db.Find(&posts)

	utils.Respond(w, utils.Message(true, "Post verileri başarıyla alındı", posts), http.StatusOK)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Geçersiz post ID", nil), http.StatusBadRequest)
		return
	}

	var post models.Post
	err = config.GetDB().First(&post, id).Error
	if err != nil {
		utils.Respond(w, utils.Message(false, "Post bulunamadı", nil), http.StatusNotFound)
		return
	}

	utils.Respond(w, utils.Message(true, "Post verisi başarıyla alındı", post), http.StatusOK)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.Respond(w, utils.Message(false, "Geçersiz istek!", nil), http.StatusBadRequest)
		return
	}

	if post.Title == "" || post.Description == "" || post.ImageURI == "" {
		utils.Respond(w, utils.Message(false, "Title, description ve imageURI alanları boş olamaz", nil), http.StatusBadRequest)
		return
	}

	post.CreatedAt = time.Now()
	db := config.GetDB()

	if err := db.Create(&post).Error; err != nil {
		utils.Respond(w, utils.Message(false, "Database hatası", nil), http.StatusBadRequest)
		return
	}

	utils.Respond(w, utils.Message(true, "Post verisi başarıyla oluşturuldu", post), http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	post := &models.Post{}
	post.UpdatedAt = time.Now()
	err := config.GetDB().Table("posts").Where("id = ?", id).First(post).Error
	if err != nil {
		utils.Respond(w, utils.Message(false, "Post bulunamadı", nil), http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Geçersiz İstek!", nil), http.StatusBadRequest)
		return
	}

	err = config.GetDB().Save(post).Error
	if err != nil {
		utils.Respond(w, utils.Message(false, "Post güncellenemedi", nil), http.StatusInternalServerError)
		return
	}

	utils.Respond(w, utils.Message(true, "Post verisi başarıyla güncellendi", post), http.StatusOK)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post models.Post
	err := config.GetDB().First(&post, id).Error
	if err != nil {
		utils.Respond(w, utils.Message(false, "Post bulunamadı", nil), http.StatusNotFound)
		return
	}

	config.GetDB().Delete(&post)

	utils.Respond(w, utils.Message(true, "Post silindi", post), http.StatusOK)
}
