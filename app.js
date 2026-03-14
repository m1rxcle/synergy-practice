const firstNumber = document.getElementById("firstNumber")
const secondNumber = document.getElementById("secondNumber")
const result = document.getElementById("result")
const plus = document.getElementById("plus")
const minus = document.getElementById("minus")
const multiply = document.getElementById("multiply")
const divide = document.getElementById("divide")
const errorContainer = document.getElementById("error-container")
const error = document.getElementById("error")
const errorText = document.getElementById("error-text")

function showError() {
	errorContainer.classList.remove("hidden")
}

function hideError() {
	errorContainer.classList.add("hidden")
}

function chooseAction(action) {
	const number1 = +firstNumber.value
	const number2 = +secondNumber.value

	if (isNaN(number1) || isNaN(number2)) {
		result.textContent = 0
		firstNumber.value = ""
		secondNumber.value = ""
		error.textContent = "Ошибка: Введите корректные числа!"
		showError()
		return
	}

	hideError()

	switch (action) {
		case "-":
			result.textContent = number1 - number2
			break

		case "*":
			result.textContent = number1 * number2
			break

		case "/":
			if (number2 === 0) {
				firstNumber.value = ""
				secondNumber.value = ""
				error.textContent = "На ноль делить нельзя!"
				showError()
				return
			}

			result.textContent = number1 / number2
			break

		default:
			result.textContent = number1 + number2
			break
	}

	firstNumber.value = ""
	secondNumber.value = ""
}

plus.addEventListener("click", () => chooseAction("+"))
minus.addEventListener("click", () => chooseAction("-"))
multiply.addEventListener("click", () => chooseAction("*"))
divide.addEventListener("click", () => chooseAction("/"))
