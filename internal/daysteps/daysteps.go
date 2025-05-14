package daysteps

import (
	"fmt"
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
		return 0, 0, fmt.Errorf("Ошибка парсинга")
	}
	steps, err := strconv.Atoi(separateData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка приведения количества шагов в целое число")
	}
	if steps == 0 {
		return 0, 0, fmt.Errorf("Количество шагов равно 0")
	}
	if steps < 0 {
		return 0, 0, fmt.Errorf("Количество шагов отрицательно")
	}
	time, err := time.ParseDuration(separateData[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка приведения времени")
	}
	if time < 0 {
		return 0, 0, fmt.Errorf("Время отрицательно")
	}
	return steps, time, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, time, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка получения данных")
		return ""
	}
	if steps == 0 {
		fmt.Println("Количество шагов равно 0")
		return ""
	}
	if steps < 0 {
		fmt.Println("Количество шагов отрицательно")
		return ""
	}
	distance := float64(steps) * stepLength
	distanceInKm := distance / mInKm
	caloriesSpent, err := spentcalories.WalkingSpentCalories(steps, weight, height, time)
	if err != nil {
		fmt.Println("Ошибка рассчёта затраченных калорий")
		return ""
	}
	message := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceInKm, caloriesSpent)
	return message
}
