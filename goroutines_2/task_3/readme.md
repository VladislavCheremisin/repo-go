результаты тестирования
vlad@vlad-LIFEBOOK:~/Рабочий стол/fast/repo-go/goroutines_2/task_3$ go test -bench=. -benchtime=10s bench_test.go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-3320M CPU @ 2.60GHz
BenchmarkMutex50w_50r-4            14055            740909 ns/op
BenchmarkRWMutex50w_50r-4          13417            889123 ns/op
BenchmarkMutex10w_90r-4            10000           1226458 ns/op
BenchmarkRWMutex10w_90r-4          10000           1568939 ns/op
BenchmarkMutex90w_10r-4            24085            527955 ns/op
BenchmarkRWMutex90w_10r-4          21508            584617 ns/op
PASS
ok      command-line-arguments  107.193s
