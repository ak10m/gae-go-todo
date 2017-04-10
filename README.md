TODO List Application on GAE
=======================================

* [go](https://github.com/golang/go)
* [dep](https://github.com/golang/dep)
* Google App Engine
* Google Cloud Datastore


Getting start
---------------------------------------

### install Golang

* https://golang.org/dl/

### install dep

```
go get -u github.com/golang/dep/...
```

#### usage

```
$ dep ensure -update
```

... https://github.com/golang/dep/issues/149

```
$ rm lock.json manifest.json
$ dep init
```


### install Google Cloud SDK

See

* https://cloud.google.com/sdk/docs
* https://cloud.google.com/appengine/docs/standard/go/download

```
$ gcloud components update
$ gcloud components install app-engine-go
```

### run dev server

See https://cloud.google.com/appengine/docs/standard/go/tools/using-local-server

```
$ dev_appserver.py [PATH_TO_YOUR_APP]
```

TODO
---------------------------------------


### Testing

* https://cloud.google.com/appengine/docs/standard/go/tools/localunittesting/
* https://github.com/stretchr/testify

### Lint

* https://github.com/alecthomas/gometalinter


