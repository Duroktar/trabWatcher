Watcher
=======

Runs a command whenever a watched file changes.

- License: MIT


Install
-------

.. code-block:: fish
    
    go get github.com/Duroktar/trabWatcher
    
That should place a `trabWatcher` executable into $GOPATH/bin


Usage
-----

These assume you are already in the desired working directory.

**fish Shell**

- Simple example

.. code-block:: fish

    trabWatcher "go run main.go" main.go


- More advanced with Wildcards

.. code-block:: fish

    # golang
    set a "go run" *.go
    trabWatcher "$a" *.go

    # python
    trabWatcher "python script.py" *.py

    # watch all files
    trabWatcher "python script.py" *.*


TODO
----

Here's a rundown of suggestions_ from the golang community on reddit. I'll mark em off as I get to em. Thanks!

- **DONE** Refactor exports (variables/functions/etc..)
- **DONE** Prefer short-cicuiting inside of functions to keep indents lean.
- **DONE** Put a stop method on watcher

.. raw:: html

    <s>

- **DONE** It was suggested I put watcher in a seperate package. But I think I'll end up just put it back in main.go. I don't plan on anyone using this *too* seriously and it seems like the way to go for something intended to be used as a binary (correct me if I'm wrong)-. 

.. raw:: html

    </s>

- **DONE** Regarding the above crossed-out section. I realize now that folders in the main package are considered seperate packages. I for some reason thought they had to be a seperate repo. So I ended up re-refactoring watcher as originally suggested. (Consequently this solved the problem I was having naming packages anything other than main)

- Comments. More of em.
- Get feedback, learn and move on. There are *many* tools already in the wild that can do this sort of thing and I'd rather not dwell.

I don't think I'll use fsnotify as suggested, because it would probably end up replacing the whole Watcher class and, this was more or less just for a first project kinda thing. I did end up taking a few points from their implementation however.

.. _suggestions: https://www.reddit.com/r/golang/comments/69j0lm/i_wrote_my_first_golang_program_may_i_ask_for_a/

Copyright & License
-------------------

Copyright 2017 Scott Doucet

Code released under the MIT License.

    