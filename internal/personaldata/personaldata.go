package personaldata

import "fmt"

// структура с персональными данными: имя, вес и рост пользователя
type Personal struct {
	// TODO: добавить поля
	Name string
	Weight float64
	Height float64
}

// метод для вывода персональных данных на экран
func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
