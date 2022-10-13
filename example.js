import exec from 'k6/x/exec';

export default function () {
  console.log(exec.command("date"));
  console.log(exec.command("ls",["-a","-l"], {
    "dir": "sub-directory" // optional directory in which the command has to be run
  }));
}