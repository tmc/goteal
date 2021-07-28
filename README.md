# goteal

goteal is a program to compile a Go program into an [AVM](https://developer.algorand.org/articles/introducing-algorand-virtual-machine-avm-09-release/)/[TEAL](https://developer.algorand.org/docs/reference/teal/specification/) program.


## Why goteal

goteal exists to simplify the process of authoring and testing Algorand Smart Contracts to people
that are familiar with the Go programming language.

As many Go programmers know, the Go programming language has a rich tooling ecosystem. This project intends to allow people that know Go to use parts of that ecosystem in the process of authoring Algorand Smart Contracts.

## How goteal works

goteal takes a Go program that follows a few constraints (see below) and transpiles/compiles it to
[TEAL](https://developer.algorand.org/docs/reference/teal/specification/) (Transaction Execution Approval Language) - the language that the Algorand system uses to validate and approve transactions.

## Constraints

A [TEAL](https://developer.algorand.org/docs/reference/teal/specification/) program must exit with
an integer value that indicates whether or not the transaction should proceed (zero for should-not-proceed,  >0 for should-proceed). `goteal` expects a package that is being compiled to TEAL to expose the following interface:

```go

func Contract(globals types.Globals, gtxn types.GroupTransaction, txn types.Transaction) (int,
error)
```

Where if any non-nil error is returned the program will halt (and reject the transaction).

## Interface

When an Algorand Smart Contract is executing there are several impoortant sources of context:

* globals - contains network-wide information including values of current global configuration
  parameters.
* txn - The current Transaction.
* gtxn - The current Transaction Group (a Transaction in Algorand may be part of a group).

