package main

import "rashirenie/register-service/internal/app"

func main() {
	app := app.New() //запуск  (арр - это название файла,New - это функция в файле)
	app.Run()

}
