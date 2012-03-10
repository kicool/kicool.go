package main

import (
	"os"
	"fmt"
	"rand"
	"time"
)

const MAX_GUESS_TIME = 10
const DEBUG = false

type guess struct {
	nmbr [4]int
}

var (
	gen *guess
	user *guess
	times = 0
	ok = false
	hintinfo[MAX_GUESS_TIME]string
)

func generate() *guess {
	rand.Seed(time.Nanoseconds())
	rnd := rand.Perm(10)
	tmp := new(guess)
	for i := range tmp.nmbr {
		tmp.nmbr[i] = rnd[i]
	}
	return tmp
}

func validate(g *guess) bool {
        for i := range g.nmbr {
		for j:= i+1; j < len(g.nmbr); j++ {
			if g.nmbr[i] == g.nmbr[j] {
				return false
			}
		}
	}

	return true
}

func dump(hint string, g *guess) {
	fmt.Print(hint)
        for i := range g.nmbr {
		fmt.Print(g.nmbr[i])
	}
	fmt.Println()
}

func getFromUser() *guess {
	tmp := new(guess)
	for {
		fmt.Printf("\n%d -> ", times)
		n, err := fmt.Scanf("%1d%1d%1d%1d\n", &tmp.nmbr[0], &tmp.nmbr[1], &tmp.nmbr[2], &tmp.nmbr[3])
		if 0 == n {
			fmt.Print("[E]:", err)
		}
		if validate(tmp) {
			break
		} else {
			fmt.Println("invalid input")
		}
		if DEBUG {
			dump("user input:", tmp)
		}
	}
	fmt.Println()
	return tmp
}

func judge(u *guess, g *guess) (bool, string) {
	a, b := 0, 0
        for i := range u.nmbr {
		for j := range g.nmbr {
			if u.nmbr[i] == g.nmbr[j] {
				if i == j {
					a++
				} else {
					b++
				}
			}

		}
	}
	if 4 == a {
		return true, "4A0B"
	}
	return false, fmt.Sprintf("%1dA%1dB", a, b)
}

/*
 * 0) generate 4bit diff String like "1234"
 * 1) get user input, promoting guessed times
 * 2) give guess hint ?A?B, if 4A0B success
 * 3) guessed times ++, goto 1)
 */
func main() {
	gen = generate()
	if !validate(gen) {
		dump("failed validate:", gen)
		os.Exit(1)
	}
	if DEBUG {
		dump("door:", gen)
	}
	//scan user 
	//hintinfo, ok = judge()
	//output hintinfo
	//if ok then Congratulations
	//user = new(guess)
	for !ok && times < MAX_GUESS_TIME {
		user = getFromUser()
		ok, hintinfo[times] = judge(user, gen);
		fmt.Print(hintinfo[times])
		times++
	}
	fmt.Println()
	if ok {
		fmt.Printf("You guessed %d times.Congratulations!\n", times)
	} else {
		fmt.Println("You lose!")
	}

}
