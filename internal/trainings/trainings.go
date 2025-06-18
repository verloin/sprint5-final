package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// пакет отвечает за парсинг строк с данными и за формирование строки с информацией о тренировках
// структура для хранения данных о количестве шагов, длительности тренировки, типа тренировки
type Training struct {
	// TODO: добавить поля
	Steps int  // количество шагов, проделанных за тренировку
	TrainingType string  // тип тренировки(бег или ходьба)
	Duration time.Duration  // длительность тренировки
	personaldata.Personal  // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}


// метод для парсинга строки с данными
func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// Разделить строку datastring на слайс строк
	parts := strings.Split(datastring, ",")
	// Проверим, что у нас три части
	if len(parts) != 3 {
		return errors.New("Training: Parse: the length is less or more than three") 
	}
	// Преобразуем первый элемент в int (количество шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("Training: Parse: failed to convert steps: %w", err)
	}
	// Проверяем, что количество шагов больше 0
	if steps <= 0 {
		return errors.New("Training: Parse: incorrect number of steps")
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	// Преобразуем третий элемент в time.Duration и обрабатываем ошибки
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("Training: Parse: failed to parse duration: %w", err)
	}
	if duration <= 0 {
		return  errors.New("Training: Parse: negative durations are not allowed")
	}
	// сохраняем полученное значение в соответствующем поле структуры Training.
	t.Duration = duration
	return nil
}

// метод для формирования строки с информацией о проведённой тренировке
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// вычисляем дистанцию, используя функцию из пакета spentenergy
	distance := spentenergy.Distance(t.Steps, float64(t.Height))
	// вычисляем среднюю скорость, используя функцию из пакета spentenergy
	meanSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)
	// проверяем вид тренировки и для каждого вида вычсляем калории
	// формируем и возвращем строку
	switch t.TrainingType {
	case "Бег":
		spentCalories, err := spentenergy.RunningSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories,
			), err

	case "Ходьба":
		spentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCalories,
			), err
	default:
		// если был передан неизвестный тип тренировки, возвращаем ошибку с текстом
		return "", errors.New("Training: ActionInfo: неизвестный тип тренировки")
	}
}
