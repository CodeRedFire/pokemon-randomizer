package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define structs to match the JSON structure from PokeAPI

// SpeciesName represents a localized name for the Pokémon species
type SpeciesName struct {
	Name     string `json:"name"`
	Language struct {
		Name string `json:"name"`
	} `json:"language"`
}

// PokemonSpeciesResponse represents the response from the pokemon-species API
type PokemonSpeciesResponse struct {
	Names []SpeciesName `json:"names"`
}

type PokemonTypeResponse struct {
	Names   []SpeciesName `json:"names"`
	Sprites struct {
		Gen9 struct {
			SV struct {
				Name string `json:"name_icon"`
			} `json:"scarlet-violet"`
		} `json:"generation-ix"`
	} `json:"sprites"`
	Name string `json:"name"`
}

type PokemonAbilityResponse struct {
	Names []SpeciesName `json:"names"`
	Name  string        `json:"name"`
}

// PokemonType represents a Pokémon's type
type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

// PokemonStat represents a Pokémon's base stat
type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

// PokemonSprites represents the Pokémon's sprites
type PokemonSprites struct {
	FrontDefault string `json:"front_default"`
}

// PokemonResponse represents the response from the pokemon API
type PokemonResponse struct {
	Types   []PokemonType  `json:"types"`
	Stats   []PokemonStat  `json:"stats"`
	Sprites PokemonSprites `json:"sprites"`
}

// Final structure for the combined Pokémon information
type PokemonInfo struct {
	Names          []SpeciesName            `json:"names"` // Keep all names
	Types          []string                 `json:"types"`
	Stats          []map[string]interface{} `json:"stats"`
	FrontSpriteURL string                   `json:"front_sprite_url"`
}

type TypeInfo struct {
	Names          []SpeciesName `json:"names"` // Keep all names
	FrontSpriteURL string        `json:"front_sprite_url"`
}

// downloadImage télécharge une image depuis une URL et l'enregistre dans un fichier
func downloadImage(url string, filepath string) error {
	// Effectuer une requête HTTP GET pour l'image
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erreur lors de la requête HTTP pour l'image : %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erreur HTTP lors du téléchargement de l'image : %d %s", resp.StatusCode, resp.Status)
	}

	// Lire le corps de la réponse
	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture des données de l'image : %w", err)
	}

	// Écrire les octets de l'image dans un fichier
	err = ioutil.WriteFile(filepath, imageBytes, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier image : %w", err)
	}

	return nil
}

func getPkmnData(index int) {
	fmt.Println(fmt.Sprintf("-> %d", index))
	pokemonSpeciesURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-species/%d/", index)
	pokemonDataURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", index)

	pokemonInfo := PokemonInfo{}

	// 1. Fetch "names" from pokemon-species/1/
	speciesResp, err := http.Get(pokemonSpeciesURL)
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données de l'espèce Pokémon : %v\n", err)
		return
	}
	defer speciesResp.Body.Close()

	if speciesResp.StatusCode != http.StatusOK {
		fmt.Printf("Erreur HTTP pour l'espèce Pokémon : %d %s\n", speciesResp.StatusCode, speciesResp.Status)
		return
	}

	speciesBody, err := ioutil.ReadAll(speciesResp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse de l'espèce Pokémon : %v\n", err)
		return
	}

	var speciesData PokemonSpeciesResponse
	err = json.Unmarshal(speciesBody, &speciesData)
	if err != nil {
		fmt.Printf("Erreur lors de l'analyse JSON de l'espèce Pokémon : %v\n", err)
		return
	}
	pokemonInfo.Names = speciesData.Names // Assign all names

	// 2. Fetch "types", "stats", "sprites.front_default" from pokemon/1/
	pokemonResp, err := http.Get(pokemonDataURL)
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données du Pokémon : %v\n", err)
		return
	}
	defer pokemonResp.Body.Close()

	if pokemonResp.StatusCode != http.StatusOK {
		fmt.Printf("Erreur HTTP pour les données du Pokémon : %d %s\n", pokemonResp.StatusCode, pokemonResp.Status)
		return
	}

	pokemonBody, err := ioutil.ReadAll(pokemonResp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse du Pokémon : %v\n", err)
		return
	}

	var pokemonData PokemonResponse
	err = json.Unmarshal(pokemonBody, &pokemonData)
	if err != nil {
		fmt.Printf("Erreur lors de l'analyse JSON du Pokémon : %v\n", err)
		return
	}

	// Extract types
	types := make([]string, len(pokemonData.Types))
	for i, t := range pokemonData.Types {
		types[i] = t.Type.Name
	}
	pokemonInfo.Types = types

	// Extract stats
	stats := make([]map[string]interface{}, len(pokemonData.Stats))
	for i, s := range pokemonData.Stats {
		stats[i] = map[string]interface{}{
			"name":  s.Stat.Name,
			"value": s.BaseStat,
		}
	}
	pokemonInfo.Stats = stats

	// Extract front sprite URL
	pokemonInfo.FrontSpriteURL = pokemonData.Sprites.FrontDefault

	// Télécharger l'image du sprite si l'URL est disponible
	if pokemonInfo.FrontSpriteURL != "" {
		spriteFilename := fmt.Sprintf("sprite_%d.png", index) // Nom de fichier pour le sprite
		err := downloadImage(pokemonInfo.FrontSpriteURL, spriteFilename)
		if err != nil {
			fmt.Printf("Erreur lors du téléchargement de l'image du sprite : %v\n", err)
		} else {
			fmt.Printf("L'image du sprite a été sauvegardée dans '%s'\n", spriteFilename)
		}
	}

	// 3. Group all information into a JSON file
	outputFilename := fmt.Sprintf("data_%d.json", index)
	outputJSON, err := json.MarshalIndent(pokemonInfo, "", "    ")
	if err != nil {
		fmt.Printf("Erreur lors de la sérialisation JSON : %v\n", err)
		return
	}

	err = ioutil.WriteFile(outputFilename, outputJSON, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier JSON : %v\n", err)
		return
	}

	fmt.Printf("Les informations du Pokémon ont été sauvegardées dans '%s'\n", outputFilename)
}

