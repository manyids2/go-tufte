// Get node-types.json from tree-sitter parser definitions
//
//	❯ cat node-types.json |\
//	   grep '^    "type"' |\
//	   cut -d':' -f2 |\
//	   cut -d',' -f1 |\
//	   xargs -I{} echo {} > core/mardown_types.go
//
// -- Root --
// document
//
// -- Heading --
// atx_heading
// atx_h1_marker
// atx_h2_marker
// atx_h3_marker
// atx_h4_marker
// atx_h5_marker
// atx_h6_marker
// setext_heading
// setext_h1_underline
// setext_h2_underline
// heading_content
//
// -- Content --
// paragraph
// text
//
// -- Epigraph --
// block_quote
//
// -- Code --
// code_span
// fenced_code_block
// code_fence_content
// indented_code_block
//
// -- link --
// link
// link_destination
// link_label
// link_reference_definition
// link_text
// link_title
// uri_autolink
// www_autolink
// email_autolink
//
// -- Styles --
// strikethrough
// emphasis
// strong_emphasis
//
// -- Breaks --
// soft_line_break
// hard_line_break
// line_break
// thematic_break
//
// -- Html --
// html_atrribute
// html_attribute_value
// html_block
// html_cdata_section
// html_close_tag
// html_comment
// html_declaration
// html_open_tag
// html_processing_instruction
// html_self_closing_tag
// html_attribute_key
// html_declaration_name
// html_tag_name
//
// -- Image --
// image
// image_description
//
// -- List --
// loose_list
// tight_list
// list_item
// list_marker
// task_list_item
// task_list_item_marker
//
// -- Table --
// table
// table_cell
// table_data_row
// table_delimiter_row
// table_header_row
// table_column_alignment
//
// -- Misc --
// info_string
// backslash_escape
// character_reference
// virtual_space
package core
