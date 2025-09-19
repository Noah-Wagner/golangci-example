This repo exists to showcase a minimal repro of an inconsistency between using a golangci-lint linter library directly and through golangci-lint itself.

## Direct library usage

```
go-check-sumtype ./...
/Users/nwagner/sources/golang-bug-repro/cmd/example/main.go:14:2: exhaustiveness check failed for sum type "ExampleSum" (from /Users/nwagner/sources/golang-bug-repro/internal/example.go:8:6): missing cases for TypeA, TypeB, TypeC
```

## golangcilint-run usage

```
golangci-lint run --enable-only gochecksumtype --no-config --verbose
INFO golangci-lint has version 2.4.0 built with go1.25.0 from 43d0339 on 2025-08-13T20:45:00Z
INFO [config_reader] Module name "github.com/Noah-Wagner/example"
INFO maxprocs: Leaving GOMAXPROCS=10: CPU quota undefined
INFO [goenv] Read go env for 11.148292ms: map[string]string{"GOCACHE":"/Users/nwagner/Library/Caches/go-build", "GOROOT":"/opt/homebrew/Cellar/go/1.25.0/libexec"}
INFO [lintersdb] Active 1 linters: [gochecksumtype]
INFO [loader] Go packages loading at mode 8767 (compiled_files|exports_file|imports|name|types_sizes|deps|files) took 223.95025ms
INFO [runner/filename_unadjuster] Pre-built 0 adjustments in 221.167µs
INFO [linters_context/goanalysis] analyzers took 791.834µs with top 10 stages: gochecksumtype: 605.667µs, typecheck: 186.167µs
INFO [runner] processing took 1.79µs with stages: max_same_issues: 333ns, exclusion_paths: 208ns, source_code: 167ns, exclusion_rules: 167ns, max_from_linter: 166ns, nolint_filter: 166ns, path_shortener: 83ns, max_per_file_from_linter: 42ns, path_absoluter: 42ns, path_relativity: 42ns, cgo: 42ns, uniq_by_line: 42ns, diff: 42ns, invalid_issue: 42ns, path_prettifier: 42ns, generated_file_filter: 41ns, fixer: 41ns, sort_results: 41ns, severity-rules: 41ns, filename_unadjuster: 0s
INFO [runner] linters took 62.741875ms with stages: goanalysis_metalinter: 62.715334ms
0 issues.
INFO File cache stats: 0 entries of total size 0B
INFO Memory: 4 samples, avg is 40.3MB, max is 55.6MB
INFO Execution took 298.538959ms
```

No findings.
