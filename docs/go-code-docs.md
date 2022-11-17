# modules

I need to use our own fork of giu, since https://github.com/AllenDang/giu/pull/533
is not merged yet

```sh
go mod edit -replace github.com/AllenDang/giu=github.com/TheGreaterHeptavirate/giu@e065652598015e1c1fc21422bc10eb18a9e0a730 # probably @motorola should work, but...
```
