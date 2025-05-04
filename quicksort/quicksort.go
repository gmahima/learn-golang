package quicksort

import "fmt"

// Sort takes a slice of integers and returns a new sorted slice
func Sort(numbers []int) []int {
    // Create a copy to avoid modifying the original slice
    arr := make([]int, len(numbers))
    copy(arr, numbers)
    fmt.Println("Initial array:", arr)
    quickSort(arr, 0, len(arr)-1)
    fmt.Println("Final sorted array:", arr)
    return arr
}

func quickSort(arr []int, low, high int) {
    fmt.Printf("quickSort called with array: %v, low: %d, high: %d\n", arr, low, high)
    if low < high {
        p := partition(arr, low, high)
        fmt.Printf("Partition returned: %d, array is now: %v\n", p, arr)
        quickSort(arr, low, p)
        quickSort(arr, p+1, high)
    }
}

func partition(arr []int, low, high int) int {
    pivot := arr[(low+high)/2]  // Middle element as pivot
    fmt.Printf("Partitioning array: %v from %d to %d with pivot: %d\n", arr, low, high, pivot)
    i := low - 1
    j := high + 1
    
    for {
        // Find left element >= pivot
        for {
            i++
            fmt.Printf("  Left scan: i=%d, arr[i]=%d\n", i, arr[i])
            if arr[i] >= pivot {
                break
            }
        }
        
        // Find right element <= pivot
        for {
            j--
            fmt.Printf("  Right scan: j=%d, arr[j]=%d\n", j, arr[j])
            if arr[j] <= pivot {
                break
            }
        }
        
        // If pointers crossed, partition complete
        if i >= j {
            fmt.Printf("  Pointers crossed (i=%d, j=%d), returning j=%d\n", i, j, j)
            return j
        }
        
        // Swap mismatched pair
        fmt.Printf("  Swapping arr[%d]=%d with arr[%d]=%d\n", i, arr[i], j, arr[j])
        arr[i], arr[j] = arr[j], arr[i]
        fmt.Printf("  After swap: %v\n", arr)
    }
}
