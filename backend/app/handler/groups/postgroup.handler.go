package groups

import (
	"backend/models"
	"backend/app/cors"
	"backend/app/service"
	"backend/utils"
	"log"
	"net/http"
	"strconv"
)

func PostGroupHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	switch r.Method {
	case "POST":
		userID, errconv := strconv.Atoi(r.FormValue("userID"))
		if errconv != nil {
			log.Println("❌ Error at converting string to int in userID in postGroupHandler. ", errconv)
			return
		}
		content := r.FormValue("content")
		groupID, err := GetGroupIDFromRequest(w, r)
		if err != nil {
			return
		}
		categories := getCategoriesFromForm(r)
		imageLink, err := utils.Uploader(w, r, 20, "image", "")
		if err != nil {
			log.Println("Error to Upload img: ", err)
			return
		}

		post := models.Post{
			ToIns: models.InsPost{
				Content:  content,
				Media:    utils.FormatImgLink(imageLink),
				User_id:  userID,
				Group_id: groupID,
				Privacy:  "public",
			},
			Categories: categories,
			Viewers:    "",
		}

		notOk, err1 := service.PostServ.CreatePost(post)
		if notOk {
			log.Println("problem after post creation in group handler ", err)
			utils.Alert(w, err1)
			return
		} else {
			msg := models.Errormessage{
				Type:       "success",
				Msg:        "post created successfully",
				StatusCode: 200,
				Display:    false,
			}
			utils.Alert(w, msg)
		}
	}
}

func getCategoriesFromForm(r *http.Request) []string {
	categories := []string{"Health", "Sport", "News", "Techno", "Others", "Music"}
	var selectedCategories []string
	for _, cat := range categories {
		if r.FormValue(cat) != "" {
			selectedCategories = append(selectedCategories, cat)
		}
	}
	return selectedCategories
}
