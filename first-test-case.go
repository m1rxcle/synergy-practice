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


func input(scanner *bufio.Scanner, text string) int {
	fmt.Print(text)
	scanner.Scan()
	
	value := scanner.Text()

	value = strings.TrimSpace(value)

	number,err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("Введены некорректные данные")
		os.Exit(1)
	}

	return number
}

func getBirthday(scanner *bufio.Scanner) (int, int, int) {
	day := input(scanner, "Введите день вашего рождения: ")

	month := input(scanner, "Введите месяц вашего рождения: ")
	year := input(scanner, "Введите год вашего рождения: ")

	return day, month, year
}

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

func countAge(day int, month int, year int)int {
	 today := time.Now()
	 age := today.Year() - year

	if today.Month() < time.Month(month) || (today.Month() == time.Month(month) && today.Day() < day){
		age --
	}

	return age
	
}


func isLeapYear(year int) string {
	if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
		return "Високосный."
	}

	return "Не високосный."
}

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
		fmt.Println("Введены некорректные данные")
		return
	}

	week_day := findWeekDayFromBirth(day, month, year)
	age := countAge(day, month, year)

	
	

	fmt.Printf("\n")
	printASCII(digits, day, month, year)
	fmt.Printf("\n")

	fmt.Printf("День недели вашего рождения: %s\n", week_day)
	fmt.Printf("Ваш год был: %s\n", isLeapYear(year))
	fmt.Printf("Ваш возраст: %d", age)
}