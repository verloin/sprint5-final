package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// пакет отвечает за дневную активность, или проще говоря за прогулки
// структура с данными, в которую будет встроена структура с персональными данными
type DaySteps struct {
	// TODO: добавить поля
	Steps int  // количество шагов
	Duration time.Duration  // длительность прогулки
	personaldata.Personal  // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

// метод парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	// Разделить строку datastring на слайс строк
	parts := strings.Split(datastring, ",")
	// Проверим, что у нас 2 части
	if len(parts) != 2 {
		return errors.New("DaySteps: Parse: the length is less or more than three") 
	}
	// Преобразуем первый элемент в int (количество шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("DaySteps: Parse: failed to convert steps: %w", err)
	}
	// Проверяем, что количество шагов больше 0
	if steps <= 0 {
		return errors.New("DaySteps: Parse: incorrect number of steps")
	}
	ds.Steps = steps
	// Преобразуем второй элемент в time.Duration и обрабатываем ошибки
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("DaySteps: Parse: failed to parse duration: %w", err)
	}
	if duration <= 0 {
		return  errors.New("DaySteps: Parse: negative durations are not allowed")
	}
	// сохраняем полученное значение в соответствующем поле структуры DaySteps.
	ds.Duration = duration
	return nil
}

// метод формирует и возвращает строку с данными о прогулке
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	// вычисляем дистанцию, используя функцию из пакета spentenergy
	distance := spentenergy.Distance(ds.Steps, float64(ds.Height))
	// вычисляем количество сожжённых калорий и обрабатываем ошибки
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Weight), float64(ds.Height), ds.Duration)
	if err != nil {
		return "", fmt.Errorf("TraiDayStepsning: ActionInfo: failed to convert steps: %w", err)
	}
	// формируем и возвращаем строку с информацией
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
