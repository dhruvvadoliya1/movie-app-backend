package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func ParseInt64(value string) (int64, error) {
	if value == "" {
		return 0, nil
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}
	return int64(floatValue), nil
}

func ParseTimeFromMillis(value string) (sql.NullTime, error) {
	if value == "" {
		return sql.NullTime{Valid: false}, nil
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return sql.NullTime{}, err
	}
	return sql.NullTime{
		Time:  time.Unix(int64(floatValue/1000), 0),
		Valid: true,
	}, nil
}

func ParseTimeToSqlNull(value time.Time) sql.NullTime {
	if value.IsZero() {
		return sql.NullTime{Valid: false}
	}

	return sql.NullTime{
		Time:  value,
		Valid: true,
	}

}

func ParseStringToTime(value string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
