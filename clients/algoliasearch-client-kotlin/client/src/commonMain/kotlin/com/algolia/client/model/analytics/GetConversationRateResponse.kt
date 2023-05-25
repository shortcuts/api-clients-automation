/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * GetConversationRateResponse
 *
 * @param rate The click-through rate.
 * @param trackedSearchCount The number of tracked search click.
 * @param conversionCount The number of converted clicks.
 * @param dates A list of conversion events with their date.
 */
@Serializable
public data class GetConversationRateResponse(

  /** The click-through rate. */
  @SerialName(value = "rate") val rate: Double,

  /** The number of tracked search click. */
  @SerialName(value = "trackedSearchCount") val trackedSearchCount: Int,

  /** The number of converted clicks. */
  @SerialName(value = "conversionCount") val conversionCount: Int,

  /** A list of conversion events with their date. */
  @SerialName(value = "dates") val dates: List<ConversionRateEvent>,
)