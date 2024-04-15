// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.recommend;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** Object ID of the recommendation you want to exclude. */
public class HideConsequenceObject {

  @JsonProperty("objectID")
  private String objectID;

  public HideConsequenceObject setObjectID(String objectID) {
    this.objectID = objectID;
    return this;
  }

  /** Unique record identifier. */
  @javax.annotation.Nullable
  public String getObjectID() {
    return objectID;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HideConsequenceObject hideConsequenceObject = (HideConsequenceObject) o;
    return Objects.equals(this.objectID, hideConsequenceObject.objectID);
  }

  @Override
  public int hashCode() {
    return Objects.hash(objectID);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HideConsequenceObject {\n");
    sb.append("    objectID: ").append(toIndentedString(objectID)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }
}