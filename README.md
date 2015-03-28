Overview
========

application tree command

Installation
------------

```
$ go get github.com/morikat/cf-plugin-tree
$ cf install-plugin $GOPATH/bin/tree
```

Usage
-----

```
$ cf tree <appname>
```

Development
-----------

```
cf uninstall-plugin tree; go get ./...; cf install-plugin $GOPATH/bin/tree
```
