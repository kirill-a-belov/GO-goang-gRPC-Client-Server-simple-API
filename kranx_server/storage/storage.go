/*
Модуль для первого приложения
Представляет из себя имитацию базы данных
*/

package storage

import (
	"fmt"
	"math/rand"
)

const MEMORY_CAPACITY  =  10000 // Максимальный размер "базы данных" - 10000
var memory= make(map[string]string, MEMORY_CAPACITY)
var memoryIndex [MEMORY_CAPACITY]string// Дополнительная стуктура для реализации удаления по времени создания (самых старых)
var memoryIndexCounter = 0

func AddToMemory(key string, value string) string{
	status := "Added"
	if (memoryIndexCounter >= MEMORY_CAPACITY) {
		rnd := rand.Intn(MEMORY_CAPACITY)

		status = "Max mem was reached, key "+ memoryIndex[rnd] + " deleted!"
		DelFromMemory(memoryIndex[rnd])
		memory[key] = value
		memoryIndex[rnd] = key
		return status
		}
	memory[key] = value
	memoryIndex[memoryIndexCounter] = key
	memoryIndexCounter++
	return status
	}

func DelFromMemory(key string)string{

	delete(memory, key)
	return "Ok"

}

func GetFromMemory(key string) string{
	if (memory[key] == "") {
		return ""
	}
	return memory[key]
}

func PrintMemory(){
	fmt.Println()
	fmt.Println("Printing memory:")
	for key, value := range memory {
		fmt.Println("Key:", key, "Value:", value)
	}
	fmt.Println("Printing keylist:")
	for i:=0;i<MEMORY_CAPACITY ;i++  {
		fmt.Println(memoryIndex[i])

	}

}

