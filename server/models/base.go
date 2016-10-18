//models初始化等操作
package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("mysql", "root:@/goshopping?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	err = x.Sync(new(User))
	if err != nil {
		log.Fatal(err)
	}
}
