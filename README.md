## SQLC
### How to Generate using Windows (Menggunakan docker)
Cara generate (harus menggunakan cmd, ga boleh powershell / bash)

    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

## Avoiding Deadlock
- Selalu menjalankan tahapan query dengan urutan Id yang konsisten.
Menerapkan data id oriented, bukan operational oriented. 
https://dev.to/techschoolguru/how-to-avoid-deadlock-in-db-transaction-queries-order-matter-oh7