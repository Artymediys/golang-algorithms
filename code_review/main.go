package main

import (
	"fmt"
)

func main() {
	// n - джуны, m - сеньоры, k - кол-во проверок одной работы
	var n, m, k int
	_, err := fmt.Scan(&n, &m, &k)
	if err != nil {
		return
	}

	// Рассчитываем время, необходимое для проверки всех программ
	minutesRequired := n * k

	// Рассчитываем время, которое потратят сеньор-разработчики на проверку
	timeSpent := minutesRequired / m

	// Если время, потраченное на проверку, не делится ровно на количество сеньоров, добавляем еще одну минуту
	if minutesRequired%m != 0 {
		timeSpent++
	}

	fmt.Println(timeSpent)
}
