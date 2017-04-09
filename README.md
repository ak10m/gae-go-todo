TODO List Application on GAE
=======================================

* go 1.6
* glide
* direnv
* goapp
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

### install goapp

```
$ brew install go-app-engine-64
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

or

```
$ goapp serve [PATH_TO_YOUR_APP]
```

