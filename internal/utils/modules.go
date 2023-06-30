package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Danny-Dasilva/fhttp"
)

func (u *Util) ReadFile(files string) ([]string, error) {
	var lines []string
	file, err := os.Open(files)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func (u *Util) LoadConfig(filename string) (Config, error) {
	var config Config
	conf, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer conf.Close()

	err = json.NewDecoder(conf).Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (u *Util) Cookie(url string) string {
	resp, err := http.Get("https://" + url)
	if err != nil {
		fmt.Println("Error while getting cookies", err)
		return ""
	}

	defer resp.Body.Close()

	if resp.Cookies() != nil {
		for _, cookie := range resp.Cookies() {
			return cookie.Value
		}
	}
	return ""
}
