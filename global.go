package beego_locks

import "github.com/beego/beego/orm"

var GlobalBeegoLockFactory *BeegoLockFactory

func InitGlobalBeegoLockFactory(db *orm.DB) error {
	factory, err := NewBeegoLockFactory(db)
	if err != nil {
		return err
	}
	GlobalBeegoLockFactory = factory
	return nil
}
