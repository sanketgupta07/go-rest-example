package model

// Smaple Respose
//  URL: https://pokeapi.co/api/v2/pokedex/kanto/
//{
// 	"name":"kanto",
// 	"region": {
// 	  "url":"http:\/\/pokeapi.co\/api\/v2\/region\/1\/",
// 	  "name":"kanto"
// 	},
// 	"version_groups":[ ... ]
// 	],
// 	"is_main_series":true,
// 	"descriptions":[ ... ],
// 	"pokemon_entries":[
// 	  {
// 		"entry_number": 1,
// 		"pokemon_species": {
// 		  "url":"http:\/\/pokeapi.co\/api\/v2\/pokemon-species\/1\/",
// 		  "name":"bulbasaur"
// 		}
// 	  }
// 	  ...
// 	]
//   }

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

//PokemonSpecies A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}
