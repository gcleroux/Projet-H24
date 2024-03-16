package config

// Since WASM doesn't get access to ENV VARS, we need to give a string
// at compile time for the SERVER_URL at compile time with the -ldflags tag
var SERVER_URL string
