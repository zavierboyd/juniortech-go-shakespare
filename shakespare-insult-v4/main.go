package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

// Words from https://github.com/codepo8/shakespeare-insult-generator/blob/master/index.html
// Define our word lists
var adjectives1 = []string{"artless", "bawdy", "beslubbering", "bootless", "churlish", "cockered", "clouted", "craven", "currish", "dankish", "dissembling", "droning", "errant", "fawning", "fobbing", "froward", "frothy", "gleeking", "goatish", "gorbellied", "impertinent", "infectious", "jarring", "loggerheaded", "lumpish", "mammering", "mangled", "mewling", "paunchy", "pribbling", "puking", "puny", "qualling", "rank", "reeky", "roguish", "ruttish", "saucy", "spleeny", "spongy", "surly", "tottering", "unmuzzled", "vain", "venomed", "villainous", "warped", "wayward", "weedy", "yeasty"}
var adjectives2 = []string{"base-court", "bat-fowling", "beef-witted", "beetle-headed", "boil-brained", "clapper-clawed", "clay-brained", "common-kissing", "crook-pated", "dismal-dreaming", "dizzy-eyed", "doghearted", "dread-bolted", "earth-vexing", "elf-skinned", "fat-kidneyed", "fen-sucked", "flap-mouthed", "fly-bitten", "folly-fallen", "fool-born", "full-gorged", "guts-griping", "half-faced", "hasty-witted", "hedge-born", "hell-hated", "idle-headed", "ill-breeding", "ill-nurtured", "knotty-pated", "milk-livered", "motley-minded", "onion-eyed", "plume-plucked", "pottle-deep", "pox-marked", "reeling-ripe", "rough-hewn", "rude-growing", "rump-fed", "shard-borne", "sheep-biting", "spur-galled", "swag-bellied", "tardy-gaited", "tickle-brained", "toad-spotted", "unchin-snouted", "weather-bitten"}
var nouns = []string{"apple-john", "baggage", "boar-pig"}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "--serve" {
			http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
				insult := makeInsult()
				w.Write([]byte(insult))
			})
			fmt.Println("listening on 8080")
			go http.ListenAndServe(":8080", nil)
		}
	}
	defer fmt.Println("Enjoy your insults!")

	fmt.Println("Enter words to use in the insult:")
	fmt.Println("Type ^D or 'exit' to close the program")
	var word string
	for {
		_, err := fmt.Scan(&word)
		if err != nil {
			break
		}
		if word == "exit" {
			break
		}
		if word != "" {
			nouns[0] = word
			fmt.Println(makeInsult())
		}
	}
}

func makeInsult() string {
	return fmt.Sprintf("Thou %s %s %s!",
		adjectives1[rand.Intn(len(adjectives1))],
		adjectives2[rand.Intn(len(adjectives2))],
		nouns[rand.Intn(len(nouns))])
}
