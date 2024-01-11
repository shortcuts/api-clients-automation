// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.ingestion;

import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.databind.annotation.*;
import java.util.Objects;

/** A destination describe how the data is indexed on the Algolia side. */
public class Destination {

  @JsonProperty("destinationID")
  private String destinationID;

  @JsonProperty("type")
  private DestinationType type;

  @JsonProperty("name")
  private String name;

  @JsonProperty("input")
  private DestinationInput input;

  @JsonProperty("createdAt")
  private String createdAt;

  @JsonProperty("updatedAt")
  private String updatedAt;

  @JsonProperty("authenticationID")
  private String authenticationID;

  public Destination setDestinationID(String destinationID) {
    this.destinationID = destinationID;
    return this;
  }

  /** The destination UUID. */
  @javax.annotation.Nonnull
  public String getDestinationID() {
    return destinationID;
  }

  public Destination setType(DestinationType type) {
    this.type = type;
    return this;
  }

  /** Get type */
  @javax.annotation.Nonnull
  public DestinationType getType() {
    return type;
  }

  public Destination setName(String name) {
    this.name = name;
    return this;
  }

  /** An human readable name describing the object. */
  @javax.annotation.Nonnull
  public String getName() {
    return name;
  }

  public Destination setInput(DestinationInput input) {
    this.input = input;
    return this;
  }

  /** Get input */
  @javax.annotation.Nonnull
  public DestinationInput getInput() {
    return input;
  }

  public Destination setCreatedAt(String createdAt) {
    this.createdAt = createdAt;
    return this;
  }

  /** Date of creation (RFC3339 format). */
  @javax.annotation.Nonnull
  public String getCreatedAt() {
    return createdAt;
  }

  public Destination setUpdatedAt(String updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

  /** Date of last update (RFC3339 format). */
  @javax.annotation.Nullable
  public String getUpdatedAt() {
    return updatedAt;
  }

  public Destination setAuthenticationID(String authenticationID) {
    this.authenticationID = authenticationID;
    return this;
  }

  /** Get authenticationID */
  @javax.annotation.Nullable
  public String getAuthenticationID() {
    return authenticationID;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Destination destination = (Destination) o;
    return (
      Objects.equals(this.destinationID, destination.destinationID) &&
      Objects.equals(this.type, destination.type) &&
      Objects.equals(this.name, destination.name) &&
      Objects.equals(this.input, destination.input) &&
      Objects.equals(this.createdAt, destination.createdAt) &&
      Objects.equals(this.updatedAt, destination.updatedAt) &&
      Objects.equals(this.authenticationID, destination.authenticationID)
    );
  }

  @Override
  public int hashCode() {
    return Objects.hash(destinationID, type, name, input, createdAt, updatedAt, authenticationID);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Destination {\n");
    sb.append("    destinationID: ").append(toIndentedString(destinationID)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    input: ").append(toIndentedString(input)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    authenticationID: ").append(toIndentedString(authenticationID)).append("\n");
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