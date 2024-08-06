# ghomfc

GitHub Organization Members' Followers Counter.

## Example

List members in [kubernetes](https://github.com/kubernetes).

```console
$ date +%Y-%m-%dT%H:%M:%S%Z
2024-08-06T20:54:30JST

$ ghomfc run -format table kubernetes | head -n 12
```

Rank | Login (Name) | Number of Followers
--- | --- | ---
1 | [ahmetb (Ahmet Alp Balkan)](https://github.com/ahmetb) | 8082
2 | [terrytangyuan (Yuan Tang)](https://github.com/terrytangyuan) | 3334
3 | [feiskyer (Pengfei Ni)](https://github.com/feiskyer) | 2479
4 | [nikhita (Nikhita Raghunath)](https://github.com/nikhita) | 1903
5 | [thockin (Tim Hockin)](https://github.com/thockin) | 1818
6 | [Pradumnasaraf (Pradumna Saraf)](https://github.com/Pradumnasaraf) | 1804
7 | [gyuho (Gyuho Lee)](https://github.com/gyuho) | 1453
8 | [AkihiroSuda (Akihiro Suda)](https://github.com/AkihiroSuda) | 1450
9 | [Xunzhuo (Xunzhuo)](https://github.com/Xunzhuo) | 1193
10 | [liggitt (Jordan Liggitt)](https://github.com/liggitt) | 1172

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
