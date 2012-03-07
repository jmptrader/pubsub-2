pubsub
=====

pubsub is a package for Go that implements publish-subscribe messaging pattern.

## Installation

    go get github.com/fzzbt/radix

To run the tests:

	cd $GOROOT/src/pkg/pubsub
	go test -v

## HACKING

If you make contributions to the project, please follow the guidelines below:

*  Maximum line-width is 110 characters.
*  Run "gofmt -tabs=true -tabwidth=4" for any Go code before committing. 
   You may do this for all code files by running "make format".
*  Any copyright notices, etc. should not be put in any files containing program code to avoid clutter. 
   Place them in separate files instead. 
*  Avoid commenting trivial or otherwise obvious code.
*  Avoid writing fancy ascii-artsy comments. 
*  Write terse code without too much newlines or other non-essential whitespace.
*  Separate code sections with "//* My section"-styled comments.

New developers should add themselves to the lists in AUTHORS and/or CONTRIBUTORS files,
when submitting their first commit. See the CONTRIBUTORS file for details.


## Copyright and licensing

*Copyright 2012 The "pubsub" Authors*. See file AUTHORS and CONTRIBUTORS.  
Unless otherwise noted, the source files are distributed under the
*BSD 2-Clause License* found in the LICENSE file.
