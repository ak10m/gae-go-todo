TODO List Application on GAE
=======================================

* go 1.6
* [direnv](https://github.com/direnv/direnv)
* [goenv](https://github.com/syndbg/goenv)
* glide
* goapp
* Google App Engine
* Google Cloud Datastore


Getting start
---------------------------------------

### install go, direnv, glide, goapp

```
$ brew install go
$ brew install direnv
$ brew install glide
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

### prepare

```
$ cp .envrc{.example,}
$ goenv install $(cat .go-version)
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

