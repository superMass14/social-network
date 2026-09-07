package handler

import (
	"backend/models"
	"backend/app/cors"
	"backend/app/service"
	"backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign up handler")
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		formatedUser := models.FormatedUser{
			First_name: r.FormValue("firstname"),
			Last_name:  r.FormValue("lastname"),
			User_name:  r.FormValue("username"),
			Birth_date: r.FormValue("dateofbirth"),
			About_me:   r.FormValue("aboutme"),
			Password:   r.FormValue("password"),
			Email:      r.FormValue("email"),
		}
		log.Printf("in register...")
		log.Printf("[debug]  firstName: %v\n", formatedUser.First_name)
		log.Printf("[debug]  lastName: %v\n", formatedUser.Last_name)
		log.Printf("[debug]  userName: %v\n", formatedUser.User_name)
		log.Printf("[debug]  birth date: %v\n", formatedUser.Birth_date)
		log.Printf("[debug]  about me: %v\n", formatedUser.About_me)
		log.Printf("[debug] password: %v\n", formatedUser.Password)
		log.Printf("[debug]  email: %v\n", formatedUser.Email)
		log.Println("checking img...")
		Avatar, errAvatar := utils.Uploader(w, r, 2, "avatar", "avatar_"+formatedUser.First_name)
		if errAvatar != nil {
			log.Printf("Error uploading avatar: %v\n", errAvatar)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": errAvatar.Error()})
			return
		}
		log.Printf("[debug]  avatar: %v\n", Avatar)
		formatedUser.Avatar = utils.FormatImgLink(Avatar)
		if err := utils.IsValidEmail(strings.TrimSpace(formatedUser.Email)); err != nil {
			log.Println("Invalid email", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid email" + err.Error()})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(formatedUser.First_name)); err != nil {
			log.Println("Invalid first name", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid first name" + err.Error()})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(formatedUser.Last_name)); err != nil {
			log.Println("Invalid last name", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid last name" + err.Error()})
			return
		} else if err := utils.VerifyUsername(formatedUser.User_name); err != nil {
			log.Println("Invalid username", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid username:" + err.Error()})
			return
		} else if err := utils.VerifyAboutMe(formatedUser.About_me); err != nil {
			log.Println("Invalid about me", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid about me" + err.Error()})
			return
		}
		formatedUser.User_type = "public"
		err := service.AuthServ.CreateUser(formatedUser)
		if err != nil {
			log.Println("Error creating user", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			return
		}
		user, err := service.AuthServ.UserRepo.GetUserByEmail(formatedUser.Email)
		if err != nil {
			log.Println("Handler Error getting user", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Error getting user" + err.Error()})
			return
		}
		session, err := service.AuthServ.CreateSession(user)
		if err != nil {
			log.Println("Error creating session", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{"error": "Error creating session" + err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		log.Printf("User %s created successfully", user.Email)
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": session.Token, "user": user, "error": "ok"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Method not allowed" + r.Method})
	}
	// ...
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		credentials := make(map[string]string, 2)
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Invalid credentials 0 :", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		err = json.Unmarshal(content, &credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 0"})
			return
		}
		email, password := credentials["email"], credentials["password"]

		if utils.IsValidEmail(strings.TrimSpace(email)) != nil || utils.IsValidPassword(strings.TrimSpace(password)) != nil {
			log.Println("Invalid credentials 1", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 1"})
			return
		}
		user, err := service.AuthServ.CheckCredentials(email, password)
		if err != nil {
			log.Println("Invalid credentials 2", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 2"})
			return
		}

		session, err := service.AuthServ.CreateSession(user)
		if err != nil {
			log.Println("Error creating session", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		// fmt.Println("User Connected", user.Email)
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": session.Token, "user": user})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SignOutHandler")

	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		token := r.Header.Get("Authorization")
		if strings.TrimSpace(token) == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}

		err := service.AuthServ.RemoveSession(token)
		if err != nil {
			log.Println("Error removing session", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"message": "success"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}

}

func VerifySessionHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Verify session handler")
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "success", "session": session})

}
