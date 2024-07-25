
# This is Pseudocode for my PiDriver Project

## PiDriver is a wardriving utility intended to automate all aspects of
# scanning, logging, and uploading via the WigleAPI.

## Below is a list of variables which will likely be contained in a piDrive.conf file & what they correspond to
# GPS settings
gps-device			-	GPS device to be used with kismet
gps-scan-rate		-	GPS baudrate settings (how fast to scan GPS)
# Wifi settings
source-adapter		-	Main wifi adapter used for general network connection
logging-adapter		-	Secondary wifi adapter used for kismet logging (Default to source-adapter if empty)
# Logging settings
log-format			-	Format for kismet output
log-keep-files		-	How many capture files to store in local storage
log-max-size		-	Maximum capture file size (Begins new file after max)
# Upload settings
upload-network		-	Identifier for network to post captures to Wigle
upload-interval		-	How often to scan for upload-network and attempt file upload (Uploads all new files)
