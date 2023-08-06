package beego_locks

import (
	"database/sql"
	"github.com/beego/beego/orm"
	sqldb_storage "github.com/storage-lock/go-sqldb-storage"
	"github.com/storage-lock/go-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type BeegoLockFactory struct {
	db *orm.DB
	*storage_lock_factory.StorageLockFactory[*sql.DB]
}

func NewBeegoLockFactory(db *orm.DB) (*BeegoLockFactory, error) {
	connectionManager := NewBeegoConnectionManager(db)

	storage, err := CreateStorageForBeego(db, connectionManager)
	if err != nil {
		return nil, err
	}

	factory := storage_lock_factory.NewStorageLockFactory[*sql.DB](storage, connectionManager)

	return &BeegoLockFactory{
		db:                 db,
		StorageLockFactory: factory,
	}, nil
}

// CreateStorageForBeego 尝试从beego orm创建Storage
func CreateStorageForBeego(db *orm.DB, connectionManager storage.ConnectionManager[*sql.DB]) (storage.Storage, error) {
	return sqldb_storage.NewStorageBySqlDb(db.DB, connectionManager)
}
