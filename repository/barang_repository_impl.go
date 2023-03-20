package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ariopri/golang_restfull_api/helper"
	"github.com/ariopri/golang_restfull_api/models/domain"
)

type BarangRepositoryImpl struct {
}

func (repository *BarangRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang {
	//TODO implement me
	SQL := "insert into barang(nama, harga, stok) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, barang.Nama, barang.Harga, barang.Stok)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	barang.ID = int(id)
	return barang

	/*
		Penjelasan:
		1.  func (repository *BarangRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
			Ini adalah definisi dari metode Save. Ini mengambil tiga parameter:
			konteks, objek transaksi SQL, dan objek domain.Barang yang akan disimpan.
			Metode ini mengembalikan objek domain.Barang yang telah disimpan.

		2. SQL := "insert into barang(nama, harga, stok) values(?, ?, ?)"
			Ini adalah string SQL yang akan digunakan untuk menyimpan objek domain.Barang ke dalam database.
			Ini adalah pernyataan INSERT yang memasukkan nilai untuk kolom nama, harga, dan
			stok pada tabel barang.

		3. result, err := tx.ExecContext(ctx, SQL, barang.Nama, barang.Harga, barang.Stok)
			Ini menjalankan pernyataan SQL pada objek transaksi yang diberikan dan mengembalikan objek hasil
			yang berisi informasi tentang operasi yang dilakukan. Objek hasil dapat digunakan untuk mengambil
			ID baris terakhir yang dimasukkan ke dalam tabel.
			Ini juga menangani kesalahan jika terjadi.

		4. helper.PanicIfError(err)
			ini adalah  kode yang memeriksa apakah terjadi kesalahan atau tidak. Jika terjadi kesalahan,
			maka kode ini akan menghasilkan panic.

		5. id, err := result.LastInsertId()
			helper.PanicIfError(err)
			barang.ID = int(id)
			return barang

			Kode ini berfungsi untuk mengambil nilai id dari record terakhir yang diinsert
			ke dalam tabel barang, dan mengeset nilai ID pada objek barang dengan nilai id tersebut.
			Kemudian, method Save akan mengembalikan objek barang.

			Lebih detailnya, baris kode tersebut melakukan hal-hal berikut:
			1. result.LastInsertId() digunakan untuk mengambil nilai id dari
				record terakhir yang diinsert ke dalam tabel barang.
			2. id, err := result.LastInsertId() digunakan untuk mengambil nilai id dan
				menangani kemungkinan kesalahan dengan cara menetapkan nilai err ke nil jika tidak ada kesalahan,
				atau menetapkan nilai err ke error yang dihasilkan jika terjadi kesalahan saat
				mengambil nilai id.
			3. barang.ID = int(id) digunakan untuk mengeset nilai ID pada objek barang
				dengan nilai id yang telah diambil sebelumnya. Dalam kasus ini, nilai id yang diambil
				dari database adalah tipe data int64, sedangkan ID pada objek barang adalah
				tipe data int, sehingga perlu diubah terlebih dahulu dengan fungsi int().
			4. return barang digunakan untuk mengembalikan objek barang setelah nilai ID diubah sesuai dengan
				nilai id terakhir yang diinsert ke dalam tabel barang
	*/
}

func (repository *BarangRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang {
	//TODO implement me
	SQL := "update barang set nama = ?, harga = ?, stok = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, barang.Nama, barang.Harga, barang.Stok, barang.ID)
	helper.PanicIfError(err)
	return barang
}

func (repository *BarangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang) {
	//TODO implement me
	SQL := "delete from barang where id = ?"
	_, err := tx.ExecContext(ctx, SQL, barang.ID)
	helper.PanicIfError(err)
}

func (repository *BarangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, barangId int) (domain.Barang, error) {
	//TODO implement me
	SQL := "select id, nama, harga, stok from barang where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, barangId)
	helper.PanicIfError(err)
	defer rows.Close()

	barang := domain.Barang{}
	if rows.Next() {
		err := rows.Scan(&barang.ID, &barang.Nama, &barang.Harga, &barang.Stok)
		helper.PanicIfError(err)
		return barang, nil
	} else {
		return barang, errors.New("barang tidak ditemukan")
	}

	/*
		penjelasan
		QueryContext dan ExecContext adalah dua fungsi yang berbeda dalam package database/sql.
		QueryContext digunakan untuk mengambil data dari database sedangkan ExecContext digunakan untuk
		menjalankan perintah SQL yang tidak mengembalikan data seperti insert, update, dan delete.
		Pada kode yang diberikan, Query
		Pada baris kode barang := domain.Barang{}, kita membuat variabel barang bertipe domain.
		Barang dengan nilai awal kosong. Variabel ini akan diisi dengan data barang yang ditemukan
		di database jika query berhasil mengembalikan data.
		Variabel barang dibuat untuk menampung data barang yang ditemukan di database.
		Jika query berhasil mengembalikan data, maka variabel barang akan diisi dengan data barang
		yang ditemukan di database menggunakan fungsi Scan. Jika query tidak mengembalikan data,
		maka variabel barang akan tetap kosong.
	*/
}

func (repository *BarangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Barang {
	//TODO implement me
	SQL := "select id, nama, harga, stok from barang"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var barangs []domain.Barang
	for rows.Next() {
		barang := domain.Barang{}
		err := rows.Scan(&barang.ID, &barang.Nama, &barang.Harga, &barang.Stok)
		helper.PanicIfError(err)
		barangs = append(barangs, barang)
	}
	return barangs

	/*
		penjelasan
		Fungsi FindAll digunakan untuk mencari semua data barang yang ada di database. Fungsi ini menerima dua parameter yaitu ctx context.Context dan tx *sql.Tx. Parameter ctx digunakan untuk mengatur waktu timeout dan pembatalan operasi database. Parameter tx digunakan untuk menjalankan transaksi database.
		Pada baris kode SQL := "select id, nama, harga, stok from barang", kita membuat variabel SQL yang berisi perintah SQL untuk mencari semua data barang yang ada di database. Perintah SQL ini akan mengembalikan data barang berupa id, nama, harga, dan stok.
		Pada baris kode rows, err := tx.QueryContext(ctx, SQL), kita menjalankan perintah SQL yang ada di variabel SQL dengan menggunakan fungsi QueryContext yang dimiliki oleh objek tx. Fungsi ini digunakan untuk mengambil data dari database. Parameter pertama adalah ctx yang digunakan untuk mengatur waktu timeout dan pembatalan operasi database. Parameter kedua adalah SQL yang berisi perintah SQL yang akan dijalankan.
		Pada baris kode helper.PanicIfError(err), kita memeriksa apakah terjadi error saat menjalankan perintah SQL. Jika terjadi error, maka program akan berhenti dan menampilkan pesan error.
		Pada baris kode defer rows.Close(), kita menutup koneksi ke database setelah selesai mengambil data dari database. Ini dilakukan agar koneksi ke database tidak terus terbuka dan membebani server database.
		Pada baris kode var barangs []domain.Barang, kita membuat variabel barangs bertipe []domain.Barang dengan nilai awal kosong. Variabel ini akan diisi dengan data barang yang ditemukan di database.
		Pada baris kode for rows.Next() {, kita memeriksa apakah query berhasil mengembalikan data. Jika query mengembalikan data, maka kita membuat variabel barang bertipe domain.Barang dengan nilai awal kosong. Variabel ini akan diisi dengan data barang yang ditemukan di database menggunakan fungsi Scan. Jika query tidak mengembalikan data, maka kita mengembalikan error “barang tidak ditemukan”.
		Pada baris kode barang := domain.Barang{}, kita membuat variabel barang bertipe domain.Barang dengan nilai awal kosong. Variabel ini akan diisi dengan data barang yang ditemukan di database menggunakan fungsi Scan.
		Pada baris kode err := rows.Scan(&barang.ID, &barang.Nama, &barang.Harga, &barang.Stok), kita mengisi variabel barang dengan data yang ditemukan di database menggunakan fungsi Scan. Fungsi Scan digunakan untuk mengisi nilai-nilai pada variabel barang dengan nilai-nilai yang ditemukan di database.
		Pada baris kode helper.PanicIfError(err), kita memeriksa apakah terjadi error saat mengisi variabel barang dengan data dari database. Jika terjadi error, maka program akan berhenti dan menampilkan pesan error.
		Pada baris kode barangs = append(barangs, barang), kita menambahkan variabel barang ke dalam variabel barangs. Fungsi append digunakan untuk menambahkan nilai ke dalam slice.
		Pada baris kode return barangs, kita mengembalikan variabel barangs yang berisi semua data barang yang ditemukan di database.
	*/
}
