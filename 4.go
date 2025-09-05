package main
import "fmt"

func countVotes(votes []string, candidates []string) {
	voteCount := make(map[string]int)
	totalVotes := len(votes)

	for _, candidate := range candidates {
		voteCount[candidate] = 0
	}

	for _, vote := range votes {
		voteCount[vote]++
	}

	fmt.Println("Результаты голосования: ")
	for _, candidate := range candidates {
		count := voteCount[candidate]
		percentage :=float64(count) / float64(totalVotes) * 100
		fmt.Printf("%s: %d Голосов (%.1f%%)\n", candidate, count, percentage)
	}
	fmt.Printf("\nВсего голосов: %d\n", totalVotes)
}

func main() {
	candidates := []string{"Анна", "Борис", "Виктор"}
	voters := map[string]string{
		"Пётр": "Анна",
		"Ваня": "Борис",
		"Артём": "Виктор",
		"Оля": "Виктор",
	}

	var votes []string
	for _, vote := range voters {
		votes = append(votes, vote)
	}

	fmt.Println("Выбор голосующих:")
	for voter, candidate :=range voters {
		fmt.Println(voter, " ---> ", candidate)
	}
	fmt.Println()

	countVotes(votes, candidates)
}
