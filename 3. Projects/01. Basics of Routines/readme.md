# GO Routines

## This is a simple project to learn about go routines

### 1. What is a go routine?

A go routine is a lightweight thread managed by the Go runtime they are multiplexed onto OS threads and are scheduled by the Go runtime. we use the ```go``` keyword to create a go routine.

### 2. Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

- We use the ```chan``` keyword to create a channel.

- We can use the ```<-``` operator to send and receive values from a channel.

- We can use the ```make``` function to create a channel.

- We can use the ```close``` function to close a channel.

By default the channel is unbuffered, meaning that it will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
