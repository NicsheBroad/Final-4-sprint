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
	slices := strings.Split(data, ",")

	if len(slices) != 2 {
		return 0, 0, fmt.Errorf("invalid data, expected 2 items, got %d", len(slices))
	}

	steps, err := strconv.Atoi(strings.TrimSpace(slices[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data, expected a number, got %s", slices[0])
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid data, expected a positive number, got %d", steps)
	}

	fullTime, err := time.ParseDuration(slices[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid data, expected a duration, got %s", slices[1])
	}

	return steps, fullTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distance := (float64(steps) * stepLength) / mInKm
	calories, err := is.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distance, calories)

	return result
}
