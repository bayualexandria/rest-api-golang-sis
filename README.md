## Sistem Informasi Siswa Via Golang



## Cara menjalankan migration 

# Membuat file migration

```
migrate create -ext sql -dir databases/migrations (nama_file)
```
# SQLITE

Up
```
migrate -path databases/migrations -database "sqlite3://database.sqlite" up
```

Down
```
migrate -path databases/migrations -database "sqlite3://database.sqlite" down
```

# Postgree(Recomended)

Up
```
migrate -path database/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

Down
```
migrate -path database/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down
```