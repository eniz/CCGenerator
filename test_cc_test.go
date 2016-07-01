package main

import(
  "testing"
)

func TestVisaCard(t *testing.T) {
  card := VisaCard();

  if(card != "4111111111111111") {
    t.Fatalf("VisaCard failed: %s", card);
  }
}

func TestMasterCard(t *testing.T) {
  card := MasterCard();

  if(card != "5105105105105100") {
    t.Fatalf("MasterCard failed: %s", card);
  }
}
