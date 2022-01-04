# College

Exploring [go ent](https://github.com/ent/ent) with the gin web framework and some extracts from DDD.

Ent is a graph-based ORM for [go](https://github.com/golang/go) that prioritizes a schema-first approach and enables static typing through code generation. This eliminates the need to have `interface{}`s everywhere due to the drawbacks of the missing piece (generics) in the language.

Comparable to another [ORM (sqlboiler)](https://github.com/volatiletech/sqlboiler) which shares similar ideas, although much more stable and battle-tested. There are some subtle differences:
   - Ent lets you write your schema as normal `go` code rather than `sql`
   - Ent is a graph-based ORM. This makes it really easy to model relationships and run queries, aggregations and traversals
   - Ent has built in migration tooling (although still very much unstable)

Having used both `sqlboiler` and `ent`, I've had really good experiences and I'm curious to see what direction other ORMs take now that [generics](https://go.dev/blog/go1.18beta1) is coming to the language.

|              | Header 1        | Header 2                       || Header 3                       ||
|              | Subheader 1     | Subheader 2.1  | Subheader 2.2  | Subheader 3.1  | Subheader 3.2  |
|==============|-----------------|----------------|----------------|----------------|----------------|
| Row Header 1 | 3row, 3col span                                 ||| Colspan only                   ||
| Row Header 2 |       ^                                         ||| Rowspan only   | Cell           |
| Row Header 3 |       ^                                         |||       ^        | Cell           |
| Row Header 4 |  Row            |  Each cell     |:   Centered   :| Right-aligned :|: Left-aligned  |
:              :  with multiple  :  has room for  :   multi-line   :    multi-line  :  multi-line    :
:              :  lines.         :  more text.    :      text.     :         text.  :  text.         :
|--------------|-----------------|----------------|----------------|----------------|----------------|
[Caption Text]