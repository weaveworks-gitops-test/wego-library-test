module github.com/weaveworks-gitops-test/wego-library-test

go 1.16

require (
	github.com/sirupsen/logrus v1.7.0
	github.com/weaveworks/weave-gitops v0.2.4
)

// Only works in a dockerfile!
replace github.com/weaveworks/weave-gitops v0.2.4 => /go/src/github.com/weaveworks/weave-gitops
