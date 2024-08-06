# ghomfc

GitHub Organization Members' Followers Counter.

## Example

List members in [terraform-linters](https://github.com/terraform-linters).

```console
$ date +%Y-%m-%dT%H:%M:%S%Z
2024-08-06T20:13:23JST

$ ghomfc run -format table terraform-linters | head -n 12
```

Rank | Login (Name) | Number of Followers
--- | --- | ---
1 | [chenrui333 (Rui Chen)](https://github.com/chenrui333) | 644
2 | [bendrucker (Ben Drucker)](https://github.com/bendrucker) | 451
3 | [suzuki-shunsuke (Shunsuke Suzuki)](https://github.com/suzuki-shunsuke) | 319
4 | [jpreese (John Reese)](https://github.com/jpreese) | 125
5 | [lonegunmanb ()](https://github.com/lonegunmanb) | 120
6 | [PatMyron (Pat Myron)](https://github.com/PatMyron) | 96
7 | [jayzes (Jay Zeschin)](https://github.com/jayzes) | 93
8 | [circa10a (Caleb Lemoine)](https://github.com/circa10a) | 89
9 | [wata727 (Kazuma Watanabe)](https://github.com/wata727) | 85
10 | [pd (Kyle Hargraves)](https://github.com/pd) | 78

## Install

`ghomfc` is a single binary written in Go.
So you only need to put the executable binary into `$PATH`.

1. [Homebrew](https://brew.sh/)

```sh
brew install suzuki-shunsuke/ghomfc/ghomfc
```

2. [Scoop](https://scoop.sh/)

```sh
scoop bucket add suzuki-shunsuke https://github.com/suzuki-shunsuke/scoop-bucket
scoop install ghomfc
```

3. [aqua](https://aquaproj.github.io/)

```sh
aqua g -i suzuki-shunsuke/ghomfc
```

4. Download a prebuilt binary from [GitHub Releases](https://github.com/suzuki-shunsuke/ghomfc/releases) and install it into `$PATH`

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
