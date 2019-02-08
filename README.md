This is a test project to show how to use Go modules when building projects and libraries with many subpackages, where you want to be able to use different versions of those subpackages.

For example, let's say you have one team who owns a pg package and another team that owns a kafka package. You want to have the pg team be able to make changes to their package without unintentionally affecting the kafka package or users of the kafka package.

The key thing is to add a go.mod for each sub package, instead of having one top-level go.mod for the whole repo. This way in your service' go.mod file you can control that sub packages version like so. Giving each sub package their own go.mod gives them their own identity from a go module perspective, as described in the [go mod wiki](https://github.com/golang/go/wiki/Modules#gomod).

Here's an example of how you'd setup the go.mod of the project that's using the libraries:

~~~
module github.com/travisjeffery/gomodservice

require (
	github.com/travisjeffery/gomodlib v0.0.0-20190208085931-4eb83f17895c
	github.com/travisjeffery/gomodlib/subpkga v0.0.0-20190208085931-4eb83f17895c // indirect
	github.com/travisjeffery/gomodlib/subpkgb v0.0.0-20190208085931-4eb83f17895c // indirect
)

replace github.com/travisjeffery/gomodlib/subpkga => github.com/travisjeffery/gomodlib/subpkga v0.0.0-20190208090204-d8f7841ac5e3
~~~

If you don't care about using specific sub package versions you don't need to work about any of this, just go get the root package and you're done.

[gomodservice](https://github.com/travisjeffery/gomodservice) is an example of a project depending on your packages, the go.mod file in that project is how you'd set yours up as well.
