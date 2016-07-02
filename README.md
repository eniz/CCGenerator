# Generate test CC numbers on the command line.

First goal of this program is to learn the basic structure of a Go program. Second goal is to eliminate the need for me to go (no pun indented) to [this site](https://www.paypalobjects.com/en_US/vhelp/paypalmanager_help/credit_card_numbers.htm) constantly when testing ecommerce sites we are developing.

The idea here is that we can pipe the output of this program directly to pbcoby to grap a test credit card number like so:

```
test_cc | pbcopy
```

or

```
test_cc -card=mastercard | pbcopy
```

If this proves useful I'd like to support all the of the credit card types on the site linked above.
