package user

import (
	"context"
	"database/sql"
	"errors"
	"eulabs/database"
	"eulabs/domain"
	"log"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

var Db *sql.DB

var (
	errProductNotFound = errors.New("product not found")
)

func GetAll(ctx context.Context) ([]domain.Product, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		id, 
		name, 
		COALESCE('description', ''), 
		value, 
		created_at, 
		updated_at 
	FROM eulabs.products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Value,
			&product.CreatedAt,
			&product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func Create(ctx echo.Context, u *domain.Product) (domain.Product, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Product{}, err
	}
	u.Id = uuid.Must(uuid.NewRandom()).String()

	query := `INSERT INTO eulabs.products (id, name, description, value, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`

	_, err = Db.ExecContext(ctx.Request().Context(), query, u.Id, u.Name, u.Description, u.Value)
	if err != nil {
		return domain.Product{}, err
	}
	defer Db.Close()

	product, err := GetById(ctx.Request().Context(), u.Id)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func GetById(ctx context.Context, id string) (domain.Product, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.Product{}, err
	}

	rows, err := Db.Query(`
	SELECT
	    id, 
		name, 
		COALESCE('description', ''), 
		value, 
		created_at, 
		updated_at  
	FROM eulabs.products WHERE id = ? limit 1`, id)
	if err != nil {
		return domain.Product{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.Product{}, errProductNotFound
	}
	var product domain.Product
	err = rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Value,
		&product.CreatedAt,
		&product.UpdatedAt)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}


