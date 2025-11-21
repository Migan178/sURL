document.getElementById("find").addEventListener("submit", e => {
	e.preventDefault();

	const url = document
		.getElementById("find-url")
		.value.replace(window.location.href, "");
	console.log(window.location);
	window.location.href = `/info/${url}`;
});
