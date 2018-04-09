package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20180319_153405 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20180319_153405{}
	m.Created = "20180319_153405"

	migration.Register("Users_20180319_153405", m)
}

// Run the migrations
func (m *Users_20180319_153405) Up() {
	m.SQL("create table users (" +
		"`id` int unsigned auto_increment," +
		"username varchar(64) not null," +
		"email varchar(64) not null ," +
		"password varchar(100) not null," +
		"role_id int unsigned not null," +
		"parent_id int unsigned not null default 0," +
		"head_img varchar(100) not null default ''," +
		"created_at int unsigned," +
		"UNIQUE unique_username (username(64))," +
		"UNIQUE unique_email (email(64))," +
		"INDEX index_role_id (role_id)," +
		"primary key (`id`)" +
		")")

}

// Reverse the migrations
func (m *Users_20180319_153405) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("drop tables users")
}
