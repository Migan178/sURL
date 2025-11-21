function changeIcon(id) {
	const icon = document.getElementById(id);

	icon.classList.remove("bi-clipboard");
	icon.classList.add("bi-clipboard-check-fill");

	setTimeout(() => {
		icon.classList.remove("bi-clipboard-check-fill");
		icon.classList.add("bi-clipboard");
	}, 2500);
}

document.getElementById("copy-original-url").addEventListener("click", e => {
	e.preventDefault();

	const a = document.getElementById("original-url");

	navigator.clipboard.writeText(a.href);
	changeIcon("original-url-copy-icon");
});

document.getElementById("copy-redirect-url").addEventListener("click", e => {
	e.preventDefault();

	const a = document.getElementById("redirect-url");

	navigator.clipboard.writeText(a.href);
	changeIcon("redirect-url-copy-icon");
});
