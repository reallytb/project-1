package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	separateData := strings.Split(data, ",")
	if len(separateData) != 2 {
		log.Println("Ошибка парсинга")
		return 0, 0, fmt.Errorf("ошибка парсинга")
	}
	steps, err := strconv.Atoi(separateData[0])
	if err != nil {
		log.Println("Ошибка приведения количества шагов в целое число")
		return 0, 0, fmt.Errorf("ошибка приведения количества шагов в целое число")
	}
	if steps == 0 {
		log.Println("Количество шагов равно 0")
		return 0, 0, fmt.Errorf("количество шагов равно 0")
	}
	if steps < 0 {
		log.Println("Количество шагов отрицательно")
		return 0, 0, fmt.Errorf("количество шагов отрицательно")
	}
	time, err := time.ParseDuration(separateData[1])
	if err != nil {
		log.Println("Ошибка приведения времени")
		return 0, 0, fmt.Errorf("ошибка приведения времени")
	}
	if time < 0 {
		log.Println("Время отрицательно")
		return 0, 0, fmt.Errorf("время отрицательно")
	}
	if time == 0 {
		log.Println("Время нулевое")
		return 0, 0, fmt.Errorf("время нулевое")
	}
	return steps, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, time, err := parsePackage(data)
	if err != nil {
		fmt.Println("ошибка получения данных")
		return ""
	}
	if steps == 0 {
		fmt.Println("количество шагов равно 0")
		return ""
	}
	if steps < 0 {
		fmt.Println("количество шагов отрицательно")
		return ""
	}
	distance := float64(steps) * stepLength
	distanceInKm := distance / mInKm
	caloriesSpent, err := spentcalories.WalkingSpentCalories(steps, weight, height, time)
	if err != nil {
		fmt.Println("ошибка рассчёта затраченных калорий")
		return ""
	}
	message := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceInKm, caloriesSpent)
	return message
}
