package config

var C *Config

type Config struct {
	Window  window
	Network network
}

type window struct {
	Title  string
	Width  int
	Height int
}

type network struct {
	Url string
}

func NewConfig() *Config {
	C = &Config{
		Window: window{
			Title:  "Multiplayer Golang Platformer",
			Width:  640,
			Height: 360,
		},
		Network: network{
			Url: "http://localhost:8888",
		},
	}
	return C
}
