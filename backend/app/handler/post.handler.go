package handler

import (
	"backend/app/cors"
	"backend/app/handler/user"
	"backend/app/service"
	"backend/models"
	"backend/utils"
	"backend/utils/seed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func POSTHandler(w http.ResponseWriter, r *http.Request) {
	// --------retrieving form values ----------
	cors.SetCors(&w)
	log.Println("gotten userID => ", r.FormValue("userID"))
	userID, errconv := strconv.Atoi(r.FormValue("userID"))
	if errconv != nil {
		log.Println("❌ Error at converting string to int in userID in post. ", errconv)
	}
	if r.Method != "POST" {
		msg := models.Errormessage{
			Type:       "Not allowed",
			Msg:        "Method not allowed",
			StatusCode: 401,
			Display:    true,
		}
		utils.Alert(w, msg)
	} else {
		log.Println("in post")
		requestType := r.FormValue("type")
		switch requestType {
		case "createPost":
			log.Println("in creation")
			log.Println("--------------------------------------------")
			log.Println("          Post Form values                 ")
			log.Println("--------------------------------------------")
			PostContent := r.FormValue("content")
			log.Println("[INFO] userID: ", userID)            //debug
			log.Println("[INFO] post content: ", PostContent) //debug
			Sport := r.FormValue("Sport")
			Health := r.FormValue("Health")
			Music := r.FormValue("Music")
			News := r.FormValue("News")
			Others := r.FormValue("Others")
			Techno := r.FormValue("Tech")

			categorie := []string{Health, Sport, News, Techno, Others, Music}
			var sortCat []string
			for _, v := range categorie {
				if v != "" {
					sortCat = append(sortCat, v)
				}
			}
			categorie = sortCat
			log.Println("[INFO] categories: ", categorie) //debug

			Privacy := r.FormValue("privacy")
			log.Println("[INFO] privacy: ", Privacy) //debug

			Viewers := fmt.Sprintf("%v, %s", userID, r.FormValue("viewers"))
			if strings.Split(Viewers, ", ")[1] == "" && Privacy == "almost" {
				msg := models.Errormessage{
					Type:       "Bad request",
					Msg:        "must selected viewers",
					StatusCode: 400,
					Display:    false,
				}
				utils.Alert(w, msg)
				return
			}
			log.Println("[INFO] viewers: ", Viewers) //debug

			Image, errimg := utils.Uploader(w, r, 20, "image", "")
			if errimg != nil {
				msg := models.Errormessage{
					Type:       "Bad request",
					Msg:        errimg.Error(),
					StatusCode: 400,
					Display:    false,
				}
				utils.Alert(w, msg)
				return
			}
			log.Println("[INFO] imagelink: ", Image) //debug
			PostToCreate := models.Post{
				ToIns: models.InsPost{
					Content:  PostContent,
					Media:    utils.FormatImgLink(Image),
					User_id:  userID,
					Group_id: 0,
					Privacy:  Privacy,
				},
				Categories: categorie,
				Viewers:    Viewers,
			}
			log.Println(PostToCreate)
			notOk, err := service.PostServ.CreatePost(PostToCreate)
			if notOk {
				log.Println("problem after create service ", err)
				utils.Alert(w, err)
				return
			} else {
				msg := models.Errormessage{
					Type:       "success",
					Msg:        "post created successfully",
					StatusCode: 200,
					Display:    false,
				}
				log.Println("post created successfully")
				utils.Alert(w, msg)
			}
		case "loadPosts":
			log.Println("[FETCHING POST DATA ◼◼◼] with ", userID)
			postTab, err := service.PostServ.GetPost(userID)
			if err != nil {
				log.Println("problem after get service ", err)
				msg := models.Errormessage{
					Type:       "Internal Servor Error",
					Msg:        "Oops something wrong occured !",
					StatusCode: 500,
					Display:    true,
				}
				utils.Alert(w, msg)
				return
			} else {
				//log.Println("Gotten => ", postTab)
				//log.Println("Gotten top => ", postTab[0])

				var db = seed.CreateDB()
				defer db.Close()
				followers, err := user.GetFollowers(db, userID)
				if err != nil {
					fmt.Println(err)
					return
				}
				utils.AlertData(w, models.WResponse{
					Type:       "loadPost",
					Data:       postTab,
					Followers:  followers,
					StatusCode: 200,
					Display:    false,
					Msg:        "posts retrieved succesfully",
				})
			}
		}
	}
}
