package solid

import (
	di "go-program/solid/dependencyinversion"
	"testing"
)

func TestDependencyInversion(t *testing.T) {
	amit := di.NewUser("amit", &di.Mysql{"localhost", 3306})
	vijay := di.NewUser("vijay", &di.Postgres{"localhost", 54321})
	amit.GetUser("get all users")
	vijay.GetUser("get all users")
}