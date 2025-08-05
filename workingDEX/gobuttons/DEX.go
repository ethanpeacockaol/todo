import subprocess
import os
subprocess.run(["tmux", "new", "-d", "-s", "WiNdOwS"])

subprocess.run(["tmux", "send-keys", "-t", "WiNdOwS", "pd sh ubuntu --termux-home", "C-m"])
subprocess.run(["tmux", "send-keys", "-t", "WiNdOwS", "vncserver -kill", "C-m"])
subprocess.run(["tmux", "send-keys", "-t", "WiNdOwS", "rm -rf /tmp/.X*", "C-m"])
subprocess.run(["tmux", "send-keys", "-t", "WiNdOwS", "vncserver -xstartup /usr/bin/startlxqt -geometry 720x760", "C-m"])


import time
time.sleep(1)
os.system("am start --user 0 -n com.realvnc.viewer.android/com.realvnc.viewer.android.app.ConnectionChooserActivity")

#subprocess.run(["tmux", "send-keys", "-t", session_name, command, "C-m"], check=True)

#subprocess.run(["tmux", "new-session", "-d", "-s", session_name], check=True)



        
