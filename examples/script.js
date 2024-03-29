import exec from "k6/x/exec";

export default function () {
  console.log(exec.command("date"));

  // Changes the directory before running the command
  console.log(
    exec.command("ls", ["-a", "-l"], {
      dir: "sub-directory", // optional directory in which the command has to be run
    }),
  );

  // Sets an environment variable
  console.log(
    exec.command("bash", ["-c", "echo $FOO"], {
      env: ["FOO=bar"],
    }),
  );

  // Allows to change behavior when command fails.
  // This will make the stderr be returned instead of the stdout
  console.log(
    exec.command("false", [], {
      fatalError: false,
    }),
  );
}
