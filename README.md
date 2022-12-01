# Go UUID benchmark

```sh
go test -v -bench=. -benchmem ./...

=== RUN   Test_Fiber_UUID
--- PASS: Test_Fiber_UUID (0.00s)
=== RUN   Test_Fiber_UUIDv4
--- PASS: Test_Fiber_UUIDv4 (0.00s)
=== RUN   Test_Custom_Crypto
--- PASS: Test_Custom_Crypto (0.00s)
=== RUN   Test_Custom_SHA256
--- PASS: Test_Custom_SHA256 (0.00s)
=== RUN   Test_Custom_SHA3
--- PASS: Test_Custom_SHA3 (0.00s)
=== RUN   Test_Custom_Blake2b
--- PASS: Test_Custom_Blake2b (0.00s)
=== RUN   Test_Custom_Blake2s
--- PASS: Test_Custom_Blake2s (0.00s)
goos: linux
goarch: amd64
pkg: github.com/leonklingele/uuidbench
cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
Benchmark_Fiber_UUID
Benchmark_Fiber_UUID-4          12300488                97.53 ns/op      48 B/op          1 allocs/op
Benchmark_Fiber_UUIDv4
Benchmark_Fiber_UUIDv4-4         1097896              1057 ns/op      64 B/op          2 allocs/op
Benchmark_Custom_Crypto
Benchmark_Custom_Crypto-4        1000000              1087 ns/op      80 B/op          2 allocs/op
Benchmark_Custom_SHA256
Benchmark_Custom_SHA256-4        3478196               359.5 ns/op      48 B/op          1 allocs/op
Benchmark_Custom_SHA3
Benchmark_Custom_SHA3-4          1048357              1310 ns/op    1040 B/op          6 allocs/op
Benchmark_Custom_Blake2b
Benchmark_Custom_Blake2b-4       4448674               265.4 ns/op      48 B/op          1 allocs/op
Benchmark_Custom_Blake2s
Benchmark_Custom_Blake2s-4       5541786               221.7 ns/op      48 B/op          1 allocs/op
PASS
ok      github.com/leonklingele/uuidbench 11.231s
```
