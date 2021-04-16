# Golang gambling all in one

一步一步記錄建置遊戲的方法

首先做好開發架構，預計部署在容器環境處理，方便部署。

環境需求: 需要安裝好 docker 以及 docker compose，且有網路的環境

執行方式

```bash
git clone <https://github.com/singer0503/Golang-Gambling-InOne.git>
cd Golang-Gambling-InOne
docker-compose up

```

測試是否正常

```bash
http://localhost:5005

http://localhost:5005/headers
```

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/014f78e1-e2bf-4714-859b-6c5a41b0817e/_2021-04-16_7.08.44.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/014f78e1-e2bf-4714-859b-6c5a41b0817e/_2021-04-16_7.08.44.png)

docker-compose 裡面有三個服務

- server：Golang Server 遊戲主程式
- db：使用 Postgres SQL 當成資料存放處，如：歷史勝敗資料，會員資料等
- redis：使用 Redis 做資料暫存處，如：當前的下注資料

![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e2841a7e-0e87-4118-95d8-3100d99698ea/_2021-04-16_7.04.29.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e2841a7e-0e87-4118-95d8-3100d99698ea/_2021-04-16_7.04.29.png)

附錄指令有時候會需要進入 container 內，指令為

```bash
docker exec -it 001-golang-gambling-inone_server_1 /bin/sh

```