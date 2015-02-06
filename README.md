pointcoind
====

pointcoind is a fork of Conformal's [btcd](https://github.com/btcsuite/btcd)

It relays newly mined blocks, maintains a transaction pool, and relays
individual transactions that have not yet made it into a block.  It ensures all
transactions admitted to the pool follow the rules required by the block chain
and also includes the same checks which filter transactions based on
miner requirements ("standard" transactions) as Bitcoin Core.

One key difference between pointcoind and Bitcoin Core is that pointcoind does *NOT* include
wallet functionality and this was a very intentional design decision.  See the
blog entry [here](https://blog.conformal.com/pointcoind-not-your-moms-bitcoin-daemon)
for more details.  This means you can't actually make or receive payments
directly with pointcoind. 

## Requirements

[Go](http://golang.org) 1.2 or newer.

## Installation

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Run the following command to obtain pointcoind, all dependencies, and install it:
  ```$ go get github.com/PointCoin/pointcoind/...```

- pointcoind (and utilities) will now be installed in either ```$GOROOT/bin``` or
  ```$GOPATH/bin``` depending on your configuration.  If you did not already
  add the bin directory to your system path during Go installation, we
  recommend you do so now.

## Updating

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following command to update pointcoind, all dependencies, and install it:
  ```$ go get -u -v github.com/btcsuite/pointcoind/...```

## Getting Started

pointcoind has several configuration options avilable to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Linux/BSD/POSIX/Source

```bash
$ ./pointcoind
````

## Documentation

You can use the btcsuite documentation for must things.  It is located in the [docs](https://github.com/btcsuite/pointcoind/tree/master/docs) folder.

## License

pointcoind is licensed under the [copyfree](http://copyfree.org) ISC License.
