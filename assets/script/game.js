const input = document.getElementById("user-value");
const button = document.getElementById("bonus-button");

// Maintenir le focus sur l'input sauf si l'utilisateur clique sur le bouton
document.addEventListener("click", (event) => {
  if (event.target !== button) {
    input.focus(); // Redonne le focus à l'input
    event.preventDefault(); // Empêche tout autre comportement
  }
});

// Maintenir le focus même si l'utilisateur tente de le perdre
input.addEventListener("blur", () => {
  setTimeout(() => input.focus(), 0); // Re-focus immédiatement après le blur
});

// Permettre au bouton de fonctionner normalement tout en gardant le focus sur l'input
button.addEventListener("click", (event) => {
  console.log("Bouton cliqué !");
  // Ici vous pouvez inclure d'autres actions pour le bouton
});
