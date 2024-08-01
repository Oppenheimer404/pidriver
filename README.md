# PiDriver

What is PiDriver
----------------
PiDriver is a utility built for the RaspberryPi Zero 2 W and written in Go programming language.

How to Set-Up PiDriver on a Raspberry Pi Zero 2 W
-------------------------------------------------

Install Go
-------------
1. Visit [The Go Doccumentation](https://go.dev/doc/install) and **download the latest version of Go**. (You can follow the steps found in the official docs - although I reccomend sticking to this unoffical guide as I will do my best to cover all steps in detail )`accurate as of go 1.22.5`
    - Ensure you download the **ARMv6** version as it is compatable with RaspberryPiOS.
2. **VERIFY YOUR HASHES**
    - Verifying a hash ensures that the file you downloaded isn't compromised.
    - To do this - run `sha256sum $filename` where `$filename` is the file you just downloaded.
    - Ensure the resulting hash matches the provided hash found [Here](https://go.dev/dl/).
3. Execute the following command to **remove any previous Go installation**. This will delete the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

```bash
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf $goinstallation
```
**Note:** You will likely need to add `sudo` before `rm` and `tar` to ensure you have the proper permissions.

4. Add /usr/local/go/bin to the PATH environment variable by adding the following line to your $HOME/.profile:

```bash
export PATH=$PATH:/usr/local/go/bin
```
**Note:** This is simply so that you can execute go binaries via terminal from anywhere locally.

5. **Reboot** to apply changes using `sudo reboot`

6. Verify go is installed by running `go version`
**Note:** You can now safely remove the .tar.gz installation file.

Setup Raspberry Pi Zero 2 W
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

```

Install PiDriver
----------------
To install onto your Raspberry Pi (so you can just run `pidriver` from any terminal), run:
