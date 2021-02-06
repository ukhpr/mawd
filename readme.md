# MAWD Altimeter Data Capture

A tool to configure a PerfectFlite MAWD altimeter and download flight data

The protocol is pretty straightforward:

Simply send the ASCII command over serial link to the MAWD, and receive the response

Serial data format:	8N1, XOn/XOff, 38,400 bps for commands+data, 9,600 bps for telemetry


## Command list

Command | Action | Implemented?
------- | ------ | ------------
A0 | Turn 1 second apogee delay Off | TODO
A1 | Turn 1 second apogee delay On | TODO
A[CR] | Report current status of apogee delay On/Off | needs testing
C | Start continuity beep sequence (send any char to end) | TODO
D | Dump data from last run | needs testing
FD | Fire drogue channel | TODO
FM | Fire main channel | TODO
I | Identify altimeter model | needs testing
Lxx | Set low voltage alarm threshold to xx/10 volts | TODO
L[CR] | Report current low voltage alarm threshold | needs testing
R | Reboot | needs testing
S | Report stats (ground, apogee, #samps, machdel, mainalt) | needs testing
T0 | Turn telemetry output during flight Off | TODO
T1	 | 	Turn telemetry output during flight On | TODO
T[CR] | Report current status of telemetry output On/Off | needs testing
V | Report firmware version number | needs testing


## Main deployment settings

SW1 | SW2 | SW3 | Altitude
--- | --- | --- | --------
Off | Off | Off | 300 | feet AGL
Off | Off | On | 500 feet AGL
Off | On | Off | 700 feet AGL
Off | On | On | 900 feet AGL
On | Off | Off | 1100 feet AGL
On | Off | On | 1300 feet AGL
On | On | Off | 1500 feet AGL
On | On | On | 1700 feet AGL


## Mach delay settings

SW4 | SW5 | SW6 | Delay
--- | --- | --- | --------
Off | Off | Off | 0 seconds
Off | Off | On | 2 seconds
Off | On | Off | 4 seconds
Off | On | On | 6 seconds
On | Off | Off | 8 seconds
On | Off | On | 10 seconds
On | On | Off | 12 seconds
On | On | On | 14 seconds
