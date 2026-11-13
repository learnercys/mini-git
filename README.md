# mini-git

A small version-control system written from scratch in Go, built to understand how Git works under the hood: content-addressable storage, trees, commits, refs, and branches.

```
$ mini-git init
Initialized empty mini-git repository in /tmp/demo/.mini-git
$ echo hello > a.txt
$ mini-git add a.txt
$ mini-git commit "first commit"
$ mini-git log
commit c89d9a2aef41f34030df66f303bf7bd2176aeffc
Author: John Doe <john@example.com>
Date: 2026-11-13 05:55:43 +0000 UTC

    first commit
$ mini-git branch feature
Created branch 'feature'
```

## Commands

| Command | What it does |
|---|---|
| `init` | Creates `.mini-git/` with `objects/`, `refs/`, and a `HEAD` pointing at `master` |
| `add <file>` | Hashes the file into a blob, stores it, and records it in the index |
| `commit <message>` | Builds a tree from the index, writes a commit object, and moves the branch ref |
| `log` | Walks the parent chain from `HEAD` |
| `status` | Compares index ↔ last commit (staged) and working tree ↔ last commit (unstaged) |
| `branch [name]` | Lists branches or creates one at the current commit |
| `checkout <branch>` | Points `HEAD` at the branch and rebuilds the working tree from its commit |
| `merge <branch>` | Fast-forward merge (checks ancestry by walking parents) |

## How it works

```
.mini-git/
├── HEAD                 ref: refs/heads/master
├── index                <blob-hash> <path>, one per line
├── refs/heads/<branch>  commit hash
└── objects/ab/cdef…     blob | tree | commit, addressed by SHA-1
```

- **Content-addressable storage.** Every object is stored at `objects/<first 2 hex>/<remaining 38>`, like Git's loose objects. Identical content is stored only once.
- **Three object types.** Each one is serialized as `<type> <size>\0<payload>`:
  - **blob:** raw file content
  - **tree:** `mode type hash\tname` entries
  - **commit:** `tree`, `parent`, `author`, `committer`, blank line, message
- **Refs are just files.** A branch is a file holding a commit hash. `HEAD` is a symbolic ref to a branch, or a raw hash when detached.
- **History is a linked list.** `log` and the fast-forward check in `merge` both walk `parent` pointers.

## Simplifications compared with real Git

These were deliberate, to keep the core ideas visible:

- Objects are stored uncompressed. Git zlib-compresses them.
- Blob hashes cover the content only, so hashes are **not** Git-compatible. Git hashes the header and content together.
- Trees are flat: one root tree holding relative paths, not nested subtrees.
- The index is a plain text file, not Git's binary index format.
- `merge` handles fast-forward only. A three-way merge needs a merge-base search and conflict handling.
- The author is hard-coded. There is no config file yet.

## Run it

Requires Go 1.22+.

```bash
go install github.com/learnercys/mini-git@latest
```

Or from a clone:

```bash
go build -o mini-git . && ./mini-git init
go test ./...
```

## Project layout

```
blob/ tree/ commit/   object types: serialize, deserialize, hash
objects/              read/write objects in the object store
repository/           init and repository-root discovery
commands/             one file per CLI command
```

## What's next

- [ ] Nested trees for subdirectories
- [ ] zlib compression and Git-compatible hashing (so `git cat-file` can read the objects)
- [ ] Three-way merge with a lowest-common-ancestor search
- [ ] `diff` between the working tree and a commit
