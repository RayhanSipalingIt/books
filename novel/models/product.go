package models

type Product struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Judul     string `gorm:"type:varchar(300)" json:"judul"`
	TipeBuku  string `gorm:"type:varchar(255)" json:"tipe_buku"`
	Genre     string `gorm:"type:varchar(255)" json:"genre"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Img       string `gorm:"type:varchar(255)" json:"img"`
}
