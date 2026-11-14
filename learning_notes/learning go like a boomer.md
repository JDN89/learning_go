
## NOTES along the way
[[notes of learning go]]
## My setup 
 - Neovim
## Study notes
[[Golang notes]]
## Golang blog notes
[[golang blog articles]]
## Read 
- [ ] A tour of go
- [ ] o reailys book 
# Projects
Tips:
- Add tests with the `testing` package
- use go dock and read the source of packages I use often

- [ ] [[Advent of code 2020 to learn the basics]] 
	- [x] done until day 7  continue after grind 75
- [x]  [backend basics article](https://eblog.fly.dev/backendbasics.html)
- [x] [how i use htmx with go](https://www.alexedwards.net/blog/how-i-use-htmx-with-go)
- [ ] [[Grind 75 leetcode in go]]
- [ ] [[Read about Go it's GC, optimization techniques and profiling]]
- [ ] [[ring buffer in go]]
- [ ] [[ring buffer in golang]]
- [ ] [[Network server tcp]]
- [ ] [[make my own SSG (should be fun)]]
- [ ] [[implement compression algo - huffman encoding]]
- [ ] [[implementing DEFLATE (what git uses)]]
- [ ] [[let's go further REST and MVC impl]]
## EASY
- Word / Line Counter
	Count lines, words, and bytes in one or more files (like a mini wc). Support flags for which stats to show.
	Key packages: os, io, bufio, flag, fmt
- Text Replacer / Searcher
	Find and replace a string (or simple pattern) in files, or just list matching lines.
	Key packages: os, bufio, strings, regexp, path/filepath
- Simple JSON Pretty-Printer / Filter
	Read JSON from a file or stdin, pretty-print it, or extract one field and print it.
	Key packages: encoding/json, os, io, flag
- CSV to JSON Converter
	Convert a CSV file to JSON (or vice versa) with basic options (headers, delimiter).
	Key packages: encoding/csv, encoding/json, os, flag

## MEDIUM

These introduce concurrency, HTTP, and more real-world patterns.

- Directory Size Calculator
	Walk a directory tree and report total size, number of files, and optionally the largest files.
	Key packages: os, path/filepath, sort
- Parallel File Hasher
	Compute SHA-256 (or other) hashes of all files in a directory using a worker pool.
	Key packages: crypto/sha256, os, path/filepath, sync, context
- URL Status Checker
	Take a list of URLs and check them concurrently, reporting status code and response time.
	Key packages: net/http, sync, context, time, flag
- In-Memory URL Shortener
	Simple HTTP server with endpoints to create a short link and to redirect. Store mappings in a map (protect with mutex).
	Key packages: net/http, sync, encoding/json, time
- Log Parser / Top-N Reporter
	Parse an access log and report the top N IPs, status codes, or user agents.
	Key packages: bufio, strings, regexp, sort, time
- Simple Static File Server
	Serve files from a directory with basic features (directory listing, custom headers, or simple auth).
	Key packages: net/http, os, path/filepath, mime

## HARDER

These combine multiple concepts, more networking, or careful concurrency/context usage.

- Concurrent Web Crawler
	Crawl a site starting from one URL, collect unique links up to a certain depth, with rate limiting and a worker pool.
	Key packages: net/http, net/url, sync, context, regexp or golang.org/x/net/html (stdlib-adjacent)
- TCP Chat Server
	Multi-client TCP chat: clients connect, send messages, and receive broadcasts. Handle disconnects cleanly.
	Key packages: net, bufio, sync, context
- HTTP Client with Retries + Caching
	A reusable client that fetches URLs with exponential backoff, respects context timeouts, and caches responses in memory.
	Key packages: net/http, context, time, sync, encoding/json
- File Watcher + Processor
	Watch a directory (or poll it) and process new/changed files (e.g., hash them or move them). Graceful shutdown.
	Key packages: os, path/filepath, time, context, sync (note: fsnotify is third-party; pure stdlib can poll)
- Minimal Reverse Proxy
	A small HTTP reverse proxy that forwards requests to a backend, adds headers, and logs them.
	Key packages: net/http, net/http/httputil, log/slog, context

# References
[based academy video about learning](https://youtu.be/czYHZRTcMeg?si=XzsnE2ePPVBdwonj)
