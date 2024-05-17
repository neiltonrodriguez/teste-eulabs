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

func GetAll(ctx context.Context) ([]domain.ProductDTO, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return nil, err
	}

	rows, err := Db.Query(`
	SELECT 
		id, 
		name, 
		description, 
		value, 
		created_at, 
		updated_at 
	FROM eulabs.products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.ProductDTO
	for rows.Next() {
		var product domain.ProductDTO
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

func Create(ctx echo.Context, u *domain.Product) (domain.ProductDTO, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.ProductDTO{}, err
	}
	u.Id = uuid.Must(uuid.NewRandom()).String()

	query := `INSERT INTO eulabs.products (id, name, description, value, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`

	_, err = Db.ExecContext(ctx.Request().Context(), query, u.Id, u.Name, u.Description, u.Value)
	if err != nil {
		return domain.ProductDTO{}, err
	}
	defer Db.Close()

	product, err := GetById(ctx.Request().Context(), u.Id)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func GetById(ctx context.Context, id string) (domain.ProductDTO, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.ProductDTO{}, err
	}

	rows, err := Db.Query(`
	SELECT
	    id, 
		name, 
		description, 
		value, 
		created_at, 
		updated_at  
	FROM eulabs.products WHERE id = ? limit 1`, id)
	if err != nil {
		return domain.ProductDTO{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.ProductDTO{}, errProductNotFound
	}
	var product domain.ProductDTO
	err = rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Value,
		&product.CreatedAt,
		&product.UpdatedAt)
	if err != nil {
		return domain.ProductDTO{}, err
	}

	return product, nil
}

func Update(ctx context.Context, payload *domain.Product, id string) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `
	UPDATE eulabs.products
	SET 
		name = ?, 
		description = ?, 
		value = ?
	WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, payload.Name, payload.Description, payload.Value, id)
	if err != nil {
		return err
	}

	defer Db.Close()

	return nil
}

func Delete(ctx context.Context, id string) error {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return err
	}

	query := `DELETE FROM eulabs.products WHERE id = ?`

	_, err = Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	defer Db.Close()

	return nil
}