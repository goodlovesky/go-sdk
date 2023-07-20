使用一下命令测试当前文件夹下的所有测试文件性能
go test ./... -bench=. -benchmem

测试结果
BenchmarkTestString1-8           2224446               533.0 ns/op            88 B/op          2 allocs/op
BenchmarkTestString2-8           2881166               410.5 ns/op            32 B/op          2 allocs/op
BenchmarkTestString3-8           3556190               340.1 ns/op            32 B/op          2 allocs/op
BenchmarkTestString4-8           3731331               320.1 ns/op            32 B/op          2 allocs/op
BenchmarkTestString5-8          10870698               107.1 ns/op            32 B/op          2 allocs/op
BenchmarkTestString6-8          12935973                89.80 ns/op           32 B/op          2 allocs/op
BenchmarkTestString7-8          13020026                90.52 ns/op           16 B/op          1 allocs/op
BenchmarkTestString8-8          16465339                69.45 ns/op           16 B/op          1 allocs/op