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
	rows, _ := db.Query("select id, films_ser from kp_people_films limit 1")
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

type leaf struct {
	stringMap   map[string]interface{}
	stringValue string
	isValue     bool
}

func modifyMap(val interface{}) leaf {
	stringMap := make(map[string]interface{})
	if map1, ok := val.(map[interface{}]interface{}); ok {
		for k, v := range map1 {
			stringKey := modifyValue(k)
			leaf1 := modifyMap(v)
			if leaf1.isValue {
				stringMap[stringKey] = leaf1.stringValue
			} else {
				stringMap[stringKey] = leaf1.stringMap
			}
		}
	} else {
		return leaf{nil, modifyValue(val), true}
	}
	return leaf{stringMap, "", false}
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
func traverseMap(key interface{}, val interface{}) {
	switch val.(type) {
	default:
		fmt.Printf("%v => %v \n", key, val)
	case map[interface{}]interface{}:
		map1 := val.(map[interface{}]interface{})
		for k, v := range map1 {
			traverseMap(k, v)
		}
	}
}
