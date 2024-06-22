# Slices vs Iterators: Benchmarking in Go v1.23

## Measurement Results

In this implementation, the slice and the iterator's memory usage increased almost linearly as the number of elements in a given slice increased. The memory usage ratio between slices and iterators varied from 2 to 3.5 times as the number of elements increased.

### Results Table

| Number of Elements | Slice Memory Usage [Byte] | Iterator Memory Usage [Byte] |
| ------------------ | ------------------------- | ---------------------------- |
| 100                | 1,920                     | 896                          |
| 1,000              | 16,384                    | 8,192                        |
| 10,000             | 210,176                   | 81,920                       |
| 100,000            | 2,757,888                 | 802,816                      |
| 1,000,000          | 29,086,976                | 8,003,584                    |
| 10,000,000         | 281,081,088               | 80,003,072                   |

### Results Graph

![Memory Usage Comparison](https://github.com/Siddhant-K-code/slice-vs-iterator-benchmarking/assets/55068936/f4c30424-fdad-4be4-9a94-126c7168f6d9)

## Conclusions and Discussion

I found that the case that returns an iterator uses less memory than the case that creates and returns a new slice. As the number of elements in the given slice increased, the memory usage increase was greater in the case of returning a slice than in the case of returning an iterator. This is likely due to the logic of memory allocation in the function when appending elements.
