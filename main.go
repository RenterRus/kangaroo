package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main(){
	//Анонимная функция для проверки валидности данных
	errorParser := func(name string, err error){
		if err != nil {
			fmt.Printf("%s entered incorrectly: %s", name, err.Error())
			os.Exit(1)
		}
	}

	//Получение данных и проверка на валидность
	kangaroo1, err := strconv.Atoi(os.Args[1])
	errorParser("kangaroo1", err)

	speed1, err := strconv.Atoi(os.Args[2])
	errorParser("speed1", err)

	kangaroo2, err := strconv.Atoi(os.Args[3])
	errorParser("kangaroo2", err)

	speed2, err := strconv.Atoi(os.Args[4])
	errorParser("speed2", err)

	//Запуск анализатора
	fmt.Println(calculating(kangaroo1, speed1, kangaroo2, speed2))
}

func calculating(x1, v1, x2, v2 int) string{
	//Дистанция между кеггуру
	distanse := math.Abs(float64(x1-x2))
	//Скорость изменения дистанции между ними
	rapprochement := math.Abs(float64(v2)) - math.Abs(float64(v1))
	//Если скорость положительная, т.е. происходит увеличение дистанции то в одно время кенгуру не будут на одном месте,
	//если ведущий не остановится, что условием задачи не предусмотрено. Со скоростью 0 дистанция не изменяется
	if rapprochement >= 0{
		return "NO"
	}
	//Тут уже происходит сближение
	rapprochement = math.Abs(rapprochement)
	//Если остаток от деления дистанции на скорость сближения будет равен 0, то кенгуру будут в одной точке в одно время
	if int(distanse) % int(rapprochement) == 0{
		return "YES"
	}
	return "NO"
}