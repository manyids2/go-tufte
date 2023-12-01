```bash
cat ./../go-tree-sitter/markdown/grammar.json |\
  # Get at proper indent level
  grep '^    "' |\
  # Massage string
  cut -d':' -f1 |\
  sed 's/^    //' |\
  sed 's/"//g' >> TOKENS.md
```

document

backslash_escape
_backslash_escape
entity_reference
numeric_character_reference

link_label
link_destination
_link_destination_parenthesis
_text_no_angle
link_title

_newline_token
_last_token_punctuation
_block
_block_not_section
section
_section1
_section2
_section3
_section4
_section5
_section6
thematic_break
_atx_heading1
_atx_heading2
_atx_heading3
_atx_heading4
_atx_heading5
_atx_heading6
_atx_heading_content
_setext_heading1
_setext_heading2

indented_code_block
_indented_chunk
fenced_code_block
code_fence_content
info_string
language

html_block
_html_block_1
_html_block_2
_html_block_3
_html_block_4
_html_block_5
_html_block_6
_html_block_7
link_reference_definition
_text_inline_no_link

block_quote

list
_list_plus
_list_minus
_list_star
_list_dot
_list_parenthesis
list_marker_plus
list_marker_minus
list_marker_star
list_marker_dot
list_marker_parenthesis
_list_item_plus
_list_item_minus
_list_item_star
_list_item_dot
_list_item_parenthesis
_list_item_content

paragraph
_blank_line
_newline
_soft_line_break
_line
_word
_whitespace

task_list_marker_checked
task_list_marker_unchecked

pipe_table
_pipe_table_newline
pipe_table_delimiter_row
pipe_table_delimiter_cell
pipe_table_row
pipe_table_cell
