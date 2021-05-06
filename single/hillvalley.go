package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

/*def solution(A):
  peaks = []

  for idx in xrange(1, len(A)-1):
      if A[idx-1] < A[idx] > A[idx+1]:
          peaks.append(idx)

  if len(peaks) == 0:
      return 0

  for size in xrange(len(peaks), 0, -1):
      if len(A) % size == 0:
          block_size = len(A) // size
          found = [False] * size
          found_cnt = 0
          for peak in peaks:
              block_nr = peak//block_size
              if found[block_nr] == False:
                  found[block_nr] = True
                  found_cnt += 1

          if found_cnt == size:
              return size

  return 0*/

func Solution(A []int) int {
	// write your code in Go 1.4
	var hill []int
	for idx := 1; idx <= len(A)-1; idx++ {
		fmt.Println(idx)
		if idx < len(A)-1 {

			if A[idx-1] < A[idx] && A[idx] > A[idx+1] {
				hill = append(hill, idx)
			}
		} else {
			if A[idx-1] < A[idx] {
				hill = append(hill, idx)
			}
		}
		fmt.Println("hill")
		fmt.Println(hill)
	}

	var valley []int
	/*for idx := 1; idx <= len(A)-1; idx++ {
		fmt.Println(idx)
		if idx < len(A)-1 {
			if A[idx-1] > A[idx] && A[idx] < A[idx+1] {
				valley = append(valley, idx)
			}
		} else {
			if A[idx-1] > A[idx] {
				hill = append(valley, idx)
			}
		}
		fmt.Println("valley")
		fmt.Println(valley)
	}*/

	for k, v := range hill {
		temp := 0
		for i := 0; i <= v; i++ {
			if temp == 0 {
				temp = A[i]
			} else if temp > A[i] {
				temp = A[i]
			}
		}

		if k != len(hill)-1 {
			for i := v; i <= hill[k+1]; i++ {
				if temp == 0 {
					temp = A[i]
				} else if temp > A[i] {
					temp = A[i]
				}
			}
		}

		valley = append(valley, temp)
	}
	fmt.Println(valley)

	return 0

}

func main() {
	A := []int{2, 2, 3, 4, 3, 3, 2, 2, 1, 1, 2, 5}

	fmt.Println(Solution(A))
}
