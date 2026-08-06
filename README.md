## Rest Full API Sistem Informasi Siswa Via Golang

Aplikasi ini menggunakan HttpOnly untuk akses token authentication nya, tanpa perlu set cookie di front end. Untuk front end saya menggunakan React [link react](https://github.com/bayualexandria/sis-react).
Dan untuk mobile app nya saya menggunakan Flutter [link flutter](https://github.com/bayualexandria/flutter-sis-getx)

### Fitur 
* Authentication Token HTTPOnly(Login/Logout)
* Authentication melalui social media gmail
* Data Guru/Wali Kelas/Admin (Insert/GetAll/GetByUsername/Update/Delete)
* Data Siswa (Insert/GetAll/GetByUsername/Update/Delete)
* Notifikasi Verifikasi Email dan Forgot Password Pada Akun Siswa dan Guru
* Data kelas (Insert/GetAll/GetById/Update/Delete)
* Data Sampah (Softdelete)


### Perintah untuk menjalankan program golang
```
go run main.go
```

### Build ke production
Windows
```
go build -o app.exe
```
Mac OS/Linux
```
go build -o app
```
##### Note:
Jika ingin mengubah nama pada saat build, ubah pada nama "app"



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

## Cara menjalankan seeder
Untuk menjalankan perintah seeder pada golang
```
go run main.go seed
```
##### Note:
Gunakan perintah di atas untuk satu kali menjalankan seeder, setelah itu jalankan "go run main.go" seperti biasa.


##### Note:
Jika terjadi error
```
"gcc" not found
```
Jalankan perintah ini di terminal/cmd/powershell
```
set CGO_ENABLED=0
``