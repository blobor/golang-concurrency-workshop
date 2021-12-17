# Select statement

The select statement is used when multiple goroutines are sending data via channels then the select statement receives data concurrently and chooses the case randomly if all are ready. If no case is ready then it simply outputs the default case if the default case is already provided before. This shows the versatility of the select statement which is used to selectively get data from multiple provider channels.

* simple select
  ```bash
  go run 03-select/01-simple_select/simple.go
  ```
* default select
  ```bash
  go run 03-select/02-default-select/default.go
  ```
* select with cancellation
  ```bash
  go run 03-select/03-cancellation/cancell.go
  ```
