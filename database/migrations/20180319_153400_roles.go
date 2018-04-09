package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Roles_20180319_153400 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Roles_20180319_153400{}
	m.Created = "20180319_153400"

	migration.Register("Roles_20180319_153400", m)
}

// Run the migrations
func (m *Roles_20180319_153400) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("create table roles (" +
		"`id` int unsigned auto_increment," +
		"`title` varchar(64)," +
		"primary key (`id`)" +
	")")

}

// Reverse the migrations
func (m *Roles_20180319_153400) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("drop table roles")
}
