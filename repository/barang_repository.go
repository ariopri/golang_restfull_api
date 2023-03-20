package repository

import (
	"context"
	"database/sql"
	"github.com/ariopri/golang_restfull_api/models/domain"
)

type BarangRepository interface {
	Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
	Update(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
	Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang)
	FindById(ctx context.Context, tx *sql.Tx, barangId int) (domain.Barang, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Barang
}

/*
	PENJELASAN
	1. Membuat interface BarangRepository
	BarangRepository merupakan interface yang berisi method-method yang akan digunakan untuk mengakses database.
	Interface ini akan diimplementasikan oleh struct yang akan mengakses database.

	1. Method Save(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
		Parameter ctx bertipe context.Context: Objek context.Context digunakan untuk mengelola konteks
		aplikasi pada saat melakukan operasi pada database. Parameter ini harus di-passing pada setiap
		method di repository untuk memastikan setiap operasi pada database berada dalam konteks aplikasi yang sama.
		Parameter tx bertipe *sql.Tx: Objek *sql.Tx digunakan untuk melakukan transaksi pada database.
		Dengan menggunakan transaksi, kita dapat memastikan operasi-operasi pada database dilakukan secara
		atomik, artinya semua operasi harus berhasil atau gagal secara bersamaan.
		Parameter ini juga harus di-passing pada setiap method di repository untuk memastikan setiap operasi pada
		database berada dalam transaksi yang sama.
		Parameter barang bertipe domain.Barang: Objek domain.Barang merepresentasikan data barang yang akan disimpan,
		diupdate, atau dihapus pada database. Parameter ini harus di-passing
		pada method Save dan Update karena kedua method tersebut membutuhkan objek barang
		untuk melakukan operasi pada database. Sementara itu, method Delete tidak memerlukan objek barang sebagai
		parameter karena data barang yang akan dihapus dapat diidentifikasi melalui ID barang.
		Return value: Objek domain.Barang yang merepresentasikan data barang yang berhasil disimpan atau diupdate
		di dalam database.

	2. Method Update(ctx context.Context, tx *sql.Tx, barang domain.Barang) domain.Barang
		Parameter ctx dan tx memiliki penjelasan yang sama dengan pada method Save.
		Parameter barang merepresentasikan data barang yang akan diupdate pada database.
		Return value: Objek domain.Barang yang merepresentasikan data barang yang berhasil diupdate
		di dalam database.

	3. Method Delete(ctx context.Context, tx *sql.Tx, barang domain.Barang)
		Parameter ctx dan tx memiliki penjelasan yang sama dengan pada method Save.
		Parameter barang merepresentasikan data barang yang akan dihapus dari database.
		Return value: tidak ada return value pada method ini.

	4. Method FindById(ctx context.Context, tx *sql.Tx, barangId int) (domain.Barang, error)
		Parameter ctx dan tx memiliki penjelasan yang sama dengan pada method Save.
		Parameter barangId bertipe int: integer ini merepresentasikan ID barang yang akan dicari pada database.
		Return value: Objek domain.Barang yang merepresentasikan data barang yang berhasil ditemukan di dalam
		database dan error jika terjadi kesalahan dalam pencarian data.

	5. Method FindAll(ctx context.Context, tx *sql.Tx) []domain.Barang
		Parameter ctx dan tx memiliki penjelasan yang sama dengan pada method Save.
		Return value: slice dari objek domain.Barang yang merepresentasikan data barang yang berhasil ditemukan
		di dalam database.

*/
