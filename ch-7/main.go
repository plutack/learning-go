// You have been asked to manage a basketball league and are write a program to help you.
// Define two types. The first one, called Team has a field for the name of the team and a
// field has a field for the player names. The second type is called League and has a field
// called Teams for the teams in the league and field called Wins that maps a team's name
// to its number of Wins.

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []string
	Wins  map[string]int
}

// Add two methods to League. The first method is called MatchResult. It takes four parameters:
// the name of first team, its score in the game, the name of the second team, ans its score in
// the game. This method should update the Wins field in League. Add a second method to League
// called Ranking that returns a slice of the team names in order of wins. Build your data
// structures and called these methods from the main function in your program using some sample
// data.

func (l *League) MatchResult(home string, homeScore int, away string, awayScore int) {
	// add checks to see if home and away teams are members of the league
	if _, ok := l.Wins[home]; !ok {
		panic("home team not a member of league")
	}
	if _, ok := l.Wins[away]; !ok {
		panic("away team not a member of league")
	}
	// home wins
	if homeScore > awayScore {
		l.Wins[home]++
	}
	// away wins
	if homeScore < awayScore {
		l.Wins[away]++
	}
}

func (l *League) Ranking() []string {
	keys := l.Teams
	sort.SliceStable(keys, func(i, j int) bool {
		return l.Wins[keys[i]] > l.Wins[keys[j]]
	})
	return keys
}

func NewLeague(teams ...string) *League {
	sort.Slice(teams, func(i, j int) bool {
		return teams[i] < teams[j]
	})
	return &League{
		Teams: teams,
		Wins: func(teams []string) map[string]int {
			wins := make(map[string]int)
			for _, team := range teams {
				wins[team] = 0
			}
			return wins
		}(teams),
	}
}

// Define an interface called Ranker that has a single method called Ranking that returns a slice of
// strings. Write a function called RankPrinter with two Parameters, the first of type Ranker and the
// second of type io.Writer. Use the io.WriteString function to write the values returned by Ranker to
// the io.Writer, with a new line seperating each result. Call this function from main.

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	teams := r.Ranking()
	for _, team := range teams {
		io.WriteString(w, team+"\n")
	}
}

func main() {
	league1 := NewLeague("Arsenal", "Manchester City", "Tottenham", "Chelsea")
	fmt.Println(league1)

	fmt.Println(league1.Ranking())

	league1.MatchResult("Arsenal", 2, "Tottenham", 7)
	fmt.Println(league1.Ranking())
	fmt.Println(league1)

	file, err := os.Create("./ranking.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	RankPrinter(league1, file)
}
