package main

import "fmt"

func main() {

	//todo if kondisi `boolean`
	point := 0
	if point == 1 {
		fmt.Println("Point bernilai 1")
	}
	if point == 0 {
		fmt.Println("Point bernilai 0")
	}

	//todo if, else if, else
	score := 10

	if score > 5 {
		fmt.Println("score yang baik")
	} else if score < 5 {
		fmt.Println("score yang lumayan")
	} else {
		fmt.Println("anda tidak punya score")
	}

	//todo variabel temporary pada if - else
	points := 1100.0

	if percent := points / 10.0; percent >= 100 {
		fmt.Printf("%.1f %s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f %s good \n", percent, "%")
	} else {
		fmt.Printf("%.1f %s nod bad \n", percent, "%")
	}

	//todo switch - case `value`
	nilai := 6

	switch nilai {
	case 5:
		fmt.Println("nilai anda 5")
	case 6:
		fmt.Println("nilai anda 6")
	case 7:
		fmt.Println("nilai anda 7")
	default:
		fmt.Println("nilai gak ada!")
	}

	//todo case banyak kondisi
	hasil := 0

	switch hasil {
	case 9, 10:
		fmt.Println("perfect")
	case 5, 6, 7, 8:
		fmt.Println("good")
	case 1, 2, 3, 4:
		fmt.Println("nod bad")
	default:
		{
			fmt.Println("so bad bro")
			fmt.Println("can you be better 👍")
		}
	}

	//todo switch style if-else `value and boolean`
	values := 5

	switch {
	case values == 8:
		fmt.Println("perfect")
	case (values < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more ")
		}
	}

	//todo keyword fallhrough di switch `fallthrough seperti do-while`

	goal := true

	switch goal {
	case true:
		fmt.Println("Player selebrasi Goal!")
		fallthrough
	case false:
		fmt.Println("tidak selebrasi!")
	}

	//todo seleksi kondisi bersarang
	// if -else - switch
	scores := 10

	if scores > 7 {
		switch scores {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if scores == 5 {
			fmt.Println("nod bad")
		} else if scores == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if scores == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
