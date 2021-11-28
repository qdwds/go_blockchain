package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main(){
	db,err := bolt.Open("my.db",0600,nil)
	defer db.Close()
	if err != nil{
		log.Fatal(err)
	}

	//	创建表
	err = db.Update(func(tx *bolt.Tx) error {
		//	创建 表
		b,err := tx.CreateBucket([]byte("blockBucket"))
		if err != nil{
			return fmt.Errorf("create blockBucket errpr")
		}

		//	表里面存储数据
		if b != nil {
			err := b.Put([]byte("key"),[]byte("value"))
			if err != nil{
				log.Panic("数据库存储数据失败")
			}
		}
		return nil
	})
	if err != nil{
		log.Panic(err)
	}

}