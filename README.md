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
	bitLen := byte(6)                                                                   
	mask := byte(1<<bitLen - 1)

	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i, _ := range bytes {
		bytes[i] = alphabet[bytes[i]&mask]
	}

	return string(bytes[:]), nil

}
```

usage
```
Goid64(21) => XjhI1USXYwb8HdTF6SIYI
```

