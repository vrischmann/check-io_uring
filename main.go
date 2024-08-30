// A generated module for CheckIoUring functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"dagger/check-io-uring/internal/dagger"
)

type CheckIoUring struct{}

func (m *CheckIoUring) Foo(
	source *dagger.Directory,
	platform dagger.Platform,
) *dagger.Container {
	ctr := dag.
		Container(dagger.ContainerOpts{
			Platform: platform,
		}).
		Build(source).
		WithExec([]string{"io_uring-cp", "/foo", "/home/vincent/foo"}, dagger.ContainerWithExecOpts{
			// InsecureRootCapabilities: true,
		}).
		WithExec([]string{"wc", "-c", "/home/vincent/foo"})

	return ctr
}
