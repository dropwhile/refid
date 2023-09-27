RefID
=====

[![Build Status](https://github.com/dropwhile/refid/workflows/unit-tests/badge.svg)][1]
[![GoDoc](https://godoc.org/github.com/dropwhile/refid?status.png)][2]
[![Go Report Card](https://goreportcard.com/badge/dropwhile/refid)](https://goreportcard.com/report/dropwhile/refid)
[![License](https://img.shields.io/github/license/dropwhile/refid.svg)](https://github.com/dropwhile/refid/blob/master/LICENSE.md)

## About

A RefID is a sortable unique identifier, similar to UUIDv7, with a few difference.

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                           unix_ts_µs                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                   unix_ts_µs                  |      tag      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
7 bytes unix_ts_µs:
    48 bits of microseconds from 1970 (about 2280 or so years worth)
1 byte tag:
    255 separate tags (0 being untagged)
8 bytes random pad:
    fill with crypto/rand random
```

## Features

*   tagging (support for 255 distinct tags)
*   unix timestamp with microsecond precision
*   supports go/sql scanner/valuer
*   multiple encodings supported: native (base32), base64, base16 (hex)
*   similar to UUIDv7, with different tradeoffs
    *    slightly larger random section
    *    tag support
    *    no UUID version field
    *    not an rfc standard

## Non-Features

*   refids, like UUIDs, do not internally perform any signature verification.
    If the validity of the encoded timestamp and tag are required for any secure
    operations, the refid SHOULD be externally verified before parsing/decoding.
    An example of this could be a wrapping encoder/decoder doing hmac signing and verification.

## Inspirations

*   https://github.com/gofrs/uuid  
    gofrs' UUIDv7 might be a better option if you need something standard.
*   https://github.com/uuid6/uuid6-ietf-draft,
    https://github.com/ietf-wg-uuidrev/rfc4122bis  
    UUIDv6/v7 drafts

## Installation
```
go get -u github.com/dropwhile/refid
```

## Usage

### Simple
```go
// generate refid
rId, err := refid.New()
// generate refid (or panic)
rId = refid.Must(refid.New())

// encoding...
// encode to native encoding (base32 with Crockford alphabet)
s := rId.String() // "0r326xw2xbpga5tya7px89m7hw"
// encode to base64 encoding
s = rId.ToBase64String() // "BgYjd4Lq7QUXXlHt1CaHjw"
// encode to hex encoding
s = rId.ToHexString() // "0606237782eaed05175e51edd426878f"
// raw bytes
b := rId.Bytes()

// decoding...
// decode from native
rId2, err := refid.Parse(s)
// decode from base64
rId2, err = refid.FromBase64String(s)
// decode from hex
rId2, err = refid.FromHexString(s)

// get the time out of a RefID (as a time.Time)
var ts time.Time = rId2.Time()
```

### Tagging

Simple tagging usage:
```go
myTag := 2

// generate a RefID with tag set to 1
rId = refid.Must(refid.NewTagged(1))
// you can also set it manually after generation
rId.SetTag(myTag)
// check if it is tagged
rId.Tagged() // true
// check if it has a specific tag
rId.HasTag(1) // false
rId.HasTag(2) // true


s := rId.String()
// require desired tag or fail parsing
r, err := refid.ParseTagged(1, s) // err != nil here, as refid was tagged 2
r, err = refid.ParseTagged(2, s) // err == nil here, as refid was tagged 2
```

#### What use is tagging?

Tag support ensures that a refid of a certain tag type can be made distinct from other refids -- those of a different tag type, or those with no tag type.  

A hypothetical example is a refid url paramater for a type named "author", can be
enforced as invalid when someone attempts to supply it as input for a different
refid url parameter for a type named "book".

Making tagging usage easier with RefIDTagger:
```go
// AuthorRefID ensures it will only succesfully generate and parse tag=2 refids
AuthorRefIDT := refid.RefIDTagger(2)
// BookRefID ensures it will only succesfully generate and parse tag=3 refids
BookRefIDT := refid.RefIDTagger(3)

authorRefID := refid.Must(AuthorRefIDT.New()) // generated with a tag of 2
authorRefID.HasTag(2) // true
bookRefID := refid.Must(BookRefIDT.New()) // generated with a tag of 3
bookRefID.HasTag(3) // true

r, err := AuthorRefIDT.Parse(authorRefID.String()) // succeeds; err == nil
r, err = bookRefID.Parse(authorRefID.String()) // fails; err != nil
```

## reftool command like utility

Installation:
```
go install github.com/dropwhile/refid/cmd/reftool@latest
```

```
# generate a refid with a tag of 5
% reftool generate -t 5
native enc:   0r326xw2xbpga5tya7px89m7hw
hex enc:      0606237782eaed05175e51edd426878f
base64 enc:   BgYjd4Lq7QUXXlHt1CaHjw
tag value:    5
time(string): 2023-09-24T23:47:38.954477Z
time(micros): 1695599258954477

# generate a refid with a tag of 5, and only output the native(base32) encoding
% reftool generate -t 5 -o
0r34ky6h51r012an8skhbsvxt0

# generate a refid with a tag of 5, and only output the hex encoding
% reftool generate -t 5 -o=hex
060649f82794f10039169e91d0696763

# generate a refid with a tag of 5, and only output the base64 encoding
% reftool generate -o=base64
BgZJ-i1F2wALdZFJrWvNzA

# genrate a refid with a tag of 2, at a specific timestamp
% reftool generate -t 2 -w "2023-01-01T00:00:11.123456Z"
native enc:   0qrjh15pzc004nzrkbpcp2v0wm
hex enc:      05f12884b6fb000257f89aeccb0b60e5
base64 enc:   BfEohLb7AAJX-Jrsywtg5Q
tag value:    2
time(string): 2023-01-01T00:00:11.123456Z
time(micros): 1672531211123456

# decode a refid and display
% reftool decode 0qrjh15pzc004nzrkbpcp2v0wm
native enc:   0qrjh15pzc004nzrkbpcp2v0wm
hex enc:      05f12884b6fb000257f89aeccb0b60e5
base64 enc:   BfEohLb7AAJX-Jrsywtg5Q
tag value:    2
time(string): 2023-01-01T00:00:11.123456Z
time(micros): 1672531211123456
```

[1]: https://github.com/dropwhile/refid/actions
[2]: https://godoc.org/github.com/dropwhile/refid
[3]: https://choosealicense.com/licenses/mit/
