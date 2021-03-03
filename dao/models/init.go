package models

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/caches"
	"github.com/xormplus/xorm/log"

	//self library..
	. "github.com/kennyzhu2013/go-os/src/dbservice/conf"
	data "github.com/kennyzhu2013/go-os/src/dbservice/models"
	// "github.com/xormplus/xorm"
)

var (
	ErrNotExist = errors.New("not exist")

	//singleton instance..
    orm *xorm.Engine
)

func SyncAllTables() error {
	//struct sync...
	//if not define then create...
	/*err = orm.Sync2(new(Setting), new(Category), new(Post), new(Image),
		new(User), new(FavoritePost), new(Follow), new(Topic), new(FollowTopic),
		new(Page), new(Notification), new(Comment), new(Bulletin))
	if err != nil {
		panic(err)
	}*/
	return orm.Sync2(new(data.Preferences), )
}

//need test promode ..
func Init(isProMode bool) {
	var err error
	orm, err = xorm.NewEngine(OrmConf.DriverName, OrmConf.DataSource)
	if err != nil {
		panic(err)
	}

	err = orm.Ping()
	if err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns(OrmConf.MaxIdle)
	orm.SetMaxOpenConns(OrmConf.MaxOpen)
	if OrmConf.DebugLog {
		orm.Logger().SetLevel(log.LOG_DEBUG)
	} else {
		orm.Logger().SetLevel(log.LOG_INFO)
	}


	if !isProMode {
		orm.ShowSQL(true)

		//simple log..
		sqlWriter,_ := os.Create("./log/sql.log")
		logger := log.NewSimpleLogger(sqlWriter)
		logger.ShowSQL(true)
		orm.SetLogger(logger)
	}

	if OrmConf.IsCached {
		ormCache := caches.NewLRUCacher2(caches.NewMemoryStore(), OrmConf.CacheTime, OrmConf.CacheCount)
		orm.SetDefaultCacher(ormCache)
	}

	err = SyncAllTables()
	if err != nil {
		panic(err)
	}

	fmt.Println("Models Init successed...")
	//orm set...
	//social.SetORM(orm)
}
