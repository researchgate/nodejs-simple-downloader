# nodejs-simple-downloader

This is a tool that simply downloads the desired version of Node.js and unpacks
it.

> ⚠️ &nbsp; **Node.JS Apple Silicon support —**
> Due to the lack of official binaries for `darwin-arm64` (Apple Silicon) architectures for Node.JS versions before 16.0.0, this tool **does not** support downloading any `darwin-arm64` binaries. Instead, it will download the binaries for the `darwin-x64` (Intel Macs) architecture regardless of if the requested version supports Apple Silicon.

## FAQ

### Why?

Although there are existing solutions like `nvm`, `n`, `nodenv` and potentially
more, all these solutions focus on installing nodejs globally and making it
available in the shell environment. In contrast what `node-simple-downloader` does is
simply downloading and extracting the desired Node.js version to a specified directory.
How the downloaded Node.js version is then used is totally up to the user.

The use-case that triggered the creation of this tool is that all our nodejs applications
specify themselves which version of Node.js they need. This allows app developers to take
care of maintaining the local Node.js installation without involvement of system operators.

### Are there more details on how the above mentioned setup looks like?

Not yet. We are currently only planing on migrating away from
[researchgate/nodejs-installer](https://github.com/researchgate/nodejs-installer/tree/researchgate)
and haven't finalized how it will work in the end. Once this is done we will share more details.

### Okay, but why golang?

I always wanted to learn `go` and the scope of this tool seemed doable for
someone who hasn't worked with `go` before. So bear in mind that the code is probably not
structured and written in the best way possible. If you have suggestions feel
free to open PRs. I'm happy to learn how it can be done better.

The other decision to go with `go` is that having a single executable without a requirement for a
runtime, interpreter or package manager should make the setup easier in the end.

## Usage

### Node.js

#### With specific version

```bash
nsd nodejs /path/ --version 12.15.0
```

#### With specific version from file

`package.json` (exact version in engines field) and `.nvmrc` are currently supported.

```bash
nsd nodejs /path/ --version-from-file ./.nvmrc
```

### Yarn (Currently only v1 supported)

#### Full distribution

```bash
nsd yarn /path/ --version 1.22.5
```

#### Single file distribution

This will download the single javascript file and move it into the download directory with the name supplied to `--single-file`

```bash
nsd yarn /path/ --version 1.22.5 --single-file yarn
```

#### With specific version from file

`package.json` (exact version in engines field) is currently supported.

```bash
nsd nodejs /path/ --version-from-file ./package.json
```

## How to build

> Requires go version 1.15 or newer

- Checkout project
- Run `go mod tidy`
- Run `cd nsd && go build`

## How to run tests

`cd nsd && go test ./...`
