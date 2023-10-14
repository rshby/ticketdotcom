package repository

import (
	"backend/model/entity"
	repository "backend/repository/interface"
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"sync"
)

type ProvinceRepo struct {
	DB *sql.DB
}

// function Provider
func NewProvinceRepo(db *sql.DB) repository.IProvinceRepo {
	return &ProvinceRepo{db}
}

func (p *ProvinceRepo) GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.Province, chanError chan error) {
	defer wg.Done()

	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceRepo GetById")
	defer span.Finish()

	span.SetTag("id", id)

	statement, err := p.DB.PrepareContext(ctxTracing, "SELECT id, name FROM province WHERE id=?")
	defer statement.Close()
	if err != nil {
		chanRes <- entity.Province{}
		chanError <- err
		return
	}

	// execute query
	row := statement.QueryRowContext(ctxTracing, id)
	if row.Err() != nil {
		chanRes <- entity.Province{}
		chanError <- errors.New("record not found")
		return
	}

	var province entity.Province
	// scan
	if err := row.Scan(&province.Id, &province.Name); err != nil {
		chanRes <- entity.Province{}
		chanError <- err
		return
	}

	// success get data
	span.SetTag("response", province)
	chanRes <- province
	chanError <- nil
	return
}

func (p *ProvinceRepo) Insert(ctx context.Context, input *entity.Province) (*entity.Province, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceRepo Insert")
	defer span.Finish()

	span.LogFields(
		log.Object("request", *input))

	// create statement prepare query
	statement, err := p.DB.PrepareContext(ctxTracing, "INSERT INTO province(name) VALUES (?)")
	if err != nil {
		return nil, err
	}

	// execute query
	res, err := statement.ExecContext(ctxTracing, input.Name)
	if err != nil {
		return nil, err
	}

	id, _ := res.LastInsertId()
	input.Id = int(id)

	// success insert
	span.LogFields(
		log.Object("response-province", *input))
	return input, nil
}

func (p *ProvinceRepo) GetAll(ctx context.Context) ([]entity.Province, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "ProvinceRepo GetAll")
	defer span.Finish()

	// prepare query
	statement, err := p.DB.PrepareContext(ctxTracing, "SELECT id, name FROM province")
	defer statement.Close()
	if err != nil {
		return nil, err
	}

	// execute query
	rows, err := statement.QueryContext(ctxTracing)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var provinces []entity.Province

	for rows.Next() {
		var province entity.Province

		// scan
		if err := rows.Scan(&province.Id, &province.Name); err != nil {
			return nil, err
		}

		// append to object response
		provinces = append(provinces, province)
	}

	// check if not found
	if len(provinces) == 0 {
		return nil, errors.New("record not found")
	}

	// success
	return provinces, nil
}
