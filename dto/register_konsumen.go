package dto

type RegisterKonsumen struct {
	Full_name    string `json:"full_name" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Nik          string `json:"nik" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	TempatLahir  string `json:"tempat_lahir" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Gaji         int64  `json:"gaji" binding:"required"`
}
