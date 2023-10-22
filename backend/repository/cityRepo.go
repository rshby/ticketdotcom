package repository

import (
	"backend/model/entity"
	repository "backend/repository/interface"
	"context"
	"database/sql"
	"errors"
	"github.com/opentracing/opentracing-go"
	"sync"
)

type CityRepo struct {
	DB *sql.DB
}

// function provider
func NewCityRepo(db *sql.DB) repository.ICityRepo {
	return &CityRepo{db}
}

func (c *CityRepo) Insert(ctx context.Context, input *entity.City) (*entity.City, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CityRepo) GetAll(ctx context.Context) ([]entity.City, error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityRepo GetAll")
	defer span.Finish()

	// create prepare statement
	statement, err := c.DB.PrepareContext(ctxTracing, "SELECT id, name, province_id FROM city")
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

	var cities []entity.City
	for rows.Next() {
		var city entity.City
		if err := rows.Scan(&city.Id, &city.Name, &city.ProvinceId); err != nil {
			return nil, err
		}

		cities = append(cities, city)
	}

	// if not found
	if len(cities) == 0 {
		return nil, errors.New("record not found")
	}

	// success
	return cities, nil
}

func (c *CityRepo) GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.City, chanError chan error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityRepo GetById")
	defer span.Finish()

	span.SetTag("request-id", id)

	wg.Add(1)
	defer func() {
		close(chanRes)
		close(chanError)
		wg.Done()
	}()

	// create statement query
	statement, err := c.DB.PrepareContext(ctxTracing, "SELECT id, name, province_id FROM city WHERE id=?")
	defer statement.Close()
	if err != nil {
		chanRes <- entity.City{}
		chanError <- err
		return
	}

	// execute query
	row := statement.QueryRowContext(ctxTracing, id)
	if row.Err() != nil {
		chanRes <- entity.City{}
		chanError <- row.Err()
		return
	}

	var city entity.City
	if err := row.Scan(&city.Id, &city.Name, &city.ProvinceId); err != nil {
		chanRes <- entity.City{}
		chanError <- err
		return
	}

	// success
	chanRes <- city
	chanError <- nil

	span.SetTag("response-object", city)
	return
}

func (c *CityRepo) GetByProvinceId(ctx context.Context, wg *sync.WaitGroup, provinceId int, chanRes chan []entity.City, chanError chan error) {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityRepo GetByProvinceId")
	defer span.Finish()

	wg.Add(1)
	defer wg.Done()

	defer func() {
		close(chanError)
		close(chanRes)
	}()

	span.SetTag("province_id", provinceId)

	// create statement query
	statement, err := c.DB.PrepareContext(ctxTracing, "SELECT id, name, province_id FROM city WHERE province_id=?")
	defer statement.Close()
	if err != nil {
		chanRes <- nil
		chanError <- err
		return
	}

	// execute query
	rows, err := statement.QueryContext(ctxTracing, provinceId)
	defer rows.Close()
	if err != nil {
		chanRes <- nil
		chanError <- err
		return
	}

	var cities []entity.City
	for rows.Next() {
		var city entity.City
		if err := rows.Scan(&city.Id, &city.Name, &city.ProvinceId); err != nil {
			chanRes <- nil
			chanError <- err
			return
		}

		// append to cities
		cities = append(cities, city)
	}

	// if not found
	if len(cities) == 0 {
		chanRes <- nil
		chanError <- errors.New("record not found")
		return
	}

	// success get data
	span.SetTag("response", cities)
	chanRes <- cities
	chanError <- nil
	return
}
