package models

import (
	"beeApi/utils"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TUser struct {
	Id        int    `orm:"column(t_id);auto" json:"id"`
	TUserName string `orm:"column(t_user_name);size(100)" json:"userName"`
	TPwd      string `orm:"column(t_pwd);size(100)" json:"pwd"`
	TSalt     string `orm:"column(t_salt);size(100)" json:"salt"`
}

func (t *TUser) TableName() string {
	return "t_user"
}

// AddTUser insert a new TUser into database and returns
// last inserted Id on success.
func AddTUser(m *TUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTUserById retrieves TUser by Id. Returns error if
// Id doesn't exist
func GetTUserById(id int) (v *TUser, err error) {
	o := orm.NewOrm()
	v = &TUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTUser retrieves all TUser matches certain condition. Returns empty list if
// no records exist
func GetAllTUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTUser updates TUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateTUserById(m *TUser) (err error) {
	o := orm.NewOrm()
	v := TUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTUser deletes TUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTUser(id int) (err error) {
	o := orm.NewOrm()
	v := TUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

/***
LoginSuccess  示范一个例子 非真用jwt
*/
type LoginReturn struct {
	Token string `json:"token"`
	*TUser
}

func Login(name, pwd string) (*LoginReturn, error) {
	o := orm.NewOrm()
	var user = &TUser{}
	tableName := user.TableName()
	err := o.QueryTable(tableName).Filter("t_user_name", name).Filter("t_pwd", pwd).One(user)
	if err != nil {
		return nil, err
	}
	var loginReturn = &LoginReturn{}
	loginReturn.TUser = user
	//设置token
	loginReturn.Token = utils.CreateTocken(strconv.Itoa(user.Id))
	return loginReturn, nil
}

func GetUser(user TUser) (*TUser, error) {
	o := orm.NewOrm()
	var returnUser = &TUser{}
	tableName := returnUser.TableName()
	err := o.QueryTable(tableName).Filter("t_user_name", user.TUserName).Filter("t_pwd", user.TPwd).One(returnUser)
	if err != nil {
		return returnUser, err
	}
	return returnUser, nil
}
