![](img/Gofare_banner.png)
Golang utility to pretty print Mifare dump, both `1k` and `4k`.

Based on the mighty [`mfdread`](https://github.com/zhovner/mfdread).

## Usage

```sh
$> gofare -dump MifareDump.bin
```

![](img/img0001.png)

To get the color codes:

```sh
$> gofare -dump MifareDump.bin -v
```

![](img/img0002.png)

To save the keys from a dump:

```sh
$> gofare -dump MifareDump.bin -keys
```

![](img/img0003.png)