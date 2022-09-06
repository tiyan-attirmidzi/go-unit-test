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

### Before and After Test
- Biasanya dalam `unit test`, kadang kita ingin melakukan sesuatu sebelum dan setelah sebuah unit test dieksekusi.
- Jikalau kode yang kita lakukan sebelum dan setelah selalu sama antara `unit test` function, maka membuat manual di `unit test` function nya adalah hal yang membosankan dan terlalu banyak kode duplikat jadinya.
- Untungnya di Golang terdapat fitur yang namnya `testing.M`.
- Fitur ini bernama `Main`, dimana digunakan untuk mengatur eksekusi `unit test`, namun hal ini juga bisa kita gunakan untuk melakukan Before dan After di `unit test`.

## `testing.M`
- Untuk mengatur eksekusi ***unit-test***, cukup dengan membuat sebuah _function_ bernama `TestMain` dengan parameter `*testing.M`.
- Jika terdapat _function_ `TestMain` tersebut, maka secara otomatis Go-Lang akan mengeksekusi _function_ ini tiap kali akan menjalankan ***unit-test*** di sabuah `package`.
- Dengan ini kita bisa mengatur `Before` dan `After` ***unit-test*** sesuai dengan keinginan.

> **Warning**
> _Function_ `TestMain` hanya dieksekusi hanya sekali per Go-Lang `package`, bukan per tiap _function_ ***unit-test***

## `Sub-Test`
- Go-Lang mendukung fitur pembuatan _function_ ***unit-test*** didalam _function_ ***unit-test***.
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki oleh Bahasa Pemrograman yang lainnya.
- Untuk membuat `sub-test`, bisa dengan menggunakan _function_ `Run()` pada `struct` yang dimiliki oleh `*testing.T`.

### Menjalankan Hanya `Sub-Test`
Untuk menjalankan sebuah `unit-test` _function_, dapat menjalankan dengan perintah:
```bash
go test -run <FunctionNameTest>
```
Untuk menjalankan hanya satu `sub-test`, dapat menjalankan dengan perintah:
```bash
go test -run <FunctionNameTest/SubTestName>
```
Atau untuk semua `test`, semua `sub-test` di semua _function_, dapat menjalankan dengan perintah:
```bash
go test -run /<SubTestName>
```

## `Table-Test`
- `Table-Test` yaitu dimana kita menyediakan data berupa `slice` yang berisi parameter dan ekspektasi hasil dari `unit-test`.
- Lalu `slice` tersebut kita iteraskan menggunakan `sub-test`.

## `MOCK`
- `Mock` adalah _object_ yang sudah diprogram dengan ekspektasi tertentu sehingga ketika dipanggil, akan menghasilkan data yang sudah diprogram diawal.
- `Mock` adalah salah satu teknik dalam `unit-test`, dimana kita bisa membuat `mock-object` dari suatu _object_ yang memang sulit untuk di-`testing`.
- Misal apabila ingin membuat `unit-test`, namun ternyata ada kode program yang harus memanggil `API Call` ke `third-party` _service_. Hal ini sangat sulit untuk di `test`, karena `unit-test`harus selalu memanggil `third-party` _service_, dan belum tentu juga _response_-nya sesuai dengan apa yang diinginkan.
- Pada kasus poin di atas (3), cocok sekali untuk menggunakan `mock-object`.

### Testify `MOCK`
- Untuk membuat `mock-object`, tidak ada fitur bawaan Go-Lang, namun dapat memanfaatkan _library_ dari `testify` yang sebelumnya digunakan untuk `assertion`.
- Testify mendukung pembuatan `mock-object`, sehingga cocok apabila ingin membuat `mock-object`.
- Namun, perlu diperhatikan, jika desain kode program jelek, akan sulit untuk melakukan `mocking`, jadi pastikan dalam pembuatan kode program memiliki desain yang baik.

#### Contoh Kasus: Aplikasi Query ke _Database_
- Untuk contoh kasus yang digunakan adalah aplikasi Go-Lang yang melakukan _query_ ke _database_.
- Dengan mambuat `Layer Service` sebagai _Business Logic_ dan `Layer Repository` sebagai jembatan ke _Database_.
- Agar kode mudah untuk di-`testing`, didasarkan agar membuat kontrak berupa _interface_