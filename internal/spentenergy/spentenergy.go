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

// Функции для расчёта потраченной энергии при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// проверяем входные параметры на корректность
	if steps <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the number of steps must be greater than 0") }
	if weight <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the weight must be greater than 0") }
	if height <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the height must be greater than 0") }
	if duration <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the duration must be greater than 0") }
	// рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)
	// переводим продолжительность в минуты 
	durationInMinutes := duration.Minutes()
	// рассчитываем количество калорий
	amountСalories := (weight * meanSpeed * durationInMinutes) / minInH
	spentCalories := amountСalories * walkingCaloriesCoefficient
	return spentCalories, nil
}

// Функции для расчёта потраченной энергии при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверяем входные параметры на корректность
	if steps <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the number of steps must be greater than 0") }
	if weight <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the weight must be greater than 0") }
	if height <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the height must be greater than 0") }
	if duration <= 0 { return 0, errors.New("spentenergy: RunningSpentCalories: the duration must be greater than 0") }
	// рассчитаем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)
	// переводим продолжительность в минуты 
	durationInMinutes := duration.Minutes()
	// рассчитываем количество калорий
	amountСalories := (weight * meanSpeed * durationInMinutes) / minInH
	return amountСalories, nil
}

// Функции для расчёта средней скорости
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	// проверяем что продолжительность больше нуля
	if duration <= 0 { return 0 }
	// получаем часы из продолжительности
	hours := duration.Hours()
	// получаем дистанцию
	distance := Distance(steps, height)
	// вычисляем среднюю скорость
	averageSpeed := distance / hours
	return averageSpeed
}

// Функции для расчёта пройденной дистанции
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
		// рассчитаем длину шага
	stepLength := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return stepLength
}
