# goteal

goteal enables writing Smart Contracts for Algorand in the Go programming language.

goteal achieves this by compiling Go programs into [AVM](https://developer.algorand.org/articles/introducing-algorand-virtual-machine-avm-09-release/)/[TEAL](https://developer.algorand.org/docs/reference/teal/specification/) programs.


## Why goteal

goteal exists to simplify the process of authoring and testing Algorand Smart Contracts for people
that are familiar with the Go programming language.

Those familiar with Go know that the language and surrounding ecosystem offers many helpful tools. `goteal` exists to bridge the Go ecosystem with the Algorand Smart Contract ecosystem. 

## How goteal works

goteal takes a Go program that follows a few constraints (see below) and transpiles/compiles it to
[TEAL](https://developer.algorand.org/docs/reference/teal/specification/) (Transaction Execution Approval Language) - the language that nodes in the Algorand network execute to validate and approve transactions.

## Constraints

A [TEAL](https://developer.algorand.org/docs/reference/teal/specification/) program must exit with
an integer value that indicates whether or not the transaction should proceed (zero for should-not-proceed,  >0 for should-proceed). `goteal` expects a package that is being compiled to TEAL to expose the following interface:

```go

func Contract(globals types.Globals, gtxn types.GroupTransaction, txn types.Transaction) (int,
error)
```

Where if any non-nil error is returned the program will halt (and reject the transaction).

## Interface

When an Algorand Smart Contract is executing there are several important sources of context:

* globals - contains network-wide information including values of current global configuration
  parameters.
* txn - The current Transaction.
* gtxn - The current Transaction Group (a Transaction in Algorand may be part of a group).


# Compiling a contract with goteal

You compile a go program to TEAL program with the following pattern:

```shell
goteal build github.com/yourname/yourproject/contract1
```

This will leverage your current Go environment to find source code in the `contract1` package, will
find the `Contract()` entrypoint (as defined above) and output to stdout the TEAL program
equivalent.

You can output direct bytecode by supplying the `-o bytecode` argument.

# goteal vet

`goteal` intends to encourage common patterns and warn contract authors about common missteps or
errors in contract development. The `goteal vet` subcommand will run your contract through several
static analysis checks to help guide you in your contract authorship.

**Example usage:** `goteal vet github.com/yourname/yourproject/contract1`

`goteal vet` will output in a standard format any issues that are detected in your contract design.

# goteal test

`goteal` leverages the rich testing ecosystem from the Go programming language and makes it easy to
write tests that verify the behavior of your contract.

By placing your test code into `*_test.go` files in your package directory you can easily test your
contract via `goteal test`:

**Example usage:** `goteal test github.com/yourname/yourproject/contract1`

