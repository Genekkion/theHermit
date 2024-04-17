# Bubblegum 🫧🍬

Insipired by quick fix lists used in Neovim, Bubblegum is intended to provide a similar experience for use within the BubbleTea environment.

It works by wrapping the main view, replacing the characters at certain areas with the content of the list. This enables the background view to continue updating while the list is shown.

![Demo with Moai app](./bubblegum.gif)

At the moment, it is tested with wrapping `fullscreen` views only, but further improvements will make use of relative sizing based off the child view.