## Known Bugs - MUTN CLI
- Various Markdown issues resulting from the use of [Glamour](https://github.com/charmbracelet/glamour)
  - Note that MUTN is currently using the older Glamour v0.7.0, as v0.8.0 introduced a line-wrapping regression
  - [Odd handling of line breaks](https://github.com/charmbracelet/glamour/issues/84)
  - [Broken quotes](https://github.com/charmbracelet/glamour/issues/172) (seems related to line break issue)
  - [Extra blank lines after nested list items](https://github.com/charmbracelet/glamour/issues/102)
  - [Incorrect wrapping of text in lists](https://github.com/charmbracelet/glamour/issues/56)
  - [Excessive binary size and increased program startup time](https://github.com/charmbracelet/glamour/issues/288)
