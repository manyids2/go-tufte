# go-tufte

Web and cli based markdown viewer and editor.

## Markdown specs

[tree-sitter-markdown](TREESITTER.md)

## Treesitter + go + markdown

- done [fork of go-tree-sitter](github.com/manyids2/go-tree-sitter-with-markdown)
- can output children of any markdown file now.

- Need to switch to [MDeiml implementation](https://github.com/MDeiml/tree-sitter-markdown/tree/split_parser)
  - !! `split_parser` branch
  - `tree-sitter-markdown`
  - `tree-sitter-markdown-inline`
- Mainly done
  - inline parsing still left


### Roadmap

- [x] use go-tree-sitter for markdown
- [x] point at any md file
- [x] print the usual hierarcy tree
- [x] print toc based on atx headers
- [ ] keep reference to sections
- [ ] render tree based on node type ( block stack from glow? )
- [ ] manage memory for rendered elements
- [ ] folding
- [ ] inline parsing

## `toc` Table of contents

- Prints following info about headers

  - indented text
  - heading level
  - start row
  - end row
  - start byte
  - end byte

- CONSTRAINT: Headers are only single line
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


## `tui` Terminal interface

Features:

- [ ] Navigate by section
- [ ] Styles
- [ ] Presentation

### Moving to tcell from tview?


### Moving to nvim from go?


## Tufte

Need examples of conversion to html.

- Data ink.
- Typography.
- Fonts.
- Spacing.

## Sidenotes

- need to find the article
