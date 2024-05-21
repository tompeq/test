package main

import (
    "fmt"
    "os"
)
// CHANGE 
// NEW CHANGE 1
// NEW CHANGE 2
func MAIN() {

    if len(os.Args) < 4 {
        fmt.Println("Usage:", os.Args[0], "input_file part_of_matrix operation")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    partOfMatrix := os.Args[2]
    operation := os.Args[3]

    file, err := os.Open(inputFile)
    if err != nil {
        fmt.Println("Unable to open file")
        os.Exit(1)
    }
    defer file.Close()

    matrix := make([][]int, 4)
    for i := range matrix {
        matrix[i] = make([]int, 4)
        for j := range matrix[i] {
            var num int
            fmt.Fscanf(file, "%d", &num)
            matrix[i][j] = num
        }
    }

    switch {
    case partOfMatrix == "upper_left" && operation == "summa":
        sum := 0
        for i := 0; i < 4; i++ {
            for j := 0; j <= i; j++ {
                sum += matrix[i][j]
            }
        }
        fmt.Println("Sum of upper left triangle da:", sum)
    case partOfMatrix == "upper_right" && operation == "sum":
        sum := 0
        for i := 0; i < 4; i++ {
            for j := i; j < 4; j++ {
                sum += matrix[i][j]
            }
        }
        fmt.Println("Sum of upper right triangle:", sum)
    case partOfMatrix == "lower_left" && operation == "min":
        var lowerLeftElements []int
        for i := 0; i < 4; i++ {
            for j := 0; j <= i; j++ {
                lowerLeftElements = append(lowerLeftElements, matrix[i][j])
            }
        }
        minElem := lowerLeftElements[0]
        for _, num := range lowerLeftElements {
            if num < minElem {
                minElem = num
            }
        }
        fmt.Println("Minimum element of lower left triangle:", minElem)
    case partOfMatrix == "lower_right" && operation == "max":
        var lowerRightElements []int
        for i := 0; i < 4; i++ {
            for j := i; j < 4; j++ {
                lowerRightElements = append(lowerRightElements, matrix[i][j])
            }
        }
        maxElem := lowerRightElements[0]
        for _, num := range lowerRightElements {
            if num > maxElem {
                maxElem = num
            }
        }
        fmt.Println("Maximum element of lower right triangle:", maxElem)
    default:
        fmt.Println("Invalid part of the matrix or operation now ")
    }
}
