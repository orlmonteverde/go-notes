let burger = document.getElementById('navbar-burger')
let menu = document.getElementById('navbar-menu')

burger.addEventListener('click', (e) => {
	burger.classList.toggle('is-active')
	menu.classList.toggle('is-active')
})
