package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Run() {
	InitMongodb()
	DoInsert()
	DoSelect()
	//DoDelete()
	DoUpdate()
}

var client *mongo.Client

func InitMongodb() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions.SetMaxPoolSize(50)
	// 连接到MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"nickname"`
	Age      int                `bson:"age"`
}

// DoSelect 查询操作
func DoSelect() {
	filter := bson.D{{}}
	// 查询单条数据
	var user UserEntity
	err := client.Database("homestead").Collection("users").FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %v, Id: %v", user, user.ID.Hex())
	// 查询多条数据
	var users []*UserEntity
	findOptions := options.Find()
	findOptions.SetLimit(100)
	cur, err := client.Database("homestead").Collection("users").Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 查找多个文档返回一个光标 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		var elem UserEntity // 创建一个值，将单个文档解码为该值
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	defer cur.Close(context.TODO())
	fmt.Println(users)
}

// DoInsert 插入操作
func DoInsert() {
	user := UserEntity{Username: "bom", Age: 22}
	// one insert
	InsertOneRes, err := client.Database("homestead").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert result id: %v \n", InsertOneRes.InsertedID)
	// batch insert
	users := []interface{}{UserEntity{Username: "li3", Age: 22}, UserEntity{Username: "li4", Age: 22}}
	InsertManyRes, err := client.Database("homestead").Collection("users").InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert result ids: %v \n", InsertManyRes.InsertedIDs)
}

// DoDelete 删除操作
func DoDelete() {
	//idPrimitive, err := primitive.ObjectIDFromHex("6151e2613e9ec078ce8d76ef")
	//if err != nil {
	//	log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	//}
	// bson.D{{"_id",idPrimitive}}

	// del one
	DeleteOneRes, err := client.Database("homestead").Collection("users").DeleteOne(context.TODO(), bson.D{{"nickname", "bom"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents\n", DeleteOneRes.DeletedCount)

	// del all
	DeleteManyRes, err := client.Database("homestead").Collection("users").DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents\n", DeleteManyRes.DeletedCount)
}

// DoUpdate 更新操作
func DoUpdate() {
	filter := bson.D{{"nickname", "bom"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	// UpdateOne
	updateOneResult, err := client.Database("homestead").Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update one success %v \n", updateOneResult.ModifiedCount)

	// UpdateMany
	UpdateManyResult, err := client.Database("homestead").Collection("users").UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update many success %v \n", UpdateManyResult.ModifiedCount)
}
