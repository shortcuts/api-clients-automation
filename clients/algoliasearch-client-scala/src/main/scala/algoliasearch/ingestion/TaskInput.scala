/** Ingestion API API powering the Data Ingestion connectors of Algolia.
  *
  * The version of the OpenAPI document: 1.0
  *
  * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
  * https://openapi-generator.tech Do not edit the class manually.
  */
package algoliasearch.ingestion

import org.json4s._

/** TaskInput
  */
sealed trait TaskInput

trait TaskInputTrait extends TaskInput

object TaskInput {}

object TaskInputSerializer extends Serializer[TaskInput] {
  override def deserialize(implicit format: Formats): PartialFunction[(TypeInfo, JValue), TaskInput] = {

    case (TypeInfo(clazz, _), json) if clazz == classOf[TaskInput] =>
      json match {
        case value: JObject => Extraction.extract[OnDemandDateUtilsInput](value)
        case value: JObject => Extraction.extract[ScheduleDateUtilsInput](value)
        case _              => throw new MappingException("Can't convert " + json + " to TaskInput")
      }
  }

  override def serialize(implicit format: Formats): PartialFunction[Any, JValue] = { case value: TaskInput =>
    value match {
      case value: OnDemandDateUtilsInput => Extraction.decompose(value)(format - this)
      case value: ScheduleDateUtilsInput => Extraction.decompose(value)(format - this)
    }
  }
}