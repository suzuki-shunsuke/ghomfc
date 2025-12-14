# ghomfc

[![DeepWiki](https://img.shields.io/badge/DeepWiki-suzuki--shunsuke%2Fghomfc-blue.svg?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAyCAYAAAAnWDnqAAAAAXNSR0IArs4c6QAAA05JREFUaEPtmUtyEzEQhtWTQyQLHNak2AB7ZnyXZMEjXMGeK/AIi+QuHrMnbChYY7MIh8g01fJoopFb0uhhEqqcbWTp06/uv1saEDv4O3n3dV60RfP947Mm9/SQc0ICFQgzfc4CYZoTPAswgSJCCUJUnAAoRHOAUOcATwbmVLWdGoH//PB8mnKqScAhsD0kYP3j/Yt5LPQe2KvcXmGvRHcDnpxfL2zOYJ1mFwrryWTz0advv1Ut4CJgf5uhDuDj5eUcAUoahrdY/56ebRWeraTjMt/00Sh3UDtjgHtQNHwcRGOC98BJEAEymycmYcWwOprTgcB6VZ5JK5TAJ+fXGLBm3FDAmn6oPPjR4rKCAoJCal2eAiQp2x0vxTPB3ALO2CRkwmDy5WohzBDwSEFKRwPbknEggCPB/imwrycgxX2NzoMCHhPkDwqYMr9tRcP5qNrMZHkVnOjRMWwLCcr8ohBVb1OMjxLwGCvjTikrsBOiA6fNyCrm8V1rP93iVPpwaE+gO0SsWmPiXB+jikdf6SizrT5qKasx5j8ABbHpFTx+vFXp9EnYQmLx02h1QTTrl6eDqxLnGjporxl3NL3agEvXdT0WmEost648sQOYAeJS9Q7bfUVoMGnjo4AZdUMQku50McDcMWcBPvr0SzbTAFDfvJqwLzgxwATnCgnp4wDl6Aa+Ax283gghmj+vj7feE2KBBRMW3FzOpLOADl0Isb5587h/U4gGvkt5v60Z1VLG8BhYjbzRwyQZemwAd6cCR5/XFWLYZRIMpX39AR0tjaGGiGzLVyhse5C9RKC6ai42ppWPKiBagOvaYk8lO7DajerabOZP46Lby5wKjw1HCRx7p9sVMOWGzb/vA1hwiWc6jm3MvQDTogQkiqIhJV0nBQBTU+3okKCFDy9WwferkHjtxib7t3xIUQtHxnIwtx4mpg26/HfwVNVDb4oI9RHmx5WGelRVlrtiw43zboCLaxv46AZeB3IlTkwouebTr1y2NjSpHz68WNFjHvupy3q8TFn3Hos2IAk4Ju5dCo8B3wP7VPr/FGaKiG+T+v+TQqIrOqMTL1VdWV1DdmcbO8KXBz6esmYWYKPwDL5b5FA1a0hwapHiom0r/cKaoqr+27/XcrS5UwSMbQAAAABJRU5ErkJggg==)](https://deepwiki.com/suzuki-shunsuke/ghomfc)

[MIT LICENSE](LICENSE) | [Install](INSTALL.md) | [Usage](USAGE.md)

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

## Usage

See also [USAGE.md](USAGE.md)

```sh
ghomfc run [-format <json|table>] "<GitHub Organization name>"
```

## Environment variables

- `GITHUB_TOKEN`: GitHub Access Token. This is required to call GitHub GraphQL API

## LICENSE

[MIT](LICENSE)
