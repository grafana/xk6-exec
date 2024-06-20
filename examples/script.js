import exec from 'k6/x/exec';

export default function () {
  // Basic example:
  console.log(exec.command("date"));
  
  // With custom error handling:
  try {
    var output = exec.command("ls",["-a", "NO_SUCH_DIR"], {
      "continue_on_error": true
    });
  } catch (e) {
        console.log("ERROR: " + e);
        if (e.value && e.value.stderr) {
                console.log("STDERR: " + String.fromCharCode.apply(null, e.value.stderr))
        }
  }

  // without error handling the test will stop when the following command fails
  console.log(exec.command("ls",["-a","-l"], {
    "dir": "sub-directory" // optional directory in which the command has to be run
  }));

  console.log("this message will not be printed")
}
