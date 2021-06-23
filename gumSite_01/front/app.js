// var ourRequest = new XMLHttpRequest();

// ourRequest.open('GET', 'http://localhost:8080/selectAllProducts')
// ourRequest.onload = function () {
// 	var ourData = JSON.parse(ourRequest.responseText)
// 	console.log(ourData[0]);
// }
// ourRequest.send();

// async function GetPruducts() {
// 	let result = await fetch('http://localhost:8080/selectAllProducts', {
// 		method: "GET",
// 		headers: "Access-Control-Allow-Origin: http://localhost:8080'"
// 		mode: "cors",
// 		credentials: 'include'
// 	})

// 	let data = await result.json();
// 	console.log(data);
// }
Access - Control - Allow - Origin
var invocation = new XMLHttpRequest();
var url = 'http://localhost:8080/selectAllProducts';

function callOtherDomain() {
	if (invocation) {
		invocation.open('GET', url, true);
		invocation.onload = function () {
			var ourData = JSON.parse(invocation.responseText)
			console.log(ourData)
		}
		invocation.setRequestHeader = 'Access-Control-Allow-Origin: http://localhost:8080'
		invocation.send();
	}
}

document.addEventListener("DOMContentLoaded", () => {
	callOtherDomain();
});