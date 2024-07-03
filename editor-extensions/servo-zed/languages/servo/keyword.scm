[
  (container_doc_comment)
  (doc_comment)

] @comment.doc

[
    (line_comment)
] @comment

(INTEGER) @number

(FLOAT) @number

[
  "true"
  "false"
] @boolean

[
  (LINESTRING)
  (STRINGLITERALSINGLE)
] @string

(CHAR_LITERAL) @string.special.symbol
(EscapeSequence) @string.escape
(FormatSequence) @string.special

[
  "struct"
  "sum"
  "enum"
  "option"
  "service"
  "alias"
  "map"
  "list"
] @keyword

[
  "="
  "->"
  "?"
] @operator

[
  ";"
  "."
  ":"
] @punctuation.delimiter

[
  "["
  "]"
  "("
  ")"
  "{"
  "}"
] @punctuation.bracket
