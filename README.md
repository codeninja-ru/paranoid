# paranoid
paranoid UID generator

Just copy and paste

## Go
```go
import (
	"crypto/rand"
)

func Goid64(length int) (string, error) {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~" // the length must be 64 bytes

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, _ := range bytes {
		bytes[i] = alphabet[bytes[i]&63]
	}

	return string(bytes[:]), nil

}

```

usage
```
uid, _ := Goid64(21)
uid => XjhI1USXYwb8HdTF6SIYI
```

