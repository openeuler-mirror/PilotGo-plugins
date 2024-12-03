# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 15:53:47 2024 +0800 
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
