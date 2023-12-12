# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from re import match
from typing import Annotated, Any, ClassVar, Dict, List, Optional

from pydantic import BaseModel, Field, StrictInt, StrictStr, field_validator

from algoliasearch.insights.models.view_event import ViewEvent

try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class ViewedObjectIDs(BaseModel):
    """
    Use this event to track when users viewed items in the search results.
    """

    event_name: Annotated[str, Field(min_length=1, strict=True, max_length=64)] = Field(
        description="Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework. ",
        alias="eventName",
    )
    event_type: ViewEvent = Field(alias="eventType")
    index: StrictStr = Field(description="Name of the Algolia index.")
    object_ids: Annotated[List[StrictStr], Field(min_length=1, max_length=20)] = Field(
        description="List of object identifiers for items of an Algolia index.",
        alias="objectIDs",
    )
    user_token: Annotated[
        str, Field(min_length=1, strict=True, max_length=129)
    ] = Field(
        description="Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens. ",
        alias="userToken",
    )
    timestamp: Optional[StrictInt] = Field(
        default=None,
        description="Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp. ",
    )
    authenticated_user_token: Optional[StrictStr] = Field(
        default=None,
        description="User token for authenticated users.",
        alias="authenticatedUserToken",
    )
    __properties: ClassVar[List[str]] = [
        "eventName",
        "eventType",
        "index",
        "objectIDs",
        "userToken",
        "timestamp",
        "authenticatedUserToken",
    ]

    @field_validator("event_name")
    def event_name_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not match(r"[\x20-\x7E]{1,64}", value):
            raise ValueError(
                r"must validate the regular expression /[\x20-\x7E]{1,64}/"
            )
        return value

    @field_validator("user_token")
    def user_token_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if not match(r"[a-zA-Z0-9_=\/+-]{1,129}", value):
            raise ValueError(
                r"must validate the regular expression /[a-zA-Z0-9_=\/+-]{1,129}/"
            )
        return value

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ViewedObjectIDs from a JSON string"""
        return cls.from_dict(loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={},
            exclude_none=True,
        )
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ViewedObjectIDs from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "eventName": obj.get("eventName"),
                "eventType": obj.get("eventType"),
                "index": obj.get("index"),
                "objectIDs": obj.get("objectIDs"),
                "userToken": obj.get("userToken"),
                "timestamp": obj.get("timestamp"),
                "authenticatedUserToken": obj.get("authenticatedUserToken"),
            }
        )
        return _obj