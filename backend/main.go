package main

import (
	"fmt"
	"furniture_shop/controllers"
	"furniture_shop/dao/data_stores"
	"furniture_shop/database"
	"net/http"
	"os"
)

func main() {
	provider := &database.DBProvider{}

	f_store := &data_stores.FurnituresStore{Provider: provider}
	c_store := &data_stores.CategoriesStore{Provider: provider}

	f_controller := &controllers.FurnituresController{Store: f_store}
	c_controller := &controllers.CategoriesController{Store: c_store}
	i_controller := &controllers.ImagesController{PathToImages: os.Getenv("IMAGES")}

	mux := http.NewServeMux()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	mux.Handle("/furnitures", f_controller)
	mux.Handle("/categories", c_controller)
	mux.Handle("/images", i_controller)

	fmt.Println("SERVER STARTED")

	http.ListenAndServe("0.0.0.0:"+os.Getenv("SERVERPORT"), corsMiddleware(mux))
}
