package main

import (
	"fmt"
)

type User struct {
	Id   int64
	Name string
}

func main() {

	// Ключ должен быть comparable(сравниваемый с другими типами)
	//  Incomparable types: 1) slices 2) maps 3) functions 4) struct with incomparable types

	// var defaultMap map[string]string // если создаем с дефолтным значением - ничего с ней не сможем сделать

	// fmt.Printf("\n%T, %#v\n", defaultMap, defaultMap)
	// fmt.Printf("len: %d\n\n", len(defaultMap))

	// slice by make
	mapByMake := make(map[string]string)
	fmt.Printf("%T, %#v\n", mapByMake, mapByMake)
	fmt.Printf("len: %d\n\n", len(mapByMake))

	// // slice by make with cap
	// mapByMakeWithCap := make(map[string]string, 3) // 3 - кол-во элементов, которые мы хотим разместить в мапе
	// fmt.Printf("%T, %#v\n", mapByMakeWithCap, mapByMakeWithCap)
	// fmt.Printf("len: %d\n\n", len(mapByMakeWithCap))

	//slice by literal
	mapByLiteral := map[string]int{"Daria": 30, "Egor": 35}
	fmt.Printf("%T, %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("len: %d\n\n", len(mapByLiteral))

	// // slice by new
	// mapWithNew := *new(map[string]string)
	// fmt.Printf("%T, %#v\n", mapWithNew, mapWithNew)
	// fmt.Printf("len: %d\n\n", len(mapWithNew))

	// insert value; вставка и обновление значений
	mapByMake["First"] = "Vova"
	fmt.Printf("%T, %#v\n", mapByMake, mapByMake)
	fmt.Printf("len: %d\n\n", len(mapByMake))

	// update value
	mapByMake["First"] = "new"
	fmt.Printf("%T, %#v\n", mapByMake, mapByMake)
	fmt.Printf("len: %d\n\n", len(mapByMake))

	// get map value
	fmt.Println(mapByLiteral["Egor"])

	// get map default value
	fmt.Println(mapByLiteral["No value"])

	// check value existence
	value, ok := mapByLiteral["No value"]
	fmt.Println()
	fmt.Printf("Value: %d IsExist: %t\n\n", value, ok)

	// delete value
	delete(mapByMake, "First")
	fmt.Printf("%T, %#v\n\n", mapByMake, mapByMake)

	// map iteration
	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}
	fmt.Printf("%T, %#v\n", mapForIteration, mapForIteration)

	for key, value := range mapForIteration {
		fmt.Printf("KEY: %s; VALUE: %d\n\n", key, value)
	}

	// unique values
	// использование map как set (фильтр уникальности)
	// User - структура

	users := []User{
		{
			Id:   1,
			Name: "Egor",
		},
		{
			Id:   34,
			Name: "Daria",
		},
		{
			Id:   13,
			Name: "Lucky",
		},
		{
			Id:   34,
			Name: "Sasha",
		},
	}

	uniqueUsers := map[int64]struct {
		Id    int64
		Name2 string
	}{
		1: {Id: 1, Name2: "Vova"},
		2: {Id: 2, Name2: "Sasha"},
		3: {Id: 2, Name2: "Sasha"},
	}

	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct {
				Id    int64
				Name2 string
			}{
				Id:    user.Id,
				Name2: user.Name,
			}
		}
	}

	fmt.Printf("Type: %T ; Value: %v\n", uniqueUsers, uniqueUsers)

	// быстрый поиск значения с помощью мапы
	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}
	fmt.Println(usersMap)
	fmt.Println(findInSLice(34, users))

	ss := findInMap(34, usersMap)
	fmt.Println(findInMap(34, usersMap))
	fmt.Println(*ss)

}

func findInSLice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}

	}
	return nil
}

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}
