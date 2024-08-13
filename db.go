package main

import (
	"fmt"
	"hash/fnv"

	"go.mills.io/bitcask/v2"
)

type db struct {
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprint(h.Sum32())
}

func (d *db) saveData(article Article) {
	articleHash := hash(article.title)
	fmt.Println("Saving data " + articleHash)
	db, _ := bitcask.Open("article.db")
	defer db.Close()
	db.Put([]byte(articleHash), []byte("true"))
}

func (d *db) checkData(article Article) bool {
	articleHash := hash(article.title)
	db, _ := bitcask.Open("article.db")
	defer db.Close()
	data, _ := db.Get([]byte(articleHash))
	if data == nil {
		return false
	}
	fmt.Println("Data already exists " + articleHash)
	return true
}

func (d *db) deleteData(article Article) {
	articleHash := hash(article.title)
	fmt.Println("Delete data " + articleHash)
	db, _ := bitcask.Open("article.db")
	defer db.Close()
	db.Delete([]byte(articleHash))
}
