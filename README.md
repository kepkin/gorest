# gorest
Tool for generating go code based on open api

# Rationale

Why not [goa][goa]

It's fun to design your API with frontend developer together! Odds are he/she knows swagger better than go.

Why not [go-swagger][go-swagger]

It is designed as a framework, while gorest is designed as library. Framework usually implies limitions for instance:
  - you can not access request object in handler
  - they have they own cmd flag parser for some reason
  - no integration with other libs like [elastic-apm][elastic-apm] while GIN does
  - you can not add custom handler without adding it to swagger, which might be insecure (like callback handler for PayPal)
  - own router, which you can't change
  
For now you have to use gin with gorest, but it's gonna to be changed in the future.

# Building

```
make
```

[goa]: https://github.com/goadesign/goa
[go-swagger]: https://github.com/go-swagger/go-swagger
[elastic-apm] https://github.com/elastic/apm-agent-go
