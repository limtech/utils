# useful utils for GO

```go
package utils // import "github.com/limtech/utils"

func HttpGet(url string) ([]byte, error)
func HttpPost(url string, data url.Values, headers map[string]string) ([]byte, error)
func HttpPostJson(url string, data interface{}, header map[string]string) ([]byte, error)

func RandomString(n int) string
func SubString(str string, start int, length int) string
```