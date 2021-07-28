# goteal's design

## Overview

This document describes the design and rationale backing goteal.

## Design

goteal is a program that converts a subset of the Go programming language to TEAL, the language that
defines the smart contract language used in the [Algorand](https://www.algorand.com/) Virtual
Machine (AVM).

goteal operates over existing Go source code by translating it into the SSA (Static Single
Assignment) representation of the source code and translating this resulting IR into AVM/TEAL
programs and bytecode.

## Limitations

Currently there are not plans to support aspects of the Go programming language that would not be
reasonable to include e when targeting such a limited environment as the AVM.

Unsupported Go Language features:

* init functions
* goroutines
* channels
* defer
* select

Given the constrained nature of the AVM, source code that imports any of these packages will fail to
compile:

* networking-related packages
* syscalls
* unsafe

Note: It's desirable for this this list to shrink if possible.
