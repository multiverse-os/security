# Shodan Search
Very quick library and example of searching `shodan.io` for values, exposing a major issue with the current state of many p2p networks which very stupidly turn on highly exploitable RPCs by default.

Hoping by making this easy to use will bring some light to the situation and help push developers in a more secure design direction istead of making the same mistakes over and over and building insecure tools like many tools from the early internet. 



### Example
Import and just call the single existing function:

Two ways two initialize the `Shodan` object, the standard `New()` for no-login style usage, and `Login(username, password)` and both return a `Shodan` struct can be used to then run searches against and will return the list of ip addresses. (Eventually should grab all the port data and CVEs.)

Look at the `cmd/ssearch` for example usage; but its just `shodan.Search("value")`.
