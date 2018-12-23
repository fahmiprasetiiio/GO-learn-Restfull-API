package handler

import "github.com/jinzhu/gorm"

// TableFahmi -
type TableFahmi struct {
	gorm.Model
	Nama   string `gorm:varchar(60)`
	Umur   int
	Alamat string
}

// CreteTable -
func (t TableFahmi) CreateTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&TableFahmi{}).Error

	return
}

// InsertData -
func (t TableFahmi) InsertData(db *gorm.DB) (err error) {
	err = db.Create(&t).Error

	return
}

// UpdateData -
func (t TableFahmi) UpdateData(db *gorm.DB) (err error) {
	dataLama := new(TableFahmi)

	db.Where("id = ?", t.ID).First(&dataLama)

	dataLama.Nama = t.Nama
	dataLama.Alamat = t.Alamat
	dataLama.Umur = t.Umur

	err = db.Save(&dataLama).Error

	return
}

// DeleteData -
func (t TableFahmi) DeleteData(db *gorm.DB) (err error) {
	err = db.Delete(&t).Error

	return
}

// GetDataByName -
func (t TableFahmi) GetDataByName(db *gorm.DB) (res TableFahmi, err error) {

	err = db.Where("nama = ?", t.Nama).First(&res).Error //first akan dapat satu saja

	return
}

// GetAllData -
func (t TableFahmi) GetAllData(db *gorm.DB) (res []TableFahmi, err error) {

	err = db.Find(&res).Error //Find akan dapat semua

	return
}
