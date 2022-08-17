package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql TODO
	"github.com/jmoiron/sqlx"
)

// UserModel SqlX框架去做数据分类
type UserModel struct {
	Id       int            `db:"id" form:"id"`
	Email    string         `db:"email" form:"email" binding:"email"`
	Password string         `db:"password" form:"password"`
	Avatar   sql.NullString `db:"avatar"`
}

var db_sqlx *sqlx.DB

// ConnectSqlxDatabases 连接数据库
func ConnectSqlxDatabases() {
	var err error
	db_sqlx, err = sqlx.Connect("mysql", "root:Xcy3329257@(localhost)/ginhello?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Printf("connect DB failed, err:%v\n", err)
		return
	}
	log.Println("sqlx数据库连接接成功")
}

// CloseSqlxDatabases TODO
func CloseSqlxDatabases() {
	db_sqlx.Close()
	log.Println("sqlx数据库关闭成功")
}

// Save 插入数据
func (user *UserModel) Save() int64 {
	result, err := db_sqlx.Exec(
		"INSERT INTO ginhello.user(email, password) VALUES (?,?)",
		user.Email, user.Password)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		return -1
	}
	// 插入成功后，获取insert id
	// mysql表如果存在自增id，则可以通过Exec返回的结果对象的LastInsertId，查询新插入数据的ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

// QueryByEmail 根据email查询单条数据
func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	sqlStr := "select * from ginhello.user where email = ?"
	err := db_sqlx.Get(&u, sqlStr, user.Email)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		log.Panicln(err)
	}
	return u
}

// QueryById 根据id查询单条数据
func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	sqlStr := "select * from ginhello.user where id = ?"
	err := db_sqlx.Get(&u, sqlStr, id)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		log.Panicln(err)
	}
	return u, err
}

// Update 根据id更新数据表
func (user *UserModel) Update(id int) error {
	_, err := db_sqlx.Exec("update user set password=?,avatar=? where id=?", user.Password, user.Avatar.String, id)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
	}
	return err
}

// User_Sqlx TODO
type User_Sqlx struct {
	UserId   int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

// QueryRow TODO:测试使用
func QueryRow() {
	u := UserModel{}
	// sqlStr := "select * from ginhello.user"
	sqlStr := "select * from ginhello.user where email = ?"
	err := db_sqlx.Get(&u, sqlStr, "30513207@qq.com")
	if err != nil {
		log.Printf("tomxiang:::get failed, err:%v\n", err)
		log.Panicln(err)
		return
	}
	log.Printf("tomxiang:::get success\n")
}
