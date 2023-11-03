package main

import "testing"

func TestRegistratePass(t *testing.T) {
	t.Run("Правильный пароль", func(t *testing.T) {
		pass := "Пароль123!"
		pass2 := "Пароль123!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)

		if err != "" {
			t.Errorf("Ошибки не ожидалось, получено: %s", err)
		}

		if !isCorrect {
			t.Errorf("Ожидалось isCorrect будет true, получено false")
		}
	})

	t.Run("Нет заглавной буквы в пароле", func(t *testing.T) {
		pass := "пароль123!"
		pass2 := "пароль123!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль не содержит заглавную букву"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Нет спецсимвола в пароле", func(t *testing.T) {
		pass := "Пароль123"
		pass2 := "Пароль123"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль не содержит спецсимвола"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Пароли не совпадают", func(t *testing.T) {
		pass := "Пароль123!"
		pass2 := "Пароль123!!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароли не совпадают"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Латиница в пароле", func(t *testing.T) {
		pass := "Pass123!"
		pass2 := "Pass123!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль содержит латиницу"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Нет цифры в пароле", func(t *testing.T) {
		pass := "Пароль!"
		pass2 := "Пароль!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль не содержит цифру"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Короткий Пароль", func(t *testing.T) {
		pass := "Пасс1!"
		pass2 := "Пасс1!"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль меньше 7 символов"

		if err != expectedError {
			t.Errorf("Ожидаемая ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Нет строчной буквы в пароле", func(t *testing.T) {
		pass := "ПАРОЛЬ123"
		pass2 := "ПАРОЛЬ123"
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пароль не содержит строчную букву"

		if err != expectedError {
			t.Errorf("Ожидаемая ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Пустая строка в качестве пароля первого пароля", func(t *testing.T) {
		pass := ""
		pass2 := ""
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пустая строка в качестве пароля"

		if err != expectedError {
			t.Errorf("Ожидаемая ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Пустая строка в качестве пароля второго пароля", func(t *testing.T) {
		pass := "Пароль123!"
		pass2 := ""
		err, isCorrect := checkRegistrate("NewUser", pass, pass2)
		expectedError := "Пустая строка в качестве пароля"

		if err != expectedError {
			t.Errorf("Ожидаемая ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})
}

func TestRegistrateLogin(t *testing.T) {
	t.Run("Логин уже существует", func(t *testing.T) {
		login := "Aldar"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Логин уже существует"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет true, получено false")
		}
	})

	t.Run("Пустая строка логина", func(t *testing.T) {
		login := ""
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Пустая строка в качества логина"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect Будет false, получено true")
		}
	})

	t.Run("Короткий логин", func(t *testing.T) {
		login := "logi"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Логин меньше 5 символов"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Некорректные символы в логине", func(t *testing.T) {
		login := "User$Name"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Логин содержит символы, отличные от латиницы, цифр и знака подчеркивания"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Некорректный номер телефона", func(t *testing.T) {
		login := "+7913-913-759-1929"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Номер телефона не удовлетворяет заданному формату +x-xxx-xxx-xxxx"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Некорректный email", func(t *testing.T) {
		login := "@dsadl"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Email не удовлетворяет общему формату xxx@xxx.xxx"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Некорректный email", func(t *testing.T) {
		login := "@dsadl"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")
		expectedError := "Email не удовлетворяет общему формату xxx@xxx.xxx"

		if err != expectedError {
			t.Errorf("Ожидалась ошибка %s, получено: %s", expectedError, err)
		}

		if isCorrect {
			t.Errorf("Ожидалось isCorrect будет false, получено true")
		}
	})

	t.Run("Правильный логин", func(t *testing.T) {
		login := "NewUser"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")

		if err != "" {
			t.Errorf("Ошибки не ожидалось, получено: %s", err)
		}

		if !isCorrect {
			t.Errorf("Ожидалось isCorrect будет true, получено false")
		}
	})

	t.Run("Правильный email", func(t *testing.T) {
		login := "aldar@aldar.ru"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")

		if err != "" {
			t.Errorf("Ошибки не ожидалось, получено: %s", err)
		}

		if !isCorrect {
			t.Errorf("Ожидалось isCorrect будет true, получено false")
		}
	})

	t.Run("Правильный телефон", func(t *testing.T) {
		login := "+7-913-759-1929"
		err, isCorrect := checkRegistrate(login, "Пароль123!", "Пароль123!")

		if err != "" {
			t.Errorf("Ошибки не ожидалось, получено: %s", err)
		}

		if !isCorrect {
			t.Errorf("Ожидалось isCorrect будет true, получено false")
		}
	})
}
