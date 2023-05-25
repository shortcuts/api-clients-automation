/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.predict

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * SegmentParentConditionOperands
 */
@Serializable(SegmentParentConditionOperandsSerializer::class)
public sealed interface SegmentParentConditionOperands {

  public companion object {

    /**
     * Nested segment conditions that only contain operands.
     *
     * @param `operator`
     * @param operands
     */
    public fun SegmentChildConditions(
      `operator`: SegmentConditionOperator,
      operands: List<SegmentChildConditionOperands>,
    ): SegmentChildConditions = com.algolia.client.model.predict.SegmentChildConditions(
      `operator` = `operator`,
      operands = operands,
    )

    /**
     * Operand for affinity model predictions.
     *
     * @param name
     * @param filters
     */
    public fun SegmentOperandAffinity(
      name: String,
      filters: List<SegmentAffinityFilter>,
    ): SegmentOperandAffinity = com.algolia.client.model.predict.SegmentOperandAffinity(
      name = name,
      filters = filters,
    )

    /**
     * Operand for funnel stage model predictions.
     *
     * @param name
     * @param filters
     */
    public fun SegmentOperandFunnelStage(
      name: String,
      filters: List<SegmentFunnelStageFilter>,
    ): SegmentOperandFunnelStage = com.algolia.client.model.predict.SegmentOperandFunnelStage(
      name = name,
      filters = filters,
    )

    /**
     * Operand for order value model predictions.
     *
     * @param name
     * @param filters
     */
    public fun SegmentOperandOrderValue(
      name: String,
      filters: List<SegmentOrderValueFilter>,
    ): SegmentOperandOrderValue = com.algolia.client.model.predict.SegmentOperandOrderValue(
      name = name,
      filters = filters,
    )

    /**
     * Operand for user profile properties.
     *
     * @param name
     * @param filters
     */
    public fun SegmentOperandProperty(
      name: String,
      filters: List<SegmentPropertyFilter>,
    ): SegmentOperandProperty = com.algolia.client.model.predict.SegmentOperandProperty(
      name = name,
      filters = filters,
    )
  }
}

internal class SegmentParentConditionOperandsSerializer : KSerializer<SegmentParentConditionOperands> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("SegmentParentConditionOperands")

  override fun serialize(encoder: Encoder, value: SegmentParentConditionOperands) {
    when (value) {
      is SegmentChildConditions -> SegmentChildConditions.serializer().serialize(encoder, value)
      is SegmentOperandAffinity -> SegmentOperandAffinity.serializer().serialize(encoder, value)
      is SegmentOperandFunnelStage -> SegmentOperandFunnelStage.serializer().serialize(encoder, value)
      is SegmentOperandOrderValue -> SegmentOperandOrderValue.serializer().serialize(encoder, value)
      is SegmentOperandProperty -> SegmentOperandProperty.serializer().serialize(encoder, value)
    }
  }

  override fun deserialize(decoder: Decoder): SegmentParentConditionOperands {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize SegmentChildConditions
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<SegmentChildConditions>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SegmentChildConditions (error: ${e.message})")
      }
    }

    // deserialize SegmentOperandAffinity
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<SegmentOperandAffinity>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SegmentOperandAffinity (error: ${e.message})")
      }
    }

    // deserialize SegmentOperandFunnelStage
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<SegmentOperandFunnelStage>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SegmentOperandFunnelStage (error: ${e.message})")
      }
    }

    // deserialize SegmentOperandOrderValue
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<SegmentOperandOrderValue>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SegmentOperandOrderValue (error: ${e.message})")
      }
    }

    // deserialize SegmentOperandProperty
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<SegmentOperandProperty>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize SegmentOperandProperty (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}