/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * DailyUsers
 *
 * @param date Date in the format YYYY-MM-DD.
 * @param count Number of unique users.
 */
@Serializable
public data class DailyUsers(

  /** Date in the format YYYY-MM-DD. */
  @SerialName(value = "date") val date: String,

  /** Number of unique users. */
  @SerialName(value = "count") val count: Int,
)