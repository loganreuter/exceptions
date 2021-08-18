package err

import "log"

func Handle(try func(), catch func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("\033[31m", "Error: ", r, "\033[0m")
			catch()
		}
	}()

	try()
}
