{
  function exec1(cmd, env = {}) {
    return jssh.exec(cmd, env, 1);
  }

  function exec2(cmd, env = {}) {
    return jssh.exec(cmd, env, 2);
  }

  jssh.ssh.exec1 = function exec1(cmd, env = {}) {
    return jssh.ssh.exec(cmd, env, 1);
  };

  jssh.ssh.exec2 = function exec2(cmd, env = {}) {
    return jssh.ssh.exec(cmd, env, 2);
  };

  jssh.exec1 = exec1;
  jssh.exec2 = exec2;
  Object.freeze(jssh.ssh);
}
