# Ferox

<p align="center">
  <img src="docs/assets/forex (2).png" alt="Ferox — Behavioral Web Reconnaissance & Fuzzing Engine">
</p>

<p align="center">
  <strong>Behavioral Web Reconnaissance & Fuzzing Engine</strong>
</p>

<p align="center">
  Built in Go • Open Source • Actively Developing
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#architecture">Architecture</a> •
  <a href="#roadmap">Roadmap</a> •
  <a href="#security">Security</a>
</p>

---

## What is Ferox?

**Ferox is a Go-based behavioral web reconnaissance and fuzzing engine designed to find meaningful response differences instead of simply producing more HTTP results.**

Traditional fuzzing can tell you:

```text
/admin      → 404
/login      → 200
/api        → 403
```

Ferox aims to ask a harder question:

> **Does this response behave differently enough from the expected baseline to deserve investigation?**

The project focuses on turning large-scale reconnaissance into a smaller, more useful set of signals for security researchers and penetration testers.

> ⚠️ **Authorized testing only**
>
> Use Ferox only against infrastructure you own or targets explicitly authorized for security testing, such as systems covered by a bug bounty or penetration-testing engagement.

---

## Why Ferox?

A large reconnaissance scan can produce thousands of responses.

The challenge isn't always finding more responses.

The challenge is identifying the responses that **actually matter**.

Ferox is being designed around:

* Behavioral response analysis
* Baseline-aware comparison
* Structural response fingerprinting
* Anomaly detection
* WAF/block-page awareness
* Discovery-driven follow-up
* High-concurrency scanning
* Explainable results

The long-term goal is simple:

```text
Thousands of requests
        ↓
Response analysis
        ↓
Behavioral signals
        ↓
Ranked findings
        ↓
Human investigation
```

---

## Features

### Core Reconnaissance

* Wordlist-based web discovery
* `FUZZ` substitution
* `FUZZ2` secondary wordlist support
* Extension expansion
* Concurrent request processing
* Global request-rate limiting
* Status-code filtering
* JSONL output

### Behavioral Analysis

Ferox is designed to go beyond status-code filtering.

The response analysis pipeline includes:

```text
Response
   ↓
Normalization
   ↓
Structural Fingerprint
   ↓
Baseline Comparison
   ↓
Clustering
   ↓
Anomaly Score
   ↓
Classification
```

Potential classifications include:

```text
baseline
anomaly
waf_suspected
```

### Security-Aware Defaults

Ferox is designed to fail safely where possible:

* TLS verification enabled by default
* Redirects not automatically followed
* Request timeout limits
* Response body size limits
* Concurrency sanity checks
* Streamed primary wordlists
* Scan output excluded from Git
* Graceful Ctrl+C handling

See [`docs/SECURITY_PRACTICES.md`](docs/SECURITY_PRACTICES.md) for the reasoning behind these decisions.

---

## Installation

### Requirements

* Go 1.22+

### Build from source

```bash
git clone https://github.com/YOUR_USERNAME/ferox.git
cd ferox

go build -o ferox ./cmd/ferox
```

Run:

```bash
./ferox --help
```

---

## Usage

### Basic discovery

```bash
./ferox \
  -u "https://target.example/FUZZ" \
  -w wordlists/sample-small.txt
```

### Concurrent scanning

```bash
./ferox \
  -u "https://target.example/FUZZ" \
  -w wordlist.txt \
  -c 40
```

### Extension discovery

```bash
./ferox \
  -u "https://target.example/FUZZ" \
  -w wordlist.txt \
  -e .php,.bak,.old
```

### Two-position fuzzing

```bash
./ferox \
  -u "https://target.example/FUZZ?param=FUZZ2" \
  -w dirs.txt \
  -w2 params.txt
```

### Rate-limited scanning

```bash
./ferox \
  -u "https://target.example/FUZZ" \
  -w wordlist.txt \
  -rps 20
```

### JSONL output

```bash
./ferox \
  -u "https://target.example/FUZZ" \
  -w wordlist.txt \
  -o results.jsonl
```

> Only scan targets for which you have explicit authorization.

---

## Architecture

