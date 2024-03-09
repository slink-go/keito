# keito - けいと - 毛糸
Security key &amp; token generator

## Generate secret key:

Generates cryptographically secure random string of given lenght (or minimum required length for given signing algorithm)

```bash
keito key -a <algorithm> -l <key length>

    -a, --algo <string>    [OPT] signing algorithm (applies minimum key lengh limit) 
    -l, --length <num>     [OPT] key length
```

## Generate JWT token:

```bash
keito token <arguments>

  -a, --algo <string>      [REQ] key signing algorithm (hs256, hs384, hs512)
  -c, --claims <string>    [OPT] token claims (comma-separated key=value pairs)
  -d, --duration <string>  [REQ] token duration (i.e. 15m, 1h, 3d)
  -i, --issuer <string>    [OPT] token issuer
  -k, --key <string>       [REQ] signature key
  -o, --onetime            [OPT] generate 'jti' claim for one-time-use token
  -s, --subject <string>   [OPT] token subject
```

## Parse JWT token:

```bash
keito parse -t <token> [-k <key>]

  -t, --token <string>     [REQ] token to be parsed
  -k, --key <string>       [OPT] signing key to verify token signature
```
