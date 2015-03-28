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

output sample
```
$ cf tree cf-dora
|--.bash_logout                                                                                                 [0/1905]
|--.bashrc
|--.profile
|--app
|  |--.buildpack-diagnostics
|  |  |--buildpack.log
|  |--.bundle
|  |  |--config
|  |  |--install.log
|  |--.gitignore
|  |--.java-buildpack.log
|  |--.profile.d
|  |  |--ruby.sh
|  |--.rspec
|  |--.ruby-version
|  |--Gemfile
|  |--Gemfile.lock
|  |--README.md
|  |--bin
|  |  |--erb
|  |  |--gem
...

```
Notice
------

* Too many call to files api endpoint

Development
-----------

```
cf uninstall-plugin tree; go get ./...; cf install-plugin $GOPATH/bin/tree
```
