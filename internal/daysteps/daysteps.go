package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	is "github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slises := strings.Split(data, ",")

	if len(slises) != 2 {
		return 0, 0, fmt.Errorf("invalid data, expected 2 items, got %d", len(slises))
	}

	steps, err := strconv.Atoi(slises[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data, expected a number, got %s", slises[0])
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid data, expected a positive number, got %d", steps)
	}

	fullTime, err := time.ParseDuration(slises[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data, expected a duration, got %s", slises[1])
	}

	return steps, fullTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, err := is.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}

	result := fmt.Sprintf("Количество шаров: %d.\nДистанция составила %.2f.\nВы сожгли %.2f ккал.",
		steps, distance, calories)

	return result
}
