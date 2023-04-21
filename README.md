# prog

`prog` is a progress meter that acts as an intermediary who counts the size of data transferred between two piped programs on *nix.

The output of `prog` is printed in stderr.

## Example

```
$ dd if=/dev/urandom of=sample bs=1G count=1
$ cat sample | prog | cat > /dev/null
```
