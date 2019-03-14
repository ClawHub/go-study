package lmdb

import (
	"encoding/json"
	"fmt"
	"github.com/bmatsuo/lmdb-go/lmdb"
	"runtime"
)

//lmdb客户端结构体
type lmdbCliet struct {
	Env *lmdb.Env
	DBI lmdb.DBI
}

//close
func (c lmdbCliet) Close() {
	if c.Env != nil {
		err := c.Env.Close()
		fmt.Println(err)
	}
}

//get
func (c lmdbCliet) Get(key string) string {
	//默认值""
	var result string
	_ = c.Env.View(func(txn *lmdb.Txn) (err error) {
		v, err := txn.Get(c.DBI, []byte(key))
		if err != nil {
			return err
		}
		//获取成功
		result = string(v)
		return nil
	})
	return result
}

//set
func (c lmdbCliet) Set(key, val string) bool {
	//默认false
	var result bool
	_ = c.Env.Update(func(txn *lmdb.Txn) error {
		err := txn.Put(c.DBI, []byte(key), []byte(val), 0)
		if err != nil {
			return err
		}
		//插入成功
		result = true
		return nil

	})
	return result
}

//Del
func (c lmdbCliet) Del(key string) bool {
	//默认false
	var result bool
	_ = c.Env.Update(func(txn *lmdb.Txn) error {
		err := txn.Del(c.DBI, []byte(key), nil)
		if err != nil {
			return err
		}
		//删除成功
		result = true
		return nil
	})
	return result

}

//stat
func (c lmdbCliet) Stat() string {
	stat, err := c.Env.Stat()
	if err != nil {
		return "stat error"
	}
	//json解析
	bytes, err := json.Marshal(stat)
	if err != nil {
		return "json marshal error"
	}
	return string(bytes)
}

//batch put
func (c lmdbCliet) BatchPut(keyValMap map[string]string, commitNum int) bool {
	//使用锁，因为采用了mdb_txn_begin
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	//默认false
	var result bool
	//开始事务
	txn, err := (c.Env).BeginTxn(nil, 0)
	if err != nil {
		return result
	}
	//打开游标
	var cur *lmdb.Cursor
	cur, err = txn.OpenCursor(c.DBI)
	if err != nil {
		return result
	}

	//入游标数量
	var num = 1
	//迭代待处理数据
	for key, val := range keyValMap {
		//入游标
		err = cur.Put([]byte(key), []byte(val), 0)
		if err != nil {
			return result
		}
		//达到提交数量
		if num%commitNum == 0 {
			//提交事务
			err = txn.Commit()
			if err != nil {
				return result
			}
			//新开启事务
			txn, err := (c.Env).BeginTxn(nil, 0)
			if err != nil {
				return result
			}
			//新开游标
			cur, err = txn.OpenCursor(c.DBI)
			if err != nil {
				return result
			}
		}

		//入游标数量+1
		num++
	}

	//最后提交一次事务
	//提交事务
	err = txn.Commit()
	if err != nil {
		return result
	}
	return true
}
