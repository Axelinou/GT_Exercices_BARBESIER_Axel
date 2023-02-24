package main

import "fmt"

func main() {

	var allumettes = 0
	var Choix_Remove1 int
	var Choix_Remove2 int

	fmt.Print("Bonjour, combien d'allumettes voulais vous ? : ")
	fmt.Scanln(&allumettes)

	for allumettes > 0 {

		fmt.Println("Il reste", allumettes, "allumettes")
		fmt.Println("Player1, combien d'allumettes voulais vous retirer ? (1,2 ou 3)")
		fmt.Scanln(&Choix_Remove1)

		if Choix_Remove1 < 1 || Choix_Remove1 > 3 || Choix_Remove1 > allumettes {
			fmt.Println("Votre choix est invalide, veuillez entrer 1,2 ou 3")
			continue
		}

		allumettes -= Choix_Remove1

		if allumettes == 0 {
			fmt.Println("Felicitation Player1 vous avez perdu")
			break
		}

		fmt.Println("Player2, combien d'allumettes voulais vous retirer ? (1,2 ou 3)")
		fmt.Scanln(&Choix_Remove2)

		if Choix_Remove2 < 1 || Choix_Remove2 > 3 || Choix_Remove2 > allumettes {
			fmt.Println("Votre choix est invalide, veuillez entrer 1,2 ou 3")
			continue
		}

		allumettes -= Choix_Remove2

		if allumettes == 0 {
			fmt.Println("Felicitation Player2 vous avez perdu")
			break
		}

	}

}
