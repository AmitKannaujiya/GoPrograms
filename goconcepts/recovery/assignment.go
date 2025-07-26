package recovery

import "log"

func recoveryUse() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovering from function")
		}
	}()

	panic("throwing error from panic")
}

func recoveryWithFatal() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Recovering from function: with fatal")
		}
	}()

	panic("throwing error from panic")
}