package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	// "reflect"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/korjavin/go-php-serialize"
	"golang.org/x/text/encoding/unicode"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const dburl = "root:@unix(/run/mysqld/mysqld.sock)/kino"

func init() {
	var err error
	db, err = sql.Open("mysql", dburl)
	if err != nil {
		log.Fatalf("mysql: %s", err)
	}
}
func main() {
	rows, err := db.Query("select id,films_ser FROM kp_people_films LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var ser string
	for rows.Next() {
		if err := rows.Scan(&id, &ser); err != nil {
			log.Fatal(err)
		}
		val, _ := phpserialize.DecodeWithEncoding(ser, unicode.UTF8)
		stringMap := modifyMap(val)
		jsonString, err := json.Marshal(stringMap)
		if err != nil {
			log.Fatal(err)
		}

		spew.Dump(stringMap, jsonString)

	}

}

func modifyMap(val interface{}) interface{} {
	stringMap := make(map[string]interface{})
	if map1, ok := val.(map[interface{}]interface{}); ok {
		for k, v := range map1 {
			stringKey := modifyValue(k)
			leaf1 := modifyMap(v)
			stringMap[stringKey] = leaf1
		}
	} else {
		return modifyValue(val)
	}
	return stringMap
}
func modifyValue(val interface{}) string {
	if s, ok := val.(string); ok {
		return s
	}
	if i, ok := val.(int64); ok {
		return strconv.FormatInt(i, 10)
	}
	return ""

}
