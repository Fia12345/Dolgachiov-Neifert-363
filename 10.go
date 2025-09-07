package main
import ("fmt"; "strings")

type Stats struct {
	CharCount int
	WordCount int
	SentanceCount int
}

func textStats(text string) Stats {
	stats := Stats{}
	stats.CharCount = len(text)
	words := strings.Fields(text)
	stats.WordCount = len(words)

	for _, char := range text {
		switch char {
		case '.', '!', '?':
			stats.SentanceCount++
		}
	}

	return stats
}

func main() {
	text := "бррр-р-р-р-р, это Красти крабс? Нет, это Патрик. уи-и-и."
	stats := textStats(text)

	fmt.Printf("Текст: %s\n", text)
	fmt.Printf("Количество символов: %d\n", stats.CharCount)
	fmt.Printf("Количество слов: %d\n", stats.WordCount)
	fmt.Printf("Количество предложений: %d\n", stats.SentanceCount)

	text2 := "Как тебя зовут? Калолик старший. Как тебя ЗВАЛИ? Максим."
	stats2 := textStats(text2)
	
	fmt.Printf("\nТекст 2: %s\n", text2)
	fmt.Printf("Количество символов: %d\n", stats2.CharCount)
	fmt.Printf("Количество слов: %d\n", stats2.WordCount)
	fmt.Printf("Количество предложений: %d\n", stats2.SentanceCount)
}