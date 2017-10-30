# Kubernetes CSI Volume Plugin

This is an example CSI plugin for Kubernetes which uses the FLEX volume type.  It's under heavy development, so at this time README.md is notes for the developers coding.  Once complete this will change to something user friendly.
**First Steps**

Setup the project source for dev:

```bash
# export GOPATH=/opt/go
mkdir -p $GOPATH/src/github.com/csi-volumes/
cd $GOPATH/src/github.com/csi-volumes/
git clone git@github.com:childsb/kubernetes-csi.git
cd $GOPATH/src/github.com/csi-volumes/kubernetes-csi/
```


**To Build**

```bash
make
```

I wasn't sure if we should split the binaries out or have one mono binary.  So for now, the provisioner binary is in cmd/flex-provisioner, and i manually build the single binary in Makefile.  I also used Glide for dep. management.


This is nowhere near finished, i just stubbed out the provisioner piece and project structure.


**To Test**

`cd test ; go test -v`  # This is BROKEN :(   glide dependencies stepping on each other..





