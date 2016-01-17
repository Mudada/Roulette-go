package main
import (
	"fmt"
	"time"
	"math/rand"
)

func main () {
	total := 200;
	fmt.Println("Welcome to lets'go, this time the program will be a Roulette, You will get 200$ a first and can gamble any $ you want on a numbers")
	fmt.Println("If the number is selected by the Roulette you will win 3x the amout of your gamble, else you loose all of it")
	for {
		gam, num := gamble(total)
		if gam == -1 {
			fmt.Println("Aurevoir et a bientot ^^")
			break
		}
		total = total-gam
		fmt.Printf("Your total is now, %d\n", total)
		fmt.Println("Lets start the Roulette !!!!")
		var rand int
		rand = getrand()
		if rand == num {
			fmt.Println("You Win  !!!!")
			total = total + gam*3
			fmt.Printf("Your total is %d$\n", total)
	}else{
		fmt.Println("You loose,  :c")
		fmt.Printf("Your total is %d$\n", total)
	}
	}
}

func gamble(total int)(gam, num int){
	start :
	fmt.Println("How many you want to gamble and on wich numbers (You can also tap '-1' for the gamble to exit the game)? ('gamble numbers')")
	fmt.Scanln(&gam, &num)
	if gam == -1{
		gam = -1
		num = -1
		return
	}
	fmt.Printf("You Gambled %d$ on the %d\n",gam,num)
	if total-gam < 0 {
		fmt.Println("You dont have this much Credit")
		goto start
	}
	return
}

func getrand()(random int){
	rand.Seed(time.Now().UnixNano())
	random = rand.Intn(5)
	fmt.Printf("The Roulette give %d\n",random)
	return
}
