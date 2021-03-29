package timescaledb

import (
	"fmt"
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	"github.com/lib/pq"
)

const (
	FPTableDaily   = "daily_stats_fp"
	FPTableOverall = "statistics_fp"
)

func (f *TimescaleFunctions) AddRowFP(table string, data []models.StatsFP) error {
	tx, err := f.Connect()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(pq.CopyIn(table, "created_at", "category", "data"))
	if err != nil {
		return err
	}

	for _, value := range data {
		_, err = stmt.Exec(value.CreatedAt, value.Category, value.Data)
		if err != nil {
			return err
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}

var DeleteFPByCategoryQuery = "DELETE FROM %s WHERE category=?"

func (f *TimescaleFunctions) DeleteFPByCategory(table string, category int) error {
	tx, err := f.Connect()
	if err != nil {
		return err
	}
	query := fmt.Sprintf(DeleteFPByCategoryQuery, table)
	result, err := tx.Exec(query, category)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return f.RollbackWithErrorStack(tx, err)
	}

	if affected == 0 {
		if errRb := tx.Rollback(); errRb != nil {
			return err
		}
		return ErrFailedToDelete
	}

	return tx.Commit()
}

var DeleteFPByTimeQuery = "SELECT drop_chunks(INTERVAL '?', '%s');"

// DeleteFPByTime is handling cleanup delete of old data after it has been backed up.
// It will delete all data in the table older, than the INTERVAL.
func (f *TimescaleFunctions) DeleteFPByTime(table string, intervalString string) error {
	tx, err := f.Connect()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(DeleteIntByTimeQuery, table)
	result, err := tx.Exec(query, intervalString)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return f.RollbackWithErrorStack(tx, err)
	}

	if affected == 0 {
		if errRb := tx.Rollback(); errRb != nil {
			return err
		}
		return ErrFailedToDelete
	}

	return tx.Commit()
}

var GetFPByCategoryAndTimeQuery = "SELECT * FROM %s WHERE viewer_id = $1 AND created_at > $2"

// GetFPByCategoryAndTime returns a chunk of data belonging to the defined viewer and starting from the defined time.
func (f *TimescaleFunctions) GetFPByCategoryAndTime(table string, category int, time time.Time) ([]models.StatsFP, error) {
	tx, err := f.Connect()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(GetFPByCategoryAndTimeQuery, table)
	rows, err := tx.Query(query, category, time)
	if err != nil {
		return nil, err
	}

	dataList := make([]models.StatsFP, 0)
	defer rows.Close()
	for rows.Next() {
		data := models.StatsFP{}
		err = rows.Scan(&data.CreatedAt, &data.Category, &data.Data)
		if err != nil {
			return nil, f.RollbackWithErrorStack(tx, err)
		}
		dataList = append(dataList, data)
	}

	err = rows.Err()
	if err != nil {
		return nil, f.RollbackWithErrorStack(tx, err)
	}

	return dataList, tx.Commit()
}
