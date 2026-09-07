package groups

import (
	"backend/models"
	"backend/app/cors"
	utils "backend/utils"
	"backend/utils/seed"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// var group models.Group

// var UserId int = 1

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)

	var db = seed.CreateDB()
	defer db.Close()
	var group models.Group = parseFormData(w, r)

	groupNameTrim := strings.TrimSpace(group.Name)

	if len(groupNameTrim) == 0 {
		// http.Error(w, "Le champ 'Name' est obligatoire", http.StatusBadRequest)

		msg := models.Errormessage{
			Type:       "Create group",
			Msg:        "Empty name",
			StatusCode: http.StatusBadRequest,
			Display:    true,
		}
		utils.Alert(w, msg)
		return
	}

	isExist, err := checkGroupExists(db, group.Name)
	if err != nil {
		log.Println("Erreur lors de la vérification du groupe: ", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)

		msg := models.Errormessage{
			Type:       "Create group",
			Msg:        "Err on create group",
			StatusCode: http.StatusConflict,
			Display:    true,
		}
		utils.Alert(w, msg)
		return
	}

	if isExist {
		// response := map[string]string{"message": "Le groupe existe déjà"}
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusConflict)
		// json.NewEncoder(w).Encode(response)
		msg := models.Errormessage{
			Type:       "Create group",
			Msg:        "Group already exist",
			StatusCode: http.StatusConflict,
			Display:    true,
		}
		utils.Alert(w, msg)
		return
	}

	insertGroupCreated(db, group)

	// response := map[string]string{}
	w.Header().Set("Content-Type", "application/json")

	msg := models.Errormessage{
		Type:       "Create group",
		Msg:        "Groupe créé",
		StatusCode: http.StatusCreated,
		Display:    false,
	}
	utils.Alert(w, msg)

	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(response)
}

func insertGroupCreated(db *sql.DB, group models.Group) {
	stmt, err := db.Prepare("INSERT INTO groups(name, description, id_user_create, avatar, creation_date) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Erreur lors de la préparation de la requête d'insertion: ", err)
		return
	}
	defer stmt.Close()
	if group.Avatar == "NULL" {
		group.Avatar = "http://localhost:8080/uploads/35a23cb2-d742-48fa-8d77-8a6e19bf571fsnk.jpg"
	}

	result, err := stmt.Exec(group.Name, group.Description, group.ID_User_Create, group.Avatar, group.Creation_Date)
	if err != nil {
		log.Println("Erreur lors de l'insertion des données: ", err)
		return
	}

	id, _ := result.LastInsertId()
	group.ID = int(id)

	stmt, err = db.Prepare("INSERT INTO group_followers(user_id, group_id) VALUES(?, ?)")
	if err != nil {
		log.Println("Erreur lors de la préparation de la requête d'insertion: ", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(group.ID_User_Create, group.ID)
	if err != nil {
		log.Println("Erreur lors de l'insertion des données: ", err)
		return
	}
}
func parseFormData(w http.ResponseWriter, r *http.Request) models.Group {
	var avatarGroup = "profil_group_avatar.jpg"
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Impossible de lire le form", http.StatusBadRequest)
		return models.Group{}
	}

	idUserStr := r.FormValue("user_id")
	userId, err := strconv.Atoi(idUserStr)
	if err != nil {
		log.Println("Cannot convert idUser to int on Create group: ", err)
		return models.Group{}
	}

	group := models.Group{
		Name:           r.FormValue("name"),
		Description:    r.FormValue("description"),
		ID_User_Create: userId,
		Creation_Date:  utils.GetCurrentDateTime(),
		Avatar:         avatarGroup,
	}

	Image, _ := utils.Uploader(w, r, 20, "image", "")
	group.Avatar = utils.FormatImgLink(Image)

	return group
}

func checkGroupExists(db *sql.DB, groupName string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM groups WHERE name = ?)", groupName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
