# xk6-exec

This is a [k6](https://go.k6.io/k6) extension using the
[xk6](https://github.com/grafana/xk6) system.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  xk6 build --with github.com/grafana/xk6-exec@latest
  ```

## Development
To make development a little smoother, use the `Makefile` in the root folder. The default target will format your code, run tests, and create a `k6` binary with your local code rather than from GitHub.

```bash
make
```
Once built, you can run your newly extended `k6` using:
```shell
 ./k6 run example.js
 ```

## Example

```javascript
// script.js
import exec from 'k6/x/exec';

export default function () {
  console.log(exec.command("date"));
  console.log(exec.command("ls",["-a","-l"]));
}
```

Result output:

```
$ ./k6 run example.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: ../example.js
     output: -

  scenarios: (100.00%) 1 scenario, 1 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1 iterations for each of 1 VUs (maxDuration: 10m0s, gracefulStop: 30s)

INFO[0000] vie 29 ene 2021 12:53:28 CET                  source=console
INFO[0000] total 27040
drwxrwxr-x 6 dgzlopes dgzlopes     4096 ene 29 12:52 .
drwxrwxr-x 3 dgzlopes dgzlopes     4096 ene 29 12:46 ..
-rw-rw-r-- 1 dgzlopes dgzlopes     8399 ene 29 12:45 builder.go
-rw-rw-r-- 1 dgzlopes dgzlopes     1871 ene 29 12:45 builder_test.go
drwxrwxr-x 3 dgzlopes dgzlopes     4096 ene 29 12:45 cmd
-rw-rw-r-- 1 dgzlopes dgzlopes     6842 ene 29 12:45 environment.go
drwxrwxr-x 8 dgzlopes dgzlopes     4096 ene 29 12:45 .git
drwxrwxr-x 3 dgzlopes dgzlopes     4096 ene 29 12:45 .github
-rw-rw-r-- 1 dgzlopes dgzlopes      118 ene 29 12:45 .gitignore
-rw-rw-r-- 1 dgzlopes dgzlopes      923 ene 29 12:45 .golangci.yml
-rw-rw-r-- 1 dgzlopes dgzlopes       85 ene 29 12:45 go.mod
-rw-rw-r-- 1 dgzlopes dgzlopes     1020 ene 29 12:45 .goreleaser.yml
-rw-rw-r-- 1 dgzlopes dgzlopes      183 ene 29 12:45 go.sum
-rwxrwxr-x 1 dgzlopes dgzlopes 27598848 ene 29 12:52 k6
-rw-rw-r-- 1 dgzlopes dgzlopes    11357 ene 29 12:45 LICENSE
-rw-rw-r-- 1 dgzlopes dgzlopes     1805 ene 29 12:45 platforms.go
-rw-rw-r-- 1 dgzlopes dgzlopes     3370 ene 29 12:45 README.md
drwxrwxr-x 3 dgzlopes dgzlopes     4096 ene 29 12:45 vendor  source=console

running (00m00.0s), 0/1 VUs, 1 complete and 0 interrupted iterations
default ✓ [======================================] 1 VUs  00m00.0s/10m0s  1/1 iters, 1 per VU

     data_received........: 0 B 0 B/s
     data_sent............: 0 B 0 B/s
     iteration_duration...: avg=2.3ms min=2.3ms med=2.3ms max=2.3ms p(90)=2.3ms p(95)=2.3ms
     iterations...........: 1   55.736622/s
```
