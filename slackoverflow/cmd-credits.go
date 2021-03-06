package slackoverflow

import "github.com/aframevr/slackoverflow/std"

// slackoverflow credits
// Display credits
type cmdCredits struct{}

// Execute
func (a *cmdCredits) Execute(args []string) error {

	// Contibutors list is auto generated by makefile
	if contributors != nil {
		contributorsList := std.NewTable("SlackOverflow contributors")
		for _, contributor := range contributors {
			contributorsList.AddRow(contributor)
		}
		contributorsList.Print()
	}
	return nil
}
