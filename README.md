# paranoid
paranoid UID generator

Simple UID generator with uniform distribution

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

## JavaScript
```javascript
function jsid64(length) {
    length = length || 21;
    var crypto = self.crypto || self.msCrypto;
    var alphabet = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_~';
    var bytes = new Int8Array(length);
    crypto.getRandomValues(bytes);
    var result = "";
    for(var i = 0; i < bytes.length; i++) {
        result = result + alphabet[bytes[i]&63];
    }

    return result;
}
```

usage
```
jsid(21)
=> "VHwd5~gXdo64UiK3uR_Zf"
```

## How it works

