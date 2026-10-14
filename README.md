# bench

How long a GitHub Actions job takes to run the same `go test` on Linux,
macOS and Windows runners, on an OpenBSD -current VM from
[chroot.dev](https://chroot.dev) with
[chroot-action](https://github.com/chroot-dev/chroot-action), and on
OpenBSD 7.9 in QEMU on the runner with
[cross-platform-actions](https://github.com/cross-platform-actions/action).

The test is the one from
[chroot-dev/example](https://github.com/chroot-dev/example): it only runs
on OpenBSD and is skipped elsewhere, so the times are what each platform
takes to get a job to the point of running Go.

The workflow is [.github/workflows/bench.yml](.github/workflows/bench.yml),
started by hand.

## License

ISC, see [LICENSE](LICENSE).
