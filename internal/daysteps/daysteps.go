package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("Неверная длина")
	}
	numbeStr, durationStr := parts[0], parts[1]
	number, err := strconv.Atoi(numbeStr)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать количество шагов '%s' в число: %w", parts[0], err)
	}
	if number <= 0 {
		return fmt.Errorf("Количество шагов должно быть положительным")
	}
	ds.Steps = number

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать длительность '%s' в time.Duration: %w", parts[1], err)
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть положительной")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", fmt.Errorf("Некорректные данные")
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return info, nil
}
