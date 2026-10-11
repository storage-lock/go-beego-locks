package beego_locks

import (
	"context"
	"database/sql"
	"testing"

	"github.com/beego/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func testBeegoDB(t *testing.T) *orm.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", "file:beego_test.db?mode=memory")
	assert.Nil(t, err)
	t.Cleanup(func() { db.Close() })
	return &orm.DB{DB: db}
}

func TestNewBeegoConnectionManager(t *testing.T) {
	ormDB := testBeegoDB(t)
	m := NewBeegoConnectionManager(ormDB)
	assert.NotNil(t, m)
	assert.Equal(t, "beego-orm-connection-manager", m.Name())

	ctx := context.Background()
	got, err := m.Take(ctx)
	assert.Nil(t, err)
	assert.Equal(t, ormDB.DB, got)

	assert.Nil(t, m.Return(ctx, got))
	assert.Nil(t, m.Shutdown(ctx))
}

func TestBeegoConnectionManager_ShutdownNil(t *testing.T) {
	m := NewBeegoConnectionManager(nil)
	assert.Nil(t, m.Shutdown(context.Background()))
}

func TestNewBeegoLockFactory(t *testing.T) {
	ormDB := testBeegoDB(t)
	factory, err := NewBeegoLockFactory(ormDB)
	assert.Nil(t, err)
	assert.NotNil(t, factory)

	// 端到端加解锁
	ctx := context.Background()
	lock, err := factory.CreateLock("test-beego-lock")
	assert.Nil(t, err)
	assert.NotNil(t, lock)
	assert.Nil(t, lock.Lock(ctx, "owner-1"))
	assert.Nil(t, lock.UnLock(ctx, "owner-1"))
}

func TestNewBeegoLockFactory_Error(t *testing.T) {
	// db.DB 为 nil：底层 NewStorageBySqlDb 返回 db can not nil
	_, err := NewBeegoLockFactory(&orm.DB{})
	assert.NotNil(t, err)
}

func TestCreateStorageForBeego(t *testing.T) {
	ormDB := testBeegoDB(t)
	m := NewBeegoConnectionManager(ormDB)
	s, err := CreateStorageForBeego(ormDB, m)
	assert.Nil(t, err)
	assert.NotNil(t, s)

	_, err = CreateStorageForBeego(&orm.DB{}, m)
	assert.NotNil(t, err)
}

func TestInitGlobalBeegoLockFactory(t *testing.T) {
	ormDB := testBeegoDB(t)
	assert.Nil(t, InitGlobalBeegoLockFactory(ormDB))
	assert.NotNil(t, GlobalBeegoLockFactory)
	// 还原全局，避免污染其它测试
	GlobalBeegoLockFactory = nil

	assert.NotNil(t, InitGlobalBeegoLockFactory(&orm.DB{}))
	assert.Nil(t, GlobalBeegoLockFactory)
}
