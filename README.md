RefId
=====

[![Build Status](https://github.com/dropwhile/refid/workflows/unit-tests/badge.svg)][1]
[![GoDoc](https://godoc.org/github.com/dropwhile/refid?status.png)][2]
[![Go Report Card](https://goreportcard.com/badge/dropwhile/refid)](https://goreportcard.com/report/dropwhile/refid)
[![License](https://img.shields.io/github/license/dropwhile/refid.svg)](https://github.com/dropwhile/refid/blob/master/LICENSE.md)

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

## Inspirations

*   https://github.com/gofrs/uuid  
    gofrs' UUIDv7 might be a better option if you need something standard.
*   https://github.com/uuid6/uuid6-ietf-draft,
    https://github.com/ietf-wg-uuidrev/rfc4122bis  
    UUIDv6/v7 drafts

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

## Installation
```
go get -u github.com/dropwhile/refid
```

## Usage

### Simple
```go
// generate refid
refId, err := refid.New()
// generate refid (or panic)
refId = refid.Must(refid.New())

// encoding...
// encode to native encoding (base32 with Crockford alphabet)
s := refId.String() // "0r326xw2xbpga5tya7px89m7hw"
// encode to base64 encoding
s = refId.ToBase64String() // "BgYjd4Lq7QUXXlHt1CaHjw"
// encode to hex encoding
s = refId.ToHexString() // "0606237782eaed05175e51edd426878f"
// raw bytes
b := refId.Bytes()

// decoding...
// decode from native
refId2, err := refid.Parse(s)
// decode from base64
refId2, err = FromBase64String(s)
// decode from hex
refId2, err = FromHexString(s)

// get the time out of a refId (as a time.Time)
var ts time.Time = refId2.Time()
```

### Tagging

Simple tagging usage:
```go
myTag := 2

// generate a refId with tag set to 1
refId = refid.Must(refid.NewTagged(1))
// you can also set it manually after generation
refId.SetTag(myTag)
// check if it is tagged
refId.Tagged() // true
// check if it has a specific tag
refId.HasTag(1) // false
refId.HasTag(2) // true


s := refId.String()
// require desired tag or fail parsing
r, err := refid.ParseTagged(1, s) // err != nil here, as refid was tagged 2
r, err = refid.ParseTagged(2, s) // err == nil here, as refid was tagged 2
```

#### What use is tagging?

Tag support ensures that a refid of a certain tag type can be made distinct from other refids -- those of a different tag type, or those with no tag type.  

A hypothetical example is a refid url paramater for a type named "author", can be
enforced as invalid when someone attempts to supply it as input for a different
refid url parameter for a type named "book".

Making tagging usage easier with RefIdTagger:
```go
// AuthorRefId ensures it will only succesfully generate and parse tag=2 refids
AuthorRefIdT := refid.RefIdTagger(2)
// BookRefId ensures it will only succesfully generate and parse tag=3 refids
BookRefIdT := refid.RefIdTagger(3)

authorRefId := refid.Must(AuthorRefIdT.New()) // generated with a tag of 2
authorRefId.HasTag(2) // true
bookRefId := refid.Must(BookRefIdT.New()) // generated with a tag of 3
bookRefId.HasTag(3) // true

r, err := AuthorRefIdT.Parse(authorRefId.String()) // succeeds; err == nil
r, err = bookRefId.Parse(authorRefId.String()) // fails; err != nil
```

## reftool command like utility

Installation:
```
go install github.com/dropwhile/refid/cmd/reftool@latest
```

```
% reftool generate -t 5
native enc:   0r326xw2xbpga5tya7px89m7hw
hex enc:      0606237782eaed05175e51edd426878f
base64 enc:   BgYjd4Lq7QUXXlHt1CaHjw
tag value:    5
time(string): 2023-09-24T23:47:38.954477Z
time(micros): 1695599258954477

% generate -t 2 -w "2023-01-01T00:00:11.123456Z"
native enc:   0qrjh15pzc004nzrkbpcp2v0wm
hex enc:      05f12884b6fb000257f89aeccb0b60e5
base64 enc:   BfEohLb7AAJX-Jrsywtg5Q
tag value:    2
time(string): 2023-01-01T00:00:11.123456Z
time(micros): 1672531211123456

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