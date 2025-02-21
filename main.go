<<<<<<< HEAD
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://kpoppingapi.netlify.app/api"

type ResponseData struct {
	Images  []string    `json:"images,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func fetchAPI(url string, w http.ResponseWriter) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Erreur API: %v", resp.Status), resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture de la réponse", http.StatusInternalServerError)
		return
	}

	var data ResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Erreur de parsing JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getIdolPics(w http.ResponseWriter, r *http.Request) {
	idol := r.URL.Query().Get("idol")
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/idolpics/%s/%s", baseURL, idol, group)
	fetchAPI(url, w)
}

func getGroupImages(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/groupImages/%s", baseURL, group)
	fetchAPI(url, w)
}

func getProfileData(w http.ResponseWriter, r *http.Request) {
	idol := r.URL.Query().Get("idol")
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/profile/%s/%s", baseURL, idol, group)
	fetchAPI(url, w)
}

func main() {
	http.HandleFunc("/idolpics", getIdolPics)
	http.HandleFunc("/groupimages", getGroupImages)
	http.HandleFunc("/profile", getProfileData)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
=======
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const baseURL = "https://kpoppingapi.netlify.app/api"

type ResponseData struct {
	Images  []string    `json:"images,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func fetchAPI(url string, w http.ResponseWriter) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Erreur API: %v", resp.Status), resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erreur de lecture de la réponse", http.StatusInternalServerError)
		return
	}

	var data ResponseData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Erreur de parsing JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getIdolPics(w http.ResponseWriter, r *http.Request) {
	idol := r.URL.Query().Get("idol")
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/idolpics/%s/%s", baseURL, idol, group)
	fetchAPI(url, w)
}

func getGroupImages(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/groupImages/%s", baseURL, group)
	fetchAPI(url, w)
}

func getProfileData(w http.ResponseWriter, r *http.Request) {
	idol := r.URL.Query().Get("idol")
	group := r.URL.Query().Get("group")
	url := fmt.Sprintf("%s/profile/%s/%s", baseURL, idol, group)
	fetchAPI(url, w)
}

func main() {
	http.HandleFunc("/idolpics", getIdolPics)
	http.HandleFunc("/groupimages", getGroupImages)
	http.HandleFunc("/profile", getProfileData)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
>>>>>>> 6e44933b7baa8375d3711fe1c63a9d78ff764883
