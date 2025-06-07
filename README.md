
How to use

1 - Install the dependencies and compile the code

```bash
go build main.go
```

2 - Create an entry in polybar config

```ini
[module/mtitle]
type = custom/script
format-foreground = ${colors.primary}
format = "<label>"
exec = /path/to/executable/main
interval = 1
```

3 - Set the module's location (left, right or center)

```
modules-center = mtitle
```

MPRIS D-BUS specification:
[https://specifications.freedesktop.org/mpris-spec/latest/Media_Player.html
](https://specifications.freedesktop.org/mpris-spec/latest/Player_Interface.html)
