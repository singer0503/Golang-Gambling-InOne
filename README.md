# Golang gambling all in one
一步一步記錄建置遊戲的方法

首先做好開發架構，預計部署在容器環境處理，方便部署。

環境需求: 需要安裝好 docker 以及 docker compose，且有網路的環境

執行方式
```bash
git pull https://github.com/singer0503/Golang-Gambling-InOne.git
cd Golang-Gambling-InOne
docker-compose up
```

docker-compose 裡面有三個服務
- server：遊戲主程式
- db：資料存放處，如：歷史勝敗資料，會員資料等
- redis：資料暫存處，如：當前的下注資料

有時候會需要進入 container 內，指令為
```bash
docker exec -it 001-golang-gambling-inone_server_1 /bin/sh
```