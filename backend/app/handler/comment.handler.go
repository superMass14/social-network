package handler

import (
	"backend/models"
	"backend/app/cors"
	"backend/app/service"
	"backend/utils"
	"log"
	"net/http"
	"strconv"
)

func COMMENTHandler(w http.ResponseWriter, r *http.Request) {
	// --------retrieving form values ----------
	//postID := "b01af696-f879-41a1-bfb0-70fa01852138" //?default value for testing purposes
	cors.SetCors(&w)
	log.Println("gotten userID in comment handler => ", r.FormValue("userID"))
	idUser, errconv := strconv.Atoi(r.FormValue("userID"))
	if errconv != nil {
		log.Println("❌ Error at converting string to int in userID of comment handler. ", errconv)
		msg := models.Errormessage{
			Type:       "Bad request",
			Msg:        "wrong id!",
			StatusCode: 400,
			Display:    true,
		}
		utils.Alert(w, msg)
		return
	}

	log.Println("--------------------------------------------")
	log.Println("          COMMENT Form values                 ")
	log.Println("--------------------------------------------")
	if r.Method == "POST" {
		log.Println("request type ", r.FormValue("type"))
		switch r.FormValue("type") {
		case "addComment":
			CommentContent := r.FormValue("content")
			postID := r.FormValue("postID")
			log.Println("[INFO] comment content: ", CommentContent) //debug
			log.Println("[INFO] postID create: ", postID)           //debug

			Image, errimg := utils.Uploader(w, r, 20, "image", "")
			if errimg != nil {
				log.Println("[INFO] imagelink error: ", errimg) //debug
			} else {
				log.Println("[INFO] imagelink: ", Image) //debug
			}

			CommentToCreate := models.Comment{
				Content: CommentContent,
				Post_id: postID,
				User_id: idUser,
				Media:   utils.FormatImgLink(Image),
			}
			log.Println(CommentToCreate)
			notOk, err := service.CommentServ.CreateComment(CommentToCreate)
			if notOk {
				log.Println("problem after create service ", err)
				utils.Alert(w, err)
				return
			} else {
				msg := models.Errormessage{
					Type:       "success",
					Msg:        "comment created successfully",
					StatusCode: 200,
					Display:    false,
				}
				utils.Alert(w, msg)
			}
		case "loadComments":
			log.Println("[FETCHING COMMENT DATA ◼◼◼]")
			postID := r.FormValue("postID")
			log.Println("debug get id: ", postID)
			commentTab, err := service.CommentServ.FetchComments(postID)
			if err != nil {
				log.Println("problem after fetch comment service ", err)
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
				utils.AlertData(w, models.WResponse{
					Type:       "loadPost",
					Data:       commentTab,
					StatusCode: 200,
					Display:    false,
					Msg:        "comments retrieved succesfully",
				})
			}
		}
		return
	}
	//-------not allowed method
	msg := models.Errormessage{
		Type:       "Not allowed",
		Msg:        "Method not allowed",
		StatusCode: 401,
		Display:    true,
	}
	utils.Alert(w, msg)
}
