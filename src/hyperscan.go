package gitleaks

import (
	"fmt"

	"github.com/flier/gohs/hyperscan"
)

func newBlockDatabase() (hyperscan.BlockDatabase, error) {
	var patterns []*hyperscan.Pattern
	for _, re := range config.Regexes {
		pattern, err := hyperscan.ParsePattern(re.regex.String())
		if err != nil {
			fmt.Printf("can use %s as regex for hyperscan\n", re.regex.String())
		}
		patterns = append(patterns, pattern)
	}

	// setup hyperscan database
	bdb, err := hyperscan.NewBlockDatabase(patterns...)
	if err != nil {
		return nil, err
	}
	return bdb, nil
}
