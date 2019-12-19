package model

type TblUnit struct {
	Id       int    `json:"id" gorm:"column:id"`
	NamaUnit string `json:"nama_unit" gorm:"column:nama_unit"`
	KodeUnit string `json:"kode_unit" gorm:"column:kode_unit"`
}
