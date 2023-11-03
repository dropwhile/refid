refid
=====

[![Build Status](https://github.com/dropwhile/refid/workflows/unit-tests/badge.svg)][1]
[![GoDoc](https://godoc.org/github.com/dropwhile/refid?status.png)][2]
[![Go Report Card](https://goreportcard.com/badge/dropwhile/refid)](https://goreportcard.com/report/dropwhile/refid)
[![License](https://img.shields.io/github/license/dropwhile/refid.svg)](https://github.com/dropwhile/refid/blob/master/LICENSE.md)

## About

A refid (short for Reference Identifier) is a unique identifier,
similar to UUIDv7, with a few difference.

There are two types of refids: TimePrefixed and RandomPrefixed.

### TimePrefixed (type:0x00)

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                           unix_ts_ms                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|           unix_ts_ms          |    rand_a   |t|      tag      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
unix_ts_ms:
    48 bits big-endian unsigned number of Unix epoch timestamp milliseconds.
    (2284 years worth... until about year 4200 or so)
rand_a:
    7 bits random pad. fill with crypto/rand random.
t:
    1 bit for type. TimePrefixed (type:0)
tag:
    8 bits tag. 255 separate tags (0 being untagged).
rand_b:
    64 bits random pad. fill with crypto/rand random.
```

### RandomPrefixed (type:0x01)

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_a                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    rand_a                   |t|      tag      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
rand_a:
    55 bits random pad. fill with crypto/rand random.
t:
    1 bit for type. RandomPrefixed (type:1)
tag:
    8 bits tag. 255 separate tags (0 being untagged).
rand_b:
    64 bits random pad. fill with crypto/rand random.
```

## Features

General:
*   tagging (support for 255 distinct tags)
*   supports go/sql scanner/valuer
*   multiple encodings supported: native (base32), base64, base16 (hex)
*   similar to UUIDv7, with different tradeoffs

TimePrefix:
*   unix timestamp with millisecond precision
*   sortable, db index friendly
*   Compared to UUIDv7
    *   tagging support
    *   48 bits of Unix timestamp milliseconds from epoch (similar to UUIDv7)
    *   slightly smaller random section (71 vs 74 bits), though still good
        collision resistance
    *   not a standard

RandomPrefix:
*   not sortable, not db index friendly
*   Compared to UUIDv7
    *   tagging support
    *   slightly smaller random section (119 vs 122 bits), though still good
        collision resistance
    *   not a standard

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
// generate a TimePrefixed refid
rID, err := refid.New()
// generate a TimePrefixed refid (or panic)
rID = refid.Must(refid.New())
// generate a RandomPrefixed refid (or panic)
rID = refid.Must(refid.NewRandom())

// encoding...
// encode to native encoding (base32 with Crockford alphabet)
s := rID.String() // "0r326xw2xbpga5tya7px89m7hw"
// encode to base64 encoding
s = rID.ToBase64String() // "BgYjd4Lq7QUXXlHt1CaHjw"
// encode to hex encoding
s = rID.ToHexString() // "0606237782eaed05175e51edd426878f"
// raw bytes
b := rID.Bytes()

// decoding...
// decode from native
rID, err := refid.Parse(s)
// decode from base64
rID, err = refid.FromBase64String(s)
// decode from hex
rID, err = refid.FromHexString(s)

// get the time out of a TimePrefixed refid (as a time.Time)
var ts time.Time = rID.Time()
```

### Tagging

Simple tagging usage:
```go
myTag := 2

// generate a refid with tag set to 1
rID = refid.Must(refid.NewTagged(1))
// you can also set it manually after generation
rID.SetTag(myTag)
// check if it is tagged
rID.Tagged() // true
// check if it has a specific tag
rID.HasTag(1) // false
rID.HasTag(2) // true


s := rID.String()
// require desired tag or fail parsing
r, err := refid.ParseTagged(1, s) // err != nil here, as refid was tagged 2
r, err = refid.ParseTagged(2, s) // err == nil here, as refid was tagged 2
```

#### What use is tagging?

Tag support ensures that a refid of a certain tag type can be made distinct from
other refids -- those of a different tag type, or those with no tag type.  

A hypothetical example is a refid url paramater for a type named "author", can
be enforced as invalid when someone attempts to supply it as input for a
different refid url parameter for a type named "book".

Making tagging usage easier with refid.IDTagger:
```go
// AuthorID ensures it will only succesfully generate and parse tag=2 refids
AuthorIDT := refid.IDTagger(2)
// BookID ensures it will only succesfully generate and parse tag=3 refids
BookIDT := refid.IDTagger(3)

authorID := refid.Must(AuthorIDT.New()) // generated with a tag of 2
authorID.HasTag(2) // true
bookID := refid.Must(BookIDT.New()) // generated with a tag of 3
bookID.HasTag(3) // true

r, err := AuthorIDT.Parse(authorID.String()) // succeeds; err == nil
r, err = bookIDT.Parse(authorID.String()) // fails; err != nil
```

## reftool command like utility

Installation:
```
go install github.com/dropwhile/refid/cmd/reftool@latest
```

```
# generate a refi.ID with a tag of 5
% reftool generate -t 5
native enc:   0r326xw2xbpga5tya7px89m7hw
hex enc:      0606237782eaed05175e51edd426878f
base64 enc:   BgYjd4Lq7QUXXlHt1CaHjw
tag value:    5
type:         TimePrefixed
time(string): 2023-09-24T23:47:38.954477Z
time(millis): 1695599258954477

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
type:         TimePrefixed
time(string): 2023-01-01T00:00:11.123456Z
time(millis): 1672531211123456

# decode a refid and display
% reftool decode 0qrjh15pzc004nzrkbpcp2v0wm
native enc:   0qrjh15pzc004nzrkbpcp2v0wm
hex enc:      05f12884b6fb000257f89aeccb0b60e5
base64 enc:   BfEohLb7AAJX-Jrsywtg5Q
tag value:    2
type:         TimePrefixed
time(string): 2023-01-01T00:00:11.123456Z
time(millis): 1672531211123456
```

[1]: https://github.com/dropwhile/refid/actions
[2]: https://godoc.org/github.com/dropwhile/refid
[3]: https://choosealicense.com/licenses/mit/
