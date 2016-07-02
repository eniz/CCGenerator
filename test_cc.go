package main

import(
    "fmt"
    "os"
    "flag"
)


func main() {
    cardType := flag.String("card", "visa", "a credit card type - visa or mastercard accepted.");
    card := "";
    flag.Parse();

    switch(*cardType) {
      case "visa":
          card = VisaCard();
          break;
      case "mastercard":
          card = MasterCard();
          break;
    }

    fmt.Fprintln(os.Stdout, card);
}


func VisaCard() string {
    return "4111111111111111";
}

func MasterCard() string {
    return "5105105105105100";
}
