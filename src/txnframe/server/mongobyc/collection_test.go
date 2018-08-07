/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mongobyc

import (
	"context"
	"fmt"
	"testing"
)

type keyval struct {
	Key string `json:"key"`
}

func TestInsertOne(t *testing.T) {

	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertOne(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

}

func TestInsertMany(t *testing.T) {

	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertMany(context.Background(), []interface{}{`{"key":"uri"}`, `{"key":"uri2"}`})
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

}

func TestUpdateOne(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertOne(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

	_, err = mongo.Collection("uri_test").UpdateOne(context.Background(), `{"key":{"$eq":"uri"}}`, `{"$set":{"key":"urid"}}`)
	if nil != err {
		fmt.Println("failed to update:", err.Error())
		return
	}

}

func TestUpdateMany(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertOne(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

	_, err = mongo.Collection("uri_test").UpdateMany(context.Background(), `{"key":{"$eq":"uri"}}`, `{"$set":{"key":"uridmany"}}`)
	if nil != err {
		fmt.Println("failed to update:", err.Error())
		return
	}

}

func TestDeleteOne(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertOne(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

	_, err = mongo.Collection("uri_test").DeleteOne(context.Background(), `{"key":{"$eq":"urid"}}`)
	if nil != err {
		fmt.Println("failed to update:", err.Error())
		return
	}

}

func TestDeleteMany(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	_, err := mongo.Collection("uri_test").DeleteMany(context.Background(), `{"key":{"$eq":"urid"}}`)
	if nil != err {
		fmt.Println("failed to update:", err.Error())
		return
	}

}

func TestDeleteCollection(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	if err := mongo.Collection("uri_test").Drop(context.Background()); nil != err {
		fmt.Println("failed to drop collection:", err.Error())
		return
	}
}

func TestCount(t *testing.T) {
	InitMongoc()
	defer CleanupMongoc()

	mongo := NewClient("mongodb://test.mongoc:27017", "db_name_uri")
	if err := mongo.Open(); nil != err {
		fmt.Println("failed  open:", err)
		return
	}
	defer mongo.Close()

	err := mongo.Collection("uri_test").InsertOne(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to insert:", err)
		return
	}

	cnt, err := mongo.Collection("uri_test").Count(context.Background(), `{"key":"uri"}`)
	if nil != err {
		fmt.Println("failed to count:", err.Error())
		return
	}
	fmt.Println("cnt:", cnt)

}
