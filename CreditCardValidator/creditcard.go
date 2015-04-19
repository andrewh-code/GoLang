package main

import (
    "fmt"
    "regexp"
    "os"
    )

// create struct to hold credit card values
type CC_info struct {
    brand string
    number_pattern string
    ccv int
    length string
}

func cc_validate(cc_brand int, cc_number string, ccv_number string) {

    var creditcards [14]CC_info
    i := cc_brand - 1

    creditcards[0]  = CC_info{brand:          "American Express",
                              number_pattern: "^3[47]",
                              ccv:            4,
                              length:         ".{15}$"};
    creditcards[1]  = CC_info{brand:          "China UnionPay",
                              number_pattern: "^62",
                              ccv:            3,
                              length:         "{16,19}"};
    creditcards[2]  = CC_info{brand:          "Diners Club Carte Blanche",
                              number_pattern: "^30[0-5]",
                              ccv:            3,
                              length:         ".{14}$"};
    creditcards[3]  = CC_info{brand:          "Diners Club International",
                              number_pattern: "^(30[0-5,9]|36|38|39)",
                              ccv:            3,
                              length:         ".{14}$"};
    creditcards[4]  = CC_info{brand:          "Diners Club United States & Canada",
                              number_pattern: "^(54|55)",
                              ccv:            3,
                              length:         ".{16}$"};
    creditcards[5]  = CC_info{brand:          "Discover",
                              number_pattern: "^6(011|[22126-22925]|[44-49]|5)",
                              ccv:            3,
                              length:         ".{16}&"};
    creditcards[6]  = CC_info{brand:          "InterPayment",
                              number_pattern: "^636",
                              ccv:            3,
                              length:         ".{16,19}$"};
    creditcards[7]  = CC_info{brand:          "JCB",
                              number_pattern: "^35[28-89]",
                              ccv:            3,
                              length:         ".{16}$"};
    creditcards[8]  = CC_info{brand:          "Maestro",
                              number_pattern: "^(50[0000-9999]|56[0000-9999])",
                              ccv:            3,
                              length:         ".{12,19}$"};
    creditcards[9]  = CC_info{brand:          "Dankort",
                              number_pattern: "^5019",
                              ccv:            3,
                              length:         ".{16}$"};
    creditcards[10] = CC_info{brand:          "MasterCard",
                              number_pattern: "^5[1-5]",
                              ccv:            3,
                              length:         ".{16}$"};
    creditcards[11] = CC_info{brand:          "Visa",
                              number_pattern: "^4",
                              ccv:            3,
                              length:         ".{13|16}$"};
    creditcards[12] = CC_info{brand:          "Visa Electron",
                              number_pattern: "^(4026|417500|4405|4508|4844|4913|4917)",
                              ccv:            3,
                              length:         ".{16}$"};
    creditcards[13] = CC_info{brand:          "UATP",
                              number_pattern: "^1",
                              ccv:            3,
                              length:         ".{15}$"};

    //validate the credit card values
    fmt.Println("Credit card carrier chosen is: ", creditcards[i].brand)

    //validate cc number pattern
    pattern_check := regexp.MustCompile(creditcards[i].number_pattern)
    pattern_result := pattern_check.MatchString(cc_number)
    if (pattern_result){
        fmt.Println("Credit card number matches brand: ", creditcards[i].brand)
    }else{
        fmt.Println("Credit card number does not match brand.")
        fmt.Println("Now exiting...")
        os.Exit(2)
    }

    //validate cc number length
    length_check := regexp.MustCompile(creditcards[i].length)
    length_result := length_check.MatchString(cc_number)
    if (length_result){
        fmt.Println("Credt card number is the correct length")
    }else{
        fmt.Println("Credit card number is not the correct length")
        fmt.Println("Now exiting...")
        os.Exit(3)
    }

    //validate ccv number length
    if (len(ccv_number) == creditcards[i].ccv){
        fmt.Println("Credt card ccv is the correct length")
    }else{
        fmt.Println("Credit card ccv is not the correct length")
        fmt.Println("Now exiting...")
        os.Exit(4)
    }


}
