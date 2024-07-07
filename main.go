package main

import (
	"fmt"
	"math/rand"
)

type StatePemain struct {
	Poin                int
	ListDadu            []int // list dadu yang didapatkan pemain di percobaan sebelumnya
	DaduCounter         int
	TempListDadu        []int
	TempExtractedNumber []int
	IsEmpty             bool
}

func main() {
	startGame(5, 8)
}

func startGame(numPemain int, numDadu int) {

	var isContinue bool = true

	var listPemain []*StatePemain

	for i := 0; i <= numPemain-1; i++ {
		var state StatePemain
		state.Poin = 0
		state.DaduCounter = numDadu

		listPemain = append(listPemain, &state)
	}

	var counterPercobaan = 0

	for isContinue {
		counterPercobaan++
		for _, pemain := range listPemain {
			var newListDadu []int
			if pemain.IsEmpty == true {
				continue
			}
			for j := 1; j <= pemain.DaduCounter; j++ {
				num := rollDadu()
				newListDadu = append(newListDadu, num)
			}

			pemain.ListDadu = newListDadu
		}
		fmt.Printf("Percobaan %d \n", counterPercobaan)
		for idx, each := range listPemain {
			fmt.Printf("Pemain #%d (%d) : %v \n", idx+1, each.Poin, each.ListDadu)
			//cek angka 6
			newListWithout6, poin := extractNumber6(each.ListDadu)
			each.ListDadu = newListWithout6
			each.Poin += poin

			// cek dadu dengan angka 1
			newListWithout1, extracted := extractNumber1(each.ListDadu)
			each.TempListDadu = newListWithout1
			each.TempExtractedNumber = extracted

		}

		for idx, each := range listPemain {
			if idx == len(listPemain)-1 {
				listPemain[0].ListDadu = append(listPemain[0].TempListDadu, each.TempExtractedNumber...)
				listPemain[0].DaduCounter = len(listPemain[0].ListDadu)
			} else {
				listPemain[idx+1].ListDadu = append(listPemain[idx+1].TempListDadu, each.TempExtractedNumber...)
				listPemain[idx+1].DaduCounter = len(listPemain[idx+1].ListDadu)
			}
		}

		fmt.Printf("Setelah Evaluasi \n")
		for idx, each := range listPemain {
			fmt.Printf("Pemain #%d (%d) : %v \n", idx+1, each.Poin, each.ListDadu)
			if len(each.ListDadu) == 0 {
				each.IsEmpty = true
			} else {
				each.IsEmpty = false
			}
		}
		fmt.Println()

		checkPemainDaduHabis := countDaduHabis(listPemain)
		if checkPemainDaduHabis >= len(listPemain)-1 {
			isContinue = false
		}

		if isContinue == false {
			var idxPemenangTemp int = 0
			for idx, each := range listPemain {
				if idx == 0 {
					continue
				}

				if each.Poin > listPemain[idxPemenangTemp].Poin {
					idxPemenangTemp = idx
				}
			}
			fmt.Printf("Pemenang => Pemain %d \n", idxPemenangTemp+1)
		}
	}
}

func countDaduHabis(lists []*StatePemain) int {
	var counter = 0

	for _, each := range lists {
		if each.IsEmpty {
			counter++
		}
	}

	return counter
}

func extractNumber1(params []int) (newList []int, extracted []int) {
	for _, each := range params {
		if each == 1 {
			extracted = append(extracted, each)
		} else {
			newList = append(newList, each)
		}
	}

	return newList, extracted
}

func extractNumber6(params []int) (newList []int, poin int) {
	poin = 0

	for _, each := range params {
		if each == 6 {
			poin++
		} else {
			newList = append(newList, each)
		}
	}

	return newList, poin
}

func rollDadu() int {
	return rand.Intn(6) + 1 // rand.Intn(6) menghasilkan angka dari 0 hingga 5, jadi tambahkan 1
}
