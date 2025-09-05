package main
import "fmt"

func main() {
	postsTags := [][]string{
		{"go","backend", "programming"},
		{"git", "go", "tools", "development"},
		{"backend", "api", "rest", "go"},
		{"tools", "git", "workflow"},
		{"go", "concurrency", "backend", "youtube"},
	}
	uniqueTags := collectUniqueTags(postsTags)

	fmt.Println("уникальные теги из всех постов:")
	for tag := range uniqueTags {
		fmt.Println("-", tag)
	}

	fmt.Printf("\n Всего уникальных тегов: %d\n", len(uniqueTags))
}

func collectUniqueTags(posts [][]string) map[string]bool {
	uniqueTags := make(map[string]bool)
	for _, post := range posts {
		for _, tag := range post {
			uniqueTags[tag] =true
		}
	}
	return uniqueTags
}