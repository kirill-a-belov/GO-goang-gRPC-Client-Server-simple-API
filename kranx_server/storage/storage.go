/*
Модуль для первого приложения
Представляет из себя имитацию базы данных
*/

package storage

import (
	"fmt"
	"container/list"
)

const MEMORY_CAPACITY  =  10000 // Максимальный размер "базы данных" - 10000
var memory= make(map[string]string, MEMORY_CAPACITY)
var memoryIndex = list.New() // Дополнительная стуктура для реализации удаления по времени создания (самых старых)

func AddToMemory(key string, value string) string{
	status := "Added"
	if (len(memory) >= MEMORY_CAPACITY) {
		status = "Max mem was reached, key "+ memoryIndex.Front().Value.(string) + " deleted!"
		DelFromMemory(memoryIndex.Front().Value.(string))
	}
	memory[key] = value;
	memoryIndex.PushBack(key)
	return status
	}

func DelFromMemory(key string)string{

	//TODO Узкое место, нужно поменять структуру данных и/или придумать как скастить ключ к элементу списка без перебора
	for e := memoryIndex.Front(); e != nil; e = e.Next() {
		if (e.Value.(string) == key){
			memoryIndex.Remove(e)
			delete(memory, key)
			return "Ok"
		}
	}
	return "Nothing to delete"
}

func GetFromMemory(key string) string{
	if (memory[key] == "") {
		return "!!!Empty value!!!"
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
	for e := memoryIndex.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

