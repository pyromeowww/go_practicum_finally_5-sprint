package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("The value must not be zero.")
	}

	midSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * midSpeed * minutes) / minInH
	correctСalories := calories * walkingCaloriesCoefficient
	return correctСalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("The value must not be zero.")
	}
	midSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * midSpeed * minutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	if steps < 0 || height < 0 {
		return 0
	}

	distance := Distance(steps, height)
	hours := duration.Hours()
	midSpeed := distance / hours
	return midSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps < 0 || height < 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient // длина шага
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}
