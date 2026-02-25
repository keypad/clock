```bash
> what is this?

 tiny timezone utility.
 local cli + local http output in plain text.
 simple responses.

> features?

 ✓ one line timestamp output
 ✓ timezone aliases (utc/local/est/cst/mst/pst/cet/ist/jst)
 ✓ alias listing at /zones
 ✓ local server mode for curl use
 ✓ zero external services

> run?

 go run ./src
 go run ./src utc
 go run ./src serve

> test?

 go test ./...

> examples?

 curl http://127.0.0.1:4173/
 curl http://127.0.0.1:4173/utc
 curl http://127.0.0.1:4173/zones
 curl http://127.0.0.1:4173/bad

> stack?

 go 1.26 stdlib

> links?

 https://github.com/keypad/clock
```
