 
document.getElementById("pokemonForm").addEventListener("submit", function(event) {
event.preventDefault();
var pokemonId = document.getElementById("pokemonId").value;

fetch("/api/pokemon")
.then(response => response.json())
.then(data => {
    var pokemonInfo = document.getElementById("pokemonInfo");
    pokemonInfo.innerHTML = "<h2>Pokémon encontrado:</h2>" +
                            "<p>ID: " + data.id + "</p>" +
                            "<p>Nombre: " + data.name + "</p>" +
                            "<p>Especie: " + data.species + "</p>";
})
.catch(error => {
    console.error("Error al obtener datos de la Pokédex:", error);
});
});

 