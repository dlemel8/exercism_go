package tournament

import (
	"io"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

func Tally(reader io.Reader, writer io.Writer) error {
	teamsMap, err := loadTeamsMap(reader)
	if err != nil {
		return err
	}

	teamsSlice := toSortedSlice(teamsMap)

	writeTable(writer, teamsSlice)

	return nil
}

const tableFormat = "%-30v | %2v | %2v | %2v | %2v | %2v\n"

type team struct {
	name  string
	wins  uint
	draws uint
	loses uint
}

func (t *team) plays() uint {
	return t.wins + t.draws + t.loses
}

func (t *team) points() uint {
	return 3*t.wins + 1*t.draws + 0*t.loses
}

func (t *team) String() string {
	return fmt.Sprintf(tableFormat, t.name, t.plays(), t.wins, t.draws, t.loses, t.points())
}

func loadTeamsMap(reader io.Reader) (map[string]*team, error) {
	res := make(map[string]*team)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}
		first, second, result := parts[0], parts[1], parts[2]
		if err := updateTeamsMap(res, first, second, result); err != nil {
			return nil, err
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func updateTeamsMap(teams map[string]*team, first, second, result string) error {
	verifyTeamExists(teams, first)
	verifyTeamExists(teams, second)
	switch result {
	case "win":
		teams[first].wins += 1
		teams[second].loses += 1
	case "loss":
		teams[first].loses += 1
		teams[second].wins += 1
	case "draw":
		teams[first].draws += 1
		teams[second].draws += 1
	default:
		return fmt.Errorf("unknown result: %s", result)
	}
	return nil
}

func verifyTeamExists(teams map[string]*team, name string) {
	if _, ok := teams[name]; !ok {
		teams[name] = &team{name: name}
	}
}

func toSortedSlice(teams map[string]*team) []team {
	res := make([]team, 0, len(teams))
	for _, team := range teams {
		res = append(res, *team)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].points() > res[j].points() ||
			(res[i].points() == res[j].points() && res[i].name < res[j].name)
	})
	return res
}

func writeTable(writer io.Writer, teams []team) {
	writer.Write([]byte(fmt.Sprintf(tableFormat, "Team", "MP", "W", "D", "L", "P")))
	for _, team := range teams {
		writer.Write([]byte(team.String()))
	}
}
