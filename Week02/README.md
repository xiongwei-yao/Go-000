学习笔记

    // package database/sql 的错误信息 sql.ErrNoRows 出现过程
    // 1、执行Conn.QueryRowContext、DB.QueryRow、DB.QueryRowContext、Stmt.QueryRow、Stmt.QueryRowContext、
    //	Tx.QueryRow、Tx.QueryRowContext等查询操作返回Row对象（结果Row对象中记录不超过一条）
    // 2、执行Row的Scan方法，发现结果集Rows不包含记录，返回sql.ErrNoRows
    // sql.ErrNoRows 的处置：
    //		返回sql.ErrNoRows说明数据库操作没有问题，只是查询结果无数据，不属于程序异常、操作异常不需要Wrap抛给上层。
    //		Scan返回的其它error需要Wrap抛给上层。