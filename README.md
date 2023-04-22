# prog

`prog` is a progress meter that acts as an intermediary who counts the size of data transferred between two piped programs on *nix.

The output of `prog` is printed in stderr. This is useful when you are fucking waiting for some data processing through piped programs to get done but got no output.

## Example

```
$ dd if=/dev/urandom of=sample bs=1G count=1
<output truncated>

$ cat sample | prog | cat > /dev/null
[prog] Elapsed 989.505407ms, transferred 1.00 GiB, avg 1.01 GiB/s

$ watch -n 1 ls -la | prog | cat > /dev/null
[prog] Elapsed 19.0s, transferred 1.00 KiB, avg 53.86 B/s
```
