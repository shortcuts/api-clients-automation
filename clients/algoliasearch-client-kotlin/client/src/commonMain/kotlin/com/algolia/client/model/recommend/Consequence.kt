/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.recommend

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Effect of the rule.
 *
 * @param hide Exclude items from recommendations.
 * @param promote Place items at specific positions in the list of recommendations.
 * @param params
 */
@Serializable
public data class Consequence(

  /** Exclude items from recommendations. */
  @SerialName(value = "hide") val hide: List<HideConsequenceObject>? = null,

  /** Place items at specific positions in the list of recommendations. */
  @SerialName(value = "promote") val promote: List<PromoteConsequenceObject>? = null,

  @SerialName(value = "params") val params: ParamsConsequence? = null,
)
