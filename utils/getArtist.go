package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Artists struct {
    Artists []string `json:"artists"`
}

func GetArtist() Artists{
	file, err := os.Open("./Data/artists.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return Artists{}
	}
	defer file.Close()

	// Lire le contenu du fichier
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return Artists{}
	}

	// Déclarer une variable pour stocker les artistes
	var artists Artists

	// Décoder le JSON
	err = json.Unmarshal(byteValue, &artists)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON:", err)
		return artists
	}
	return artists
}