Current high-level pipeline:

```text
                    Wordlist
                       │
                       ▼
              ┌─────────────────┐
              │ Job Generation  │
              │ FUZZ/FUZZ2/ext  │
              └────────┬────────┘
                       │
                       ▼
                  Job Channel
                       │
                       ▼
              ┌─────────────────┐
              │ Worker Pool     │
              │ + Rate Limiter  │
              └────────┬────────┘
                       │
                       ▼
                  HTTP Client
                       │
                       ▼
                    Response
                       │
                       ▼
              Behavioral Analysis
                       │
                ┌──────┴──────┐
                ▼             ▼
             Baseline       Anomaly
                │             │
                └──────┬──────┘
                       ▼
                    Output
```

Ferox separates **request execution** from **response intelligence**, allowing the behavioral analysis layer to evolve independently from the scanning engine.

See [`docs/ARCHITECTURE.md`](docs/ARCHITECTURE.md) for deeper design decisions.

---

## Roadmap

Ferox is being developed incrementally.

### v0.1 — Core Engine

* [x] Wordlist parsing
* [x] HTTP requests
* [x] Basic filtering

### v0.2 — Concurrent Engine

* [x] Worker pool
* [x] Rate limiting
* [x] `FUZZ` / `FUZZ2`
* [x] Extension expansion
* [x] JSONL output
* [x] Panic recovery
* [x] Test coverage

### v0.3 — Smart Response Diffing

* [ ] Structural fingerprinting
* [ ] Response normalization
* [ ] Baseline calibration
* [ ] Response clustering
* [ ] Anomaly scoring
* [ ] WAF block-page detection

### v0.4 — Cascading Discovery

* [ ] Discovery-triggered sub-jobs
* [ ] JavaScript route extraction
* [ ] JSON endpoint extraction
* [ ] ID-like parameter discovery
* [ ] Discovery graph

### v0.5 — Adaptive Reconnaissance

* [ ] Adaptive rate limiting
* [ ] WAF-aware backoff
* [ ] Resume support
* [ ] TUI
* [ ] Discovery graph export

### v0.6 — Coverage & Methodology

* [ ] Coverage tracking
* [ ] Scan methodology report
* [ ] WAF events
* [ ] Request statistics
* [ ] Anomaly summary

---

## Project Structure

```text
ferox/
├── cmd/
│   └── ferox/
│       └── main.go
│
├── internal/
│   ├── wordlist/
│   ├── httpclient/
│   ├── engine/
│   └── output/
│
├── docs/
│   ├── ARCHITECTURE.md
│   ├── SECURITY_PRACTICES.md
│   └── LEARNING_PLAN.md
│
├── wordlists/
│
├── .github/
│   └── workflows/
│
├── .gitignore
├── LICENSE
├── README.md
└── go.mod
```

---

## Development

Run the test suite:

```bash
go test ./...
```

Run with the race detector:

```bash
go test ./... -race
```

Run static analysis:

```bash
go vet ./...
```

---

## Security

Ferox is a security testing tool.

Please read [`docs/SECURITY_PRACTICES.md`](docs/SECURITY_PRACTICES.md) before using or contributing to the project.

Important design principles include:

* Safe-by-default HTTP behavior
* Explicit TLS exceptions
* Controlled concurrency
* Request timeouts
* Response-size limits
* No automatic redirects
* No secrets in source control
* Scan output treated as sensitive data

---

## Responsible Use

Ferox is intended for:

* Authorized penetration testing
* Bug bounty programs
* Security research on owned infrastructure
* Local security labs
* CTF environments

**Do not use Ferox to scan systems without authorization.**

---

## Contributing

Ferox is being developed in public.

Contributions, issues, discussions, and security-focused feedback are welcome.

Before contributing:

1. Read the project documentation.
2. Understand the security implications of your change.
3. Add or update tests where appropriate.
4. Keep security-sensitive behavior explicit.
5. Document important design decisions.

---

## License

MIT License.

See [`LICENSE`](LICENSE) for details.

---

<p align="center">
  <strong>Ferox</strong>
  <br>
  Behavioral reconnaissance. Less noise. More signal.
</p>
