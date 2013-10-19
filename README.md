Janitor
=======

- J is for Javascript.
- Jan is for [Janne](http://www.github.com/janne/).
- Janitor is for modal editing.


Overview
--------
Janitor is a modal editor, just like Vim, but the similarity end there.

The core of Janitor is built in [Go](http://golang.org/) but, just as the web
browser, it is scripted in Javascript. It has a Javascript console built in for
inspecting and affecting the current state and a big part of the code is
written in Javascript, in the form of plugins.

A major principle concept in Janitor is Layers (thank you [Gary
Bernhardt](https://www.destroyallsoftware.com/talks/a-whole-new-world)). A
layer is a modification of the representation of the current file. Some
examples of possible layers are:

- Syntax highlighting
- Line numbers
- Git diff
- Git blame
- Folds
