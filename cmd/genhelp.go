package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// genhelpCmd represents the genhelp command
var genhelpCmd = &cobra.Command{
	Use:   "genhelp",
	Short: "Output help from all the sub commands",
	Long: `EXAMPLE:
  This command will generate markdown for all of the cobra commands ktrouble
  supports.

  > ktrouble genhelp > HELP.md

EXAMPLE:
  This command will generate a wiki compatible file that can be submitted to
  confluence via the REST api.  See the 'scruffy publish' command.

  > ktrouble genhelp --format confluence > HELP.cf
`,
	Run: func(cmd *cobra.Command, args []string) {

		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "markdown":
			dumpToc(RootCmd, true, "")
			fmt.Printf("\n")
			dumpHelp(RootCmd, true, "")
		case "confluence":
			dumpTocConfluence(RootCmd, true, "")
			fmt.Printf("\n")
			dumpHelpConfluence(RootCmd, true, "")
		}
	},
}

// dumpToc creates a markdown based table of contents from the list of commands
func dumpToc(c *cobra.Command, root bool, parent string) {
	fullCommand := c.Use
	fullLink := c.Use
	if len(parent) > 0 {
		fullCommand = fmt.Sprintf("%s %s", parent, c.Use)
		fullLink = fmt.Sprintf("%s-%s", strings.ReplaceAll(parent, " ", "-"), c.Use)
	}
	if root {
		fmt.Printf("# %s help for all commands\n\n", c.Use)
		fmt.Printf("## TOC\n\n")
		fmt.Printf("- [%s](#%s)\n", "_main_", c.Use)
		fullCommand = ""
	} else {
		fmt.Printf("- [%s](#%s)\n", fullCommand, fullLink)
	}
	for _, child := range c.Commands() {
		if child.Hidden || child.Name() == "completion" || child.Name() == "help" || child.Name() == "genhelp" {
			continue
		}
		dumpToc(child, false, fullCommand)
	}
}

// dumpHelp adds a markdown section for each of the commands that have been
// added to the rootCmd.
func dumpHelp(c *cobra.Command, root bool, parent string) {
	fullCommand := c.Use
	if len(parent) > 0 {
		fullCommand = fmt.Sprintf("%s %s", parent, c.Use)
	}
	if root {
		fullCommand = ""
		fmt.Printf("## %s\n\n", c.Use)
	} else {
		fmt.Printf("## %s\n\n", fullCommand)
	}
	fmt.Println("```plaintext")
	c.Help()
	fmt.Printf("```\n\n")
	fmt.Printf("[TOC](#TOC)\n\n")

	for _, child := range c.Commands() {
		if child.Hidden || child.Name() == "completion" || child.Name() == "help" || child.Name() == "genhelp" {
			continue
		}
		dumpHelp(child, false, fullCommand)
	}
}

// dumpToc creates a markdown based table of contents from the list of commands
func dumpTocConfluence(c *cobra.Command, root bool, parent string) {
	fullCommand := c.Use
	fullLink := c.Use
	if len(parent) > 0 {
		fullCommand = fmt.Sprintf("%s %s", parent, c.Use)
		fullLink = fmt.Sprintf("%s-%s", strings.ReplaceAll(parent, " ", "-"), c.Use)
	}
	if root {
		fmt.Printf("h1. %s help for all commands\n\n", c.Use)
		fmt.Printf("h2. TOC\n\n")
		fmt.Printf("* [%s|#%s]\n", "_main_", c.Use)
		fullCommand = ""
	} else {
		fmt.Printf("* [%s|#%s]\n", fullCommand, strings.ReplaceAll(fullLink, "-", ""))
	}
	for _, child := range c.Commands() {
		if child.Hidden || child.Name() == "completion" || child.Name() == "help" || child.Name() == "genhelp" {
			continue
		}
		dumpTocConfluence(child, false, fullCommand)
	}
}

// dumpHelp adds a markdown section for each of the commands that have been
// added to the rootCmd.
func dumpHelpConfluence(c *cobra.Command, root bool, parent string) {
	fullCommand := c.Use
	if len(parent) > 0 {
		fullCommand = fmt.Sprintf("%s %s", parent, c.Use)
	}
	if root {
		fullCommand = ""
		fmt.Printf("h2. %s\n", c.Use)
	} else {
		fmt.Printf("h2. %s\n", fullCommand)
	}
	fmt.Println("{code:language=plaintext}")
	c.Help()
	fmt.Printf("{code}\n\n")
	fmt.Printf("[TOC|#TOC]\n\n")

	for _, child := range c.Commands() {
		if child.Hidden || child.Name() == "completion" || child.Name() == "help" || child.Name() == "genhelp" {
			continue
		}
		dumpHelpConfluence(child, false, fullCommand)
	}
}
func init() {
	RootCmd.AddCommand(genhelpCmd)
	genhelpCmd.Flags().String("format", "markdown", "Specify the format for the doc file: markdown|confluence")
}
