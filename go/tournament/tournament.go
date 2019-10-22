package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// TeamRecord tracks matches played/won/drawn/lost, as well as the points earned.
type TeamRecord struct {
	name          string
	matchesPlayed int
	matchesWon    int
	matchesDrawn  int
	matchesLost   int
	points        int
}

// Tally takes an input of result rows and outputs a formatted table of results.
func Tally(reader io.Reader, writer io.Writer) error {
	recordsMap := make(map[string]TeamRecord, 5)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 || text[0] == '#' {
			continue
		}

		values := strings.Split(text, ";")
		if len(values) != 3 {
			return errors.New("Invalid entry; must contain two teams and one result")
		}

		recordA := recordsMap[values[0]]
		recordB := recordsMap[values[1]]

		switch values[2] {
		case "win":
			recordWin(&recordA, &recordB)
		case "loss":
			recordWin(&recordB, &recordA)
		case "draw":
			recordA.matchesPlayed++
			recordA.matchesDrawn++
			recordA.points++

			recordB.matchesPlayed++
			recordB.matchesDrawn++
			recordB.points++
		default:
			return errors.New("Invalid result")
		}

		recordsMap[values[0]] = recordA
		recordsMap[values[1]] = recordB
	}

	records := make([]TeamRecord, 0, len(recordsMap))
	for name, record := range recordsMap {
		record.name = name // for sortability by name when all we have is the slice of records
		records = append(records, record)
	}

	sort.Slice(records, func(i, j int) bool {
		if records[i].points == records[j].points {
			return records[i].name < records[j].name
		}
		return records[i].points > records[j].points
	})

	fmt.Fprintln(writer, "Team                           | MP |  W |  D |  L |  P")
	for _, r := range records {
		fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n", r.name, r.matchesPlayed,
			r.matchesWon, r.matchesDrawn, r.matchesLost, r.points)
	}
	return nil
}

func recordWin(recordA, recordB *TeamRecord) {
	recordA.matchesPlayed++
	recordA.matchesWon++
	recordA.points += 3

	recordB.matchesPlayed++
	recordB.matchesLost++
}
