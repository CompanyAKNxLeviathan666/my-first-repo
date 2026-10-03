package main
import "fmt"
func main(){
	//1. Ввод данных от пользователя
	var num intfmt.Print("Введите целое число: ")
	//2. Логика обработка данных
	if num > 10 {
		num = num * 2
	} else {
		num = num / 2
	}
	//3. Вывод результатов
	fmt.Println("Результат:", num)
	fmt.Println("Исходное число:", num)
}