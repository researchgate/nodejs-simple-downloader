# nodejs-simple-downloader

This is a tool that simply downloads the desired version of nodejs and unpacks
it.

## FAQ

### Why?

Although there are existing solutions like `nvm`, `n`, `nodenv `and potentially
more, all these solutions focus on installing nodejs globally and making it
available in the environment. In contrast what `node-simple-downloader` does is
simply download and extract the desired nodejs version to a specified directory.
How the downloaded NodeJS is then used is totally up to the user. The use-case
that triggered the creation of this tool is that all our nodejs applications
specify themselves which version of NodeJS they run and also take care of NodeJS
updates. This allows app developers to take care of maintaining the local NodeJS
installation without involvement of system operators.

### Okay, but why golang?

I always wanted to learn `go` and the scope of this tool seemed doable for
someone who hasn't worked with `go` before. So bear in mind that the code is probably not
structured and written in the best way possible. If you have suggestions feel
free to open PRs. I'm happy to learn how it can be done better.

## How to build

- Checkout project
- Run `go mod tidy`
- Run `cd nsd && go build`
