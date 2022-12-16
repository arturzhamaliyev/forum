package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"forum/internal/entity"
	"forum/internal/service"
	"forum/internal/tool/customErr"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	log.Println("| | post handler is done!")
	return &postHandler{
		service: service,
	}
}

func (p *postHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	userID := r.Context().Value(userCtx)
	post := entity.Post{
		UserID: userID.(uint64),
	}

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	postID, err := p.service.CreatePost(r.Context(), post)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(postID); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (p *postHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	posts, err := p.service.GetAllPosts(r.Context())
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}

func (p *postHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	postID, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, customErr.InvalidData, http.StatusBadRequest)
		return
	}

	post, err := p.service.GetPostByID(r.Context(), postID)
	if err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, customErr.InvalidContract, http.StatusInternalServerError)
		return
	}
}
