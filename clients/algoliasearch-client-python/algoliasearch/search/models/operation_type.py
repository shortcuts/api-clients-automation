# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class OperationType(str, Enum):
    """
    Operation to perform (_move_ or _copy_).
    """

    """
    allowed enum values
    """
    MOVE = "move"
    COPY = "copy"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of OperationType from a JSON string"""
        return cls(loads(json_str))