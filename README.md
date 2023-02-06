# go-line-message-analyzer

<!-- https://sc0vu.medium.com/%E9%AB%98%E6%95%88%E8%83%BD-golang-%E7%A8%8B%E5%BC%8F-%E6%95%88%E8%83%BD%E6%AF%94%E8%BC%83-f84bb4fc390a -->
# Russ Cox 的 benchstat 工具，benchstat 可以檢視並分析 benchmark 的結果

#  go test -bench /. -benchmem -memprofile=memprofile.out -cpuprofile=profile.out -blockprofile=blockprofile.out
# go tool pprof profile.out
# go tool pprof -http :8080  profile.out
<!-- 
-cpuprofile=$FILE 將 CPU 分析結果存到檔案 $FILE.
-memprofile=$FILE 將記憶體分析結果存到檔案 $FILE，-memprofilerate=N 調整分析的頻率為 1/N.
-blockprofile=$FILE 將 block 分析結果存到檔案 $FILE. -->

# 改善效能的時候，要確定程式有通過原先的測試

# 每次執行 benchmark 都需要初始化物件，如果初始化會花比較多時間，那可以用 b.ResetTimer() 重設計時器
<!-- func BenchmarkExpensive(b *testing.B) {
        boringAndExpensiveSetup()
        b.ResetTimer() 
        for n := 0; n < b.N; n++ {
                // do something
        }
} -->
# 每次回圈都要執行，但又不需要計算 benchmark，這時可以使用 b.StartTimer() 以及 b.StopTimer() 來暫停計時器
<!-- func BenchmarkComplicated(b *testing.B) {
        for n := 0; n < b.N; n++ {
                b.StopTimer() 
                complicatedSetup()
                b.StartTimer() 
                // do something
        }
} -->
# 併發基準測試
<!-- func BenchmarkCombinationParallel(b *testing.B) {
    // 測試一個對象或者函數在多線程的場景下面是否安全
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            m := rand.Intn(100) + 1
            n := rand.Intn(m)
            combination(m, n)
        }
    })
} -->