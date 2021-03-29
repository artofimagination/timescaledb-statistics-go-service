package dbcontrollers

import (
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	timescaledb "github.com/artofimagination/timescaledb-statistics-go-service/timescaledb"
)

func (c *TimescaleController) AddDailyDataInt(data []models.StatsInt) error {
	if err := c.DBFunctions.AddRowInt(timescaledb.IntTableDaily, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyIntByCategory(category int) error {
	if err := c.DBFunctions.DeleteIntByCategory(timescaledb.IntTableDaily, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteDailyIntByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteIntByTime(timescaledb.IntTableDaily, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetDailyIntByCategoryAndTime(category int, time time.Time) ([]models.StatsInt, error) {
	data, err := c.DBFunctions.GetIntByCategoryAndTime(timescaledb.IntTableDaily, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}

func (c *TimescaleController) AddOverallDataInt(data []models.StatsInt) error {
	if err := c.DBFunctions.AddRowInt(timescaledb.IntTableOverall, data); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallIntByCategory(category int) error {
	if err := c.DBFunctions.DeleteIntByCategory(timescaledb.IntTableOverall, category); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) DeleteOverallIntByTime(intervalString string) error {
	if err := c.DBFunctions.DeleteIntByTime(timescaledb.IntTableOverall, intervalString); err != nil {
		return ErrFailedToAddData
	}
	return nil
}

func (c *TimescaleController) GetOverallIntByCategoryAndTime(category int, time time.Time) ([]models.StatsInt, error) {
	data, err := c.DBFunctions.GetIntByCategoryAndTime(timescaledb.IntTableOverall, category, time)
	if err != nil {
		return nil, ErrFailedToAddData
	}
	return data, nil
}
