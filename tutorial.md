# Tutorial

## description

The subscan plugin gives more possibilities to subscan-essentials. Subscan sends the Extrinsic or Event received by the fetch block to the plugin, which can be customized and processed into formatted data, stored in the database, and presented through http api.

## Quick start

### Install gen tool

```shell script
$ go get github.com/subscan-explorer/subscan-plugin/tools/gen-plugin
```

### Gen plugin

```shell script
# assume need generate balance module
$ gen-plugin balance
```

The directory structure is as follows:
```
├── dao
│   └── dao.go
├── http
│   └── http.go
├── main.go
├── model
│   └── model.go
└── service
    └── service.go

4 directories, 5 files
```

The logic can refer to https://github.com/subscan-explorer/subscan-essentials/tree/master/plugins/balance

### Register plugin

```go
package plugin
func init() {
	registerNative(balance.New())
	registerNative(system.New())
}
```


