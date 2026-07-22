### Build a Pokedex in Golang
This project was to build an interactive command-line Pokedex REPL using Go to learn HTTP and file tree / heirarchy of a Go project.
This tool interacts with the [PokéAPI](https://pokeapi.co/) to allow users to explore areas, catch and inspect wild Pokemon, and attempt to catch 'em all to build out their personal Pokedex.

#### Features
* **Interactive REPL Interface**: CLI command handling with built-in help and exit commands for users.
* **Built-in Cache**: Custom cache to avoid redundant network requests and respect PokeAPI rate limits.
* **Map Navigation**: Listing areas and allowing users to go forward and back pages of areas. (`map` and `mapb`)
* **Area Exploration**: Inspect specific location areas to see a list of what wild pokemon reside there. (`explore`)
* **Catching and Inspect Mechanics**: Dynamic catch rate calculation using `math/rand` and a Pokemon's `base_experience`. (`catch`)
* **Pokedex**: Tracks and stores all successfully caught Pokemon across a users session. (`pokedex`)

##### Learned Skills
* How to use HTTP within code to make requests and how to use the JSON responses to collect data into Go structs for use in other areas of the project.
* How to set up my own Go projects file tree and make internal packages.
* How to set up a REPL in Go and set up a cache for quicker stored data access rather than doing redundant network calls.

##### Possible Additions
* More unit tests could never hurt.
* Maybe rework the catch system to use the actual catch rate of the pokemon.
* Add a shop that would allow users to use different Pokeballs to influence catch rates.
* Random encounters with wild pokemon.
* Add a party mechanic to allow caught pokemon to level up over time.
* Allow pokemon to evolve based off of level / a set amount of time.
* Battle simulations with wild pokemon.
* Persist a user's pokedex to disk so they can save their progress between sessions.

This projects was very enjoyable all the way through, as much as the internal packages and the separation of concerns threw me for a bit of a loop.