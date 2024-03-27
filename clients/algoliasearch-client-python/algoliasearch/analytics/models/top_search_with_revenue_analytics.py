# coding: utf-8

"""
Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
"""
from __future__ import annotations

from json import loads
from typing import Annotated, Any, Dict, List, Optional, Self, Union

from pydantic import BaseModel, Field, StrictInt, StrictStr

from algoliasearch.analytics.models.click_positions_inner import ClickPositionsInner
from algoliasearch.analytics.models.currencies_value import CurrenciesValue


class TopSearchWithRevenueAnalytics(BaseModel):
    """
    TopSearchWithRevenueAnalytics
    """

    search: StrictStr = Field(description="Search query.")
    count: StrictInt = Field(description="Number of searches.")
    click_through_rate: Optional[
        Union[
            Annotated[float, Field(le=1, strict=True, ge=0)],
            Annotated[int, Field(le=1, strict=True, ge=0)],
        ]
    ] = Field(
        description="Click-through rate, calculated as number of tracked searches with at least one click event divided by the number of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true. ",
        alias="clickThroughRate",
    )
    average_click_position: Optional[
        Union[
            Annotated[float, Field(strict=True, ge=1)],
            Annotated[int, Field(strict=True, ge=1)],
        ]
    ] = Field(
        description="Average position of a clicked search result in the list of search results. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true. ",
        alias="averageClickPosition",
    )
    click_positions: Annotated[
        List[ClickPositionsInner], Field(min_length=12, max_length=12)
    ] = Field(
        description="List of positions in the search results and clicks associated with this search.",
        alias="clickPositions",
    )
    conversion_rate: Optional[
        Union[
            Annotated[float, Field(le=1, strict=True, ge=0)],
            Annotated[int, Field(le=1, strict=True, ge=0)],
        ]
    ] = Field(
        description="Conversion rate, calculated as number of tracked searches with at least one conversion event divided by the number of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true. ",
        alias="conversionRate",
    )
    tracked_search_count: StrictInt = Field(
        description="Number of tracked searches. Tracked searches are search requests where the `clickAnalytics` parameter is true.",
        alias="trackedSearchCount",
    )
    click_count: Annotated[int, Field(strict=True, ge=0)] = Field(
        description="Number of clicks associated with this search.", alias="clickCount"
    )
    conversion_count: Annotated[int, Field(strict=True, ge=0)] = Field(
        description="Number of conversions from this search.", alias="conversionCount"
    )
    nb_hits: StrictInt = Field(description="Number of results (hits).", alias="nbHits")
    currencies: Dict[str, CurrenciesValue] = Field(
        description="Revenue associated with this search, broken-down by currencies."
    )
    add_to_cart_rate: Optional[
        Union[
            Annotated[float, Field(le=1, strict=True, ge=0)],
            Annotated[int, Field(le=1, strict=True, ge=0)],
        ]
    ] = Field(
        description="Add-to-cart rate, calculated as number of tracked searches with at least one add-to-cart event divided by the number of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true. ",
        alias="addToCartRate",
    )
    add_to_cart_count: Annotated[int, Field(strict=True, ge=0)] = Field(
        description="Number of add-to-cart events from this search.",
        alias="addToCartCount",
    )
    purchase_rate: Optional[
        Union[
            Annotated[float, Field(le=1, strict=True, ge=0)],
            Annotated[int, Field(le=1, strict=True, ge=0)],
        ]
    ] = Field(
        description="Purchase rate, calculated as number of tracked searches with at least one purchase event divided by the number of tracked searches. If null, Algolia didn't receive any search requests with `clickAnalytics` set to true. ",
        alias="purchaseRate",
    )
    purchase_count: StrictInt = Field(
        description="Number of purchase events from this search.", alias="purchaseCount"
    )

    model_config = {"populate_by_name": True, "validate_assignment": True}

    def to_json(self) -> str:
        return self.model_dump_json(by_alias=True, exclude_unset=True)

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of TopSearchWithRevenueAnalytics from a JSON string"""
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
        _items = []
        if self.click_positions:
            for _item in self.click_positions:
                if _item:
                    _items.append(_item.to_dict())
            _dict["clickPositions"] = _items
        _field_dict = {}
        if self.currencies:
            for _key in self.currencies:
                if self.currencies[_key]:
                    _field_dict[_key] = self.currencies[_key].to_dict()
            _dict["currencies"] = _field_dict
        # set to None if click_through_rate (nullable) is None
        # and model_fields_set contains the field
        if (
            self.click_through_rate is None
            and "click_through_rate" in self.model_fields_set
        ):
            _dict["clickThroughRate"] = None

        # set to None if average_click_position (nullable) is None
        # and model_fields_set contains the field
        if (
            self.average_click_position is None
            and "average_click_position" in self.model_fields_set
        ):
            _dict["averageClickPosition"] = None

        # set to None if conversion_rate (nullable) is None
        # and model_fields_set contains the field
        if self.conversion_rate is None and "conversion_rate" in self.model_fields_set:
            _dict["conversionRate"] = None

        # set to None if add_to_cart_rate (nullable) is None
        # and model_fields_set contains the field
        if (
            self.add_to_cart_rate is None
            and "add_to_cart_rate" in self.model_fields_set
        ):
            _dict["addToCartRate"] = None

        # set to None if purchase_rate (nullable) is None
        # and model_fields_set contains the field
        if self.purchase_rate is None and "purchase_rate" in self.model_fields_set:
            _dict["purchaseRate"] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of TopSearchWithRevenueAnalytics from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate(
            {
                "search": obj.get("search"),
                "count": obj.get("count"),
                "clickThroughRate": obj.get("clickThroughRate"),
                "averageClickPosition": obj.get("averageClickPosition"),
                "clickPositions": (
                    [
                        ClickPositionsInner.from_dict(_item)
                        for _item in obj.get("clickPositions")
                    ]
                    if obj.get("clickPositions") is not None
                    else None
                ),
                "conversionRate": obj.get("conversionRate"),
                "trackedSearchCount": obj.get("trackedSearchCount"),
                "clickCount": obj.get("clickCount"),
                "conversionCount": obj.get("conversionCount"),
                "nbHits": obj.get("nbHits"),
                "currencies": (
                    dict(
                        (_k, CurrenciesValue.from_dict(_v))
                        for _k, _v in obj.get("currencies").items()
                    )
                    if obj.get("currencies") is not None
                    else None
                ),
                "addToCartRate": obj.get("addToCartRate"),
                "addToCartCount": obj.get("addToCartCount"),
                "purchaseRate": obj.get("purchaseRate"),
                "purchaseCount": obj.get("purchaseCount"),
            }
        )
        return _obj