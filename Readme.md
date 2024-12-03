## Install
```
go get github.com/hogesako/go-generics-syncmap
```

## Usage
ISUCON用なので内部で初期化チェックはしていない。必ずInitを呼ぶ必要がある
```
import (
	syncmap "github.com/hogesako/go-generics-syncmap"
)

rwmap := syncmap.RWSyncMap[int, string]
rwmap.Init()

rwmap.Store(1, "one")
require.Equal(t, 1, rwmap.Len())
val, ok := rwmap.Load(1)
```