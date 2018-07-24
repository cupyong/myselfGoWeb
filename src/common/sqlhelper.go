package common

import (
	"database/sql"
	"./log"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

/**
读取全部
 */
func SqlReader(sqlStr string) (*sql.Rows, error) {
	db, err := sql.Open(DATABASE_TYPE, SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return nil, err
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	rows, err := db.Query(sqlStr)
	db.Close()
	return rows, err
}
func ThirdSqlReader(sqlStr string) (*sql.Rows, error) {
	db, err := sql.Open(DATABASE_TYPE, THIRD_SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return nil, err
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	rows, err := db.Query(sqlStr)
	db.Close()
	return rows, err
}

func SqlReadOne(sqlStr string) (*sql.Row,error) {
	db, err := sql.Open(DATABASE_TYPE,SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return nil,err
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	row := db.QueryRow(sqlStr)
	db.Close()
	return row,err
}

func SqlUpdate(sqlStr string) bool {
	db, err := sql.Open(DATABASE_TYPE,SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	db.Close()
	return true
}

func SqlInsert(sqlStr string) bool {
	db, err := sql.Open(DATABASE_TYPE,SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	_, err = result.LastInsertId()
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	db.Close()
	return true
}

func SqlDelete(sqlStr string) bool {
	db, err := sql.Open(DATABASE_TYPE,SQL_CONNECT)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	log.Log(log.Notice, "[Sql:%s][%s]", sqlStr, log.File_line())
	result, err := db.Exec(sqlStr)
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Log(log.Err, "[%s][%s]", err, log.File_line())
		return false
	}
	db.Close()
	return true
}
