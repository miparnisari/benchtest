## Running benchmarks reliably with GitHub action

This repo shows this workflow:

1. Run benchmarks for your Go code in the `main` branch and in PRs.
2. Compare PR benchmarks against `main` benchmarks:
   1. if no `main` benchmarks are available, just publish benchmarks in the summary of the PR.
   2. if there `main` benchmarks, and they got worse in the PR, **fail the PR**.