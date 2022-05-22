package main

import (
  "fmt"
  "strings"
)

// This functions checks the byte length 
//   length - length to be checked
func checkByteLength(length int) bool{
  if (length > 8) {
    fmt.Print("Bytes cannot be bigger thant 8 bits");
    return false;
  }

  return true;
}

// This functions checks the digits of the bits given by parameter
//   bits - bits to be checked
func checkBitsDigits(bits string) bool {
  for i := 0; i < len(bits); i++{
    if(bits[i] != '0' && bits[i] != '1'){
      fmt.Println("Bits must be 0 or 1");
      return false;
    }
  }
  return true;
}

func main() {
	fmt.Println("Program Started")
  
  var binaryLength int;
  var binary string;
  decimal := 0;
  decimalExponent := 1;

  for true{
    fmt.Println("Enter the binary byte that you want to convert\nTo exit the program, write exit");
    
    // Scan string and string length
    fmt.Scanln(&binary);
    binaryLength = len(binary) - 1;
    
    // Compare if the user want to exit the program
    if (strings.Compare(binary, "exit") == 0){
      fmt.Print("Exiting the program");
      return;
    }
  
    // Check the length
    if (checkByteLength(len(binary)) == false){
      continue;
    }
    // Check the digits
    if (checkBitsDigits(binary) == false){
      continue;
    }
  
    //  Evaluating decimal
    for binaryLength >= 0 {
      decimal += int(binary[binaryLength] - '0') * decimalExponent;
      decimalExponent *= 2;
      binaryLength--;
    }
  
    // Printing decimal value
    fmt.Printf("The result of the binary is %d\n\n", decimal);
  }
}
