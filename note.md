## RUN ON TERMINAL

Run test all on package
```bash
go test
```

Use option `-v` for see detail func test
```bash 
go test -v
```

Run test spesific func
```bash
go test -run <funcTestName>
```

Run test all package (from root dir)
```bash
go test ./...
```

## Fail The Test

- sangat tidak disarankan menggunakan `panic()` karena ketika error dan `panic()` dieksekusi, func selanjutnya tidak akan dieksekusi atau pengeksekusian tidak dapat dilanjutkan.
- `t.Fail()` akan meninggalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal.
- `t.FailNow()` akan menggagalkan unit test pada saat itu juga (langsung), tanpa melanjutkan eksekusi unit test.
- `t.Error(args..)` lebih seperti melakukan log (print) error, namun ketika telah melakukan log error, akan secara otomatis memanggil `Fail()`. Sehingga menganggap unit test dianggao gagal. Namun karena hanya memanggil `Fail()`, artinya eksekusi unit test akan tetap berjalan sampai selesai.
- `t.Fatal(args..)` mirip dengan `Error()` hanya saja ketika telah melakukan log error, akan memanggil `FailNow()`. Sehingga mengakibatkan eksekusi unit test terhenti.

## Test with Assertion (mod with `Testify`)

Get `Testify`
```bash
go get github.com/stretchr/testify
```

### Assert vs Require

- Testify menyediakan dua package untuk assertion, yaitu `assert` dan `require`.
- Saat menggunakan `assert`, jika pengecekan gagal maka `assert` akan memanggil `Fail()` artinya eksekusi unit test akan tetap dilakukan.
- Namun jika menggunakan `require`, jika pengecekan gagal maka `require` akan memanggil `FailNow()` artinya eksekusi unit test tidak akan dilanjutkan.

### Skip Test
- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi `unit test`.
- Di Golang juga dapat membatalkan eksekusi `unit test` jika ingin.
- Untuk membatalkan `unit test` bisa menggunakan function `Skip()`