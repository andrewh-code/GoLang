/*
* Credit Card Validation program
* GoLang program to validate credit card numbers
*/

package main

import (
  // import standard libraries
  "fmt"
  "os"
  "bufio"
  "strings"
  "regexp"
  /* import custom libraries
     don't need to import custom lirbary/files, just put package main in
     the other file and add the file to the compilation command ie
     go run main.go creditcard.go
     "GoLang/CreditCardValidator/creditcard"
  */

  )

  // use regular expression to validate only numbers in the credit card number
func validatenumber(input string) {

    // declare variables
    check := regexp.MustCompile("^[0-9]+$")
    result := check.MatchString(input)

    if (!result){
        fmt.Println("incorrect input. Expecting integer as input.")
        fmt.Println("Now exiting...")
        os.Exit(1)
    }
}

func main() {

  // declare variables
  var cc_brand int
  var ccv_number, cc_number string

  // main screen
  fmt.Println("Please enter the number corresponding to the credit card brand" +
              " you own: \n")
  // from wikipedia (http://en.wikipedia.org/wiki/Bank_card_number)
  fmt.Println("1. American Express")
  fmt.Println("2. China UnionPay")
  fmt.Println("3. Diners Club Carte Blanche")
  fmt.Println("4. Diners Club International")
  fmt.Println("5. Diners Club United States & Canada")
  fmt.Println("6. Discover")
  fmt.Println("7. InterPayment")
  fmt.Println("8. JCB")
  fmt.Println("9. Maestro")
  fmt.Println("10. Dankort")
  fmt.Println("11. MasterCard")
  fmt.Println("12. Visa")
  fmt.Println("13. Visa Electron")
  fmt.Println("14. UATP")
  fmt.Println("\n")

  fmt.Println("Please enter the number: ")
  if _, err := fmt.Scanf("%d\n", &cc_brand); err != nil {
    fmt.Println("Input Error:", err)
    os.Exit(1)
  }

  fmt.Println("Please enter your credit card number (without spaces)")
  //use bufio library to handle input with spaces (ie, 4532 3118 8116 3067)
  cc_num_input := bufio.NewReader(os.Stdin)
  cc_number, err := cc_num_input.ReadString('\n')
  if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
  }

  fmt.Println("Please enter the CCV number")
  ccv_input := bufio.NewReader(os.Stdin)
  ccv_number, err = ccv_input.ReadString('\n')
  if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
  }

  // trim the new line character from the input, remove the '\' and the 'n'
  // the newline character is appended onto the input string from stdin
  // remember to chomp it
  //This could also be used: cc_number = cc_number[0:len(cc_number)-2]
  cc_number = strings.TrimSpace(cc_number)
  ccv_number = strings.TrimSpace(ccv_number)

  //remove all the spaces in the credit card number
  cc_number = strings.Replace(cc_number, " ", "", -1)

  //call function to validate ccnumber
  validatenumber(cc_number)
  validatenumber(ccv_number)
  fmt.Println("=============================================")

  //go through cc brands and validate
  cc_validate(cc_brand, cc_number, ccv_number)
  
  //use luhn's algorithm to verify cc number is valid
  luhns_algorithm(cc_number)
}
