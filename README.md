rpidinsync
==========

Raspberry PI GPIO DIN sync example using Go. This code accompanies a blog post that I wrote (), that you should almost certainly read before you do anything else.

Connect the Raspberry Pi's GPIO25 pin to the CLOCK pin of your DIN sync device. Connect the GPIO24 pin to the START/STOP pin of your DIN sync device. Connect one of the Raspberry Pi GROUND pins to the GND pin of your DIN sync device.

There are two ways to run the example code. 

Download the 'run' binary and use `sudo` to execute it on your Raspberry Pi. Root privileges are necessary because the Pi uses Linux's /sys/class exporting to read and write from the GPIO pins.

To run the code from source, first clone this repository onto your Raspberry Pi or just download src/main.go

The example code depends on (and was inspired by) https://github.com/davecheney/gpio

```
go get github.com/davecheney/gpio
sudo go run /path-to-your-file/main.go
```

If all goes well, your DIN sync'd device should now be playing whatever is currently in its sequencer at 120 bpm.

All of this code is available under an MIT license (see accompanying [LICENSE](https://github.com/ceberly/rpidinsync/blob/master/LICENSE) file). Please feel free to do what you want with it. 
I am especially interested in contributions that will help turn the Raspberry Pi into a MIDI/DIN sync bridge.
