package main

import ( . "fmt"; "os"; "bufio"; "strings" )

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mr := make(map[string][]int)
	docs := make([]string, 0)

	l := 0

	/* Scandiamo i documenti (le righe del file di input) e
	   indicizziamo le parole. */
	for scanner.Scan() {
		doc := scanner.Text()
		docs = append(docs, doc)
		a := strings.Split(strings.ToLower(doc), " ")
		for _, s := range a {
			if (len(s)) == 0 { continue }
			if lmr := len(mr[s]); lmr == 0 || mr[s][lmr -1] != l {
				mr[s] = append(mr[s], l)
			}
		}
		l++
	}

	keyword := strings.ToLower(os.Args[1])

	for _, l := range mr[keyword] {
		doc := docs[l]
		Print(l, "\t→\t")
		// Stampiamo il documento con la keyword evidenziata
		for _, t := range strings.Split(doc, " ") {
			if keyword == strings.ToLower(t) {
				Print("\033[31;1;4m", t, "\033[0m")
			} else {
				Print(t)
			}
			Print(" ")
		}

		Println()
	}
}
