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

Install PiDriver
----------------
To install onto your Raspberry Pi (so you can just run `pidriver` from any terminal), run:

Using `dpkg -i`
-----------

```bash
git clone https://github.com/Oppenheimer404/PiDriver.git
cd PiDriver
sudo dpkg -i pidriver.deb
```

**Note:** To uninstall use `sudo dpkg -P pidriver` to purge everything - or use `sudo dpkg -r pidriver` to only remove the pidriver package.

Using `apt`
-----------
**Note:** Not yet working

```bash
sudo apt install NOTAVALIABLE
```
