package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)


func yhendus(w http.ResponseWriter, req *http.Request) {
	//db, err := sql.Open("postgres", "user=bfmgwbvnvpqnps dbname=d3b1ljbjr0e123 sslmode=disable 
	//	password=D9jvuDlVzqp9BELsCAt0eNzesm host=ec2-54-225-201-25.compute-1.amazonaws.com ")
	db, err := sql.Open("postgres", "postgres://bfmgwbvnvpqnps:D9jvuDlVzqp9BELsCAt0eNzesm@ec2-54-225-201-25.compute-1.amazonaws.com:5432/d3b1ljbjr0e123")
	if err != nil {
		log.Fatal(err)
	}
stmt, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
if err != nil {
	log.Fatal(err)
}
res, err := stmt.Exec("Dolly")
if err != nil {
	log.Fatal(err)
}
lastId, err := res.LastInsertId()
if err != nil {
	log.Fatal(err)
}
rowCnt, err := res.RowsAffected()
if err != nil {
	log.Fatal(err)
}
log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

func main() {
	http.HandleFunc("/yhendus.go", yhendus)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}