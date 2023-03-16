package dbconn

import (
	"database/sql"
	"fmt"
	"time"
)

type Member struct {
	Id         int
	Username   string
	Password   string
	Created_at string
	Updated_at string
	Deleted_at sql.NullString
}

var db *sql.DB

func InitDB() (err error) {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1)/upm")
	if err != nil {
		fmt.Printf("db open err : %s\n", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("db ping err : %s\n", err)
		return err
	}

	return nil
}

func Findone(username string) Member {
	var m Member
	err := db.QueryRow("select id, username,password,created_at,updated_at,deleted_at from members where username = ?", username).Scan(&m.Id, &m.Username, &m.Password, &m.Created_at, &m.Updated_at, &m.Deleted_at)
	if err != nil {
		fmt.Printf("findone data failed err :%s\n", err)
	}

	fmt.Printf("findone member info %v\n", m)
	return m
}

// 查询多条数据
func FindsData() []Member {
	var s []Member
	rows, err := db.Query("select id,username,password,created_at,updated_at,deleted_at from `members`")
	if err != nil {
		fmt.Printf("findsData failed err:%s\n", err)
		return nil
	}
	for rows.Next() {
		var m Member
		err = rows.Scan(&m.Id, &m.Username, &m.Password, &m.Created_at, &m.Updated_at, &m.Deleted_at)
		if err != nil {
			fmt.Printf("findsData scan failed err:%s\n", err)
			return nil
		}
		fmt.Printf("findsData User info %v\n", m)
		s = append(s, m)
	}
	fmt.Printf("%v\n", s)

	return s
}

// 插入一条数据
func InsertData(username string, password string) (err error) {
	// 增、改、删 使用Exec方法
	exec, err := db.Exec("insert into members(username,password,created_at,updated_at,deleted_at) values (?,?,?,?,?)", username, password, time.Now(), time.Now(), nil)
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return err
	}
	id, err := exec.LastInsertId() // 往表中最后追加一条数据
	if err != nil {
		fmt.Printf("exec insert failed err:%s\n", err)
		return err
	}
	fmt.Printf("insert data id is : %d\n", id)
	return nil
}

// 更新数据
func UpdateData(id int, username, password string) {
	ret, err := db.Exec("update members set username = ?,password = ? where id = ?", username, password, id)
	if err != nil {
		fmt.Printf("update failed err:%s\n", err)
	}
	affected, err := ret.RowsAffected() // 返回受影响的函数
	if err != nil {
		fmt.Printf("update RowsAffected failed err:%s\n", err)
	}
	fmt.Printf("update success rows:%d\n", affected)
}

// 删除数据
func DelData(id int) {
	ret, err := db.Exec("delete from members where id = ?", id)
	if err != nil {
		fmt.Printf("del failed err:%s\n", err)
		return
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed err:%s\n", err)
		return
	}
	fmt.Printf("update success rows:%d\n", affected)
}

func Close() {
	defer db.Close()
}
