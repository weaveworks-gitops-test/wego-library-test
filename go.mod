module github.com/weaveworks-gitops-test/wego-library-test

go 1.16

require (
	github.com/sirupsen/logrus v1.7.0
	github.com/weaveworks/weave-gitops v0.2.5
)

// Only works in a dockerfile!
replace github.com/weaveworks/weave-gitops => /go/src/github.com/weaveworks/weave-gitops

replace github.com/go-logr/logr v1.1.0 => github.com/go-logr/logr v0.4.0

replace github.com/go-logr/zapr v1.1.0 => github.com/go-logr/zapr v0.4.0
