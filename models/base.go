package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER = "mysql"
	//DB_URL="root:cxx15755599659@tcp(127.0.0.1:3306)/myblog"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	DB_URL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, url, db)
	orm.RegisterDriver(DB_DRIVER, orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", DB_URL)
	if err != nil {
		fmt.Println(err)
	}
	orm.RegisterModel(new(TUser))

}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
