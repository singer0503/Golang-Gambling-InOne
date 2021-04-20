一步一步記錄建置遊戲的方法

首先做好開發架構，預計部署在容器環境處理，方便部署。

環境需求: 需要安裝好 docker 以及 docker compose，且有網路的環境

執行方式

```bash
git clone https://github.com/singer0503/Golang-Gambling-InOne.git
cd Golang-Gambling-InOne
docker-compose up

```

測試是否正常

```bash
http://localhost:5005

http://localhost:5005/headers
```

![https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F014f78e1-e2bf-4714-859b-6c5a41b0817e%2F_2021-04-16_7.08.44.png?table=block&id=100be06d-139c-4187-9e0b-692f5cd91fe1&width=3250&userId=e8aa0888-ca7b-4216-869b-435a8115d8eb&cache=v2](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F014f78e1-e2bf-4714-859b-6c5a41b0817e%2F_2021-04-16_7.08.44.png?table=block&id=100be06d-139c-4187-9e0b-692f5cd91fe1&width=3250&userId=e8aa0888-ca7b-4216-869b-435a8115d8eb&cache=v2)

docker-compose 裡面有三個服務

- server：Golang Server 遊戲主程式
- db：使用 Postgres SQL 當成資料存放處，如：歷史勝敗資料，會員資料等
- redis：使用 Redis 做資料暫存處，如：當前的下注資料

![https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fe2841a7e-0e87-4118-95d8-3100d99698ea%2F_2021-04-16_7.04.29.png?table=block&id=9b07acaf-5103-4d5f-bbb5-2075c341bca3&width=3250&userId=e8aa0888-ca7b-4216-869b-435a8115d8eb&cache=v2](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fe2841a7e-0e87-4118-95d8-3100d99698ea%2F_2021-04-16_7.04.29.png?table=block&id=9b07acaf-5103-4d5f-bbb5-2075c341bca3&width=3250&userId=e8aa0888-ca7b-4216-869b-435a8115d8eb&cache=v2)

附錄指令有時候會需要進入 container 內，指令為

```bash
docker exec -it 001-golang-gambling-inone_server_1 /bin/sh

```