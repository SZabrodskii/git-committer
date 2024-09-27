package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type AnekdotService struct{}

func NewAnekdotService() *AnekdotService {
	return &AnekdotService{}
}

func (as *AnekdotService) GetRandomAnekdot() (string, error) {
	resp, err := http.Get("https://www.anekdot.ru/random/anekdot/")
	if err != nil {
		return "", errors.New("failed to fetch anekdot")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	re := regexp.MustCompile(`<div class="text">(.+?)</div>`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) < 2 {
		return "", errors.New("anekdot not found")
	}

	anekdot := matches[1]
	anekdot = regexp.MustCompile(`<[^>]*>`).ReplaceAllString(anekdot, "")
	return anekdot, nil
}

func (as *AnekdotService) SaveAnekdotToFile(anekdot, repoName, fileName string) error {
	filePath := filepath.Join(repoName, fmt.Sprintf("%s.txt", fileName))

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(anekdot)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
