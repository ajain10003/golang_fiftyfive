# Problem
Description

Let’s implement the code for a supermarket checkout that calculates the total price of a number
of items.
An item has the following attributes:
● SKU
● Unit Price
Our goods are priced individually. Some items are multi priced: buy n of them, and they’ll cost
you less than buying them individually. For example, item ‘A’ might cost $50 individually, but this
week we have a special offer: buy three ‘A’s and they’ll cost you $130.
Here is an example of prices:
SKU Unit Price Special Price
A $50 3 for $130
B $30 2 for $45
C $20
D $15
Our checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll
recognize the two B’s and price them at 45 (for a total price so far of 95). Because the pricing
changes frequently, we need to be able to pass in a set of pricing rules each time we start
handling a checkout transaction.


The interface to the checkout should look like:

co = new CheckOut(pricing_rules);
co.scan(item);
co.scan(item);
price = co.total();
Here are some examples of cases:

Items Total
A, B $80
A, A $100
A, A, A $130
C, D, B, A $115

# File Structure
I selected a cobra library to create this so we can scale to be manageble command based application where we can add more flags as an argument and complete the operation.
-- main.go 
    This is the root file which calls file in the cmd direcotry
-- cmd/shop.go
    This is the main file where we added Execute function to receive the command, setup the default price Rules and call different function to generate Total calculation
    It depends upon many services , for this example one of the service may be checkout or create Cart. so service related code was handled with separate repository
-- service/checkout
    This directory will include all related services used for the microservice. Here i created Checkout package to basically scan the input then check if this exist in the Sku list and finally it is generating the total amount to be paid with the calcualtion of special offer/ Price
-- pkg
    This folder can be utilize for helper libraries so for example i used here uber zap logger to log the things in the console. We can create custom packages here so that common utility methods can be shared with different microservices.

# How to execute the build
move to root direcotry where main.go file exist then execute
## Go run build 
it will create a binary, i am attaching generated binary as well



