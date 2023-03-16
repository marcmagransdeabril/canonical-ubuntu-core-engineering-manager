#!/bin/bash

# Fail on error and show debugging information

set -ex

# Trap errors

function handle_error {
  echo "Error occurred in line ${1}. Exiting."
  exit 1
}

trap 'handle_error $LINENO' ERR

# Install dependencies

sudo apt-get -y install git curl jq make qemu gcc bison flex libelf-dev libssl-dev qemu-system-x86-64 busybox-static

# Get the latest stable kernel version
latest_stable=$(curl -s https://www.kernel.org/releases.json | jq -r '.latest_stable.version')

# Get latest Linux kernel source code 
wget https://cdn.kernel.org/pub/linux/kernel/v6.x/linux-${latest_stable}.tar.xz
tar -xf linux-${latest_stable}.tar.xz

# Build the Linux kernel using the default configuration
cd linux-${latest_stable}
make arch=x86_64 defconfig
make arch=x86_64  -j $(nproc)
cd ..

# Create minimal initramfs image: https://wiki.gentoo.org/wiki/Custom_Initramfs#Prerequisites

sudo rm -rf initrd
mkdir initrd
cd initrd
sudo mkdir -p {bin,dev,etc,lib,lib64,mnt/root,proc,root,sbin,sys}
sudo cp -a /dev/{null,console,zero,ptmx,ttyS0,random,urandom} dev/

cat >> init << EOF
#!/bin/sh
mount -t proc none /proc
mount -t sysfs none /sys
echo
echo "Hello World!"
echo
setsid  cttyhack sh
exec /bin/sh
EOF

sudo chmod +x init

sudo cp /bin/busybox bin/busybox
sudo ln -s /bin/busybox bin/echo
sudo ln -s /bin/busybox bin/sh
sudo ln -s /bin/busybox bin/mount

find . -print0 | cpio --null --create --verbose --format=newc | gzip --best > ../initrd.img
cd ..


# Run QEMU with the newly built kernel
qemu-system-x86_64 \
-kernel linux-${latest_stable}/arch/x86_64/boot/bzImage  \
-nographic \
-append "console=ttyS0" \
-initrd  initrd.img \
-m 2G 

