/*
* Run test cases against creditcard.go file
*
*/
package main

import (
    "fmt"
  )

// function to convert boolean input to integer
func btoi (input bool) int{

  if input{
    return 1
  }else
  {
    return 0
  }
}

func main() {

    //declare variables
    var sum = 0
    var count = 0

    sum += tc1()
    count = count + 1
    sum += tc2()
    count = count + 1
    sum += tc3()
    count = count + 1
    sum += tc4()
    count = count + 1
    sum += tc5()
    count = count + 1
    sum += tc6()
    count = count + 1
    sum += tc7()
    count = count + 1
    sum += tc8()
    count = count + 1
    sum += tc9()
    count = count + 1
    sum += tc10()
    count = count + 1

    fmt.Println("==================================================")
    fmt.Println("Total test cases run: ", count)
    fmt.Println("Total test cases PASSED: ", sum)
    fmt.Println("Total test cases FAILED: ", count - sum)
    fmt.Println("Percent successfully completed: ", sum/count)
    fmt.Println("==================================================")
}


func tc1() int{

    input := "4716257953423954"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 1: Visa - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc2() int{

    input := "5515375668895736"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 2: Mastercard - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc3() int{

    input := "348167913925010"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 3: Amex - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc4() int{

    input := "6011333489233261"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 4: Discover - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc5() int{

    input := "3112643331182036"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 5: JCB - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc6() int{

    input := "5468730100666292"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 6: Diner's Club NA - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc7() int{

    input := "30158618387674"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 7: Diner's Club Carte Blance - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc8() int{

    input := "36546674028899"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 8: Diner's Club International - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc9() int{

    input := "5018415654248367"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 9: Maestro - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}

func tc10() int{

    input := "4175007091216479"
    var result bool
    var result_msg string

    result = luhns_algorithm(input)

    if (result){
        result_msg = "PASSED"
    }else {
        result_msg = "FAILED"
    }

    fmt.Println("==================================================")
    fmt.Println("Test Case 10: Visa Electron - ", input)
    fmt.Println("Result:", result_msg)
    fmt.Println("==================================================")

    return btoi(result)
}
