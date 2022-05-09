# Server

This server to serve merkle proof to client

## Build

```
  go build -v .
```

## Config

Sample

```
[
  {
    "name": "privateB",
    "contract": "0x1649a8909B39cfb23ae34d274653FAd4Bd9BAcC7",
    "merkleRoot": "0x8bbc059b2d155670357b70c2aa67d6f3e271bf2f4404a832b9d58068842c9f75",
    "tokenTotal": "0xd3c21bcecceda0000000",
    "claims": {
      "0x50fAcafA6D2222De862519b959B4d11dda81B94d": {
        "index": 0,
        "amount": "0x69e10de76676d0000000",
        "proof": [
          "0x1536d908ef93d1fcb8c53c21403ca5f4f3b5a5cdfb20f99a8ac8a2b0a281966b"
        ]
      },
      "0x955884E6d067Cd9FA3066541BAcC3d4Ddf9982d1": {
        "index": 1,
        "amount": "0x69e10de76676d0000000",
        "proof": [
          "0xe05d68d346e2808af5eb089cca4d0bc51b19a71e15ae92ab2771747032e0fdd1"
        ]
      }
    }
  },
  {
    "name": "privateA",
    "contract": "0xDD8aE4537EB0937BA46C08654e41f0Bf7fc32aE5",
    "merkleRoot": "0x9a1776cbc2514c76120957a785b56e6eccbd7bf1ec772a2b853ca5108c1fa76b",
    "tokenTotal": "0x69e10de76676d0000000",
    "claims": {
      "0x64f1487dcf33702B32ddDe3F2342b28BD6b635F9": {
        "index": 0,
        "amount": "0x34f086f3b33b68000000",
        "proof": [
          "0xd3e70e1642b65d62e5c5ec4f9b67b57921e0239a52de7f758aa9504dc5917dd9"
        ]
      },
      "0x991eAF95Df6339C8d4bdD8Ef3f9F799dEC3531B0": {
        "index": 1,
        "amount": "0x34f086f3b33b68000000",
        "proof": [
          "0x58049dd80d79c11efe59381ee9651084279a26c3cbb87d29ad24b1887e1fc702"
        ]
      }
    }
  },
  {
    "name": "seedB",
    "contract": "0xd6b7De41C1f835bc366D7eAc4EB56fB8d5424e45",
    "merkleRoot": "0xd2c4147d69ab9aba121e2ad044022c7a5a3e4ce4d4b22f16b754706cd478c37a",
    "tokenTotal": "0x7f0e10af47c1c8000000",
    "claims": {
      "0x7a618b048F61A3DCf8a1B9bc04A5384f6DCd5964": {
        "index": 0,
        "amount": "0x3f870857a3e0e4000000",
        "proof": [
          "0x2cf288944de00b020e6eb79ae95dda922f8790e1875e171ff7544e89294e6c26"
        ]
      },
      "0x8D61aB7571b117644A52240456DF66EF846cd999": {
        "index": 1,
        "amount": "0x3f870857a3e0e4000000",
        "proof": [
          "0xc90329580f526368bde110e580f467162d3f60c2a18b57a6bfa485693355d905"
        ]
      }
    }
  }
]
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
