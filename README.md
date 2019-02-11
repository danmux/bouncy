Bouncy
======

Elasticsearch in Go

Distributed, replicated, full text search, using [Bleve](http://blevesearch.com/) and [FoundationDB](https://www.foundationdb.org/)

Design
------

Instead of considering the following hard problems:
* Rebalancing strategy
* Shard count/size for reasonable rebalances vs not too heavy scatter gather searches
* Add a time based shard dimension for optimal time range limits (e.g. log searches)

We implement a Bleve KVStore backed by FoundationDB

We use the older of Bleve's index types `upside_down` as this is structured around a key value store model of persistence, which maps well to foundationDB's sorted key value behavior.

Development
-----------

[Download FoundationDB](https://www.foundationdb.org/download/) including client libraries, and run the single [default development node](https://apple.github.io/foundationdb/local-dev.html)

Clone this repo, `go install` and run `bouncy`

### Warnings
because FoundationDB client bindings for go use cgo, you may well get some linker warnings such as...

`ld: warning: object file (...000025.o) was built for newer OSX version (10.14) than being linked (10.13)`



### Notes

Bleve provides a clean interface for its backing store, and this interface maps quite nicely to FoundationDB's API, which is provided by their go bindings https://github.com/apple/foundationdb/tree/master/bindings/go which are vendored into this repo, as is the whole of the blevesearch repo.

Vendoring is done with [go dep](https://github.com/golang/dep)
