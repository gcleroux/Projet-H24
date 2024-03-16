package networking

// Since WASM doesn't get access to ENV VARS, we need to give a string
// at compile time for the SERVER_URL at compile time with the -ldflags tag
const SERVER_URL string = "ws://localhost:8888/position"
