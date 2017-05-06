Watcher
=======

Runs a command whenever a watched file changes.

- License: MIT


Install
-------

Clone into `$GOPATH/src/github.com/Duroktar/watcher` then run..

.. code-block:: fish
    
    go install github.com/Duroktar/watcher
    
That should place watcher into $GOPATH/bin


Usage
-----

These assume you are already in the desired working directory.

**fish Shell**

- Simple example

.. code-block:: fish

    watcher "go run main.go" main.go


- More advanced with Wildcards

.. code-block:: fish

    # golang
    set a "go run" *.go
    watcher "$a" *.go

    # python
    watcher "python script.py" *.py

    # watch all files
    watcher "python script.py" *.*


Copyright & License
-------------------

Copyright 2017 Scott Doucet

Code released under the MIT License.

    