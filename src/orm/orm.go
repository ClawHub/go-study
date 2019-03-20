package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func Demo() {
	var err error
	db, err = gorm.Open("mysql", "root:minivision@tcp(114.55.36.16:3307)/test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Ping
	err = db.DB().Ping()
	if nil != err {
		log.Fatal(err)
	}
	//设置连接池
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)

	//创建表
	//check has table or not
	//if !db.HasTable(&Like{}) {
	//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
	//		panic(err)
	//	}
	//}
	//插入
	//ip := "192.168.0.1"
	//ua := "ua2"
	//title := "title"
	//like := &Like{
	//	Ip:        ip,
	//	Ua:        ua,
	//	Title:     title,
	//	Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
	//	CreatedAt: time.Now(),
	//}

	//if err := db.Create(like).Error; err != nil {
	//	fmt.Println(err)
	//}

	//删除
	//hash := uint64(4429108899809532157)
	//if err := db.Where(&Like{Hash: hash}).Delete(Like{}).Error; err != nil {
	//	fmt.Println(err)
	//}
	//查找数量
	//var count int
	//err = db.Model(&Like{}).Where(&Like{Ip: ip, Ua: ua, Title: title}).Count(&count).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(count)
	//查找所有
	//var likes []Like
	//db.Model(&Like{}).Find(&likes)
	//fmt.Println(likes)
	//查询第一个
	//var likeFirst Like
	//db.Model(&Like{}).First(&likeFirst)
	//fmt.Println(likeFirst)
	//修改
	//db.Model(&Like{}).Update("ua", "hello ua")
	//db.Model(&Like{}).Updates(Like{Ip: "192.167.0.0"})
	//db.Model(&Like{}).Updates(Like{Ua: "", Ip: "6"}) // nothing update
	//事务
	//CreateAnimals(db)

	//执行原生SQL
	//db.Exec("UPDATE likes SET ip=? WHERE hash IN (?)", "sss", []int64{89316562029100364,1381825585541421043})

	//db.Exec("DROP TABLE demo;")
	//db.Raw("SELECT * FROM likes").Scan(&likeFirst)
	//rows, err := db.Raw("SELECT ip FROM likes").Rows() // (*sql.Rows, error)
	//defer rows.Close()
	//for rows.Next() {
	//	var ip string
	//	rows.Scan(&ip)
	//	fmt.Println(ip)
	//	// ScanRows scan a row into user
	//	db.ScanRows(rows, &likeFirst)
	//	fmt.Println(likeFirst)
	//}
	// Get generic database object sql.DB to use its functions

}

//事务
//func CreateAnimals(db *gorm.DB) error {
//	ip := "192.168.0.9"
//	ip2 := "192.168.0.8"
//	ua := "ua2"
//	title := "title"
//	like := &Like{
//		Ip:        ip,
//		Ua:        ua,
//		Title:     title,
//		Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
//		CreatedAt: time.Now(),
//	}
//	like2 := &Like{
//		Ip:        ip2,
//		Ua:        ua,
//		Title:     title,
//		Hash:      murmur3.Sum64([]byte(strings.Join([]string{ip2, ua, title}, "-"))) >> 1,
//		CreatedAt: time.Now(),
//	}
//
//	tx := db.Begin()
//	if err := tx.Create(like).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	if err := tx.Create(like2).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}
