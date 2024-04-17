# Bubblegum 🫧🍬

Insipired by quick fix lists used in Neovim, Bubblegum is intended to provide a similar experience for use within the BubbleTea environment.

It works by wrapping the main view, replacing the characters at certain areas with the content of the list. This enables the background view to continue updating while the list is shown.

![Demo with Moai app](./bubblegum.gif)

At the moment, it is tested with wrapping `fullscreen` views only, but further improvements will make use of relative sizing based off the child view. An example of how to use it is shown below.

```
func (model Model) View() string {
    // Flag to check for list spawning
	if model.listSpawned {

		// model.list is the Bubblegum struct
        model.list.SetView(model.childModel.View())
        
        // The view has to be set on the list to get the most
        // updated view for rendering.

		return model.list.View()
	}
    
	return model.childModel.View()
}
```

More examples will be available soon.

## TODO List:

1. Make it flexible for child component dimensions instead of fullscreen

1. More testing for bugs