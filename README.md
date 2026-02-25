# Cerret

<ins>**C**</ins>asual <ins>**Er**</ins>ror <ins>**Ret**</ins>urns

A Go linter that detects leaking internal errors in RPC method responses.

`cerret` finds RPC-style methods that return raw `err` variables, which can leak internal implementation details to users. Instead, errors should be wrapped with appropriate status codes or user-facing messages.

## Usage

```bash
go install github.com/snehilshah/cerret/cmd/cerret@latest
go vet -vettool=$(which cerret) ./...
```