func getTypes(index int) {
	fmt.Println(fmt.Sprintf("-> %d", index))
	typeURL := fmt.Sprintf("https://pokeapi.co/api/v2/type/%d/", index)

	typeInfo := TypeInfo{}

	// 1. Fetch "names" from type/1/
	speciesResp, err := http.Get(typeURL)
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données du type : %v\n", err)
		return
	}
	defer speciesResp.Body.Close()

	if speciesResp.StatusCode != http.StatusOK {
		fmt.Printf("Erreur HTTP pour le type : %d %s\n", speciesResp.StatusCode, speciesResp.Status)
		return
	}

	speciesBody, err := ioutil.ReadAll(speciesResp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse du type : %v\n", err)
		return
	}

	var typeData PokemonTypeResponse
	err = json.Unmarshal(speciesBody, &typeData)
	if err != nil {
		fmt.Printf("Erreur lors de l'analyse JSON du type : %v\n", err)
		return
	}
	typeInfo.Names = typeData.Names // Assign all names

	// Extract front sprite URL
	typeInfo.FrontSpriteURL = typeData.Sprites.Gen9.SV.Name

	// Télécharger l'image du sprite si l'URL est disponible
	if typeInfo.FrontSpriteURL != "" {
		spriteFilename := fmt.Sprintf("%s.png", typeData.Name) // Nom de fichier pour le sprite
		err := downloadImage(typeInfo.FrontSpriteURL, spriteFilename)
		if err != nil {
			fmt.Printf("Erreur lors du téléchargement de l'image du sprite : %v\n", err)
		} else {
			fmt.Printf("L'image du sprite a été sauvegardée dans '%s'\n", spriteFilename)
		}
	}

	// 3. Group all information into a JSON file
	outputFilename := fmt.Sprintf("%s.json", typeData.Name)
	outputJSON, err := json.MarshalIndent(typeInfo, "", "    ")
	if err != nil {
		fmt.Printf("Erreur lors de la sérialisation JSON : %v\n", err)
		return
	}

	err = ioutil.WriteFile(outputFilename, outputJSON, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier JSON : %v\n", err)
		return
	}

	fmt.Printf("Les informations du Pokémon ont été sauvegardées dans '%s'\n", outputFilename)
}
func getAbilities(index int) {
	fmt.Println(fmt.Sprintf("-> %d", index))
	abilityURL := fmt.Sprintf("https://pokeapi.co/api/v2/ability/%d/", index)

	typeInfo := TypeInfo{}

	// 1. Fetch "names" from ability/1/
	abilityResp, err := http.Get(abilityURL)
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données de la capacité : %v\n", err)
		return
	}
	defer abilityResp.Body.Close()

	if abilityResp.StatusCode != http.StatusOK {
		fmt.Printf("Erreur HTTP pour la capacité : %d %s\n", abilityResp.StatusCode, abilityResp.Status)
		return
	}

	abilityBody, err := ioutil.ReadAll(abilityResp.Body)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture de la réponse de la capacité : %v\n", err)
		return
	}

	var abilityData PokemonAbilityResponse
	err = json.Unmarshal(abilityBody, &abilityData)
	if err != nil {
		fmt.Printf("Erreur lors de l'analyse JSON de la capacité : %v\n", err)
		return
	}
	typeInfo.Names = abilityData.Names // Assign all names

	// 3. Group all information into a JSON file
	outputFilename := fmt.Sprintf("%s.json", abilityData.Name)
	outputJSON, err := json.MarshalIndent(typeInfo, "", "    ")
	if err != nil {
		fmt.Printf("Erreur lors de la sérialisation JSON : %v\n", err)
		return
	}

	err = ioutil.WriteFile(outputFilename, outputJSON, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture du fichier JSON : %v\n", err)
		return
	}

	fmt.Printf("Les informations du Pokémon ont été sauvegardées dans '%s'\n", outputFilename)
}

func main() {
	//for i := 1; i <= 19; i++ {
	//	getPkmnData(i)
	//}
	//for i := 1; i <= 19; i++ {
	//	getTypes(i)
	//}
	for i := 1; i <= 367; i++ {
		getAbilities(i)
	}
}
