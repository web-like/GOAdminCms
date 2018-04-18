package models

import (
	"strings"

	"github.com/astaxie/beego/orm"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/astaxie/beego/logs"
)

type Users struct {
	Id       int64
	Username string `orm:"size(128)"`
	Password string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
	RoleId   int64
	ParentId int64
	HeadImg  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Users))
}

// AddUsers insert a new Users into database and returns
// last inserted Id on success.
func AddUsers(m *Users) (id int64, err error) {

	m.getPasswordString()
	return

	//o := orm.NewOrm()
	//id, err = o.Insert(m)
	//return
}

// GetUsersById retrieves Users by Id. Returns error if
// Id doesn't exist
func GetUsersById(id int64) (v *Users, err error) {
	o := orm.NewOrm()
	v = &Users{Id: id}
	if err = o.QueryTable(new(Users)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsers retrieves all Users matches certain condition. Returns empty list if
// no records exist
func GetAllUsers(query map[string]string, limit int64,offset int64 ) (UserData []Users, number int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Users))

	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}

	//var l []Users
	number, _ = qs.Count()
	_, err = qs.OrderBy("-id").Limit(limit,offset).All(&UserData)
	//if err == nil {
	//	for _, v := range l {
	//		UserData = append(UserData,v)
	//	}
	//}
	return UserData,number,err
}

// UpdateUsers updates Users by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersById(m *Users) (err error) {
	o := orm.NewOrm()
	v := Users{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsers deletes Users by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsers(id int64) (err error) {
	o := orm.NewOrm()
	v := Users{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Users{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func (this *Users) getPasswordString() {
	hash, err := bcrypt.GenerateFromPassword([]byte(this.Password),bcrypt.DefaultCost)
	if err != nil {
		logs.Info("get user password fail",err)
	}
	this.Password = string(hash)
}

func (this *Users) checkPassword (inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(this.Password),[]byte(inputPassword))

	if err == nil {
		return  true
	}
	return  false
}
