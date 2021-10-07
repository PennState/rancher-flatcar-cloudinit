# rancher-flatcar-cloudinit

This is a simple tool intended to be run via guestinfo when installing Flatcar Linux in vSphere using Rancher. It serves as a workaround for a problem at the intersection of Flatcar's cloud-init implementation and Rancher v2.6.

## Problem statement

Rancher v2.6 uses the `NoCloud` option for bootstrapping nodes under the vSphere driver. Because vSphere has no concept similar to a "metadata url" like most cloud providers, the only other options are `guestinfo`, which is exposed via VMWare Tools, and ISOs mounted via virtual CDROM drives.

However, the `coreos-cloudinit` service built-in to Flatcar is a stripped down, no-longer-maintained tool. Flatcar and other distros of CoreOS lineage much prefer Ignition Config over cloud-config. 

`coreos-cloudinit` does not detect and mount/read the `user-data` and `meta-data` files provided on the ISO because the device is not labeled correctly and the files are not at the expected path. Even if it did try to use these files, not all of the options in the cloud-config are supported by flatcar's minimal coreos-cloudinit.

As a workaround, `rancher-flatcar-cloudinit` was purpose-built to be run by `coreos-cloudinit`, which still supports executing scripts passed in via `guestinfo.coreos.config.data`. `rancher-flatcar-cloudinit` supports the very minimal number of options in Rancher's cloud-config, enough to allow Rancher to continue bootstrapping the node.

### Related issues

* https://github.com/flatcar-linux/Flatcar/issues/334
* https://github.com/rancher/rancher/issues/33374
* https://github.com/rancher/rancher/issues/26735
* https://github.com/rancher/rancher/issues/25336
* https://github.com/rancher/rancher/issues/24948
* 

## Usage

### guestinfo

Add a snippet like this to `guestinfo.coreos.config.data`:

```bash

#!/bin/bash
wget --timeout 300 -q https://github.com/PennState/rancher-flatcar-cloudinit/releases/download/v0.2.1/rancher-flatcar-cloudinit_0.2.1_linux_amd64 -O /tmp/rancher-flatcar-cloudinit
chmod 755 /tmp/rancher-flatcar-cloudinit
/tmp/rancher-flatcar-cloudinit
```

### OVA settings

If you are using the [official Flatcar Linux OVA](https://kinvolk.io/docs/flatcar-container-linux/latest/installing/cloud/vmware/) file, you must disable `vApp Options` after converting the VM to a Template. The `coreos-cloudinit` tool seems to prefer the OVF Environment data over data contained in `guestinfo`.

### ssh user

You will also need to edit the nodeTemplate in the API. Set `vmwarevsphereConfig.sshUser=rancher`. The default `docker` user already exists and cannot be modified.
