package dao

import (
	"database/sql"
	"domain"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	username string = "root"
	password string = "root123!@#"
	host     string = "127.0.0.1"
	dbname   string = "jenkins"
)

func init() {
	dburl := username + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8"
	var err error
	//db, err := sql.Open("mysql", dburl) //不要这样，这样的话，db是新的局部变量了，不是全局变量
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
}

func Add(user domain.User) (id int64, err error) {
	stmt, err := db.Prepare("insert into t_user(name,password,register_date_time) values(?,?,?);")
	if err != nil {
		return id, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Name, user.Password, user.RegisterDateTime)
	if err != nil {
		return id, err
	}
	id, err = result.LastInsertId()

	return id, err
}

func QueryById(id int) (user domain.User, err error) {
	stmt, err := db.Prepare("select * from t_user where id=?;")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&(user.Id), &(user.Name), &(user.Password), &(user.RegisterDateTime)); err != nil {
		return user, err
	}
	return user, nil
}

func Delete(id int) (count int64, err error) {
	stmt, err := db.Prepare("delete from t_user where id = ?;")
	if err != nil {
		return count, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return count, err
	}

	count, err = result.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

func Update(user domain.User) (count int64, err error) {
	stmt, err := db.Prepare("update t_user set name=?, password=? where id = ?;")
	if err != nil {
		return count, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Name, user.Password, user.Id)
	if err != nil {
		return count, err
	}

	count, err = result.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}
