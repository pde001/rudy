package main

import (
	"github.com/darkweak/rudy/commands"
	"github.com/spf13/cobra"
)

func main() {
	var root cobra.Command

	commands.Prepare(&root)

	if err := root.Execute(); err != nil {
		panic(err)
	}package main

	import (
		"github.com/darkweak/rudy/commands"
		"github.com/spf13/cobra"
	)
	
	func main() {
		var root cobra.Command
	
		commands.Prepare(&root)
	
		if err := root.Execute(); err != nil {
			panic(err)
		}
	}package main

	import (
		"github.com/darkweak/rudy/commands"
		"github.com/spf13/cobra"
	)
	
	func main() {
		var root cobra.Command
	
		commands.Prepare(&root)
	
		if err := root.Execute(); err != nil {
			panic(err)
		}
	}git clone https://github.com/darkweak/rudy
	cd rudy
	go build rudy.go -o rudy
	rudy [command]git clone https://github.com/darkweak/rudy
	cd rudy
	go build rudy.go -o rudy
	rudy [command]
}rudy run -u https://bagrada.net