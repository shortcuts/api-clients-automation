// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import type { HideConsequenceObject } from './hideConsequenceObject';
import type { ParamsConsequence } from './paramsConsequence';
import type { PromoteConsequenceObject } from './promoteConsequenceObject';

/**
 * Effect of the rule.
 */
export type Consequence = {
  /**
   * Exclude items from recommendations.
   */
  hide?: HideConsequenceObject[];

  /**
   * Place items at specific positions in the list of recommendations.
   */
  promote?: PromoteConsequenceObject[];

  params?: ParamsConsequence;
};
