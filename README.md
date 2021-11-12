# oci-runtime

 Simple oci client how implement oci runtime lifecycle operations


* [Installation](#installation)
* [Compile](#compile)
* [Usage](#usage)
* [Contribution](#Contribution)


## Installation

```shell
git clone git@github.com:chen-keinan/oci-runtime.git
```

## Compile
```shell
go build 
```

## Usage
Create 2 folders bundles :
```shell
mkdir ~/bundles
mkdir ~/containers
```
Copy bundle file to `~/bundles` folder example:
```shell
cp redis.tgz ~/bundles/redis.tgz
```
### Operations
Create redis runtime (bundle must exist in bundle folder otherwise an error will be thrown)
```shell
./oci-runtime create <container id> <bundle name>
```
example:
```shell
/oci-runtime create 12345 redis
```
Start redis runtime (container must have created state otherwise an error will be thrown)
```shell
./oci-runtime start <container id> 
```
example:
```shell
/oci-runtime start 12345 
```
List redis runtime
```shell
./oci-runtime state <container id> 
```
example:
```shell
/oci-runtime state 12345 
```
```shell
ID  	STATUS 	BUNDLE	        PID        	VERSION
12345	created	redis 	5577006791947779410	    1.0
```

Stop redis runtime (container must have running state otherwise an error will be thrown)
```shell
./oci-runtime kill <container id> 
```
example:
```shell
/oci-runtime kill 12345 
```
Delete redis runtime (container must have stopped state otherwise an error will be thrown)
```shell
./oci-runtime delete <container id> 
```
example:
```shell
/oci-runtime delete 12345 
```
