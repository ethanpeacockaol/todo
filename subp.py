import subprocess

# Define the Zsh command you want to run
zsh_command = "echo 'Hello from Zsh!'"

# Start Zsh and execute the command
subprocess.run(["pd sh ubuntu --termux-home", "-c", zsh_command])
