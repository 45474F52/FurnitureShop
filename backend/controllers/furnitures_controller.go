package controllers

import (
	"encoding/json"
	"fmt"
	"furniture_shop/dao"
	"furniture_shop/dao/data_stores"
	"net/http"
)

type FurnituresController struct {
	Store *data_stores.FurnituresStore
}

func (controller *FurnituresController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		category := r.URL.Query().Get("category")

		var (
			items []dao.Furniture
			err   error
		)

		if category != "" {
			items, err = controller.Store.GetByCategory(category)
		} else {
			items, err = controller.Store.GetAll()
		}

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("err: %v\n", err)
			return
		}

		if err = json.NewEncoder(w).Encode(items); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("err: %v\n", err)
		}
	}
}
