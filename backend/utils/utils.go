package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
)

func ParseArrayInt(a []string) (tab []int, err error) {
	for _, v := range a {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		tab = append(tab, r)
	}
	return tab, nil
}

func GetCurrentDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

func FormatDuration(duration time.Duration) string {
	// Get the individual components of the duration
	days := duration / (24 * time.Hour)
	duration = duration % (24 * time.Hour)
	hours := duration / time.Hour
	duration = duration % time.Hour
	minutes := duration / time.Minute
	duration = duration % time.Minute
	seconds := duration / time.Second

	// Build the formatted string
	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d day", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d mn", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d s", seconds))
	}

	// Join the parts with ", " to create the final string
	return strings.Join(parts, ", ")
}

func IsOrderParam(orderBy string) bool {
	var orderParams = []string{
		"TIME-ASC",
		"TIME-DESC",
		"MOSTLIKED-ASC",
		"MOSTLIKED-DESC",
	}

	for _, v := range orderParams {
		if v == orderBy {
			return true
		}
	}
	return false
}

func IsAlphanumeric(input string) bool {
	pattern := "^[a-zA-Z0-9]*$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func IsAlpha(input string) bool {
	pattern := "^[a-zA-Z]*$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func VerifyUsername(username string) error {
	if username == "" {
		return nil
	}
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{3,15}$`
	regex := regexp.MustCompile(pattern)
	ok := regex.MatchString(username)
	if !ok {
		return fmt.Errorf("%v : not a valid username", username)
	}
	return nil
}

func VerifyName(s string) error {
	var maxChars = 25
	if s == "" {
		return fmt.Errorf("%v: as", s)
	}
	if len(s) > maxChars || !IsAlpha(s) {
		return fmt.Errorf("%v: invalid name", s)
	}
	return nil
}

func VerifyAboutMe(s string) error {
	var maxChars = 1500
	if len(s) > maxChars {
		return fmt.Errorf("%v: invalid about me", s)
	}
	return nil
}

func IsValidEmail(email string) error {
	emailRegex := `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`

	re := regexp.MustCompile(emailRegex)
	ok := re.MatchString(email)
	if !ok {
		return fmt.Errorf("%v: invalid email", email)

	}
	return nil
}

func IsValidImageType(s string) bool {
	s = strings.ToLower(s)
	return strings.HasSuffix(s, ".jpeg") ||
		strings.HasSuffix(s, ".png") ||
		strings.HasSuffix(s, ".jpg") ||
		strings.HasSuffix(s, ".bmp") ||
		strings.HasSuffix(s, ".webp") ||
		strings.HasSuffix(s, ".gif")
}

func GenImageName(image string) (string, error) {
	idImg, errImg := uuid.NewV4()
	if errImg != nil {
		return "", errors.New("cannot generate id for img name")
	}
	return fmt.Sprintf("%s%s", idImg, image), nil
}

func IsValidPassword(password string) error {
	if len(password) < 4 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

func GenerateToken() string {
	token, err := uuid.NewV4()
	if err != nil {
		log.Println("Error GenerateToken: ", err)
	}
	return token.String()

}
