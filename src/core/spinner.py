import os
import sys
import time
import shutil
import random
import subprocess
from src.core.bcolors import *

class progress(object):
    def __init__(self, total_steps):
        self.total_steps = total_steps
        self.current_step = 0

    def spinner_update(self):
        self.current_step += 1
        progress = (self.current_step / self.total_steps) * 100
        print(bcolors.GREEN + f"\n("+ bcolors.ENDC + f"Installing:" + bcolors.DARKCYAN + "africana" + bcolors.GREEN + ")# " + bcolors.ENDC + f"[{int(progress)}%]" + bcolors.GREEN + f" {'█' * (int(progress) // 1)}", end="" + bcolors.ENDC)
        time.sleep(0.9)

    def spinner_updates(self):
        self.current_step += 1
        progress = (self.current_step / self.total_steps) * 100
        print(bcolors.GREEN + f"\n("+ bcolors.ENDC + f"Installing:" + bcolors.DARKCYAN + "africana" + bcolors.GREEN + ")# " + bcolors.ENDC + f"[{int(progress)}%]" + bcolors.GREEN + f" {'█' * (int(progress) // 1)}", end="" + bcolors.ENDC)
        time.sleep(9.0)

    def run_subprocess(self, command):
        process = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        while process.poll() is None:
            spinner.spinner_update()

    def run_subprocesx(self, command):
        process = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        while process.poll() is None:
            spinner.spinner_updates()

spinner = progress(total_steps=101)
if __name__ == "__main__":
    sys.exit(spiner())
