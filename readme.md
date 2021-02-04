# MAWD Altimeter Data Capture

A tool to configure a PerfectFlite MAWD altimeter and download flight data

The protocol is pretty straightforward:
Simply send the ASCII command over serial link to the MAWD, and receive the response
Serial data format:	8N1, XON/XOFF, 38,400 bps for commands+data, 9,600 bps for telemetry


## Command list

Command | Action | Implemented?
------- | ------ | ------------
A0 | Turn 1 second apogee delay OFF | TODO
A1 | Turn 1 second apogee delay ON | TODO
A[CR] | Report current status of apogee delay ON/OFF | needs testing
C | Start continuity beep sequence (send any char to end) | TODO
D | Dump data from last run | needs testing
FD | Fire drogue channel | TODO
FM | Fire main channel | TODO
I | Identify altimeter model | needs testing
Lxx | Set low voltage alarm threshold to xx/10 volts | TODO
L[CR] | Report current low voltage alarm threshold | needs testing
R | Reboot | needs testing
S | Report stats (ground, apogee, #samps, machdel, mainalt) | needs testing
T0 | Turn telemetry output during flight OFF | TODO
T1	 | 	Turn telemetry output during flight ON | TODO
T[CR] | Report current status of telemetry output ON/OFF | needs testing
V | Report firmware version number | needs testing

