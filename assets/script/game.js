const input = document.getElementById("user-value");

// Maintenir le focus sur l'input même en cas de clic ailleurs
document.addEventListener("click", (event) => {
  input.focus(); // Redonne le focus à l'input
  event.preventDefault(); // Empêche tout autre comportement
});

// Maintenir le focus également si l'utilisateur tente de le perdre
input.addEventListener("blur", () => {
  setTimeout(() => input.focus(), 0); // Re-focus immédiatement après le blur
});
