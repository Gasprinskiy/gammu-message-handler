package sql_gen

import (
	"database/sql"
	"tgsms/internal/entity/global"

	"github.com/jmoiron/sqlx"
)

func ExecNamed[T any](tx *sqlx.Tx, sqlQuery string, data T) error {
	_, err := tx.NamedExec(sqlQuery, data)
	return err
}

func ExecNamedReturnLastInsterted[T any](tx *sqlx.Tx, sqlQuery string, data T) (int64, error) {
	var id int64

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return id, err
	}
	defer stmt.Close()

	err = stmt.Get(&id, data)
	if err != nil {
		return id, err
	}

	return id, nil
}

func Get[T any](tx *sqlx.Tx, sqlQuery string, params ...any) (T, error) {
	var data T

	err := tx.Get(&data, sqlQuery, params...)

	return data, HandleError(err)
}

func Select[T any](tx *sqlx.Tx, sqlQuery string, params ...any) ([]T, error) {
	var data []T

	err := tx.Select(&data, sqlQuery, params...)

	if err == nil && len(data) == 0 {
		err = sql.ErrNoRows
	}

	return data, HandleError(err)
}

func SelectNamed[T any](tx *sqlx.Tx, sqlQuery string, params map[string]any) ([]T, error) {
	data := make([]T, 0)

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&data, params)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		err = sql.ErrNoRows
	}

	return data, HandleError(err)
}

func SelectNamedStruct[T any, P any](tx *sqlx.Tx, sqlQuery string, params P) ([]T, error) {
	var data []T

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&data, params)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		err = sql.ErrNoRows
	}

	return data, HandleError(err)
}

func HandleError(err error) error {
	if err == sql.ErrNoRows {
		return global.ErrNoData
	}

	return err
}
