document.getElementById("form").addEventListener("submit", e => {
	if (e.submitter.id === "create") {
		return;
	}

	e.preventDefault();

	const url = document
		.getElementById("url")
		.value.replace(window.location.href, "");
	console.log(window.location);
	window.location.href = `/info/${url}`;
});
