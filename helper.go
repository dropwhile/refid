// Copyright (c) 2023 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package refid

// Must is a helper that wraps a call to a function returning (RefId, error)
// and panics if the error is non-nil. It is intended for use in variable initializations
// such as
//
//	var (
//		refA = refid.Must(refid.New())
//		refB = refid.Must(refid.NewTagged(2))
//		refC = refid.Must(refid.Parse("0r2nbq0wqhjg186167t0gcd1gw"))
//		refD = refid.Must(refid.ParseTagged("0r2nbq0wqhjg186167t0gcd1gw", 2))
//	)
func Must(r RefId, err error) RefId {
	if err != nil {
		panic(err)
	}
	return r
}
