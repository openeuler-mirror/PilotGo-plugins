import subprocess


def ExecuteCommand(command):
    try:
        result = subprocess.run(
            command, shell=True, text=True, capture_output=True, check=True
        )
        return (result.stdout, result.returncode)
    except subprocess.CalledProcessError as e:
        print(f"命令执行失败: {e}")
        return (e.stderr, e.returncode)
