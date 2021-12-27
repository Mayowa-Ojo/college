# College

Exploring the [go entity framework](https://github.com/ent/ent) with the gin web framework and some extracts from DDD.

Ent is a graph-based ORM for [go]() that prioritizes a schema-first approach and enables static typing through code generation. This eliminates the need to have `interface{}`s everywhere due to the drawbacks of the missing piece (generics) in the language.

Comparable to another [ORM (sqlboiler)](https://github.com/volatiletech/sqlboiler) which shares similar ideas, although much more stable and battle-tested. There are some subtle differences:
   - Ent lets you write your schema as normal `go` code rather than `sql`
   - Ent is a graph-based ORM. This makes it really easy to model relationships and run queries, aggregations and traversals
   - Ent has built in migration tooling (although still very much unstable)

Having used both `sqlboiler` and `ent`, I've had really good experiences and I'm curious to see what direction other ORMs take now that [generics](https://go.dev/blog/go1.18beta1) is coming to the language.