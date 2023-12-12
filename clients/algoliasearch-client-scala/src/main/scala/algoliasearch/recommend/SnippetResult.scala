/** Recommend API The Recommend API lets you generate recommendations with several AI models. > **Note**: You should use
  * Algolia's [libraries and
  * tools](https://www.algolia.com/doc/guides/getting-started/how-algolia-works/in-depth/ecosystem/) to interact with
  * the Recommend API. Using the HTTP endpoints directly is not covered by the
  * [SLA](https://www.algolia.com/policies/sla/).
  *
  * The version of the OpenAPI document: 1.0.0
  *
  * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
  * https://openapi-generator.tech Do not edit the class manually.
  */
package algoliasearch.recommend

import algoliasearch.recommend.MatchLevel._

import org.json4s._

/** SnippetResult
  */
sealed trait SnippetResult

trait SnippetResultTrait extends SnippetResult

object SnippetResult {

  case class MapOfStringSnippetResultOption(value: Map[String, SnippetResultOption]) extends SnippetResult

  def apply(value: Map[String, SnippetResultOption]): SnippetResult = {
    SnippetResult.MapOfStringSnippetResultOption(value)
  }
}

object SnippetResultSerializer extends Serializer[SnippetResult] {
  override def deserialize(implicit format: Formats): PartialFunction[(TypeInfo, JValue), SnippetResult] = {

    case (TypeInfo(clazz, _), json) if clazz == classOf[SnippetResult] =>
      json match {
        case value: JObject => Extraction.extract[SnippetResultOption](value)
        case value: JObject => SnippetResult.apply(Extraction.extract[Map[String, SnippetResultOption]](value))
        case _              => throw new MappingException("Can't convert " + json + " to SnippetResult")
      }
  }

  override def serialize(implicit format: Formats): PartialFunction[Any, JValue] = { case value: SnippetResult =>
    value match {
      case value: SnippetResultOption => Extraction.decompose(value)(format - this)
    }
  }
}