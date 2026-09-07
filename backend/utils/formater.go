package utils

import (
	"backend/models"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func Time() (string, string) {
	current := time.Now()
	state := "am"

	// check wether the time is past or after morning
	if current.Hour() >= 12 {
		state = "pm"
	}

	date := time.Now().Format("Jan 2, 2006")
	hour := time.Now().Format("03:04" + " " + state)
	return date, hour
}

func FormatCategory(categories []string, postID string) []models.Category {
	var tab []models.Category
	for _, v := range categories {
		tab = append(tab, models.Category{Post_id: postID, Category: v})
	}
	return tab
}

func FormatViewers(viewers string, postID string) []models.Viewers {
	var tab []models.Viewers
	for _, v := range strings.Split(viewers, ",") {
		id, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			log.Println("❌ error while converting viewer id in format viewer")
			return nil
		}
		tab = append(tab, models.Viewers{Post_id: postID, User_id: id})
	}
	return tab
}

func FormatImgLink(name string) string {
	if name != "" {
		return fmt.Sprintf("http://localhost:8080/uploads/%s", name)
	}
	return "NULL"
}
