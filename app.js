const output = document.getElementById("output")
const increment = document.getElementById("increment")
const decrement = document.getElementById("decrement")
const bg = document.getElementById("counter-container")
const message = document.getElementById("message-container")

let count = 0

output.textContent = count

function printCount(count) {
	output.textContent = count

	if (count > 0) {
		bg.style.background = "yellow"
	} else if (count === 0) {
		bg.style.background = "red"
	} else {
		bg.style.background = "green"
	}
}

function disableButton(count) {
	increment.disabled = count > 9
	decrement.disabled = count < -9

	increment.classList.toggle("disabled", count > 9)
	decrement.classList.toggle("disabled", count < -9)
}

function showWarningMessage(count) {
	if (count > 9 || count < -9) {
		message.classList.remove("closed")
	} else {
		message.classList.add("closed")
	}
}

function update(count) {
	printCount(count)
	disableButton(count)
	showWarningMessage(count)
}

increment.addEventListener("click", () => {
	count++
	update(count)
})

decrement.addEventListener("click", () => {
	count--
	update(count)
})
