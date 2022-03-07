"""
Date: 2022.02.06 08:39
Description: Version Information
LastEditors: Rustle Karl
LastEditTime: 2022.02.08 18:48:27
"""
from datetime import datetime

# https://packaging.python.org/en/latest/guides/single-sourcing-package-version/

__version__ = "0.9.1"

GENERATOR_HEADER = """\
# Generated by [Toolkit-Py](https://github.com/fujiawei-dev/toolkit-py) Generator
# Created at {created_at}, Version {version}\
""".format(
    created_at=datetime.now(), version=__version__
)
