import exec from 'k6/x/exec';

export default function () {
  console.log(exec.command("date"));
  console.log(exec.command("ls",["-a","-l"]));
}