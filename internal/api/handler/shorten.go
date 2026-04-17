package handler

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var count uint64 = 0

const Alphabet = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ_abcdefghijkmnopqrstuvwxyz"

const BASE = len(Alphabet)

var keys = [4]uint32{
	0xA3B1C2D3,
	0xF1E2D3C4,
	0x12345678,
	0x9ABCDEF0,
}

func feistel(r uint32, key int) uint32 {
	x := r ^ keys[key]
	x *= 0x45d9f3b
	x ^= x >> 16
	return x
}

func obsufication(baseNumber uint64) uint64 {
	// shift by 32
	L := uint32(baseNumber >> 32)
	// mask statin 32 bit
	R := uint32(baseNumber & 0xffffffff)
	for i := range 4 {
		L, R = R, L^feistel(R, i)
	}
	return (uint64(L) << 32) | uint64(R)
}

func encode(number uint64) string {
	runes := []rune{}
	for number > 0 {
		rem := number % uint64(BASE)
		runes = append(runes, rune(Alphabet[rem]))
		number = number / uint64(BASE)
	}
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[j], runes[i] = runes[i], runes[j]
	}
	return string(runes)
}

func (h *URLHandler) Register(w http.ResponseWriter, r *http.Request) {
	count = atomic.AddUint64(&count, 1)
	fmt.Fprintf(w, "<h1>%s</h1>", encode(obsufication(count)))
}
