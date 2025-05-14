package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	separateData := strings.Split(data, ",")
	if len(separateData) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка парсинга")
	}
	steps, err := strconv.Atoi(separateData[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка приведения шагов в целое число")
	}
	time, err := time.ParseDuration(separateData[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка приведения времени")
	}
	return steps, separateData[1], time, nil
}

func distance(steps int, height float64) float64 {
	stepLenght := height * stepLengthCoefficient
	distance := float64(steps) * stepLenght
	distanceInKm := distance / mInKm
	return distanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration == 0 {
		return 0.0
	}
	distance := distance(steps, height)
	avgSpeed := distance / duration.Hours()
	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, time, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("Ошибка парсинга")
	}
	if steps == 0 {
		return "", fmt.Errorf("Количество шагов равно 0")
	}
	if steps < 0 {
		return "", fmt.Errorf("Количество шагов отрицательно")
	}
	if weight == 0 {
		return "", fmt.Errorf("Вес равен 0")
	}
	if height == 0 {
		return "", fmt.Errorf("Рост равен 0")
	}
	switch trainingType {
	case "Бег":
		distance := distance(steps, height)
		avgSpeed := meanSpeed(steps, height, time)
		caloriesSpent, err := RunningSpentCalories(steps, weight, height, time)
		if err != nil {
			return "", fmt.Errorf("Ошибка рассчёта затраченных калорий")
		}
		message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, time.Hours(), distance, avgSpeed, caloriesSpent)
		return message, nil
	case "Ходьба":
		distance := distance(steps, height)
		avgSpeed := meanSpeed(steps, height, time)
		caloriesSpent, err := WalkingSpentCalories(steps, weight, height, time)
		if err != nil {
			return "", fmt.Errorf("Ошибка рассчёта затраченных калорий")
		}
		message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, time.Hours(), distance, avgSpeed, caloriesSpent)
		return message, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0.0, fmt.Errorf("Количество шагов равно 0")
	}
	if weight == 0 {
		return 0.0, fmt.Errorf("Вес равен 0")
	}
	if height == 0 {
		return 0.0, fmt.Errorf("Рост равен 0")
	}
	if duration == 0 {
		return 0.0, fmt.Errorf("Длительность равна 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := float64(duration.Minutes())
	caloriesSpent := (weight * avgSpeed * durationInMinutes) / minInH
	return caloriesSpent, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 {
		return 0.0, fmt.Errorf("Количество шагов равно 0")
	}
	if weight == 0 {
		return 0.0, fmt.Errorf("Вес равен 0")
	}
	if height == 0 {
		return 0.0, fmt.Errorf("Рост равен 0")
	}
	if duration == 0 {
		return 0.0, fmt.Errorf("Длительность равна 0")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := float64(duration.Minutes())
	caloriesSpent := (weight * avgSpeed * durationInMinutes) / minInH
	return caloriesSpent * walkingCaloriesCoefficient, nil
}
