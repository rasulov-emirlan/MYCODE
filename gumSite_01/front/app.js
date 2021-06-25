const productsDOM = document.querySelector('.products')
const cartItemsDOM = document.querySelector('.cart-items')
const cartShowBtn = document.querySelector('.cart-icon')
const cartOverlay = document.querySelector('.cart-overlay')
const cartDOM = document.querySelector('.cart')


let cart = []

cartShowBtn.addEventListener("click", function () {
	cartOverlay.className = '.cart-overlay-show'

	cartDOM.className = '.cart-show'
	console.log("log");
})
function setupBtns() {
	const buttons = [...document.querySelectorAll(".bag-btn")];
	buttonsDOM = buttons;
	buttons.forEach(button => {
		let id = button.dataset.id;
		let inCart = cart.find(item => item.id === id);
		if (inCart) {
			button.innerHTML = "In Cart";
			button.disabled = true;
			return
		}
		button.addEventListener("click", (event) => {
			event.target.innerText = "In Cart";
			event.target.disabled = true;

			let cartItem = { ...getProductByName(event), amount: 1 };

			cart = [...cart, cartItem];

			Storage.saveCart(cart);
		});
	});
}

function getProductByName(thename) {
	fetch('http://localhost:8080/selectProdcutByName', {
		method: "GET",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			name: thename
		}),
	}).then(res => res.json())
		.then(data => {
			let elements = `
				<div class="product">
					<img src="./images/${data.name}.png" al="Food" style="height: 200px;">
					<div class="product_description">
						<h1>${data.name}</h1>
						<p class="price">${data.cost} сом</p>
						<button class="bag-btn" value="${data.name}">Добавить в корзину</button>
					</div>
	
				</div>
				`
			return elements
		});
}

function getAllProducts() {
	fetch('http://localhost:8080/selectAllProducts', {
		method: "GET",
	}).then(res => res.json())
		.then(data => {
			let elements = ""
			data.forEach(element => {
				elements += `
				<div class="product">
					<img src="./images/${element.name}.png" al="Food" style="height: 200px;">
					<div class="product_description">
						<h1>${element.name}</h1>
						<p class="price">${element.cost} сом</p>
						<button class="bag-btn">Добавить в корзину</button>
					</div>
	
				</div>
				`
			});
			productsDOM.innerHTML = elements;
		})
}

getAllProducts()
document.addEventListener("DOMContentLoaded", () => {

});