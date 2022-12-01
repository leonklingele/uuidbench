package main_test

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"testing"

	fiberutils "github.com/gofiber/utils/v2"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
)

// UUID

func Test_Fiber_UUID(t *testing.T) {
	rid1 := fiberutils.UUID()
	rid2 := fiberutils.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Fiber_UUID(b *testing.B) {
	var rid string
	for n := 0; n < b.N; n++ {
		rid = fiberutils.UUID()
	}
	_ = rid
}

// UUIDv4

func Test_Fiber_UUIDv4(t *testing.T) {
	rid1 := fiberutils.UUIDv4()
	rid2 := fiberutils.UUIDv4()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Fiber_UUIDv4(b *testing.B) {
	var rid string
	for n := 0; n < b.N; n++ {
		rid = fiberutils.UUIDv4()
	}
	_ = rid
}

// Crypto
func Test_Custom_Crypto(t *testing.T) {
	gen, err := newGenerator_Crypto()
	if err != nil {
		t.Fatal(err)
	}

	rid1 := gen.UUID()
	rid2 := gen.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Custom_Crypto(b *testing.B) {
	gen, err := newGenerator_Crypto()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var rid string
	for n := 0; n < b.N; n++ {
		rid = gen.UUID()
	}
	_ = rid
}

const (
	sizeCrypto = 32
)

type GeneratorCrypto struct {
}

func (g *GeneratorCrypto) next() [sizeCrypto]byte {
	var cur [sizeCrypto]byte

	if _, err := rand.Read(cur[:]); err != nil {
		panic(fmt.Errorf("failed to generate next: %w", err))
	}

	return cur
}

func (g *GeneratorCrypto) UUID() string {
	uuid := g.next()

	b := make([]byte, 36)
	_ = hex.Encode(b, uuid[:4])
	b[8] = '-'
	_ = hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	_ = hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	_ = hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	_ = hex.Encode(b[24:], uuid[10:16])

	return string(b)
}

func newGenerator_Crypto() (*GeneratorCrypto, error) {
	return &GeneratorCrypto{}, nil
}

// SHA256

func Test_Custom_SHA256(t *testing.T) {
	gen, err := newGenerator_SHA256()
	if err != nil {
		t.Fatal(err)
	}

	rid1 := gen.UUID()
	rid2 := gen.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Custom_SHA256(b *testing.B) {
	gen, err := newGenerator_SHA256()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var rid string
	for n := 0; n < b.N; n++ {
		rid = gen.UUID()
	}
	_ = rid
}

const (
	sizeSHA256 = sha256.Size
)

type GeneratorSHA256 struct {
	mu sync.Mutex

	cur [sizeSHA256]byte
}

func (g *GeneratorSHA256) next() [sizeSHA256]byte {
	g.mu.Lock()
	cur := g.cur
	cur = sha256.Sum256(cur[:])
	g.cur = cur
	g.mu.Unlock()

	return cur
}

func (g *GeneratorSHA256) UUID() string {
	uuid := g.next()

	b := make([]byte, 36)
	_ = hex.Encode(b, uuid[:4])
	b[8] = '-'
	_ = hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	_ = hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	_ = hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	_ = hex.Encode(b[24:], uuid[10:16])

	return string(b)
}

func newGenerator_SHA256() (*GeneratorSHA256, error) {
	var cur [sizeSHA256]byte
	if _, err := rand.Read(cur[:]); err != nil {
		return nil, fmt.Errorf("failed to generate seed: %w", err)
	}

	return &GeneratorSHA256{
		cur: cur,
	}, nil
}

// SHA3-256

func Test_Custom_SHA3(t *testing.T) {
	gen, err := newGenerator_SHA3()
	if err != nil {
		t.Fatal(err)
	}

	rid1 := gen.UUID()
	rid2 := gen.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Custom_SHA3(b *testing.B) {
	gen, err := newGenerator_SHA3()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var rid string
	for n := 0; n < b.N; n++ {
		rid = gen.UUID()
	}
	_ = rid
}

const (
	sizeSHA3 = 32
)

type GeneratorSHA3 struct {
	mu sync.Mutex

	cur [sizeSHA3]byte
}

func (g *GeneratorSHA3) next() [sizeSHA3]byte {
	g.mu.Lock()
	cur := g.cur
	cur = sha3.Sum256(cur[:])
	g.cur = cur
	g.mu.Unlock()

	return cur
}

func (g *GeneratorSHA3) UUID() string {
	uuid := g.next()

	b := make([]byte, 36)
	_ = hex.Encode(b, uuid[:4])
	b[8] = '-'
	_ = hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	_ = hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	_ = hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	_ = hex.Encode(b[24:], uuid[10:16])

	return string(b)
}

func newGenerator_SHA3() (*GeneratorSHA3, error) {
	var cur [sizeSHA3]byte
	if _, err := rand.Read(cur[:]); err != nil {
		return nil, fmt.Errorf("failed to generate seed: %w", err)
	}

	return &GeneratorSHA3{
		cur: cur,
	}, nil
}

// Blake2b

func Test_Custom_Blake2b(t *testing.T) {
	gen, err := newGenerator_Blake2b()
	if err != nil {
		t.Fatal(err)
	}

	rid1 := gen.UUID()
	rid2 := gen.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Custom_Blake2b(b *testing.B) {
	gen, err := newGenerator_Blake2b()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var rid string
	for n := 0; n < b.N; n++ {
		rid = gen.UUID()
	}
	_ = rid
}

const (
	sizeBlake2b = blake2b.Size256
)

type GeneratorBlake2b struct {
	mu sync.Mutex

	cur [sizeBlake2b]byte
}

func (g *GeneratorBlake2b) next() [sizeBlake2b]byte {
	g.mu.Lock()
	cur := g.cur
	cur = blake2b.Sum256(cur[:])
	g.cur = cur
	g.mu.Unlock()

	return cur
}

func (g *GeneratorBlake2b) UUID() string {
	uuid := g.next()

	b := make([]byte, 36)
	_ = hex.Encode(b, uuid[:4])
	b[8] = '-'
	_ = hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	_ = hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	_ = hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	_ = hex.Encode(b[24:], uuid[10:16])

	return string(b)
}

func newGenerator_Blake2b() (*GeneratorBlake2b, error) {
	var cur [sizeBlake2b]byte
	if _, err := rand.Read(cur[:]); err != nil {
		return nil, fmt.Errorf("failed to generate seed: %w", err)
	}

	return &GeneratorBlake2b{
		cur: cur,
	}, nil
}

// Blake2s

func Test_Custom_Blake2s(t *testing.T) {
	gen, err := newGenerator_Blake2s()
	if err != nil {
		t.Fatal(err)
	}

	rid1 := gen.UUID()
	rid2 := gen.UUID()

	if rid1 == rid2 {
		t.Fatalf("%s == %s", rid1, rid2)
	}
}

func Benchmark_Custom_Blake2s(b *testing.B) {
	gen, err := newGenerator_Blake2s()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	var rid string
	for n := 0; n < b.N; n++ {
		rid = gen.UUID()
	}
	_ = rid
}

const (
	sizeBlake2s = blake2s.Size
)

type GeneratorBlake2s struct {
	mu sync.Mutex

	cur [sizeBlake2s]byte
}

func (g *GeneratorBlake2s) next() [sizeBlake2s]byte {
	g.mu.Lock()
	cur := g.cur
	cur = blake2s.Sum256(cur[:])
	g.cur = cur
	g.mu.Unlock()

	return cur
}

func (g *GeneratorBlake2s) UUID() string {
	uuid := g.next()

	b := make([]byte, 36)
	_ = hex.Encode(b, uuid[:4])
	b[8] = '-'
	_ = hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	_ = hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	_ = hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	_ = hex.Encode(b[24:], uuid[10:16])

	return string(b)
}

func newGenerator_Blake2s() (*GeneratorBlake2s, error) {
	var cur [sizeBlake2s]byte
	if _, err := rand.Read(cur[:]); err != nil {
		return nil, fmt.Errorf("failed to generate seed: %w", err)
	}

	return &GeneratorBlake2s{
		cur: cur,
	}, nil
}
