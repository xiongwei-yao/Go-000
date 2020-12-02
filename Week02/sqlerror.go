//
// package database/sql 的错误信息 sql.ErrNoRows 出现过程
// 1、执行Conn.QueryRowContext、DB.QueryRow、DB.QueryRowContext、Stmt.QueryRow、Stmt.QueryRowContext、
//		Tx.QueryRow、Tx.QueryRowContext等查询操作返回Row对象（结果Row对象中记录不超过一条）
// 2、执行Row的Scan方法，发现结果集Rows不包含记录，返回sql.ErrNoRows
// sql.ErrNoRows 的处置：
//		返回sql.ErrNoRows说明数据库操作没有问题，只是查询结果无数据，不属于程序异常、操作异常不需要Wrap抛给上层。
//		Scan返回的其它error需要Wrap抛给上层。

package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// QueryUserName QueryUserName
func QueryUserName(uid int) (uname string, found bool, err error) {
	db := *sql.DB
	var name string
	err = db.QueryRow(`select username from user where user.id = :id;`, sql.Named("id", uid)).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return string(""), false, nil
		}
		return string(""), false, errors.Wrap(err, "Database error")
	}
	return name, true, nil
}

// Operator Operator
func Operator() error {
	uid := 1
	un, found, err := QueryUserName(uid)
	if err != nil {
		return errors.WithMessage(err, "Operator")
	}
	if found {
		fmt.Printf(un)
	} else {
		fmt.Printf("Not found user")
	}
	return nil
}
