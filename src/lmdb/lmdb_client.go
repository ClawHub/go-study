package lmdb

import (
	"encoding/json"
	"fmt"
	"github.com/bmatsuo/lmdb-go/lmdb"
	"go-study/src/log"
	"go.uber.org/zap"
	"runtime"
)

//lmdb客户端结构体  只创建一个lmdb环境，一个DB库
type lmdbCliet struct {
	Env *lmdb.Env
	DBI lmdb.DBI
}

/**
 * 获取lmdb客户端
 *size: sets the size of the environment memory map.
 *path: Open an environment handle path
 *name: dbi name
 */
func NewClient(size int64, path string, name string) *lmdbCliet {
	//获取lmdb环境
	env, err := lmdb.NewEnv()
	if err != nil {
		log.LmdbLogger.Error("获取lmdb环境失败", zap.Error(err))
	}

	//设置最大库数量
	err = env.SetMaxDBs(1)
	if err != nil {
		log.LmdbLogger.Error("设置最大库数量失败", zap.Int("maxDBs", 1), zap.Error(err))
	}
	//设置最大存储
	err = env.SetMapSize(size)
	if err != nil {
		log.LmdbLogger.Error("设置最大存储失败", zap.Int64("mapSize", size), zap.Error(err))
	}
	//环境打开,Open an environment handle
	//modo:UNIX权限，用于设置创建的文件和信号量。
	err = env.Open(path, 0, 0644)
	if err != nil {
		log.LmdbLogger.Error("环境打开失败", zap.String("path", path), zap.Error(err))
	}
	//ReaderCheck从reader lock表中清除陈旧的条目，并返回清除的条目数量。
	readers, err := env.ReaderCheck()
	if err != nil {
		log.LmdbLogger.Error("ReaderCheck失败", zap.Int("readers", readers), zap.Error(err))
	} else {
		log.LmdbLogger.Info("ReaderCheck", zap.Int("readers", readers))
	}
	//创建DBI
	var dbi lmdb.DBI
	err = env.Update(func(txn *lmdb.Txn) (err error) {
		//dbi, err = txn.CreateDBI("example")
		dbi, err = txn.CreateDBI(name)
		fmt.Println(err)
		return
	})
	if err != nil {
		log.LmdbLogger.Error("创建DBI失败", zap.String("dbName", name), zap.Error(err))
	}
	return &lmdbCliet{env, dbi}
}

//close
func (c *lmdbCliet) Close() {
	if c.Env != nil {
		err := c.Env.Close()
		if err != nil {
			log.LmdbLogger.Error("Close fail", zap.Error(err))
		}
	}
}

//get
func (c *lmdbCliet) Get(key string) string {
	//默认值""
	var result string
	_ = c.Env.View(func(txn *lmdb.Txn) (err error) {
		v, err := txn.Get(c.DBI, []byte(key))
		if err != nil {
			if err != nil {
				log.LmdbLogger.Error("Get fail", zap.String("key", key), zap.Error(err))
			}
			return err
		}
		//获取成功
		result = string(v)
		return nil
	})
	return result
}

//set
func (c *lmdbCliet) Set(key, val string) bool {
	//默认false
	var result bool
	_ = c.Env.Update(func(txn *lmdb.Txn) error {
		err := txn.Put(c.DBI, []byte(key), []byte(val), 0)
		if err != nil {
			if err != nil {
				log.LmdbLogger.Error("Set fail", zap.String("key", key), zap.String("val", val), zap.Error(err))
			}
			return err
		}
		//插入成功
		result = true
		return nil

	})
	return result
}

//Del
func (c *lmdbCliet) Del(key string) bool {
	//默认false
	var result bool
	_ = c.Env.Update(func(txn *lmdb.Txn) error {
		err := txn.Del(c.DBI, []byte(key), nil)
		if err != nil {
			if err != nil {
				log.LmdbLogger.Error("Del fail", zap.String("key", key), zap.Error(err))
			}
			return err
		}
		//删除成功
		result = true
		return nil
	})
	return result

}

//stat
func (c *lmdbCliet) Stat() string {
	stat, err := c.Env.Stat()
	if err != nil {
		if err != nil {
			log.LmdbLogger.Error("Stat fail", zap.Error(err))
		}
		return "stat error"
	}
	//json解析
	bytes, err := json.Marshal(stat)
	if err != nil {
		log.LmdbLogger.Error("Stat json Marshal fail", zap.Error(err))
		return "json marshal error"
	}
	return string(bytes)
}

//batch put
func (c *lmdbCliet) BatchPut(keyValMap map[string]string, commitNum int) bool {
	//使用锁，因为采用了mdb_txn_begin
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	//默认false
	var result bool
	//开始事务
	txn, err := (c.Env).BeginTxn(nil, 0)
	if err != nil {
		log.LmdbLogger.Error("BeginTxn fail", zap.Error(err))
		return result
	}
	//打开游标
	var cur *lmdb.Cursor
	cur, err = txn.OpenCursor(c.DBI)
	if err != nil {
		log.LmdbLogger.Error("OpenCursor fail", zap.Error(err))
		return result
	}

	//入游标数量
	var num = 1
	//迭代待处理数据
	for key, val := range keyValMap {
		//入游标
		err = cur.Put([]byte(key), []byte(val), 0)
		if err != nil {
			log.LmdbLogger.Error("cur.Put fail", zap.String("key", key), zap.String("val", val), zap.Error(err))
			return result
		}
		//达到提交数量
		if num%commitNum == 0 {
			//提交事务
			err = txn.Commit()
			if err != nil {
				log.LmdbLogger.Error("txn.Commit() fail", zap.Error(err))
				return result
			}
			//新开启事务
			txn, err := (c.Env).BeginTxn(nil, 0)
			if err != nil {
				log.LmdbLogger.Error("BeginTxn fail", zap.Error(err))
				return result
			}
			//新开游标
			cur, err = txn.OpenCursor(c.DBI)
			if err != nil {
				log.LmdbLogger.Error("OpenCursor fail", zap.Error(err))
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
		log.LmdbLogger.Error("txn.Commit() fail", zap.Error(err))
		return result
	}
	return true
}
