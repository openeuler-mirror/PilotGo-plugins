# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 15:53:47 2024 +0800 
import logging


def setup_logger():
    logging.basicConfig(
        format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S",
        level=logging.INFO,
    )
    logger = logging.getLogger(__name__)
    return logger
