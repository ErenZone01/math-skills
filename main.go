package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		file := os.Args[1]
		textes, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Votre fichier n'existe pas")
			return
		}
		texte := string(textes)
		if len(texte) == 0 {
			return
		}

		texte = strings.Trim(texte, "\n")
		tabTexte := strings.Split(texte, "\n")

		var tabInt []int
		for i := 0; i < len(tabTexte); i++ {
			nombre, err := strconv.Atoi(tabTexte[i])
			if err != nil {
				fmt.Println("fichier incorrect")
				return
			}
			tabInt = append(tabInt, nombre)
		}

		sort.Ints(tabInt)
		averrageF := math.Round(Averrage(tabInt))
		averrage := int(averrageF)
		median := Median(tabInt)

		varianceF := math.Round(Variance(tabInt))
		variance := int(varianceF)
		standardDevF := math.Round(StandardDev(Variance(tabInt)))
		standardDev := int(standardDevF)

		fmt.Print("Average: ")
		fmt.Println(averrage)
		fmt.Print("Median: ")
		fmt.Println(median)
		fmt.Print("Variance: ")
		fmt.Println(variance)
		fmt.Print("Standard Deviation: ")
		fmt.Println(standardDev)
	}

}

func Averrage(tabInt []int) float64 {
	var somme int
	for i := 0; i < len(tabInt); i++ {
		somme += tabInt[i]
	}
	moyenne := float64(somme) / float64(len(tabInt))
	return moyenne
}

func Median(tabInt []int) int {
	n := len(tabInt)
	if n%2 == 1 {
		median := tabInt[(n-1)/2]
		return median
	} else {
		milieu1 := tabInt[n/2-1]
		milieu2 := tabInt[n/2]
		median := (float64(milieu1) + float64(milieu2)) / 2.0
		return int(math.Round(median))
	}
}

func Variance(tabInt []int) float64 {
	var moyenne = Averrage(tabInt)
	var diff []float64
	var carre []float64
	for i := 0; i < len(tabInt); i++ {
		reste := float64(tabInt[i]) - moyenne
		diff = append(diff, reste)
	}
	for i := 0; i < len(diff); i++ {
		carre = append(carre, (diff[i] * diff[i]))
	}
	var somme float64
	for i := 0; i < len(carre); i++ {
		somme += carre[i]
	}
	Variance := float64(somme) / float64(len(carre))
	return Variance

}

func StandardDev(variance float64) float64 {
	return math.Sqrt(variance)
}
