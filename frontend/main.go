package main

import (
	"frontend/webchat"
	"net/http"
)

func main() {
	http.HandleFunc("/chatui", chatui)
	http.HandleFunc("/profile/show", profileShow)
	http.HandleFunc("/profile/edit", profileEdit)
	http.ListenAndServe(":8080", nil)
}

func chatui(w http.ResponseWriter, r *http.Request) {
	p := webchat.ChatParams{
		User: webchat.User{
			ID:       1,
			Username: "johndoe",
			Email:    "john@example.com",
		},
		Query:  "Enter your Query",
		Answer: "Response from AI Knowledge Worker",
	}
	webchat.Chat(w, p, partial(r))
}

func profileShow(w http.ResponseWriter, r *http.Request) {
	p := webchat.ProfileShowParams{
		Title: "Profile Show",
		ProfileInfo: webchat.ProfileInfo{
			UserID:    1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Phone:     "1234567890",
			Company:   "Acme Inc",
		},
	}
	webchat.ProfileShow(w, p, partial(r))
}

func profileEdit(w http.ResponseWriter, r *http.Request) {
	p := webchat.ProfileEditParams{
		Title: "Profile Edit",
		ProfileInfo: webchat.ProfileInfo{
			UserID:    1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Phone:     "1234567890",
			Company:   "Acme Inc",
		},
	}
	webchat.ProfileEdit(w, p, partial(r))
}

func partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
