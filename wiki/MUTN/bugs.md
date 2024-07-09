# Known Bugs - MUTN CLI
- Various Markdown rendering issues, including odd handling of line breaks and broken quotes
  - These bugs exist within the upstream [glamour](https://github.com/charmbracelet/glamour) package used for Markdown rendering
  - For line breaks, see [here](https://github.com/charmbracelet/glamour/issues/84)
  - For quotes, see [here](https://github.com/charmbracelet/glamour/issues/172) (seems related to line break issue)
- Some [weird Windows behavior](https://github.com/rwinkhart/MUTN/blob/main/wiki/MUTN/quirks.md)