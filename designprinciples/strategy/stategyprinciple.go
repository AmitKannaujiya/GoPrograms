package strategy

import "fmt"

type IDriveStrategy interface {
	driveVehicle(string)
}

type NormalDriveStrategy struct{}

type SportDriveStrategy struct{}

type OffRoadDriveStrategy struct{}

func (nds NormalDriveStrategy) driveVehicle(ty string) {
	fmt.Printf("Driving Normally : %s \n", ty)
}

func (sds SportDriveStrategy) driveVehicle(ty string) {
	fmt.Printf("Driving Sprots : %s \n", ty)
}

func (ods OffRoadDriveStrategy) driveVehicle(ty string) {
	fmt.Printf("Driving Offroad : %s \n", ty)
}

type Vehicle struct {
	Ty    string
	IDS IDriveStrategy
}

func (v Vehicle) Drive() {
	v.IDS.driveVehicle(v.Ty)
}
func (v Vehicle) GetType() string {
	return v.Ty
}

type SportVehicle struct {
	Vehicle Vehicle
}

type NormalVehicle struct {
	Vehicle Vehicle
}

type OffroadVehicle struct {
	Vehicle Vehicle
}

type GoodsVehicle struct {
	Vehicle Vehicle
}