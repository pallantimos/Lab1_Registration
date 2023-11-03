package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func checkRegistrate(login string, pass string, pass2 string) (string, bool) {
	err := ""
	isCorrect := true

	err, isCorrect = checkPass(pass, pass2)

	if isCorrect {
		err, isCorrect = checkLogin(login)
	}

	return err, isCorrect
}

func checkLogin(login string) (string, bool) {
	listLogin := [5]string{"Aldar", "Aleksey", "Ivan", "Mikhail", "Krug"}

	for i := 0; i < len(listLogin); i++ {
		if login == listLogin[i] {
			return "Логин уже существует", false
		}

	}

	if login == "" {
		return "Пустая строка в качества логина", false
	}

	if utf8.RuneCountInString(login) < 5 {
		return "Логин меньше 5 символов", false
	}

	reLogin := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(login)
	reMail := regexp.MustCompile(`^\w+@\w+\.\w+$`).MatchString(login)
	rePhone := regexp.MustCompile(`^\+\d{1,3}-\d{3}-\d{3}-\d{4}$`).MatchString(login)

	if containsPlus := strings.Contains(login, "+"); containsPlus && !rePhone {
		return "Номер телефона не удовлетворяет заданному формату +x-xxx-xxx-xxxx", false
	}

	if containsAt := strings.Contains(login, "@"); containsAt && !reMail {
		return "Email не удовлетворяет общему формату xxx@xxx.xxx", false
	}

	if !reLogin && !reMail && !rePhone {
		return "Логин содержит символы, отличные от латиницы, цифр и знака подчеркивания", false
	}

	return "", true
}

func checkPass(pass string, pass2 string) (string, bool) {
	isUpperLetter := false
	isDownLetter := false
	isDigit := false
	isSymbol := false

	if pass == "" || pass2 == "" {
		return "Пустая строка в качестве пароля", false
	}

	for _, r := range pass {
		if unicode.Is(unicode.Latin, r) {
			return "Пароль содержит латиницу", false
		}
		if unicode.IsUpper(r) {
			isUpperLetter = true
		} else if unicode.IsLetter(r) {
			isDownLetter = true
		} else if unicode.IsDigit(r) {
			isDigit = true
		} else {
			isSymbol = true
		}
	}

	if !isDownLetter {
		return "Пароль не содержит строчную букву", false
	}

	if !isUpperLetter {
		return "Пароль не содержит заглавную букву", false
	}

	if !isSymbol {
		return "Пароль не содержит спецсимвола", false
	}

	if !isDigit {
		return "Пароль не содержит цифру", false
	}

	if utf8.RuneCountInString(pass) < 7 {
		return "Пароль меньше 7 символов", false
	}

	if pass != pass2 {
		return "Пароли не совпадают", false
	}

	return "", true
}

func main() {
	log := logrus.New()
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл логов: %v", err)
	}
	log.SetOutput(file)
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{})

	log.Info("Приложение запущено")
	log.Info("Логгер сконфигурирован")

	var pass, pass2, login string
	print("Введите логин:\n")
	fmt.Scan(&login)
	log.Info("Введено значение логина: ", login)

	print("Введите пароль:\n")
	fmt.Scan(&pass)
	log.Info("Введено значение пароля: ", pass)

	print("Повторите пароль:\n")
	fmt.Scan(&pass2)
	log.Info("Введено значение повторного пароля: ", pass2)

	errorRegistrate, isCorrect := checkRegistrate(login, pass, pass2)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if isCorrect {
		log.Info("Логин ", login, " Успешная регистрация")
		print("Успешная Регистрация")
	} else {
		log.Error("Логин = ", login, " Пароль = ", hashedPassword, " Ошибка = ", errorRegistrate)
		print(errorRegistrate)
	}

	defer file.Close()
}
