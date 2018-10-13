package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"os"
)

func main() {

	client, err := mongo.Connect(context.Background(), "mongodb://root:pass@ixiaotang.cn:27017", nil)

	if err != nil {
		fmt.Println("err Connect ", err)
		os.Exit(1)
	}
	db := client.Database("test_db")

	person := db.Collection("person")

	//  查询数据
	if err := Find(person); err != nil {
		fmt.Println("err Find ", err)
		os.Exit(1)
	}

	//插入数据
	data := map[string]interface{}{"name": "xiaogang2", "age": 22,
		"department": map[string]interface{}{"$id": "1", "$ref": "department"}}

	if err := InsertOne(person, data); err != nil {
		fmt.Println("err InsertOne ", err)
		os.Exit(1)
	}

	// 再次查询
	if err := Find(person); err != nil {
		fmt.Println("err Find ", err)
		os.Exit(1)
	}

}

// 查询
func Find(person *mongo.Collection) error {
	if cursor, err := person.Find(context.Background(), bson.NewDocument()); err != nil {

		fmt.Println("err Find ", err)
		return err
	} else {
		for cursor.Next(context.Background()) {

			var m = make(map[string]interface{})
			cursor.Decode(&m)
			for key, value := range m {
				fmt.Println(key, value)
			}
			fmt.Println(" ====== ")
		}
	}
	return nil
}

// 插入一条
func InsertOne(person *mongo.Collection, data map[string]interface{}) error {
	_, err := person.InsertOne(context.Background(), data)
	return err
}
