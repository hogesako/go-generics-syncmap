### Usage
ISUCON用なので内部で初期化チェックはしていない。必ずInitを呼ぶ必要がある
```
var rwmap RWSyncMap[int, string]
rwmap.Init()

rwmap.Store(1, "one")
require.Equal(t, 1, rwmap.Len())
val, ok := rwmap.Load(1)
```