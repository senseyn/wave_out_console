package main

import (
	"math/rand"
	"fmt"
	"time"
	"os"
	"strings"
)

func main() {
	fmt.Println("Начало работы...\nЧтобы завершить работу программы 'CTRL+C'")
	time.Sleep(5 * time.Second)

	file, _ := os.Create("num.txt")
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	for{
		rundom := rand.Intn(101)
		file.WriteString(fmt.Sprintf("%d\n", rundom))
		count := 40
		tx := "Записываю."
		for i:=0; i<count; i++ {
			count -= 1
			tx += "."
			fmt.Println(tx)
			time.Sleep(50 * time.Millisecond)
		}

		for i:=60; i>count; i-- {
			count += 1
			tx = strings.Replace(tx, ".", "", 1)
			fmt.Println(tx)
			time.Sleep(50 * time.Millisecond)
		}
	}	
}