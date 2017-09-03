# [WIP] Atomic Registry Image Manager for Openshift - ARIMO

# Purpose:

1. Help sync images from third party registries
2. Bootstrap images using your template - TBC
3. Automate image syncing on time basis - TBC


# Run:

Create config file:

config.yaml

```
sourceregistry:
  url: registry.access.redhat.com
  username: 
  password:
  type: docker
destinationregistry:
  url: docker.io
  username: mangirdas
  password: asdasd
  type: docker
insecurepolicy: true
registriesdirpath: registries.d
dockerinsecureskiptlsverify: true
namespacemap:
- source: openshift3
  destination: mangirdas
imagelist:
- openshift3/ose-ansible:latest
- openshift3/jenkins-2-rhel7
- openshift3/jenkins-slave-maven-rhel7
- openshift3/jenkins-slave-nodejs-rhel7

#running the binary

./arimo

```

# Compile:

pre-reqs:
```
sudo dnf install gpgme-devel libassuan-devel glibc-devel glib2-devel rpm-ostree gcc ostree device-mapper-devel btrfs-progs-devel ostree-devel libostree-dev -y
go build -tags="containers_image_openpgp"

```

# Vendoring

For vendoring we use `vndr` package

# TODO:

```
1. Add tests
2. Add image bootstrap (add one more custom organization layer)
3. Add paralelism for copy (add into regman libary)
4. Add signing, re-signing, security, certificate, policies.

```