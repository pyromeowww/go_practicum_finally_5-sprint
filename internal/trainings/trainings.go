package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps                 int           // количество шагов
	TrainingType          string        // тип тренировки (бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенная структура с метотдом Print()
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("Неверная длина")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать количество шагов '%s' в число: %w", parts[0], err)
	}
	if steps <= 0 {
		return fmt.Errorf("Количество шагов должно быть положительным")
	}
	t.Steps = steps

	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать длительность '%s' в time.Duration: %w", parts[2], err)
	}
	if duration <= 0 {
		return fmt.Errorf("Продолжительность должна быть положительной")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.Steps <= 0 || t.Duration <= 0 {
		return "", fmt.Errorf("Некорректные данные")
	}
	var calories float64
	var err error
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories)
	return info, nil
}
