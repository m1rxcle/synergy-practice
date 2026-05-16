package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var digits = map[rune][5]string{
	'0': {
		"***",
		"* *",
		"* *",
		"* *",
		"***",
	},
	'1': {
		" * ",
		" * ",
		" * ",
		" * ",
		" * ",
	},
	'2': {
		"***",
		"  *",
		"***",
		"*  ",
		"***",
	},
	'3': {
		"***",
		"  *",
		"***",
		"  *",
		"***",
	},
	'4': {
		"* *",
		"* *",
		"***",
		"  *",
		"  *",
	},
	'5': {
		"***",
		"*  ",
		"***",
		"  *",
		"***",
	},
	'6': {
		"***",
		"*  ",
		"***",
		"* *",
		"***",
	},
	'7': {
		"***",
		"  *",
		"  *",
		"  *",
		"  *",
	},
	'8': {
		"***",
		"* *",
		"***",
		"* *",
		"***",
	},
	'9': {
		"***",
		"* *",
		"***",
		"  *",
		"***",
	},
}

// Функция для печати ASCII-символов
// Принимает в себя объект digits и переменные day, month, year
// Ничего не возвращает 
// Печатает в консоль
func printASCII(digits map[rune][5]string, day int, month int, year int) {
	dateStr := fmt.Sprintf("%02d.%02d.%04d", day, month, year)

	for i :=0; i < 5; i++ {
		for _, char := range dateStr {
			if char == '.' {
				fmt.Print("   ")
			}else{
				fmt.Print(digits[char][i], " ")
			}
		}

		fmt.Println()
	}
}

// Функция для ввода данных
// Принимает в себя объект scanner и текст ввода
// Возвращает введенные данные
func input(scanner *bufio.Scanner, text string) int {
	exitKey := strings.ToLower("q") 
	
	fmt.Print(text)
	scanner.Scan()
	
	value := scanner.Text()

	value = strings.TrimSpace(value)
	
	if strings.ToLower(value) == exitKey {
		fmt.Println("Программа завершена...")
		os.Exit(0)
	}

	number,err := strconv.Atoi(value)

	

	if err != nil {
		fmt.Println("Введены некорректные данные.")
		os.Exit(1)
	}

	return number
}
// Функция для получения даты рождения
// Принимает в себя объект scanner
// Возвращает введенные данные
func getBirthday(scanner *bufio.Scanner) (int, int, int) {
	day := input(scanner, "Введите день вашего рождения: ")

	month := input(scanner, "Введите месяц вашего рождения: ")
	year := input(scanner, "Введите год вашего рождения: ")

	return day, month, year
}

// Функция для определения дня недели
// Принимает в себя день, месяц, год
// Возвращает день недели по средствам выбора через  switch case
func findWeekDayFromBirth(day int, month int, year int) string {
	weekday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday()

	switch weekday {
		case time.Monday:
			return "Понедельник."
		case time.Tuesday:
			return "Вторник."
		case time.Wednesday:
			return "Среда."
		case time.Thursday:
			return "Четверг."
		case time.Friday:
			return "Пятница."
		case time.Saturday:
			return "Суббота."
		case time.Sunday:
			return "Воскресенье."
		default:
			return "Ошибка в определении дня недели."
	}
	
}

// Функция для подсчета возраста
// Принимает в себя день, месяц, год
// Возвращает возраст
func countAge(day int, month int, year int)int {
	 today := time.Now()
	 age := today.Year() - year

	if today.Month() < time.Month(month) || (today.Month() == time.Month(month) && today.Day() < day){
		age --
	}

	return age
	
}

// Функция для определения високосного года
// Принимает в себя год
// Возвращает високосный или не високосный год
func isLeapYear(year int) string {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return "Високосный."
	}

	return "Не високосный."
}

// Функция для валидации даты
// Принимает в себя день, месяц, год
// Возвращает true или false
func validateDate(day int, month int, year int) bool {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)

	 if t.After(time.Now()) {
        return false
    }

	return t.Day() == day && int(t.Month()) == month && t.Year() == year
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	day, month, year := getBirthday(scanner)

	valid := validateDate(day, month, year)

	if !valid {
		fmt.Println("Введены некорректные данные.")
		return
	}

	week_day := findWeekDayFromBirth(day, month, year)
	age := countAge(day, month, year)

	
	

	fmt.Printf("\n")
	printASCII(digits, day, month, year)
	fmt.Printf("\n")

	fmt.Printf("День недели вашего рождения: %s\n", week_day)
	fmt.Printf("Ваш год был: %s\n", isLeapYear(year))
	fmt.Printf("Ваш возраст: %d\n", age)
	input(scanner, "Введите 'q' для выхода... ")
}