package web

type BarangUpdateRequest struct {
	Id    int    `validate:"required" json:"id"`
	Nama  string `validate:"required,min=1,max=200" json:"nama"`
	Harga int    `validate:"required" json:"harga"`
	Stok  int    `validate:"required" json:"stok"`
}
