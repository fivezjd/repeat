使用 `import  github.com/stretchr/testify/assert ` 进行单元测试

```
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  var a string = "Hello"
  var b string = "Hello"

  assert.Equal(t, a, b, "The two words should be the same.")

}
```

[文档地址](https://pkg.go.dev/github.com/stretchr/testify/assert)

