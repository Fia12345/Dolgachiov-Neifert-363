package main
import "fmt"

type LogEntry struct {
	IP string
	HTTPcode int
	Time string
}

	func filterErrorsLogs(entries []LogEntry) []LogEntry {
		var errorEntries []LogEntry
		for _, entry := range entries {
			if (entry.HTTPcode >= 400 && entry.HTTPcode <500) ||(entry.HTTPcode >= 500 && entry.HTTPcode < 600) {
				errorEntries = append(errorEntries, entry)
			}
		}
		return errorEntries
	}

	func printLogEntries(entries []LogEntry) {
		for i, entry := range entries {
			fmt.Printf("%d. IP: %s, Код: %d, Время: %s\n", i + 1, entry.IP, entry.HTTPcode, entry.Time)
		}
	}

	func main() {
		logEntries := []LogEntry{
			{"133.788.0.0", 200, "2025-01-12 10:12:32"},
			{"420.420.4.2", 420, "2025-11-14 10:52:52"},
			{"194.119.4.5", 523, "2025-02-13 12:13:12"},
			{"202.202.2.4", 301, "2025-02-24 04:12:32"},
		}
		fmt.Println("Все логи:")
		printLogEntries(logEntries)
		fmt.Println()

		errorLogs := filterErrorsLogs(logEntries)

		fmt.Println("Все записи с ошибками 4xx и 5xx:")
		printLogEntries(errorLogs)
	}
