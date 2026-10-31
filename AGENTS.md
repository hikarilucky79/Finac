# Repository Guidelines

## Project Structure

This repository contains the Portuguese design contract in `arquitetura-final-trading.md` and a Go test harness. Keep suites separated under `tests/development`, `tests/preview`, and `tests/production`; shared test helpers belong in `tests/internal/contract`.

## Build and Test

Run every suite with `go test -count=1 -v ./tests/...`. Run one suite with `go test ./tests/development/...`, replacing `development` with `preview` or `production`. Run `go vet ./tests/...` before submitting Go changes. GitHub Actions runs the suites and commit-prefix check on pushes and pull requests. Configure these checks as required branch protections so a failed check blocks merging.

## Style and Test Design

Format Go files with `gofmt`. Use descriptive test names beginning with `Test`, keep helpers small, and avoid external dependencies. Each suite has at least 50 named table-driven cases using global market contexts. They validate rules in the written contract, not runtime behavior. Keep the Production suite offline: never use live broker connections, credentials, or real orders. When application code exists, add behavioral tests with fakes and sandbox evidence without weakening these contract checks.

## Safety and Phase Gates

Preserve fail-closed behavior, deterministic order decisions, broker-position reconciliation after disconnects, idempotent orders, and the independent Position Guardian and human kill switch. Keep LLM analysis outside the order path. Do not advance phases until their documented gates pass, including Phase 0 sandbox checks and Phase 1's 200 demo signals with positive expectancy after costs. Whenever a use case or test fails, correct the underlying error or bug and rerun the affected tests. Repeat this cycle until every test suite passes; do not skip, mask, or accept a failing test.

## Commits and Pull Requests

Every commit subject must start with a short bracketed prefix accepted by CI: `[test]`, `[docs]`, `[ci]`, `[fix]`, `[feat]`, `[refactor]`, or `[chore]`. Commit every repository change before considering the work complete; leave no intended changes uncommitted or untracked. Keep commits focused. Pull requests should describe the rules covered and report test commands and results. Never claim live-trading safety based on document-only checks.
