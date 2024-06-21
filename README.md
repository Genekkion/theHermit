# The Hermit 🐚

Insipired by quick fix lists used in Neovim, The Hermit is intended to provide a similar experience for use within the BubbleTea environment.

It works by wrapping the main view, replacing the characters at certain areas with the content of the list. This enables the background view to continue updating while the list is shown.

![Demo with Moai app](./theHermit.gif)

To use this module, download it using

```
go get github.com/genekkion/theHermit/list
```

And import it into your code as such

```
import listy "github.com/genekkion/theHermit/list"
```

You may want to import it under a different name than `list` as Bubble uses `list` as its package name as well.


At the moment, it is tested with wrapping `fullscreen` views only, but further improvements will make use of relative sizing based off the child view. An example of how to use it is shown below.

```Go
func (model Model) View() string {
    // Always set the child view of the list model before returning the view
    // The View() function will automatically render the list or not depending
    // on the boolean flag isShown.
	model.list.SetView("<Insert content here!>")
	return model.list.View()
}
```

Examples are available in the examples folder. Clone this repository, and run `go build` before running the binaries to try them out.

Feel free to raise issues or pull requests if you would like to contribute towards this project, thanks!

## Coming Soon:

1. Fuzzy finder plugin


## TODO List:

1. Make it flexible for child component dimensions instead of fullscreen

1. More testing for bugs

