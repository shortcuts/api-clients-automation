// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { EventsItems } from './eventsItems';

export type InsightsEvents = {
  /**
   * List of click and conversion events.  An event is an object representing a user interaction. Events have attributes that describe the interaction, such as an event name, a type, or a user token. Some attributes require other attributes to be declared, and some attributes can\'t be declared at the same time.  **All** events must be valid, otherwise the API returns an error.
   */
  events: EventsItems[];
};