# goteal's design

## Overview

This document describes the design and rationale backing goteal.

## Design

goteal operates over existing Go source code by translating it into the SSA (Static Single
Assignment) representation of the source code and translating this resulting IR into AVM/TEAL
programs and bytecode.

## Limitations

Currently there are not plans to support aspects of the Go programming langauge that would not be
reasonble to includee when targetting such a limited environment as the AVM.

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
