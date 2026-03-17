const input = document.getElementById("input")

input.addEventListener("input", () => {
	setTimeout(() => {
		input.classList.remove("form-input-bounce")
	}, 100)
	input.classList.add("form-input-bounce")
})
