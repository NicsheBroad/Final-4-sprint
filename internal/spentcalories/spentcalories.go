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
	lenStep                    = 0.65 // средняя длина шага.
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
	return ((float64(steps) * (height * stepLengthCoefficient)) * lenStep) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, str, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch str {
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		message := fmt.Sprintf("\nТип тренировки: %s.\nДлительность: %v.\nДистанция: %f.\nСкорость: %f\nСожгли калорий: %f", str, duration, dist, speed, calories)
		if err != nil {
			return "", err
		}
		return message, nil
	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		message := fmt.Sprintf("\nТип тренировки: %s.\nДлительность: %v.\nДистанция: %f.\nСкорость: %f\nСожгли калорий: %f", str, duration, dist, speed, calories)
		if err != nil {
			return "", err
		}
		return message, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
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
		return 0, fmt.Errorf("invalid data, expected a positive duration, got %d", duration)
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
		return 0, fmt.Errorf("invalid data, expected a positive number, got %f", height)
	}

	result := (duration.Minutes() * meanSpeed(steps, height, duration) * weight) / 60 / walkingCaloriesCoefficient

	return result, nil
}
