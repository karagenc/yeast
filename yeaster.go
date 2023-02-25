package yeast

import (
	"fmt"
	"math"
	"sync"
	"time"
)

const Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

var m = make(map[rune]int, len(Alphabet))

func init() {
	for i, c := range Alphabet {
		m[c] = i
	}
}

type Yeaster struct {
	seed int64
	prev string
	mu   sync.Mutex
}

func New() *Yeaster {
	return new(Yeaster)
}

func (y *Yeaster) Yeast() string {
	y.mu.Lock()
	defer y.mu.Unlock()
	now := y.Encode(time.Now().Unix())

	if now != y.prev {
		y.seed = 0
		y.prev = now
		return now
	}

	y.seed += 1
	return now + "." + y.Encode(y.seed)
}

func (y *Yeaster) Encode(num int64) string {
	encoded := ""
	length := int64(len(Alphabet))
	for {
		encoded = string(Alphabet[num%length]) + encoded
		num = int64(math.Floor(float64(num) / float64(length)))
		if num <= 0 {
			break
		}
	}
	return encoded
}

var ErrInvalidCharacter = fmt.Errorf("yeast: invalid character")

func (y *Yeaster) Decode(s string) (int64, error) {
	var (
		decoded int64
		length  = int64(len(Alphabet))
	)
	for _, c := range s {
		m, ok := m[c]
		if !ok {
			return 0, ErrInvalidCharacter
		}
		decoded = decoded*length + int64(m)
	}
	return decoded, nil
}
