TODO List Application on GAE
=======================================

* go 1.6
* glide
* Google App Engine
* Google Cloud Datastore


Getting start
---------------------------------------

### use goenv

```
$ goenv install $(cat .go-version)
$ goenv rehash
$ direnv reload
$ go version
go version go1.6 darwin/amd64
```

### install glide

```
$ brew install glide
```

### Google Cloud SDK

See

* https://cloud.google.com/sdk/docs
* https://cloud.google.com/appengine/docs/standard/go/download

```
$ gcloud components update
$ gcloud components install app-engine-go
```
