hbxss is a heartbeat for XScreenSaver. While it's running, XScreenSaver
will not be triggered. xscreensaver-command must be installed and in
the PATH. As this is a program to suppress automatic locking of a Linux
laptop, the use of the '-t' flag to limit how long the prorgam runs for
is recommended. It is intended to discourage disabling xscreensaver while,
for example, watching movies.

Usage:
	
	hbxss [-i interval] [-t time] [-v]

	-i interval	Specify the interval between heartbeats. This
			should follow the form <number><unit>, where
			unit should be one of 's', 'm', or 'h' for
			seconds, minutes, or hours, respectively.

	-t time		Specify how long the program should run for;
			if not set, hbxss will run indefinitely.

	-v		Print each heartbeat as it occurs.


LICENSE

xbhss is released under an ISC license.

