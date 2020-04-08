package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alochym01/web-w-golang/project/driver"
	"github.com/alochym01/web-w-golang/project/models"
	"github.com/alochym01/web-w-golang/project/repository"
	"github.com/alochym01/web-w-golang/project/repository/post"
	"github.com/go-chi/chi"
)

type PostController struct {
	repo repository.PostRepo
}

func NewPostHandler(db *driver.DB) *PostController {
	return &PostController{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

func (p *PostController) Fetch(res http.ResponseWriter, req *http.Request) {
	payload, _ := p.repo.Fetch(req.Context(), 5)
	respondwithJSON(res, http.StatusOK, payload)
}

// Create a new post
func (p *PostController) Create(res http.ResponseWriter, req *http.Request) {
	post := models.Post{}
	json.NewDecoder(req.Body).Decode(&post)

	newID, err := p.repo.Create(req.Context(), &post)
	fmt.Println(newID)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(res, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a post by id
func (p *PostController) Update(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	data := models.Post{ID: int64(id)}
	json.NewDecoder(req.Body).Decode(&data)
	payload, err := p.repo.Update(req.Context(), &data)

	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(res, http.StatusOK, payload)
}

// GetByID returns a post details
func (p *PostController) GetByID(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	fmt.Println(id)
	payload, err := p.repo.GetByID(req.Context(), int64(id))

	if err != nil {
		respondWithError(res, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(res, http.StatusOK, payload)
}

// Delete a post
func (p *PostController) Delete(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	_, err := p.repo.Delete(req.Context(), int64(id))

	if err != nil {
		respondWithError(res, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(res, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
