# SQLC
### How to Generate using Windows (Menggunakan docker)
Cara generate (harus menggunakan cmd, ga boleh powershell / bash)

    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

# Transactional
## Avoiding Deadlock
- Jika terdapat opeperasi dua arah secara concurent (contohnya rekening A melakukan transfer ke rekening B, dan sebaliknya), Selalu menjalankan tahapan query dengan urutan Id yang konsisten.
Menerapkan data id oriented, bukan operational oriented. 
https://dev.to/techschoolguru/how-to-avoid-deadlock-in-db-transaction-queries-order-matter-oh7

# Framework
## Gin

Installing gin :

    go get -u github.com/gin-gonic/gin

# Tools
## Docker

Open bash docker container :

    docker exec -it <mycontainer> bash