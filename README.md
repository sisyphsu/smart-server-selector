# smart-server-selector

`smart-server-selector` is an efficiency terminal tool for backend system development engineer,
especially for someone who have bunch of servers to switch.

It was developed with `golang`, which means it's very clean, no dependency, cross-platform.

In terms of features, it support keyword search, which could help you find server by hostname, ip, description, etc quickly.

Hope it could save your time.

# Demo

![demo](./demo.gif)

# Install

You can find the download links at [this page](https://github.com/sisyphsu/smart-server-selector/releases).

`smart-server-selector` support these platforms (**May not be updated timely**): 

+ [`darwin-amd64`](https://github.com/sisyphsu/smart-server-selector/releases/download/1.0.0/smart-server-selector-darwin-10.6-amd64)
+ [`linux-amd64`](https://github.com/sisyphsu/smart-server-selector/releases/download/1.0.0/smart-server-selector-linux-amd64)
+ [`linux-386`](https://github.com/sisyphsu/smart-server-selector/releases/download/1.0.0/smart-server-selector-linux-386)
+ [`windows-amd64`](https://github.com/sisyphsu/smart-server-selector/releases/download/1.0.0/smart-server-selector-windows-4.0-amd64.exe)
+ [`windows-386`](https://github.com/sisyphsu/smart-server-selector/releases/download/1.0.0/smart-server-selector-windows-4.0-386.exe)

Installation is very easy.

You can use `wget` or `curl` download it, example for `linux-amd64`:

```bash
wget https://github.com/sisyphsu/smart-server-selector/releases/download/{version}/smart-server-selector-linux-amd64

curl https://github.com/sisyphsu/smart-server-selector/releases/download/{version}/smart-server-selector-linux-amd64 > smart-server-selector-linux-amd64 

chmod +x smart-server-selector-linux-amd64

mv smart-server-selector-linux-amd64 ~/.local/bin/sss 
```

Above steps download the `smart-server-selector` in to `~/.local/bin`, and name it `sss`, 
which is more convenient for keyboard inputting. 

You should add `~/.local/bin` directory into your `$PATH`, or using other `PATH` directory, 
this isn't a prerequisite, but you better do it. 

The whole process didn't need any additional permission like `root`.

# Configuration

After started, `smart-server-selector` will load servers from `~/.servers`, you can edit this file directly.

These two format configurations are valid:

```base
# comments, empty line is ok.
test    10.10.10.1   description
test    10.10.10.1   22     username   description
```

Explain:

+ `test`: environment name, like test/pre/prod, no limit.
+ `10.10.10.1`: host name, could `ip` or `hostname`.
+ `22`: ssh port
+ `username`: ssh login name
+ `description`: any text

If your configuration is invalid, `smart-server-selector` will ignore it and print notice info.

# Advice

If your server's `ssh port` or `ssh username` don't match the default value, 
then you should config it in global ssh config(`~/.ssh/config`), for example:

```bash
Host *
        Port 9876
        User other-name
        TCPKeepAlive yes
        ServerAliveInterval 60
        StrictHostKeyChecking no
```

This way could keep the server-list clean, and no need to config `port` or `user` for every server.

For more details, [check this link](https://www.ssh.com/ssh/config/).

# License

MIT