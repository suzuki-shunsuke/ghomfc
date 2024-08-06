# ghomfc

GitHub Organization Members' Followers Counter.

## Install

`ghomfc` is a single binary written in Go.
So you only need to put the executable binary into `$PATH`.

1. [Homebrew](https://brew.sh/)

Coming soon.

2. [Scoop](https://scoop.sh/)

Coming soon.

3. [aqua](https://aquaproj.github.io/)

Coming soon.

4. Download a prebuilt binary from [GitHub Releases](https://github.com/suzuki-shunsuke/ghomfc/releases) and install it into `$PATH`

Coming soon.

5. Go

```sh
go install github.com/suzuki-shunsuke/ghomfc/cmd/ghomfc@latest
```

## Usage

```sh
ghomfc run [-format <json|table>] "<GitHub Organization name>"
```

## Environment variables

- `GITHUB_TOKEN`: GitHub Access Token. This is required to call GitHub GraphQL API

## LICENSE

[MIT](LICENSE)
