#!/bin/bash

# Removes everything that is created by the build.sh script.

set -e -x

rm -rf out
rm -rf libplist
rm -rf libusbmuxd
rm -rf libimobiledevice
rm -rf ifuse
rm -rf ideviceinstaller
rm -rf usbmuxd
