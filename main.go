package main

import (
	_rouletteHandlerHttpDelivery "Golang-Gambling-InOne/roulette/delivery/http"
	_rouletteRepo "Golang-Gambling-InOne/roulette/repository/postgresql"
	_rouletteUsecase "Golang-Gambling-InOne/roulette/usecase"

	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	melody "gopkg.in/olahol/melody.v1"

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
		Addr: viper.GetString("REDIS_ADDRESS"), // "redis:6379", //因為在 docker-compose 的服務名稱為redis，不然一般來說是 localhost 或 ip
		//Password: "a12345", // no password set
		DB: 0, // use default DB
	})
	pong, err := redisClient.Ping().Result()
	if err == nil {
		logrus.Println("redis 回應成功， Ping / ", pong)
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

	// 啟動前先刪除整個 redis db
	deleteMsg, err2 := redisClient.FlushDB().Result()
	fmt.Println("===================")
	fmt.Println(deleteMsg) // 回傳是否刪除成功
	fmt.Println(err2)      // 是否有錯誤訊息
	fmt.Println("===================")

	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")

	server.GET("/", func(c *gin.Context) {
		// 在http包使用的時候，註冊了/這個根路徑的模式處理，瀏覽器會自動的請求 favicon.ico，要注意處理，否則會出現兩次請求
		if c.Request.RequestURI == "/favicon.ico" {
			return
		}
		c.HTML(http.StatusOK, "index.html", nil)
	})

	webSocket := melody.New()

	// 建立 repository
	fmt.Println("=================== Create repository Instance")
	rouletteRepo := _rouletteRepo.NewPostgresqlRouletteRepository(db)

	// 建立 usecase
	rouletteUsecase := _rouletteUsecase.NewRouletteUsecase(rouletteRepo)

	// 建立路由, gin , melody, redis 是從外部丟進去的
	_rouletteHandlerHttpDelivery.NewRouletteHandler(server, webSocket, redisClient, rouletteUsecase)

	logrus.Info("HTTP server started")
	restfulHost := viper.GetString("RESTFUL_HOST")
	restfulPort := viper.GetString("RESTFUL_PORT")
	logrus.Fatal(server.Run(restfulHost + ":" + restfulPort))
}
