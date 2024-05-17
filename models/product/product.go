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

func GetAll(ctx context.Context) ([]domain.ProductOutputDTO, error) {
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

	var products []domain.ProductOutputDTO
	for rows.Next() {
		var product domain.ProductOutputDTO
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

func Create(ctx echo.Context, input *domain.ProductInputDTO) (domain.ProductOutputDTO, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.ProductOutputDTO{}, err
	}
	input.Id = uuid.Must(uuid.NewRandom()).String()

	query := `INSERT INTO eulabs.products (id, name, description, value, created_at, updated_at) VALUES(?, ?, ?, ?, NOW(), NOW())`

	_, err = Db.ExecContext(ctx.Request().Context(), query, input.Id, input.Name, input.Description, input.Value)
	if err != nil {
		return domain.ProductOutputDTO{}, err
	}
	defer Db.Close()

	product, err := GetById(ctx.Request().Context(), input.Id)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func GetById(ctx context.Context, id string) (domain.ProductOutputDTO, error) {
	var err error
	Db, err = database.ConnectToDB()
	if err != nil {
		return domain.ProductOutputDTO{}, err
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
		return domain.ProductOutputDTO{}, err
	}

	defer rows.Close()

	rowExist := rows.Next()
	if !rowExist {
		return domain.ProductOutputDTO{}, errProductNotFound
	}
	var product domain.ProductOutputDTO
	err = rows.Scan(
		&product.Id,
		&product.Name,
		&product.Description,
		&product.Value,
		&product.CreatedAt,
		&product.UpdatedAt)
	if err != nil {
		return domain.ProductOutputDTO{}, err
	}

	return product, nil
}

func Update(ctx context.Context, input *domain.ProductInputDTO, id string) error {
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

	_, err = Db.ExecContext(ctx, query, input.Name, input.Description, input.Value, id)
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
