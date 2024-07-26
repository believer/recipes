package main

import (
	"log"
	"net/http"
	"os"

	"github.com/believer/recipes/data"
	"github.com/believer/recipes/model"
	"github.com/believer/recipes/views"
)

func main() {
	// DB
	err := data.InitDB()

	defer data.DB.Close()

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Routing
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		var recipes []model.Recipe

		err := data.DB.Select(&recipes, "SELECT id, url, name, type FROM recipe ORDER BY name")

		if err != nil {
			log.Println(err)
		}

		views.Index(recipes).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		var recipe model.Recipe
		var ingredients []model.Ingredient

		id := r.PathValue("id")

		err := data.DB.Select(&ingredients, "SELECT i.name, ri.amount FROM recipe as r INNER JOIN recipe_ingredient AS ri ON ri.recipe_id = r.id INNER JOIN ingredient as i ON ri.ingredient_id = i.id WHERE r.id = $1;", id)

		if err != nil {
			log.Println(err)
		}

		err = data.DB.Get(&recipe, "SELECT id, url, name, type, description FROM recipe WHERE id = $1;", id)

		if err != nil {
			log.Println(err)
		}

		views.Recipe(recipe, ingredients).Render(r.Context(), w)
	})

	// Static files
	dir := http.Dir("./public")
	fs := http.FileServer(dir)
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// Start router
	err = http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
