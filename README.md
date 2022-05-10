# Server

This server to serve merkle proof to client

## Build

```
  go build -v .
```

## API

### Get contracts list

```
    curl http://localhost:8080/contracts
```

### Get proof

```
    curl http://localhost:8080/get-proof/:contract_name/:user_address
```
