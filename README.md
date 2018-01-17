# get-acme-registration

Will try to find your ACME v1 registration ID based upon a PEM-encoded RSA private key file.

**Not compatible with ACME v2**.

## Usage:

```
go get github.com/alexzorin/get-acme-registration/...
$GOPATH/bin/get-acme-registration key.pem
```

### When key registration exists:

```
$ $GOPATH/bin/get-acme-registration user.key
Found existing registration: 27778401
```

### When key is not registered:
(Please diregard the agremeent URL error, it is intentional to prevent new registrations).
```
$ $GOPATH/bin/get-acme-registration user.key
Couldn't find key, probably not registered: Provided agreement URL [intentionally_failing] does not match current agreement URL [https://letsencrypt.org/documents/LE-SA-v1.2-November-15-2017.pdf]
```

### Use a different ACME endpoint
Pass in a directory URL via `ACME_DIRECTORY` env variable.
```
ACME_DIRECTORY=https://acme-staging.api.letsencrypt.org/directory $GOPATH/bin/get-acme-registration user.key
```