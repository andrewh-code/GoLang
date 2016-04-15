//start the show
package main

//import packages
import (
    "fmt"
    "io"
    "image"
    "image/png"
    //"image/jpeg"
    //"image/gif"
)

//convert image file to PNG (use snorlax_minimalist.jpg to test)
func convertToPNG(w io.Writer, r io.Reader) error {
    //:= means declaration with implicit type
    img, _, err := image.Decode(r)
    
    if err != nil {
        return err
    }
    
    return png.Encode(w, img)
}

func main(){
    
    var input_img, output_img string
    input_img = "../imgs/snorlax_minimalist.jpg"
    output_img = "../imgs/snorlax_minimalist.PNG"
    
    
    fmt.Printf(img_path)
    convertToPNG(input_img, output_img)
}