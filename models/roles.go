package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Roles struct {
	Id    int64
	Title string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Roles))
}

// AddRoles insert a new Roles into database and returns
// last inserted Id on success.
func AddRoles(m *Roles) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRolesById retrieves Roles by Id. Returns error if
// Id doesn't exist
func GetRolesById(id int64) (v *Roles, err error) {
	o := orm.NewOrm()
	v = &Roles{Id: id}
	if err = o.QueryTable(new(Roles)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRoles retrieves all Roles matches certain condition. Returns empty list if
// no records exist
func GetAllRoles(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Roles))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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
					return nil,0, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	// count num:
	if count, err := qs.Count(); err == nil {
		var l []Roles
		qs = qs.OrderBy(sortFields...).RelatedSel()
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
			return ml, count, nil
		}
	}


	return nil, 0, err
}

// UpdateRoles updates Roles by Id and returns error if
// the record to be updated doesn't exist
func UpdateRolesById(m *Roles) (err error) {
	o := orm.NewOrm()
	v := Roles{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}

	return
}

// DeleteRoles deletes Roles by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRoles(id int64) (err error) {
	o := orm.NewOrm()
	v := Roles{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Roles{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 检查用户提交的分组信息是否符合要求
func RoleValidatorFields(this *Roles) (flag bool,errInfo map[string]string){
	valid := validation.Validation{}
	errInfo = make(map[string]string,1)
	valid.MinSize(this.Title,2,"title").Message("用户组最少2个字符")
	valid.MaxSize(this.Title,20,"title").Message("用户组最多20个字符")
	flag = true
	if valid.HasErrors() {
		flag = false
		for _,err := range valid.Errors {
			errInfo[err.Key] = err.Message
		}
	}else if 0 == this.Id  {
		o := orm.NewOrm()
		role := Roles{Title:this.Title}
		if err := o.Read(&role,"Title"); err != orm.ErrNoRows {
			flag = false
			errInfo["title"] = "该用户组已经存在"
		}
	}
	return
}

