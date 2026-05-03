## Sistem Informasi Siswa Via Golang

## Clean cache ketika package error

Gunakan perintah dibawah ini untuk membersihkan package
```
go clean -modcache
```
Dan memperbaharui kembali dengan perintah ini
```
go mod tidy
```

## Cara menjalankan migration 

### Membuat file migration

```
migrate create -ext sql -dir databases/migrations (nama_file)
```
### SQLITE

Up
```
migrate -path databases/migrations -database "sqlite3://database.sqlite" up
```

Down
```
migrate -path databases/migrations -database "sqlite3://database.sqlite" down
```

### Postgree(Recomended)

Up
```
migrate -path databases/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

Down
```
migrate -path databases/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down
```
### MySQL(Recomended)

Up
```
migrate -path databases/migrations -database "mysql://user:pass@tcp(localhost:3306)/dbname" up
```

Down
```
migrate -path databases/migrations -database "mysql://user:pass@tcp(localhost:3306)/dbname" down
```

##### Note:
Jika terjadi error
```
"gcc" not found
```
Jalankan perintah ini di terminal/cmd/powershell
```
set CGO_ENABLED=0
``