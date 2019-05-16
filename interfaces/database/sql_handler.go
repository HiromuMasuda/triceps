package database

type (
	SqlHandler interface {
		Execute(string, ...interface{}) (Result, error)
		Query(string, ...interface{}) (Row, error)
	}

	Result interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	Row interface {
		Scan(...interface{}) error
		Next() bool
		Close() error
	}
)
