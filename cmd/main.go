package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strings"

	starter "github.com/radenrishwan/golang-http-starter"
	"github.com/radenrishwan/golang-http-starter/migrations/query"
)

var (
	db   = starter.NewDatabase(context.Background()) // TODO: you need to move this one, this is just for example
	PORT = flag.String("PORT", "", "set port (if you set port from env, you can ignore this flag)")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /articles", createArticle)
	mux.HandleFunc("GET /articles", getArticle)

	if os.Getenv("PORT") != "" {
		*PORT = os.Getenv("PORT")
	}

	http.ListenAndServe(*PORT, mux)
}

type createArticleRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var createArticleRequest createArticleRequest
	err := json.NewDecoder(r.Body).Decode(&createArticleRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slug := strings.ReplaceAll(strings.ToLower(createArticleRequest.Title), " ", "-")
	result, err := query.New(db).CreateArticle(r.Context(), query.CreateArticleParams{
		Title: createArticleRequest.Title,
		Body:  createArticleRequest.Body,
		Slug:  slug,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"id":         result.ID,
		"title":      result.Title,
		"body":       result.Body,
		"slug":       result.Slug,
		"created_at": result.CreatedAt,
		"updated_at": result.UpdatedAt,
	}

	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(js))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	result, err := query.New(db).GetArticleBySlug(r.Context(), slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := map[string]interface{}{
		"id":         result.ID,
		"title":      result.Title,
		"body":       result.Body,
		"slug":       result.Slug,
		"created_at": result.CreatedAt,
		"updated_at": result.UpdatedAt,
	}

	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(js))
}
