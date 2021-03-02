# Go HashMap Benchmark

## Usage

```bash
cd main
go build
main.exe #or ./main
```

## Arguments

```bash
# choose the hashmap, default: all
--hashmap [BasicHashMap, CuckooHashMap, BinaryTreeHashMap, GoLangMap, all]
# number of elements to insert, MUST BE PRIME AND > 1000537!
--elements 1048575
# get call multiplier (results in elements * multiplier get calls)
-- multiplier 100
```

## Benchmark results (AMD Ryzen 1700 3.6GHz 8-Core)

The following values are averages of multiple runs and multiplier = 100.

Please be aware that memory is the Heap size that Go allocated in order to run the benchmark.

### BasicHashMap

(Grow multiplier 2x)
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 1048575       | 199.794ms     | 4.366s  | 24.400ms  | 56mb   |
| 10485751      | 3.289s        | 65.027s | 171.306ms | 608mb  |
| 104857351     | 37.56s        | 916.67s | 2.578s    | 5824mb |

(Grow multiplier 4x)
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 9973          | 1.9996ms      | 9.643ms | 1ms       | 1mb    |
| 1048575       | 218.009ms     | 4.55s   | 18ms      | 64mb   |
| 10485751      | 2.814s        | 64.303s | 227.080ms | 736mb  |
| 104857351     | 27.07s        | 685.24s | 3.432s    | 5824mb |

### CuckooHashMap

(Grow multiplier 2x)
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 1048575       | 183.634ms     | 4.246s  | 18.400ms  | 57.6mb |
| 10485751      | 1.968s        | 51.927s | 202.819ms | 730mb  |
| 104857351     | 27.2s         | 571.55s | 3.29s     | 5248mb |

(Grow multiplier 4x)
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 9973          | 1.2413ms      | 8.740ms | 1ms       | 1mb    |
| 1048575       | 192.280ms     |  4.695s | 18.210ms  | 60mb   |
| 10485751      | 2.11s         | 52.727s | 202.999ms | 730mb  |
| 104857351     | 27.7s         | 580.72s | 3.14s     | 5248mb |

### BinaryTreeHashMap
(Grow multiplier 2x)
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 1048575       | 252.753ms     | 6.623s  | 55.001ms  | 56mb   |
| 10485751      | 3.978s        | 98.876s | 400.898ms | 608mb  |
| 104857351     | 41.758s       | 1276.5s | 7.24s     | 5824mb |

Grow multiplier 4x
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 9973          | 1.5858ms      | 10.74ms | 1ms       | 1mb    |
| 1048575       | 225.000ms     | 6.916s  | 53.554ms  | 64mb   |
| 10485751      | 3.241s        | 101.31s | 486.206ms | 736mb  |
| 104857351     | 32.862s       | 930.88s | 12.994s   | 5824mb |

### GoLangMap
| Elements      | Insert        | Get     | K-V-Pairs | Memory |
| ------------- | ------------- | -----   | --------- | ------ |
| 9973          | 1.503ms       | 30.62ms | 1ms       | 1mb    |
| 1048575       | 217.614ms     | 10.76s  | 21.804ms  | 76mb   |
| 10485751      | 2.721s        | 129.63s | 195.506ms | 649mb  |
| 104857351     | 32.680s       | 1772.7s | 2.569s    | 5824mb |