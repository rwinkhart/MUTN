package cli

func glamourStyle() []byte {
	return []byte(`{
  "document": {
    "block_prefix": "\n",
    "block_suffix": "\n",
    "color": "#ebdbb2",
    "margin": 0
  },
  "block_quote": {
    "indent": 2,
    "indent_token": "│ ",
    "color": "109"
  },
  "paragraph": {},
  "list": {
    "level_indent": 4,
    "block_suffix": ""
  },
  "heading": {
    "block_suffix": "\n",
    "color": "#b16286",
    "bold": true
  },
  "h1": {
    "prefix": " ",
    "suffix": " ",
    "color": "#b16286",
    "background_color": "229",
    "bold": true
  },
  "h2": {
    "prefix": "## "
  },
  "h3": {
    "prefix": "### "
  },
  "h4": {
    "prefix": "#### "
  },
  "h5": {
    "prefix": "##### "
  },
  "h6": {
    "prefix": "###### ",
    "bold": false
  },
  "text": {},
  "strikethrough": {
    "crossed_out": true
  },
  "emph": {
    "italic": true
  },
  "strong": {
    "bold": true
  },
  "hr": {
    "color": "246",
    "format": "\n--------\n"
  },
  "item": {
    "block_prefix": "• "
  },
  "enumeration": {
    "block_prefix": ". "
  },
  "task": {
    "ticked": "[✓] ",
    "unticked": "[ ] "
  },
  "link": {
    "color": "108",
    "underline": true
  },
  "link_text": {
    "color": "72",
    "bold": true
  },
  "image": {
    "color": "208",
    "underline": true
  },
  "image_text": {
    "color": "246",
    "format": "Image: {{.text}} →"
  },
  "code": {
    "prefix": " ",
    "suffix": " ",
    "color": "203",
    "background_color": "236"
  },
  "code_block": {
    "color": "248",
    "margin": 4,
    "chroma": {
      "text": {
        "color": "#C4C4C4"
      },
      "error": {
        "color": "#F1F1F1",
        "background_color": "#F05B5B"
      },
      "comment": {
        "color": "#676767"
      },
      "comment_preproc": {
        "color": "#FF875F"
      },
      "keyword": {
        "color": "#00AAFF"
      },
      "keyword_reserved": {
        "color": "#FF5FD2"
      },
      "keyword_namespace": {
        "color": "#FF5F87"
      },
      "keyword_type": {
        "color": "#6E6ED8"
      },
      "operator": {
        "color": "#EF8080"
      },
      "punctuation": {
        "color": "#E8E8A8"
      },
      "name": {
        "color": "#C4C4C4"
      },
      "name_builtin": {
        "color": "#FF8EC7"
      },
      "name_tag": {
        "color": "#B083EA"
      },
      "name_attribute": {
        "color": "#7A7AE6"
      },
      "name_class": {
        "color": "#F1F1F1",
        "underline": true,
        "bold": true
      },
      "name_constant": {},
      "name_decorator": {
        "color": "#FFFF87"
      },
      "name_exception": {},
      "name_function": {
        "color": "#00D787"
      },
      "name_other": {},
      "literal": {},
      "literal_number": {
        "color": "#6EEFC0"
      },
      "literal_date": {},
      "literal_string": {
        "color": "#C69669"
      },
      "literal_string_escape": {
        "color": "#AFFFD7"
      },
      "generic_deleted": {
        "color": "#FD5B5B"
      },
      "generic_emph": {
        "italic": true
      },
      "generic_inserted": {
        "color": "#00D787"
      },
      "generic_strong": {
        "bold": true
      },
      "generic_subheading": {
        "color": "#777777"
      },
      "background": {
        "background_color": "#373737"
      }
    }
  },
  "table": {
    "center_separator": "┼",
    "column_separator": "│",
    "row_separator": "─"
  },
  "definition_list": {},
  "definition_term": {},
  "definition_description": {
    "block_prefix": "\n🠶 "
  },
  "html_block": {},
  "html_span": {}
}
`)
}