# nodejs-simple-downloader

This is a tool that simply downloads the desired version of Node.js and unpacks
it.

> :warning: **This tool is in a very early stage and is not polished for easy usage yet.**

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
[researchgate/nodejs-installer](https://github.com/researchgate/nodejs-installer/tree/researchgate`)
and haven't finalized how it will work in the end. Once this is done we will share more details.

### Okay, but why golang?

I always wanted to learn `go` and the scope of this tool seemed doable for
someone who hasn't worked with `go` before. So bear in mind that the code is probably not
structured and written in the best way possible. If you have suggestions feel
free to open PRs. I'm happy to learn how it can be done better.

## Usage

#### With specific version

```bash
nsd nodejs /path/ --version 12.15.0
```

#### With specific version from file

```bash
nsd nodejs /path/ --from-file ./node_version
```

#### No specific version, will try reading from CWD/.nvmrc

```bash
nsd nodejs /path/
```

## How to build

- Checkout project
- Run `go mod tidy`
- Run `cd nsd && go build`
