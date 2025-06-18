package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Println("DataParser: Info: ошибка парсинга:", err)
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("DataParser: Info: ошибка генерации отчета:", err)
		}
		// Выводим информацию
		log.Println(info)
	}
}
