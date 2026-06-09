package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slices := strings.Split(data, ",")

	if len(slices) != 3 {
		return 0, "", 0, errors.New("invalid data format: expected 3 items")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(slices[0]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps value: %w", err)
	}

	activity := strings.TrimSpace(slices[1])

	duration, err := time.ParseDuration(strings.TrimSpace(slices[2]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration value: %w", err)
	}
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLengthMeters := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLengthMeters
	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64

	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration) // Исправлено!
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	message := fmt.Sprintf("\nТип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), dist, speed, calories)

	return message, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive duration, got %v", duration)
	}

	result := (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return result, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %f", weight)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive duration, got %d", duration)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid data, expected a positive number, got %v", height)
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	result := (weight * speed * minutes) / 60 * walkingCaloriesCoefficient

	return result, nil
}
