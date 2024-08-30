# dagger insecure root capabilities

This reproducer contains a Dockerfile that buids an image containing the binary `io_uring-cp` that
can be used to test if io_uring works or not.

*Note*: this assumes linux/arm64 because I'm testing this on a M1 MacBook, change the architecture if needed.
