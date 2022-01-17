# Benchmark

To evaluate the performance of this library, we measured the following benchmarks for `gcf.Filter`.

## [`BenchmarkFilter_Volumes`](filter_test.go#L10)

By changing the number of elements and the number of filters, we are measuring how the processing time changes with respect to the amount of data and the number of processing.  
Please refer to [filter-volumes.txt](filter-volumes.txt) for results.

As a result of the measurement, it was confirmed that the processing time was roughly proportional to the amount of data and the number of processing.

## [`BenchmarkFilter_Compare`](filter_test.go#L38)

By comparing the processing of the same logic in different implementation methods, the difference in processing speed relative to other implementations is measured.  
The logic is as follows.

- The Iterable source is a number from 1 to 100
- Finds elements whose 13 remainder, 11 remainder, and 7 remainder are non-zero

The target of the benchmark is as follows.

- filter
  - Calculate each remainder with `gcf.Filter`
  - Allocation only occurs when Filter is generated
- if-func
  - The result of each remainder calculation by the external function by if statement condition
  - No allocation has occurred
- if-inline
  - The result of each remainder calculation by if statement condition directly
  - No allocation has occurred
- chan
  - Calculate each remainder on the channel
  - Allocation occurs only channel and maybe goroutine

Please refer to [filter-compare-8.txt](filter-compare-8.txt) for results.

As a result of the measurement, the processing time is too slow incomparable to the inline evaluation, but the processing time is about 4 times as long as the evaluation by the function.  
In addition, the processing time is overwhelmingly shorter than using a channel.

In addition, the processing result when 1 CPU core is used is shown in [filter-compare-1.txt](filter-compare-1.txt).  
As a result, it was confirmed that `gcf.Filter` is almost unaffected by the number of CPU cores.  
(I'm not sure why the channel implementation takes longer to process with more CPU cores)
