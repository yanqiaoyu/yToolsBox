package model

type Demo_table struct {
	Tel2             string `form:"tel2" json:"tel2" gorm:"column:tel2"`
	Tel              string `form:"tel" json:"tel" gorm:"column:tel"`
	Idcard           string `form:"idcard" json:"idcard" gorm:"column:idcard"`
	Id_com           string `form:"id_com" json:"id_com" gorm:"column:id_com"`
	Passport         string `form:"passport" json:"passport" gorm:"column:passport"`
	Hkmacao          string `form:"hkmacao" json:"hkmacao" gorm:"column:hkmacao"`
	Bank_id          string `form:"bank_id" json:"bank_id" gorm:"column:bank_id"`
	Email            string `form:"email" json:"email" gorm:"column:email"`
	Address          string `form:"address" json:"address" gorm:"column:address"`
	Name             string `form:"name" json:"name" gorm:"column:name"`
	Ip               string `form:"ip" json:"ip" gorm:"column:ip"`
	Mac              string `form:"mac" json:"mac" gorm:"column:mac"`
	Car              string `form:"car" json:"car" gorm:"column:car"`
	Assets           string `form:"assets" json:"assets" gorm:"column:assets"`
	Sex              string `form:"sex" json:"sex" gorm:"column:sex"`
	Nationality      string `form:"nationality" json:"nationality" gorm:"column:nationality"`
	Province         string `form:"province" json:"province" gorm:"column:province"`
	Religious_belief string `form:"religious_belief" json:"religious_belief" gorm:"column:religious_belief"`
	Soldier          string `form:"soldier" json:"soldier" gorm:"column:soldier"`
}
