#keito - けいと - 毛糸
Security key &amp; token generator

##Generate secret key:

Generates cryptographically secure random string of given lenght (or minimum required length for given signing algorithm)

```bash

keito key -a <algorithm> -l <key length>
    -a [optional] signing algorithm (applies minimum key lengh limit) 
    -l [optional] key length

```

##Generate JWT token:

```bash

keito token -a <...> -l <key length>
  -a, --algo 				[REQ] key signing algorithm (hs256, hs384, hs512)
  -c, --claims 			[OPT] token claims (comma-separated key=value pairs)
  -d, --duration    	[REQ] token duration (i.e. 15m, 1h, 3d)
  -i, --issuer     		[OPT] token issuer
  -k, --key				[REQ] signature key
  -o, --onetime		[OPT] generate 'jti' claim for one-time-use token
  -s, --subject		[REQ] token subject
  
```

##Parse JWT token:

```bash

keito parse -t <token> [-k <key>]
  -t, --token 			[REQ] token to be parsed
  -k, --key				[OPT] signing key to verify token signature
  
```
