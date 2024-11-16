package appctx

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
}

// A struct that implements "AppContext" interface
type appCtx struct {
	db *gorm.DB
}

// Setters for setting 2 fields of "db", "uploadprovider"
func NewAppContext(
	db *gorm.DB,
) *appCtx {
	return &appCtx{
		db: db,
	}
}

// Getters to get 2 fields of "db", "uploadprovider"
func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }
