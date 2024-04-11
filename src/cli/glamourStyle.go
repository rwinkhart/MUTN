package cli

import "strings"

func glamourStyle(styleName string) []byte {
	quarterHRString := strings.Repeat("─", width/4)
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
    "bold": true
  },
  "h1": {
    "prefix": "──────",
    "suffix": "──────",
    "color": "#FBF1C7",
    "background_color": "#b16286"
  },
  "h2": {
    "prefix": "+++++",
    "suffix": "+++++",
    "color": "#689d6a"
  },
  "h3": {
    "prefix": "────",
    "suffix": "────",
    "color": "#458588"
  },
  "h4": {
    "prefix": "+++",
    "suffix": "+++",
    "color": "#d79921"
  },
  "h5": {
    "prefix": "──",
    "suffix": "──",
    "color": "#98971a"
  },
  "h6": {
    "prefix": "+",
    "suffix": "+",
    "color": "#cc241d"
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
    "format": "\n` + strings.Repeat(quarterHRString, 4) + `\n"
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
    "prefix": "",
    "suffix": "",
    "color": "203",
    "background_color": "236"
  },
  "code_block": {
    "block_prefix": "` + quarterHRString + "code" + quarterHRString + "\\n" + `",
    "block_suffix": "` + strings.Repeat(quarterHRString, 2) + "────\\n" + `",
    "color": "248",
    "chroma": {
      "text": {
        "color": "#A89984"
      },
      "error": {
        "color": "#F1F1F1",
        "background_color": "#F05B5B"
      },
      "comment": {
        "color": "#928374"
      },
      "comment_preproc": {
        "color": "#FF875F"
      },
      "keyword": {
        "color": "#D65E5E"
      },
      "keyword_reserved": {
        "color": "#FF5FD2"
      },
      "keyword_namespace": {
        "color": "#D65E5E"
      },
      "keyword_type": {
        "color": "#FABD2F"
      },
      "operator": {
        "color": "#FBF1C7"
      },
      "punctuation": {
        "color": "#E8E8A8"
      },
      "name": {
        "color": "#A89984"
      },
      "name_builtin": {
        "color": "#FE8019"
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
        "color": "#FBF1C7"
      },
      "name_other": {},
      "literal": {},
      "literal_number": {
        "color": "#D3869B"
      },
      "literal_date": {},
      "literal_string": {
        "color": "#B8BB26"
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
        "color": "#FBF1C7"
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
}`)
}
