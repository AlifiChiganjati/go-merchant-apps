package di

import (
	"database/sql"
	"fmt"

	"github.com/AlifiChiganjati/go-merchant-apps/config"
	_ "github.com/lib/pq"
)

type InfraDI interface {
	Conn() *sql.DB
}

type infraDI struct {
	db  *sql.DB
	cfg *config.Config
}

func NewInfraDI(cfg *config.Config) (InfraDI, error) {
	conn := &infraDI{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (i *infraDI) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.Name)
	db, err := sql.Open(i.cfg.Driver, dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	i.db = db
	return nil
}

func (i *infraDI) Conn() *sql.DB {
	return i.db
}
