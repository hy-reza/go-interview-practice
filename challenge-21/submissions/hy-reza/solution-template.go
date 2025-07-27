package main

import (
	"fmt"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	targetNotFound := 6
	indexNotFound := BinarySearch(arr, targetNotFound)
	fmt.Printf("BinarySearch: %d found at index %d\n", targetNotFound, indexNotFound)


	// Test recursive binary search
	recursiveTarget := 11
	recursiveIndex := BinarySearchRecursive(arr, recursiveTarget, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", recursiveTarget, recursiveIndex)

	recursiveTargetNotFound := 20
	recursiveIndexNotFound := BinarySearchRecursive(arr, recursiveTargetNotFound, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", recursiveTargetNotFound, recursiveIndexNotFound)

	// Test find insert position
	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)

	insertTargetStart := 0
	insertPosStart := FindInsertPosition(arr, insertTargetStart)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTargetStart, insertPosStart)

	insertTargetEnd := 20
	insertPosEnd := FindInsertPosition(arr, insertTargetEnd)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTargetEnd, insertPosEnd)

	insertTargetExisting := 9
	insertPosExisting := FindInsertPosition(arr, insertTargetExisting)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d (or found at %d)\n", insertTargetExisting, insertPosExisting, insertPosExisting)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // To prevent potential overflow
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1 // Base case: target not found
	}

	mid := left + (right-left)/2 // To prevent potential overflow

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	left, right := 0, len(arr)-1
	insertPos := len(arr) // Default to inserting at the end if target is largest

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid // Target found, can be inserted at this position
		} else if arr[mid] < target {
			left = mid + 1
		} else { // arr[mid] > target
			insertPos = mid // This could be our insert position
			right = mid - 1
		}
	}
	return insertPos
}