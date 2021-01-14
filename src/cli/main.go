package cli

import (
	"fmt"
	"github.com/liampm/journal/src/app"
)

func ListFilesWithTag(searchTag string) {
	matchedFiles := app.FindEntryFilesWithTag(searchTag)

	fmt.Printf("%d entry with the tag '%s'\n", len(matchedFiles), searchTag)

	for _, fileName := range matchedFiles {
		fmt.Printf("\t- %s\n", fileName)
	}
}
