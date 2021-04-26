package model

import (
	"fmt"
	"os"

	"github.com/tinyci/ci-agents/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Model is the outer layer of our internal database model, which will
// primarily be used by the data service.
type Model struct {
	*gorm.DB
}

// New returns the model structure after the db connection work has taken place.
func New(sqlURL string) (*Model, error) {
	db, err := gorm.Open(postgres.Open(sqlURL), &gorm.Config{
		Logger:               nil, // FIXME mute this
		PrepareStmt:          true,
		FullSaveAssociations: false,
	})
	if err != nil {
		return nil, err
	}

	if os.Getenv("SQL_DEBUG") != "" {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxIdleTime(-1)
	sqlDB.SetConnMaxLifetime(-1)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	return &Model{DB: db}, nil
}

var errorMapping = map[string]error{
	"record not found": utils.ErrNotFound,
}

// SetConnPoolSize sets the connection pool size
func (m *Model) SetConnPoolSize(size int) {
	sqlDB, _ := m.DB.DB()
	sqlDB.SetMaxIdleConns(size)
	sqlDB.SetMaxOpenConns(size)
}

func preload(tx *gorm.DB) *gorm.DB {
	return tx.Statement.Preload(clause.Associations)
}

// paginate is shorthand for limit+offset GORM code.
func paginate(db *gorm.DB, page, perPage int64) *gorm.DB {
	return db.Offset(int(page * perPage)).Limit(int(perPage))
}

func (m *Model) paginate(page, perPage int64) *gorm.DB {
	return paginate(m.DB, page, perPage)
}

// MapError finds an error by string and returns an appropriate Error for it.
// The stack will NOT be preserved in the error and you will want to Wrap() it.
// If there is no potential mapping, a new Error is returned.
func MapError(err error) error {
	if err == nil {
		return nil
	}

	// FIXME this is terrible.
	if e, ok := errorMapping[err.Error()]; ok {
		return e
	}

	return err
}

// WrapError is a tail call for db transactions; it will return a wrapped and
// stack-annotated error with the msg if there is one; otherwise it will return
// nil. It also uses the errors package to normalize common errors returned
// from the DB.
func (m *Model) WrapError(call *gorm.DB, msg string) error {
	if call.Error == nil {
		return nil
	}

	fmt.Println(call.Error)

	return utils.WrapError(MapError(call.Error), "%v", msg)
}
