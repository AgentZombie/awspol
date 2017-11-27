# awspol

`awspol` is intended allow handling AWS Policy Documents received as JSON via native Go structs.

## Background

AWS Policy Document JSON doesn't reliably map 1:1 with Go data structures. Most specifically there are fields with JSON that can either be a string or an array of strings:

```
    "Resource": "this resource",
```

or:

```
    "Resource": ["resource 1", "resource 2"],
```

A go struct member can only have one type so `encoding/json` can't natively handle direct serialization/deserialization for these documents.

## How do?

This package implements a `MultiString` type which is a slice of strings. When the JSON document refers to a string the slice has one element. When the JSON document refers to an array of strings the slice has all of those elements. When element is missing from the JSON document the `MultiString` will be `nil`.

## Status

This is woefully incomplete, Fields it's known to not deal with properly are left as `json.RawMessage`. PRs with tests are welcome.
