# tab2spaces

The opinionated-Gofmt can always save developer's precious time, but in some case, we need to read those formatted code in non-editor place, the "tab" becomes so annoying! The indents maybe 8 spaces or 3 spaces or whatever, and you have no easy way to change Gofmt configurations, that such pain in the neck.

For the love of god, we need this.


```commandline
$ GO111MODULE=off go get github.com/HuuLane/tab2spaces
$ tab2spaces file1 file2
will read [file1 file2]
...
...
Successfully write to your clipboard!
# All the tabs have been replaced to 4 spaces, just paste it.
```