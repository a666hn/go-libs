package randomize

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type randomizeOptions struct {
	length int
	rule   string
}

type RandomOptions func(r *randomizeOptions)

func SetLength(v int) RandomOptions {
	return func(r *randomizeOptions) {
		r.length = v
	}
}

type IRandomize interface {
	Random() string
	RandomNumber() int
	RandomAlphabet() string
	RandomNumberString() string
}

func NewRandomize(opts ...RandomOptions) IRandomize {
	o := &randomizeOptions{}
	o.length = 15
	o.rule = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_+=!@#$%^&"

	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (rn *randomizeOptions) Random() string {
	j := random(rn, rn.rule)
	return j
}

func (rn *randomizeOptions) RandomNumber() int {
	r := "1234567890"
	j := random(rn, r)
	k, _ := strconv.Atoi(j)

	return k
}

func (rn *randomizeOptions) RandomAlphabet() string {
	r := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	j := random(rn, r)

	return j
}

func (rn *randomizeOptions) RandomNumberString() string {
	r := "0123456789"
	j := random(rn, r)

	return j
}

func random(
	rn *randomizeOptions,
	rule string,
) string {
	b := make([]rune, rn.length)
	r := []rune(rule)
	for i := range b {
		b[i] = r[rand.Intn(len(rule))]
	}

	return string(b)
}
