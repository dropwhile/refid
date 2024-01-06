refid
=====

[![Build Status](https://github.com/dropwhile/refid/workflows/unit-tests/badge.svg)][1]
[![GoDoc](https://godoc.org/github.com/dropwhile/refid/v2?status.png)](https://godoc.org/github.com/dropwhile/refid/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/dropwhile/refid/v2)](https://goreportcard.com/report/github.com/dropwhile/refid/v2)
[![License](https://img.shields.io/github/license/dropwhile/refid.svg)](https://github.com/dropwhile/refid/blob/master/LICENSE.md)

## About

A refid (short for Reference Identifier) is a unique identifier,
similar to a UUID, with some differences.

There are two types of refids:

*   TimePrefixed - similar to UUIDv7
*   RandomPrefixed - similar to UUIDv4

### TimePrefixed (type:0x00)

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                           unix_ts_ms                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|        unix_ts_ms       |       rand_a      |t|      tag      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             rand_b                            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
unix_ts_ms:
    45 bits big-endian unsigned number of Unix epoch timestamp milliseconds.
    (over 2000 years worth... until about year 3084 or so)
rand_a:
    10 bits random pad. fill with crypto/rand random.
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

TimePrefix:
*   unix timestamp with millisecond precision
*   sortable, db index friendly
*   Compared to UUIDv7
    *   tagging support
    *   slightly smaller Unix timestamp (45 vs 48 bits), though still
        enough bits until about year 3084
    *   same size random pad section (74 bits)
    *   not a standard

RandomPrefix:
*   not sortable, not db index friendly
*   Compared to UUIDv4
    *   tagging support
    *   slightly smaller random section (119 vs 122 bits), though still
        good collision resistance
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
go get -u github.com/dropwhile/refid/v2
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
s := rID.String() // "1hh3eb5563000h8ab7mbsq1gq0"
// encode to base64 encoding
s = rID.ToBase64String() // "DGI3LKUwwABFClnovNwwuA"
// encode to hex encoding
s = rID.ToHexString() // "0c62372ca530c000450a59e8bcdc30b8"
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

#### Tagging with Custom Types

A more complicated tagging example using embedded types:

```go
import (
    "github.com/dropwhile/refid/v2"
    "github.com/dropwhile/refid/v2/reftag"
)

type NoteID struct {
    // the reftag.IDt* types are generated to provide a handy means to
    // wrap with your own struct type, to provide type constraints in function
    // parameters and struct members
    reftag.IDt8
}

// some helpful shortcut function aliases
var (
    // generate a new time-prefixed refid of tag type 8
    NewNoteID       = reftag.New[NoteID]
    // generate a new random-prefixed refid of tag type 8
    NewRandomNoteID = reftag.NewRandom[NoteID]
    // handy matcher for sql mock libraries (gomock, pgxmock, etc)
    NoteIDMatcher   = reftag.NewMatcher[NoteID]()
    // some handy parsing aliases
    NoteIDFromBytes = reftag.FromBytes[NoteID]
    ParseNoteID     = reftag.Parse[NoteID]
)

func ParseNote(ctx context.Context, db dbHandle, noteStr string) (*NoteID, error) {
    noteID, err := ParseNoteID(noteStr)
    // error will be non-nil if the tag in the RefID does not match the expectation (`8`)
    ...
    return NoteID, err
}

func DBLookupNote(ctx context.Context, db dbHandle, noteID NoteID) (*DBNote, error) {
    // noteID is now a compile time checked type ensuring that RefIDs of a different
    // tag are not accidentally allowed.
    ...
}
```

## refidtool CLI utility

Installation:
```
go install github.com/dropwhile/refid/v2/cmd/refidtool@latest
```

```
# generate a refi.ID with a tag of 5
% refidtool generate -t 5
native enc:   1hh3ecgwmg40ahtaf41qc1dz88
hex enc:      0c6237321ca40805474a79037605bf42
base64 enc:   DGI3MhykCAVHSnkDdgW_Qg
tag value:    5
type:         TimePrefixed
time(string): 2023-12-08T00:49:04.916Z
time(millis): 1701996544916

# generate a refid with a tag of 5, and only output the native(base32) encoding
% refidtool generate -t 5 -o native
1hh3ecvn3dt0bs66r6p6q4gqe0

# generate a refid with a tag of 5, and only output the hex encoding
% refidtool generate -t 5 -o hex
0c6237347678a60554f8b73f96992fed

# generate a random-prefixed refid with a tag of 4, and only output the base64 encoding
% refidtool generate -r -t 4 -o base64
KRUV6EEACQRsivT1_pNr4w

# genrate a refid with a tag of 2, at a specific timestamp
% refidtool generate -t 2 -w "2023-01-01T00:00:11.123456Z"
native enc:   1gnna1wvkbx047v48dhxeh7c5g
hex enc:      0c2b55079b9afa021f644363d744ec2c
base64 enc:   DCtVB5ua-gIfZENj10TsLA
tag value:    2
type:         TimePrefixed
time(string): 2023-01-01T00:00:11.123Z
time(millis): 1672531211123

# decode a refid and display
% refidtool parse 1hh3eehsh2p05ycz44y9erhvhm
native enc:   1hh3eehsh2p05ycz44y9erhvhm
hex enc:      0c62373a3988ac02f99f213c97623b8d
base64 enc:   DGI3OjmIrAL5nyE8l2I7jQ
tag value:    2
type:         TimePrefixed
time(string): 2023-12-08T00:50:11.377Z
time(millis): 1701996611377

# here is what a random-prefixed refid looks like
# note the time is the zero value for time.Time
% refidtool parse 8xqzbn495crg5mxjztjczfk0xr
native enc:   8xqzbn495crg5mxjztjczfk0xr
hex enc:      476ff5d4892b3102d3b2fea4cfbe60ee
base64 enc:   R2_11IkrMQLTsv6kz75g7g
tag value:    2
type:         RandomPrefixed
time(string): 1970-01-01T00:00:00Z
time(millis): 0
```


[1]: https://github.com/dropwhile/refid/actions
[3]: https://choosealicense.com/licenses/mit/
