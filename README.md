rpidinsync
==========

Raspberry PI GPIO DIN Sync example using Go. This code accompanies a blog post that I wrote (), that you should almost certainly read before you do anything else.

Connect the Raspberry Pi's GPIO25 pin to the CLOCK pin of your DIN sync device. Connect the GPIO24 pin to the START/STOP pin of your DIN sync device. Connect one of the Raspberry Pi GROUND pins to the GND pin of your DIN sync device.

There are two ways to run the example code. You can use the precompiled binary here or run from source.

To do the latter, you will need to:

`go get github.com/davecheney/gpio`


Then run it!
