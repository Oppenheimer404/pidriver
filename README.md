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

Raspberry Pi Zero 2 W Setup
---------------------------
1. Download the latest version of `Raspberry Pi Imager` from [The Official Raspberry Pi Website](https://www.raspberrypi.com/software/)
2. Flash your selected microSD card with the latest version of Raspberry Pi OS.
**Note:** We will be using SSH to connect via another computer. Please enable SSH and provide network details during flash process.
3. SSH into your machine. `ssh foo@ipAddress`
4. Run `sudo apt update && sudo apt upgrade` to ensure everything is up to date. (This might take a while - Go refill your coffee or something...)
5. That's it! (Stay logged into SSH as all setup will be complete via CLI)

Run PiDriver from git Repository
--------------------------------
```bash
git clone https://github.com/Oppenheimer404/PiDriver.git
cd PiDriver
sudo ./PiDriver.py
```

Install PiDriver Using `pip`
--------------------------
To install onto your computer (so you can just run 'pidriver' from any terminal), run:

```bash
sudo pip install git+https://github.com/Oppenheimer404/PiDriver.git
```

This will install 'pidriver' to `/usr/sbin/pidriver` which should be in your `$PATH` variable using `python setup.py install`.

To uninstall, simply `sudo pip uninstall PiDriver`!

**Note:** Should you have any issues you may want to attempt to use the following command, although uninstalling is much more involved;

```bash
git clone https://github.com/Oppenheimer404/PiDriver.git
cd PiDriver
sudo python setup.py install 
```