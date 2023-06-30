package utils

type Util struct {
}
type Config struct {
	Mode struct {
		Network struct {
			Ja3      string `json:"JA3"`
			Redirect bool   `json:"Redirect"`
			TimeOut  int    `json:"TimeOut"`
			Proxy    string `json:"Proxy"`
			Agent    string `json:"Agent"`
		} `json:"Net"`
	} `json:"Modes"`
}

const (
	ChaceLoading = "Loading Cache: \u001B[36m[\u001B[39m%s\u001B[36m]\u001B[39m | \033[36m[\u001B[39m%s\u001B[36m]\u001B[39m"
	Initializing = "\033[36mInitializing...\033[39m "
)
