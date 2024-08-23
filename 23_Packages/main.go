package main

import (
	. "fmt"

	"ex124/myPackage"

	_ "ex124/myPackage"
)

// Типы пакетов:

// 1. std - стандартные пакеты
// env:
// set GOROOT=C:\Program Files\Go
// set GOPATH=C:\Users\xzx96\go
// например src/fmt

// 2. external
// гитхаб (go get)
// устанавливается в папку src

// 3. local
// которые создаем сами
// в папке src/ourPackage (C:\Program Files\Go\src или C:\Users\xzx96\go\pkg\mod)

func main() {
	Println(myPackage.Sum(1, 2))

}

// Импорт (alias, _, .)
// 1. alias
// -  mP "ex124/myPackage"
// -  f "fmt"

// 2. _ сайд эффект
// - не собираемся обращаться к ним
// - вызывается его функция init
// - например для подключения базы данных

// 3. Через точку
// -
