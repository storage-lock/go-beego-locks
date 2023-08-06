package beego_locks

import (
	"context"
	"database/sql"
	"github.com/beego/beego/orm"
	"github.com/storage-lock/go-storage"
)

// BeegoOrmConnectionManager 用来复用beego orm（https://github.com/beego/beego）的连接
// TODO 2023-8-4 01:31:00 单元测试
type BeegoOrmConnectionManager struct {
	db *orm.DB
}

var _ storage.ConnectionManager[*sql.DB] = &BeegoOrmConnectionManager{}

func NewBeegoConnectionManager(db *orm.DB) *BeegoOrmConnectionManager {
	return &BeegoOrmConnectionManager{
		db: db,
	}
}

func (x *BeegoOrmConnectionManager) Name() string {
	return "beego-orm-connection-manager"
}

func (x *BeegoOrmConnectionManager) Take(ctx context.Context) (*sql.DB, error) {
	return x.db.DB, nil
}

func (x *BeegoOrmConnectionManager) Return(ctx context.Context, db *sql.DB) error {
	return nil
}

func (x *BeegoOrmConnectionManager) Shutdown(ctx context.Context) error {
	if x.db != nil {
		return x.db.DB.Close()
	}
	return nil
}
