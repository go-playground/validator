Package go-validate-yourself
================
[![Build Status](https://travis-ci.org/joeybloggs/go-validate-yourself.svg?branch=v2)](https://travis-ci.org/joeybloggs/go-validate-yourself)

Package validator implements value validations for structs and individual fields based on tags.

Installation
============

Just use go get.

	go get gopkg.in/joeybloggs/go-validate-yourself.v2

or to update

	go get -u gopkg.in/joeybloggs/go-validate-yourself.v2

And then just import the package into your own code.

	import "gopkg.in/joeybloggs/go-validate-yourself.v2"

Usage
=====

Please see http://godoc.org/gopkg.in/joeybloggs/go-validate-yourself.v2 for detailed usage docs.

Contributing
============

There will be a development branch for each version of this package i.e. v1-development, please
make your pull requests against those branches.

If changes are breaking please create an issue, for discussion and create a pull request against
the highest development branch for example this package has a v1 and v1-development branch
however, there will also be a v2-development brach even though v2 doesn't exist yet.

I strongly encourage everyone whom creates a custom validation function to contribute them and
help make this package even better.
