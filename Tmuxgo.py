import subprocess
import os
subprocess.run(["tmux", "new", "-d", "-s", "WiNdOwS"])

subprocess.run(["tmux", "send-keys", "-t", "WiNdOwS", "pd sh ubuntu", "C-m"])

os.system("am start --user 0 -n com.realvnc.viewer.android/com.realvnc.viewer.android.app.ConnectionChooserActivity")

#subprocess.run(["tmux", "send-keys", "-t", session_name, command, "C-m"], check=True)

#subprocess.run(["tmux", "new-session", "-d", "-s", session_name], check=True)



        