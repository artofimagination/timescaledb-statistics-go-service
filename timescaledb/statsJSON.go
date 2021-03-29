package timescaledb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/artofimagination/timescaledb-statistics-go-service/models"
	"github.com/lib/pq"
)

const (
	JSONTableDaily   = "daily_stats_json"
	JSONTableOverall = "statistics_json"
)

func (f *TimescaleFunctions) AddRowJSON(table string, data []models.StatsJSON) error {
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

var DeleteJSONByCategoryQuery = "DELETE FROM %s WHERE category=?"

func (f *TimescaleFunctions) DeleteJSONByCategory(table string, category int) error {
	tx, err := f.Connect()
	if err != nil {
		return err
	}
	query := fmt.Sprintf(DeleteJSONByCategoryQuery, table)
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

var DeleteJSONByTimeQuery = "SELECT drop_chunks(INTERVAL '?', '%s');"

// DeleteByTime is handling cleanup delete of old data after it has been backed up.
// It will delete all data in the table older, than the INTERVAL.
func (f *TimescaleFunctions) DeleteJSONByTime(table string, intervalString string) error {
	tx, err := f.Connect()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(DeleteJSONByTimeQuery, table)
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

var GetJSONByCategoryAndTimeQuery = "SELECT * FROM %s WHERE viewer_id = $1 AND created_at > $2"

// GetDataByViewerAndTime returns a chunk of data belonging to the defined viewer and starting from the defined time.
func (f *TimescaleFunctions) GetJSONByCategoryAndTime(table string, category int, time time.Time) ([]models.StatsJSON, error) {
	tx, err := f.Connect()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(GetJSONByCategoryAndTimeQuery, table)
	rows, err := tx.Query(query, category, time)
	if err != nil {
		return nil, err
	}

	dataList := make([]models.StatsJSON, 0)
	defer rows.Close()
	for rows.Next() {
		data := models.StatsJSON{}
		dataMap := []byte{}
		err = rows.Scan(&data.CreatedAt, &data.Category, &dataMap)
		if err != nil {
			return nil, f.RollbackWithErrorStack(tx, err)
		}
		if err := json.Unmarshal(dataMap, &data.Data); err != nil {
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
