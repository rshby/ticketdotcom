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
	//TODO implement me
	panic("implement me")
}

func (c *CityRepo) GetById(ctx context.Context, wg *sync.WaitGroup, id int, chanRes chan entity.City, chanError chan error) {
	//TODO implement me
	panic("implement me")
}

func (c *CityRepo) GetByProvinceId(ctx context.Context, wg *sync.WaitGroup, provinceId int, chanRes chan []entity.City, chanError chan error) {
	wg.Add(1)
	defer wg.Done()

	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "CityRepo GetByProvinceId")
	defer span.Finish()

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
