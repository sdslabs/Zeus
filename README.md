# Zeus

## Lightning Fast Containers

### Description

Zeus is a container runtime built in Golang. The low-level runtime is `runz`.

### Features

- **Lightweight**: Minimal overhead for container operations.

### Issues

- **No Networking**: Containers cannot communicate with the outside world.
- **UTS Namespace**: Containers cannot change their hostname.
- **IPC Namespace**: Containers cannot communicate with each other.
- **Issues found**: Check TODOs in the code.

### Installation

- Get inside `filesystems` directory and run the arch or ubuntu script after geting docker.
- Run the following to get the script to a workable state
```bash
./install.sh
```
- Then,to test the installation, run the following
- This tests resource isolation, PID namespace, and memory limits.
- Shows the completion of my minor
```bash
sudo cp ./testing/cpumemeater ./filesystems/arch/bin/   # Eats up CPU for testing
sudo cp ./testing/forker ./filesystems/arch/bin/        # Makes multiple child processes to fill up PID namespace and memory
sudo cp ./testing/leach ./filesystems/arch/bin/         # Eats up memory for testing
```

