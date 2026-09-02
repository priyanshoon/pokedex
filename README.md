# pokedex

a dead simple cil game to explore, catch and inspect the pokemon.
it uses 3rd party api for fetching pokemon details. it also caches the
result so that we don't have to make another request for same value.

I didn't used redis for cache, instead build a simple cache (key-value) with dict
