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

	// Static files
	dir := http.Dir("./public")
	fs := http.FileServer(dir)
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		var recipes []model.Recipe

		err := data.DB.Select(&recipes, "SELECT id, url, name FROM recipe ORDER BY name")

		if err != nil {
			log.Println(err)
		}

		views.Index(recipes).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /recipe/{id}", func(w http.ResponseWriter, r *http.Request) {
		var recipe model.Recipe
		var ingredients []model.Ingredient

		id := r.PathValue("id")

		err := data.DB.Select(&ingredients, "SELECT ingredient, amount, serving_size FROM recipe_ingredient WHERE recipe_id = $1 ORDER BY amount IS NULL;", id)

		if err != nil {
			log.Println(err)
		}

		err = data.DB.Get(&recipe, "SELECT id, url, name, course, description FROM recipe WHERE id = $1;", id)

		if err != nil {
			log.Println(err)
		}

		views.Recipe(recipe, ingredients).Render(r.Context(), w)
	})

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
