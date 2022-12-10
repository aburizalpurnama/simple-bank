# SQLC
### How to Generate using Windows (Menggunakan docker)
Cara generate (kalau di windows harus menggunakan cmd, ga boleh powershell / bash)

Terlebih dahulu masuk ke direktori project.

    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

Bisa digunakan untuk generate, init, dll. tinggal ganti perintahnya.

# Transactional
## Avoiding Deadlock
- Jika terdapat opeperasi dua arah secara concurent (contohnya rekening A melakukan transfer ke rekening B, dan sebaliknya), Selalu menjalankan tahapan query dengan urutan Id yang konsisten.
Menerapkan data id oriented, bukan operational oriented.
source:
https://dev.to/techschoolguru/how-to-avoid-deadlock-in-db-transaction-queries-order-matter-oh7

# Framework & Libraries
## Gin
Framework untuk membuat RESTful HTTP API

Instalation :

    go get -u github.com/gin-gonic/gin

## Viper
Library untuk load configurasi dari file dan environment variable

Instalation :

    go get github.com/spf13/viper

## Mockgen
Library untuk membuat mocking sebuah objek saat melakukan unit testing

Instalation :

    go install github.com/golang/mock/mockgen$Version

Instruction :

- Create Mock : mockgen -package 'package-name' --build_flags=--mod=mod -destination 'dest package/file.go' 'module-name/package' 'interface-name'

    example:

        mockgen -package mockdb --build_flags=--mod=mod -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

Note :

- Saat dicek dengan 'which mockgen' tidak muncul, kemungkinan file mockgen.nya ada di inner bin folder ~/go/bin/bin.
cukup pundahkan ke ~/go/bin, dan lakukan pengecekan kembali dengan 'which mockgen'

# Tools
## Docker

Open bash docker container :

    docker exec -it <mycontainer> bash