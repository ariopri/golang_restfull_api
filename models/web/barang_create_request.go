package web

type BarangCreateRequest struct {
	Nama  string `validate:"required,min=1,max=200" json:"nama"`
	Harga int    `validate:"required" json:"harga"`
	Stok  int    `validate:"required" json:"stok"`
}
