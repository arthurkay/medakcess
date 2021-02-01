package models

import (
	"gorm.io/gorm"
)

// 0 Blood
// 1 Urine
// 2 Stool
// 3 Other
type Specimen struct {
	gorm.Model
	Name string `gorm:"type:int;enum(0,1,2,3)"`
}

type Blood struct {
	gorm.Model
	FBCID uint
	FBC   FBC
	HbsAg bool `gorm:"type:bool;enum(true,false)"`
	RDT   bool `gorm:"type:bool;enum(true,false)"`
	MPS   bool `gorm:"type:bool;enum(true,false)"`
	RPR   bool `gorm:"type:bool;enum(true,false)"`
	CAT   bool `gorm:"type:bool;enum(true,false)"`
	TAT   bool `gorm:"type:bool;enum(true,false)"`
}

type FBC struct {
	gorm.Model
	HB  float64
	WBC float64
	RBC float64
	PLT int
	MCV float64
	HCT float64
}

type Urine struct {
	gorm.Model
	UrinalysisID uint
	Urinalysis   Urinalysis
	UrineMCSID   uint
	UrineMCS     UrineMCS
	GravityIndex bool `gorm:"type:bool;enum(true,false)"`
}

type Urinalysis struct {
	gorm.Model
	ph              float64
	blood           string `gorm:"type:int;enum(0,1,2)"`
	SpecificGravity float64
	Leucocyte       bool `gorm:"type:bool;enum(true,false)"`
	Protein         bool `gorm:"type:bool;enum(true,false)"`
	Glucose         float64
	Ketones         int `gorm:"type:int;enum(1,2,3,4)"`
	Nitrates        int `gorm:"type:int;enum(0,1,2)"`
	Bilirubin       int `gorm:"type:int;enum(0,1,2,3)"`
	Urobiligin      int `gorm:"type:int;enum(0,1,2)"`
	Color           string
}

type UrineMCS struct {
	gorm.Model
	PusCells      string `gorm:"type:string;enum('small', 'large', 'many')"`
	Blood         bool   `gorm:"type:bool;enum(true,false)"`
	MicroOrganism string
}

type Stool struct {
	gorm.Model
	StoolMCSID uint
	StoolMCS   StoolMCS
}

type StoolMCS struct {
	gorm.Model
	PusCells      string `gorm:"type:string;enum('small', 'large', 'many')"`
	Blood         bool   `gorm:"type:bool;enum(true,false)"`
	MicroOrganism string
	Color         string
}

type Other struct {
	gorm.Model
	Name  string
	Value string
}
