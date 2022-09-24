package connection

import (
	"gitlab.com/abiewardani/cepot/internal/system/config"
	"gitlab.com/abiewardani/cepot/internal/system/connection/database"
)

// Connection ...
type Connection struct {
	db database.GormDatabase
}

// LoadConnection ...
func LoadConnection(cfg *config.Config) Connection {
	return Connection{
		db: database.InitGorm(cfg),
	}
}

// DB func
func (c *Connection) DB() database.GormDatabase {
	return c.db
}
