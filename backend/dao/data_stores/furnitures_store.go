package data_stores

import (
	"database/sql"
	"fmt"
	"furniture_shop/dao"
	"furniture_shop/database"
)

type FurnituresStore struct {
	Provider *database.DBProvider
}

func (s *FurnituresStore) GetAll() ([]dao.Furniture, error) {
	db, err := s.Provider.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(
		"SELECT " +
			"f.id, f.furniture_name, c.category_name, f.price, f.image " +
			"FROM furnitures AS f " +
			"LEFT JOIN categories AS c " +
			"ON c.id = f.category")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []dao.Furniture
	for rows.Next() {
		var item dao.Furniture
		err = rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Image)

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

func (s *FurnituresStore) FindByName(name string) (dao.Furniture, error) {
	var item dao.Furniture

	db, err := s.Provider.OpenConnection()

	if err != nil {
		return item, err
	}

	defer db.Close()

	row := db.QueryRow(
		"SELECT "+
			"f.id, f.furniture_name, c.category_name, f.price, f.image "+
			"FROM furnitures AS f "+
			"LEFT JOIN categories AS c "+
			"ON c.id = f.category "+
			"WHERE f.furniture_name = ? "+
			"LIMIT 1", name)

	if err := row.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Image); err != nil {
		if err == sql.ErrNoRows {
			return item, fmt.Errorf("furniture with name \"%s\" Not Found", name)
		}
		return item, err
	}

	return item, nil
}

func (s *FurnituresStore) GetByCategory(category string) ([]dao.Furniture, error) {
	db, err := s.Provider.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query(
		"SELECT "+
			"f.id, f.furniture_name, c.category_name, f.price, f.image "+
			"FROM furnitures AS f "+
			"LEFT JOIN categories AS c "+
			"ON c.id = f.category "+
			"WHERE c.category_name = ?", category)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []dao.Furniture
	for rows.Next() {
		var item dao.Furniture
		err = rows.Scan(&item.ID, &item.Name, &item.Category, &item.Price, &item.Image)

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
