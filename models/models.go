package models

import "github.com/astaxie/beego/orm"

func init()  {
	//注册model
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
}

type User struct {
	Id      int      `orm:"column(id);auto" description:"id"`
	Name    string   `orm:"column(name)" description:"name"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
	//Posts   []*Post  `orm:"reverse(many)"`  //one2many情况下，不要写reverse
}

type Profile struct {
	Id  int   `orm:"column(id);auto" description:"id"`
	Age int `orm:"column(age)" description:"age"`
	//User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

type Post struct {
	Id    int    `orm:"column(id);auto" description:"id"`
	Title string `orm:"column(title)" description:"title"`
	User  *User  `orm:"rel(fk)"`  // OneToMany relation
	Tags  []*Tag `orm:"rel(m2m)"` // m2m relation
}

type Tag struct {
	Id    int     `orm:"column(id);auto" description:"id"`
	Name  string  `orm:"column(name)" description:"name"`
	Posts []*Post `orm:"reverse(many)"`
}

