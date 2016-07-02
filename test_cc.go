package main

import(
    "fmt"
    "os"
    "flag"
    "math/rand"
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
    cards := []string{
        "4111111111111111",
        "4012888888881881",
        "4222222222222",
        "4242424242424242",
        "4012888888881881",
    }

    return cards[rand.Intn(len(cards))]
}

func MasterCard() string {
    cards := []string{
        "5555555555554444",
        "5105105105105100",
        "378282246310005",
    }
    return cards[rand.Intn(len(cards))]
}
