# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from enum import Enum
from json import loads
from typing import Self


class ApiKeyOperation(str, Enum):
    """
    ApiKeyOperation
    """

    """
    allowed enum values
    """
    ADD = "add"
    DELETE = "delete"
    UPDATE = "update"

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ApiKeyOperation from a JSON string"""
        return cls(loads(json_str))