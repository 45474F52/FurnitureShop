package controllers

import (
	"encoding/json"
	"fmt"
	"furniture_shop/dao"
	"furniture_shop/dao/data_stores"
	"net/http"
)

type CategoriesController struct {
	Store *data_stores.CategoriesStore
}

func (controller *CategoriesController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		var (
			items []dao.Category
			err   error
		)

		items, err = controller.Store.GetAll()

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
