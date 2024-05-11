package maria

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultConnMaxLifeTime = time.Minute * 3
	_defaultMaxOpenConns = 10
	_defaultMaxIdleConns = 10
)

type Maria struct {
	ConnMaxLifeTime time.Duration
	MaxOPenConns int
	MaxIdleConns int
	Builder squirrel.StatementBuilderType
	Db *sql.DB
}

func New(url string, opts ...Option) (*Maria, error) {
	ma := &Maria{
		ConnMaxLifeTime: _defaultConnMaxLifeTime,
		MaxOPenConns: _defaultMaxOpenConns,
		MaxIdleConns: _defaultMaxIdleConns,
	}

	for _, opt := range opts {
		opt(ma)
	}

	ma.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	db, err := sql.Open("mysql", url)

	if err != nil {
		return &Maria{}, fmt.Errorf("maria - new - sql.open: %v", err)
	}
	
	ma.Db = db
	ma.Db.SetConnMaxLifetime(ma.ConnMaxLifeTime)
	ma.Db.SetMaxOpenConns(ma.MaxOPenConns)
	ma.Db.SetMaxIdleConns(ma.MaxIdleConns)
	fmt.Println("SUCCESSS")
	return ma, nil
}

func(m *Maria) Close() {
	if m.Db != nil {
		m.Db.Close()
	}
}
