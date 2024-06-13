
### Sebelum menjalankan aplikasi, disarankan untuk menginstall beberapa environment dibawah ini

- Install WSL
- Install Make
- Install Golang
- Install Sqlc
- Install Docker
- Pull Postgres to Docker
- [Install Golang Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Detail config untuk data source postgre
- Host : localhost
- Port : 5432
- User : root
- Password : secret
- Database : bank_app

### Buka directory project di dalam command prompt dan jalankan perintah berikut sesuai dengan urutan
- make postgres => untuk membuat container postgres
- make migrateup => membuat migrasi tabel ke dalam database
- go mod tidy => untuk mengintall semua package yang digunakan
- make server => start server golang

## gRPC API Reference

#### Get User Balance

```http
  localhost:9090 | BankApp/GetAccount
```

| Parameter | Type    | Description                |
|:----------|:--------| :------------------------- |
| `id`       | `int64` | **Required** |

#### Create Transfer

```http
  localhost:9090 | BankApp/CreateTransfer
```

| Parameter         | Type    | Description                 |
|:------------------|:--------| :-------------------------- |
| `from_account_id` | `int64` | **Required**. |
| `to_account_id`   | `int64` | **Required**.                       |         |               |
| `amount`          | `int64` | **Required**.                     |         |                                     |

#### Get User Transaction

```http
  localhost:9090 | BankApp/GetEntry
```

| Parameter    | Type    | Description                |
|:-------------|:--------| :------------------------- |
| `account_id` | `int64` | **Required** |
| `page_id`    | `int64` | **Required**             |         |              |
| `page_size`    | `int64` | **Required**         |         |                          |


