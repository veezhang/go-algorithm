
# LeetCode

leetcode 练习题

## 100% 覆盖测试

```shell
go test -coverpkg=./... -coverprofile=coverage.data -timeout=5s ./...
go tool cover -func=coverage.data -o coverage.txt
go tool cover -html=coverage.data -o coverage.html
```
