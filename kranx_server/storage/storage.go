/*
Модуль для первого приложения
Представляет из себя имитацию базы данных
*/

package storage

import (
	"math/rand"
	"sync"
)

const MEMORY_CAPACITY  =  10 // Максимальный размер "базы данных" - 10000
var memory = struct{
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string, MEMORY_CAPACITY)}

//var memory= make(map[string]string, MEMORY_CAPACITY)


var memoryIndex [MEMORY_CAPACITY]string// Дополнительная стуктура для реализации удаления по времени создания (самых старых)
var memoryIndexCounter = 0

func AddToMemory(key string, value string) string{
	status := "Added"
	memory.Lock()
	if (memoryIndexCounter >= MEMORY_CAPACITY) {
		rnd := rand.Intn(MEMORY_CAPACITY)

		status = "Max mem was reached, key "+ memoryIndex[rnd] + " deleted!"
		delete(memory.m, memoryIndex[rnd])
		memory.m[key] = value
		memoryIndex[rnd] = key
		memory.Unlock()
		return status
		}
	memory.m[key] = value
	memoryIndex[memoryIndexCounter] = key
	memoryIndexCounter++
	memory.Unlock()
	return status
	}

func DelFromMemory(key string)string{
	memory.Lock()
	delete(memory.m, key)
	memory.Unlock()
	return "Ok"

}

func GetFromMemory(key string) string{
	memory.RLock()
	s := memory.m[key];
	memory.RUnlock()
	return s
}

/*
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
*/
