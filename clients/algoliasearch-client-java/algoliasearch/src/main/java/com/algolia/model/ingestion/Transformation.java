// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.ingestion;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** Transformation */
public class Transformation {

  @JsonProperty("transformationID")
  private String transformationID;

  @JsonProperty("code")
  private String code;

  @JsonProperty("name")
  private String name;

  @JsonProperty("description")
  private String description;

  @JsonProperty("createdAt")
  private String createdAt;

  @JsonProperty("updatedAt")
  private String updatedAt;

  public Transformation setTransformationID(String transformationID) {
    this.transformationID = transformationID;
    return this;
  }

  /** Universally unique identifier (UUID) of a transformation. */
  @javax.annotation.Nonnull
  public String getTransformationID() {
    return transformationID;
  }

  public Transformation setCode(String code) {
    this.code = code;
    return this;
  }

  /** The source code of the transformation. */
  @javax.annotation.Nonnull
  public String getCode() {
    return code;
  }

  public Transformation setName(String name) {
    this.name = name;
    return this;
  }

  /** The uniquely identified name of your transformation. */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public Transformation setDescription(String description) {
    this.description = description;
    return this;
  }

  /** A descriptive name for your transformation of what it does. */
  @javax.annotation.Nonnull
  public String getDescription() {
    return description;
  }

  public Transformation setCreatedAt(String createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /** Date of creation in RFC 3339 format. */
  @javax.annotation.Nonnull
  public String getCreatedAt() {
    return createdAt;
  }

  public Transformation setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /** Date of last update in RFC 3339 format. */
  @javax.annotation.Nullable
  public String getUpdatedAt() {
    return updatedAt;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Transformation transformation = (Transformation) o;
    return (
      Objects.equals(this.transformationID, transformation.transformationID) &&
      Objects.equals(this.code, transformation.code) &&
      Objects.equals(this.name, transformation.name) &&
      Objects.equals(this.description, transformation.description) &&
      Objects.equals(this.createdAt, transformation.createdAt) &&
      Objects.equals(this.updatedAt, transformation.updatedAt)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(transformationID, code, name, description, createdAt, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Transformation {\n");
    sb.append("    transformationID: ").append(toIndentedString(transformationID)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
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