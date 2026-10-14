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

## Results, 2026-10-14

18 runs one after the other, 16:26 to 17:15 UTC. Seconds from the job
API's timestamps, which have 1 s resolution; median (min-max).

| Job | Job time | Getting Go | `go test` |
| --- | --- | --- | --- |
| chroot.dev, copy of `ci-go` | 11 (10-17) | 2, the copy | 5 |
| ubuntu-24.04 | 18 (12-27) | 11, setup-go | 0, skipped |
| macos-15 | 31 (18-45) | 20, setup-go | 1, skipped |
| windows-2025 | 37.5 (26-49) | 20, setup-go | 1, skipped |
| chroot.dev, new VM, 2 cpus, 1 GiB (6 runs) | 71 (70-77) | 3 + 40.5 pkg_add | 22 |
| QEMU, OpenBSD 7.9, 2 cpus, 6 GB | 75.5 (60-141) | 29 + 25.5 pkg_add | 16 |
| chroot.dev, new VM, 1 cpu, 128 MiB (6 runs) | 264.5 (222-319) | 2 + 54 pkg_add | 204.5 |

- `ci-go` had 1 cpu and 128 MiB during these runs, with Go 1.27.2 and
  warm module and build caches.
- The new chroot.dev VMs installed Go from ftp.openbsd.org: that day,
  cdn.openbsd.org served the VMs a snapshot listing that named a package
  already gone.
- Go: 1.26.0 on the GitHub runners (setup-go), 1.26.2 on OpenBSD 7.9,
  1.27.2 on OpenBSD -current.

Every job and step is in [results/2026-10-14.tsv](results/2026-10-14.tsv):
the first 6 runs had no new chroot.dev VM job, the next 6 ran it with
1 cpu and 128 MiB, the last 6 with 2 cpus and 1 GiB. Left out: run
38599379607, where the new VM failed on the stale listing, and run
38599591751, cancelled.

## License

ISC, see [LICENSE](LICENSE).
