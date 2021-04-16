package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq" // 這是 postgres 的 driver 需要做初始化
)

var redisClient *redis.Client
var db *sql.DB

func init() {
	logrus.Println("init()")
	//初始化 viper，設定預設讀取環境變數
	viper.SetConfigFile(".env")   //設定檔名
	viper.SetConfigType("dotenv") //設定型態 例如：dotenv, yaml
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("讀取設定檔案錯誤 Fatal error config file: %v\n", err)
	}

	//初始化 redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379", //因為在 docker-compose 的服務名稱為redis，不然一般來說是 localhost 或 ip
		//Password: "a12345", // no password set
		DB: 0, // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err == nil {
		logrus.Println("redis 回應成功，", pong)
	} else {
		logrus.Fatal("redis 無法連線，錯誤為", err)
	}

	//初始化 postgres sql 連線 sslmode=disable , sslmode=require
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_DATABASE"))
	//開啟連線
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logrus.Fatal("postgres 無法開啟，錯誤為", err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal("postgres Ping，錯誤為", err)
	} else {
		logrus.Println("postgres 回應成功")
	}
}

func main() {
	logrus.Info("HTTP server started")
	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	server := gin.Default()
	server.GET("/", GetTest)
	server.GET("/headers", GetHeaders)
	logrus.Fatal(server.Run(restfulHost + ":" + restfulPort))
}

// 測試是否能正常連接
func GetTest(c *gin.Context) {
	result := "{'msg':'test ok!'}"
	c.JSON(http.StatusOK, result)
}

// 回傳主機看到的 headers
func GetHeaders(c *gin.Context) {
	for name, headers := range c.Request.Header {
		for _, h := range headers {
			fmt.Fprintf(c.Writer, "%v: %v\n", name, h)
		}
	}
}
