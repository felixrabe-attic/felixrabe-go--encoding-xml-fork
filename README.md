Based on [Go](http://golang.org/) changeset 21503:ad9e191a5194.

This is mostly a drop-in replacement for 'encoding/xml' with minor changes:

-   White space will encode back to white space.

-   BOM is ignored.

Usage:

    import "gopkg.in/felixrabe-go/encoding-xml-fork.v0"

It will import as "xml".

Unless otherwise noted, the Go source files are distributed
under the BSD-style license found in the LICENSE file.
