package controllers

import (
	"CMS/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
)

// RoleController operations for Role
type RoleController struct {
	beego.Controller
}

// URLMapping ...
func (c *RoleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Role
// @Param	body		body 	models.Role	true		"body for Role content"
// @Success 201 {object} models.Role
// @Failure 403 body is empty
// @router / [post]
func (c *RoleController) Post() {
	var role models.Roles
	json.Unmarshal(c.Ctx.Input.RequestBody, &role)

	flag,errMap := models.RoleValidatorFields(&role)
	if !flag {
		c.Data["json"] = map[string]map[string]string{"err_field":errMap}
		c.ServeJSON()
		return
	}

	if role.Id != 0 {
		err := models.UpdateRolesById(&role)
		var str string
		if err == nil {
			str = "更新成功"
		} else {
			str = "更新失败"
		}
		c.Data["json"] = map[string]string{"msg": str}
	} else {
		roleId, err := models.AddRoles(&role)
		if err == nil {
			c.Data["json"] = map[string]string{"role_id": strconv.FormatInt(roleId, 10)}
		} else {
			c.Data["json"] = map[string]error{"err": err}
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Role by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Role
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RoleController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Role
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Role
// @Failure 403
// @router / [get]
func (c *RoleController) GetAll() {
	limit,err := c.GetInt64("limit")
	if err != nil {
		limit = 20
	}
	offset, err := c.GetInt64("offset")
	if err != nil {
		offset = 0
	}
	fmt.Println(limit)
	roles, count, err := models.GetAllRoles(nil,nil,[]string{"id"},[]string{"desc"},offset, limit)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"data":roles, "count": count}
	} else {
		c.Data["json"] = map[string]error{"err":err}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Role
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Role	true		"body for Role content"
// @Success 200 {object} models.Role
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RoleController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Role
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RoleController) Delete() {
	var msg string
	id, ok := strconv.ParseInt(c.Ctx.Input.Param(":id"),10 ,64)
	if ok != nil {
		msg = "参数错误，请检查！"
	} else {
		if err:= models.DeleteRoles(id);err == nil {
			msg = "执行成功"
		}else {
			msg = "执行失败"
		}
	}
	c.Data["json"] = map[string]string{"msg" : msg}
	c.ServeJSON()
}
