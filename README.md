# go-tufte

Web and cli based markdown viewer and editor.

## Markdown specs

- Metadata
- Headers (`H1`, `H2`, `H3`, ...)

- [Overview](https://dave.autonoma.ca/blog/2019/05/22/typesetting-markdown-part-1/)
  Provided enough interest, the series will include the following parts:
  - Build Script – create user-friendly shell scripts
  - Tool Review – describe how the toolset works
  - Automagicify – continuously integrate typesetting
  - Theme Style – define colours, fonts, and layout
  - Interpolation – define and use external variables
  - Computation – leverage R for calculations
  - Mathematics – beautifully typeset equations
  - Annotations – apply different styles to annotated text
  - Figures – draw figures using MetaPost

Basics of build script

## Treesitter + go + markdown

- done [fork of go-tree-sitter](github.com/manyids2/go-tree-sitter-with-markdown)
- can output children of any markdown file now.

### Roadmap

- [x] use go-tree-sitter for markdown
- [x] point at any md file
- [x] print the usual hierarcy tree
- [ ] print toc based on atx headers
- [ ] create tui to test go-tree-sitter queries

## Table of contents

- Prints following info about headers

  - indented text
  - heading level
  - start row
  - end row
  - start byte
  - end byte

- CONSTRAINT: Headers are only single line
- [ ] Print toc with proper indent for sidebar
- [ ] Make it take a io.Writer interface
- [ ] Options for output format to json, yaml, md, etc.

How to use awk to trim the output?

```
❯ ./go-tufte toc -p ./data/headers-hash.md
Hi |> 1 [0, 1] (0, 5)
  hello |> 2 [2, 5] (6, 24)
    hi |> 3 [6, 9] (25, 36)

❯ ./go-tufte toc -p ./data/headers-hash.md |\
       awk -F '|>' '{print substr($1, 1, length($1)-1)}' |\
       awk '{ gsub(/^[ \t]+|[ \t]+$/, ""); print }'
Hi
hello
hi

❯ ./go-tufte toc -p ./data/headers-hash.md |\
       awk -F '|>' '{print $2}'
 1 [0, 1] (0, 5)
 2 [2, 5] (6, 24)
 3 [6, 9] (25, 36)

```

## Tufte

Need examples of conversion to html.

- Data ink.
- Typography.
- Fonts.
- Spacing.

## Sidenotes

- need to find the article
