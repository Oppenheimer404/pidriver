# PiDriver

What is PiDriver
----------------
PiDriver is a utility built for the RaspberryPi Zero 2 W. The goal of PiDriver is to enable a RPIZero2W to autonomously wardrive. This includes being able to;

- Scan for Networks
- Determine Approximate GPS Location
- Log Network SSID & Juicy info + GPS Location
- Connect to one (or more) 'Base Station' ssid's to preform file exfiltration

Since those features seem relativley easy to impliment, I want to go a step or two further. I also want to add;

- Geofencing to prevent scanning in specific areas (e.g. while at work, school, home)
- Automatic power

Prerequisites
-------------
Please ensure all following steps are complete prior to running 

Run PiDriver
------------
```
git clone https://github.com/Oppenheimer404/PiDriver.git
cd PiDriver
sudo ./PiDriver.py
```

Install PiDriver
----------------
To install 

this will install 'pidriver' to /usr/sbin/pidriver
