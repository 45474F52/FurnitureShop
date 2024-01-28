package data_stores

import (
	"database/sql"
	"fmt"
	"furniture_shop/dao"
	"furniture_shop/database"
)

type CategoriesStore struct {
	Provider *database.DBProvider
}

func (s *CategoriesStore) GetAll() ([]dao.Category, error) {
	db, err := s.Provider.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []dao.Category
	for rows.Next() {
		var item dao.Category
		err = rows.Scan(&item.ID, &item.Name, &item.Image)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *CategoriesStore) FindByName(name string) (dao.Category, error) {
	var item dao.Category

	db, err := s.Provider.OpenConnection()

	if err != nil {
		return item, err
	}

	defer db.Close()

	row := db.QueryRow(
		"SELECT * FROM categories"+
			"WHERE category_name = ? "+
			"LIMIT 1", name)

	if err := row.Scan(&item.ID, &item.Name, &item.Image); err != nil {
		if err == sql.ErrNoRows {
			return item, fmt.Errorf("category with name \"%s\" Not Found", name)
		}
		return item, err
	}

	return item, nil
}
