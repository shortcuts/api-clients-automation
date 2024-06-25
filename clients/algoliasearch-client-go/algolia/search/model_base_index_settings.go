// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// BaseIndexSettings struct for BaseIndexSettings.
type BaseIndexSettings struct {
	// Attributes used for [faceting](https://www.algolia.com/doc/guides/managing-results/refine-results/faceting/).  Facets are attributes that let you categorize search results. They can be used for filtering search results. By default, no attribute is used for faceting. Attribute names are case-sensitive.  **Modifiers**  - `filterOnly(\"ATTRIBUTE\")`.   Allows using this attribute as a filter, but doesn't evalue the facet values.  - `searchable(\"ATTRIBUTE\")`.   Allows searching for facet values.  - `afterDistinct(\"ATTRIBUTE\")`.   Evaluates the facet count _after_ deduplication with `distinct`.   This ensures accurate facet counts.   You can apply this modifier to searchable facets: `afterDistinct(searchable(ATTRIBUTE))`.
	AttributesForFaceting []string `json:"attributesForFaceting,omitempty"`
	// Creates [replica indices](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/replicas/).  Replicas are copies of a primary index with the same records but different settings, synonyms, or rules. If you want to offer a different ranking or sorting of your search results, you'll use replica indices. All index operations on a primary index are automatically forwarded to its replicas. To add a replica index, you must provide the complete set of replicas to this parameter. If you omit a replica from this list, the replica turns into a regular, standalone index that will no longer by synced with the primary index.  **Modifier**  - `virtual(\"REPLICA\")`.   Create a virtual replica,   Virtual replicas don't increase the number of records and are optimized for [Relevant sorting](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/in-depth/relevant-sort/).
	Replicas []string `json:"replicas,omitempty"`
	// Only present if the index is a [virtual replica](https://www.algolia.com/doc/guides/managing-results/refine-results/sorting/how-to/sort-an-index-alphabetically/#virtual-replicas).
	Virtual *bool `json:"virtual,omitempty"`
	// Maximum number of search results that can be obtained through pagination.  Higher pagination limits might slow down your search. For pagination limits above 1,000, the sorting of results beyond the 1,000th hit can't be guaranteed.
	PaginationLimitedTo *int32 `json:"paginationLimitedTo,omitempty"`
	// Attributes that can't be retrieved at query time.  This can be useful if you want to use an attribute for ranking or to [restrict access](https://www.algolia.com/doc/guides/security/api-keys/how-to/user-restricted-access-to-data/), but don't want to include it in the search results. Attribute names are case-sensitive.
	UnretrievableAttributes []string `json:"unretrievableAttributes,omitempty"`
	// Words for which you want to turn off [typo tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/). This also turns off [word splitting and concatenation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/splitting-and-concatenation/) for the specified words.
	DisableTypoToleranceOnWords []string `json:"disableTypoToleranceOnWords,omitempty"`
	// Attributes, for which you want to support [Japanese transliteration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/#japanese-transliteration-and-type-ahead).  Transliteration supports searching in any of the Japanese writing systems. To support transliteration, you must set the indexing language to Japanese. Attribute names are case-sensitive.
	AttributesToTransliterate []string `json:"attributesToTransliterate,omitempty"`
	// Attributes for which to split [camel case](https://wikipedia.org/wiki/Camel_case) words. Attribute names are case-sensitive.
	CamelCaseAttributes []string `json:"camelCaseAttributes,omitempty"`
	// Searchable attributes to which Algolia should apply [word segmentation](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/how-to/customize-segmentation/) (decompounding). Attribute names are case-sensitive.  Compound words are formed by combining two or more individual words, and are particularly prevalent in Germanic languages—for example, \"firefighter\". With decompounding, the individual components are indexed separately.  You can specify different lists for different languages. Decompounding is supported for these languages: Dutch (`nl`), German (`de`), Finnish (`fi`), Danish (`da`), Swedish (`sv`), and Norwegian (`no`).
	DecompoundedAttributes map[string]any `json:"decompoundedAttributes,omitempty"`
	// Languages for language-specific processing steps, such as word detection and dictionary settings.  **You should always specify an indexing language.** If you don't specify an indexing language, the search engine uses all [supported languages](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/supported-languages/), or the languages you specified with the `ignorePlurals` or `removeStopWords` parameters. This can lead to unexpected search results. For more information, see [Language-specific configuration](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/language-specific-configurations/).
	IndexLanguages []SupportedLanguage `json:"indexLanguages,omitempty"`
	// Searchable attributes for which you want to turn off [prefix matching](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/override-search-engine-defaults/#adjusting-prefix-search). Attribute names are case-sensitive.
	DisablePrefixOnAttributes []string `json:"disablePrefixOnAttributes,omitempty"`
	// Whether arrays with exclusively non-negative integers should be compressed for better performance. If true, the compressed arrays may be reordered.
	AllowCompressionOfIntegerArray *bool `json:"allowCompressionOfIntegerArray,omitempty"`
	// Numeric attributes that can be used as [numerical filters](https://www.algolia.com/doc/guides/managing-results/rules/detecting-intent/how-to/applying-a-custom-filter-for-a-specific-query/#numerical-filters). Attribute names are case-sensitive.  By default, all numeric attributes are available as numerical filters. For faster indexing, reduce the number of numeric attributes.  If you want to turn off filtering for all numeric attributes, specify an attribute that doesn't exist in your index, such as `NO_NUMERIC_FILTERING`.  **Modifier**  - `equalOnly(\"ATTRIBUTE\")`.   Support only filtering based on equality comparisons `=` and `!=`.
	NumericAttributesForFiltering []string `json:"numericAttributesForFiltering,omitempty"`
	// Controls which separators are indexed.  Separators are all non-letter characters except spaces and currency characters, such as $€£¥. By default, separator characters aren't indexed. With `separatorsToIndex`, Algolia treats separator characters as separate words. For example, a search for `C#` would report two matches.
	SeparatorsToIndex *string `json:"separatorsToIndex,omitempty"`
	// Attributes used for searching. Attribute names are case-sensitive.  By default, all attributes are searchable and the [Attribute](https://www.algolia.com/doc/guides/managing-results/relevance-overview/in-depth/ranking-criteria/#attribute) ranking criterion is turned off. With a non-empty list, Algolia only returns results with matches in the selected attributes. In addition, the Attribute ranking criterion is turned on: matches in attributes that are higher in the list of `searchableAttributes` rank first. To make matches in two attributes rank equally, include them in a comma-separated string, such as `\"title,alternate_title\"`. Attributes with the same priority are always unordered.  For more information, see [Searchable attributes](https://www.algolia.com/doc/guides/sending-and-managing-data/prepare-your-data/how-to/setting-searchable-attributes/).  **Modifier**  - `unordered(\"ATTRIBUTE\")`.   Ignore the position of a match within the attribute.  Without modifier, matches at the beginning of an attribute rank higher than matches at the end.
	SearchableAttributes []string `json:"searchableAttributes,omitempty"`
	// An object with custom data.  You can store up to 32kB as custom data.
	UserData map[string]any `json:"userData,omitempty"`
	// Characters and their normalized replacements. This overrides Algolia's default [normalization](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/handling-natural-languages-nlp/in-depth/normalization/).
	CustomNormalization *map[string]map[string]string `json:"customNormalization,omitempty"`
	// Attribute that should be used to establish groups of results. Attribute names are case-sensitive.  All records with the same value for this attribute are considered a group. You can combine `attributeForDistinct` with the `distinct` search parameter to control how many items per group are included in the search results.  If you want to use the same attribute also for faceting, use the `afterDistinct` modifier of the `attributesForFaceting` setting. This applies faceting _after_ deduplication, which will result in accurate facet counts.
	AttributeForDistinct *string `json:"attributeForDistinct,omitempty"`
}

type BaseIndexSettingsOption func(f *BaseIndexSettings)

func WithBaseIndexSettingsAttributesForFaceting(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AttributesForFaceting = val
	}
}

func WithBaseIndexSettingsReplicas(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.Replicas = val
	}
}

func WithBaseIndexSettingsVirtual(val bool) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.Virtual = &val
	}
}

func WithBaseIndexSettingsPaginationLimitedTo(val int32) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.PaginationLimitedTo = &val
	}
}

func WithBaseIndexSettingsUnretrievableAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.UnretrievableAttributes = val
	}
}

func WithBaseIndexSettingsDisableTypoToleranceOnWords(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DisableTypoToleranceOnWords = val
	}
}

func WithBaseIndexSettingsAttributesToTransliterate(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AttributesToTransliterate = val
	}
}

func WithBaseIndexSettingsCamelCaseAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.CamelCaseAttributes = val
	}
}

func WithBaseIndexSettingsDecompoundedAttributes(val map[string]any) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DecompoundedAttributes = val
	}
}

func WithBaseIndexSettingsIndexLanguages(val []SupportedLanguage) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.IndexLanguages = val
	}
}

func WithBaseIndexSettingsDisablePrefixOnAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.DisablePrefixOnAttributes = val
	}
}

func WithBaseIndexSettingsAllowCompressionOfIntegerArray(val bool) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AllowCompressionOfIntegerArray = &val
	}
}

func WithBaseIndexSettingsNumericAttributesForFiltering(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.NumericAttributesForFiltering = val
	}
}

func WithBaseIndexSettingsSeparatorsToIndex(val string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.SeparatorsToIndex = &val
	}
}

func WithBaseIndexSettingsSearchableAttributes(val []string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.SearchableAttributes = val
	}
}

func WithBaseIndexSettingsUserData(val map[string]any) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.UserData = val
	}
}

func WithBaseIndexSettingsCustomNormalization(val map[string]map[string]string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.CustomNormalization = &val
	}
}

func WithBaseIndexSettingsAttributeForDistinct(val string) BaseIndexSettingsOption {
	return func(f *BaseIndexSettings) {
		f.AttributeForDistinct = &val
	}
}

// NewBaseIndexSettings instantiates a new BaseIndexSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBaseIndexSettings(opts ...BaseIndexSettingsOption) *BaseIndexSettings {
	this := &BaseIndexSettings{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewEmptyBaseIndexSettings return a pointer to an empty BaseIndexSettings object.
func NewEmptyBaseIndexSettings() *BaseIndexSettings {
	return &BaseIndexSettings{}
}

// GetAttributesForFaceting returns the AttributesForFaceting field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAttributesForFaceting() []string {
	if o == nil || o.AttributesForFaceting == nil {
		var ret []string
		return ret
	}
	return o.AttributesForFaceting
}

// GetAttributesForFacetingOk returns a tuple with the AttributesForFaceting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAttributesForFacetingOk() ([]string, bool) {
	if o == nil || o.AttributesForFaceting == nil {
		return nil, false
	}
	return o.AttributesForFaceting, true
}

// HasAttributesForFaceting returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAttributesForFaceting() bool {
	if o != nil && o.AttributesForFaceting != nil {
		return true
	}

	return false
}

// SetAttributesForFaceting gets a reference to the given []string and assigns it to the AttributesForFaceting field.
func (o *BaseIndexSettings) SetAttributesForFaceting(v []string) *BaseIndexSettings {
	o.AttributesForFaceting = v
	return o
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetReplicas() []string {
	if o == nil || o.Replicas == nil {
		var ret []string
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetReplicasOk() ([]string, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given []string and assigns it to the Replicas field.
func (o *BaseIndexSettings) SetReplicas(v []string) *BaseIndexSettings {
	o.Replicas = v
	return o
}

// GetVirtual returns the Virtual field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetVirtual() bool {
	if o == nil || o.Virtual == nil {
		var ret bool
		return ret
	}
	return *o.Virtual
}

// GetVirtualOk returns a tuple with the Virtual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetVirtualOk() (*bool, bool) {
	if o == nil || o.Virtual == nil {
		return nil, false
	}
	return o.Virtual, true
}

// HasVirtual returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasVirtual() bool {
	if o != nil && o.Virtual != nil {
		return true
	}

	return false
}

// SetVirtual gets a reference to the given bool and assigns it to the Virtual field.
func (o *BaseIndexSettings) SetVirtual(v bool) *BaseIndexSettings {
	o.Virtual = &v
	return o
}

// GetPaginationLimitedTo returns the PaginationLimitedTo field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetPaginationLimitedTo() int32 {
	if o == nil || o.PaginationLimitedTo == nil {
		var ret int32
		return ret
	}
	return *o.PaginationLimitedTo
}

// GetPaginationLimitedToOk returns a tuple with the PaginationLimitedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetPaginationLimitedToOk() (*int32, bool) {
	if o == nil || o.PaginationLimitedTo == nil {
		return nil, false
	}
	return o.PaginationLimitedTo, true
}

// HasPaginationLimitedTo returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasPaginationLimitedTo() bool {
	if o != nil && o.PaginationLimitedTo != nil {
		return true
	}

	return false
}

// SetPaginationLimitedTo gets a reference to the given int32 and assigns it to the PaginationLimitedTo field.
func (o *BaseIndexSettings) SetPaginationLimitedTo(v int32) *BaseIndexSettings {
	o.PaginationLimitedTo = &v
	return o
}

// GetUnretrievableAttributes returns the UnretrievableAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetUnretrievableAttributes() []string {
	if o == nil || o.UnretrievableAttributes == nil {
		var ret []string
		return ret
	}
	return o.UnretrievableAttributes
}

// GetUnretrievableAttributesOk returns a tuple with the UnretrievableAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetUnretrievableAttributesOk() ([]string, bool) {
	if o == nil || o.UnretrievableAttributes == nil {
		return nil, false
	}
	return o.UnretrievableAttributes, true
}

// HasUnretrievableAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasUnretrievableAttributes() bool {
	if o != nil && o.UnretrievableAttributes != nil {
		return true
	}

	return false
}

// SetUnretrievableAttributes gets a reference to the given []string and assigns it to the UnretrievableAttributes field.
func (o *BaseIndexSettings) SetUnretrievableAttributes(v []string) *BaseIndexSettings {
	o.UnretrievableAttributes = v
	return o
}

// GetDisableTypoToleranceOnWords returns the DisableTypoToleranceOnWords field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDisableTypoToleranceOnWords() []string {
	if o == nil || o.DisableTypoToleranceOnWords == nil {
		var ret []string
		return ret
	}
	return o.DisableTypoToleranceOnWords
}

// GetDisableTypoToleranceOnWordsOk returns a tuple with the DisableTypoToleranceOnWords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDisableTypoToleranceOnWordsOk() ([]string, bool) {
	if o == nil || o.DisableTypoToleranceOnWords == nil {
		return nil, false
	}
	return o.DisableTypoToleranceOnWords, true
}

// HasDisableTypoToleranceOnWords returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDisableTypoToleranceOnWords() bool {
	if o != nil && o.DisableTypoToleranceOnWords != nil {
		return true
	}

	return false
}

// SetDisableTypoToleranceOnWords gets a reference to the given []string and assigns it to the DisableTypoToleranceOnWords field.
func (o *BaseIndexSettings) SetDisableTypoToleranceOnWords(v []string) *BaseIndexSettings {
	o.DisableTypoToleranceOnWords = v
	return o
}

// GetAttributesToTransliterate returns the AttributesToTransliterate field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAttributesToTransliterate() []string {
	if o == nil || o.AttributesToTransliterate == nil {
		var ret []string
		return ret
	}
	return o.AttributesToTransliterate
}

// GetAttributesToTransliterateOk returns a tuple with the AttributesToTransliterate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAttributesToTransliterateOk() ([]string, bool) {
	if o == nil || o.AttributesToTransliterate == nil {
		return nil, false
	}
	return o.AttributesToTransliterate, true
}

// HasAttributesToTransliterate returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAttributesToTransliterate() bool {
	if o != nil && o.AttributesToTransliterate != nil {
		return true
	}

	return false
}

// SetAttributesToTransliterate gets a reference to the given []string and assigns it to the AttributesToTransliterate field.
func (o *BaseIndexSettings) SetAttributesToTransliterate(v []string) *BaseIndexSettings {
	o.AttributesToTransliterate = v
	return o
}

// GetCamelCaseAttributes returns the CamelCaseAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetCamelCaseAttributes() []string {
	if o == nil || o.CamelCaseAttributes == nil {
		var ret []string
		return ret
	}
	return o.CamelCaseAttributes
}

// GetCamelCaseAttributesOk returns a tuple with the CamelCaseAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetCamelCaseAttributesOk() ([]string, bool) {
	if o == nil || o.CamelCaseAttributes == nil {
		return nil, false
	}
	return o.CamelCaseAttributes, true
}

// HasCamelCaseAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasCamelCaseAttributes() bool {
	if o != nil && o.CamelCaseAttributes != nil {
		return true
	}

	return false
}

// SetCamelCaseAttributes gets a reference to the given []string and assigns it to the CamelCaseAttributes field.
func (o *BaseIndexSettings) SetCamelCaseAttributes(v []string) *BaseIndexSettings {
	o.CamelCaseAttributes = v
	return o
}

// GetDecompoundedAttributes returns the DecompoundedAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDecompoundedAttributes() map[string]any {
	if o == nil || o.DecompoundedAttributes == nil {
		var ret map[string]any
		return ret
	}
	return o.DecompoundedAttributes
}

// GetDecompoundedAttributesOk returns a tuple with the DecompoundedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDecompoundedAttributesOk() (map[string]any, bool) {
	if o == nil || o.DecompoundedAttributes == nil {
		return nil, false
	}
	return o.DecompoundedAttributes, true
}

// HasDecompoundedAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDecompoundedAttributes() bool {
	if o != nil && o.DecompoundedAttributes != nil {
		return true
	}

	return false
}

// SetDecompoundedAttributes gets a reference to the given map[string]any and assigns it to the DecompoundedAttributes field.
func (o *BaseIndexSettings) SetDecompoundedAttributes(v map[string]any) *BaseIndexSettings {
	o.DecompoundedAttributes = v
	return o
}

// GetIndexLanguages returns the IndexLanguages field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetIndexLanguages() []SupportedLanguage {
	if o == nil || o.IndexLanguages == nil {
		var ret []SupportedLanguage
		return ret
	}
	return o.IndexLanguages
}

// GetIndexLanguagesOk returns a tuple with the IndexLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetIndexLanguagesOk() ([]SupportedLanguage, bool) {
	if o == nil || o.IndexLanguages == nil {
		return nil, false
	}
	return o.IndexLanguages, true
}

// HasIndexLanguages returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasIndexLanguages() bool {
	if o != nil && o.IndexLanguages != nil {
		return true
	}

	return false
}

// SetIndexLanguages gets a reference to the given []SupportedLanguage and assigns it to the IndexLanguages field.
func (o *BaseIndexSettings) SetIndexLanguages(v []SupportedLanguage) *BaseIndexSettings {
	o.IndexLanguages = v
	return o
}

// GetDisablePrefixOnAttributes returns the DisablePrefixOnAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetDisablePrefixOnAttributes() []string {
	if o == nil || o.DisablePrefixOnAttributes == nil {
		var ret []string
		return ret
	}
	return o.DisablePrefixOnAttributes
}

// GetDisablePrefixOnAttributesOk returns a tuple with the DisablePrefixOnAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetDisablePrefixOnAttributesOk() ([]string, bool) {
	if o == nil || o.DisablePrefixOnAttributes == nil {
		return nil, false
	}
	return o.DisablePrefixOnAttributes, true
}

// HasDisablePrefixOnAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasDisablePrefixOnAttributes() bool {
	if o != nil && o.DisablePrefixOnAttributes != nil {
		return true
	}

	return false
}

// SetDisablePrefixOnAttributes gets a reference to the given []string and assigns it to the DisablePrefixOnAttributes field.
func (o *BaseIndexSettings) SetDisablePrefixOnAttributes(v []string) *BaseIndexSettings {
	o.DisablePrefixOnAttributes = v
	return o
}

// GetAllowCompressionOfIntegerArray returns the AllowCompressionOfIntegerArray field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAllowCompressionOfIntegerArray() bool {
	if o == nil || o.AllowCompressionOfIntegerArray == nil {
		var ret bool
		return ret
	}
	return *o.AllowCompressionOfIntegerArray
}

// GetAllowCompressionOfIntegerArrayOk returns a tuple with the AllowCompressionOfIntegerArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAllowCompressionOfIntegerArrayOk() (*bool, bool) {
	if o == nil || o.AllowCompressionOfIntegerArray == nil {
		return nil, false
	}
	return o.AllowCompressionOfIntegerArray, true
}

// HasAllowCompressionOfIntegerArray returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAllowCompressionOfIntegerArray() bool {
	if o != nil && o.AllowCompressionOfIntegerArray != nil {
		return true
	}

	return false
}

// SetAllowCompressionOfIntegerArray gets a reference to the given bool and assigns it to the AllowCompressionOfIntegerArray field.
func (o *BaseIndexSettings) SetAllowCompressionOfIntegerArray(v bool) *BaseIndexSettings {
	o.AllowCompressionOfIntegerArray = &v
	return o
}

// GetNumericAttributesForFiltering returns the NumericAttributesForFiltering field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetNumericAttributesForFiltering() []string {
	if o == nil || o.NumericAttributesForFiltering == nil {
		var ret []string
		return ret
	}
	return o.NumericAttributesForFiltering
}

// GetNumericAttributesForFilteringOk returns a tuple with the NumericAttributesForFiltering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetNumericAttributesForFilteringOk() ([]string, bool) {
	if o == nil || o.NumericAttributesForFiltering == nil {
		return nil, false
	}
	return o.NumericAttributesForFiltering, true
}

// HasNumericAttributesForFiltering returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasNumericAttributesForFiltering() bool {
	if o != nil && o.NumericAttributesForFiltering != nil {
		return true
	}

	return false
}

// SetNumericAttributesForFiltering gets a reference to the given []string and assigns it to the NumericAttributesForFiltering field.
func (o *BaseIndexSettings) SetNumericAttributesForFiltering(v []string) *BaseIndexSettings {
	o.NumericAttributesForFiltering = v
	return o
}

// GetSeparatorsToIndex returns the SeparatorsToIndex field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetSeparatorsToIndex() string {
	if o == nil || o.SeparatorsToIndex == nil {
		var ret string
		return ret
	}
	return *o.SeparatorsToIndex
}

// GetSeparatorsToIndexOk returns a tuple with the SeparatorsToIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetSeparatorsToIndexOk() (*string, bool) {
	if o == nil || o.SeparatorsToIndex == nil {
		return nil, false
	}
	return o.SeparatorsToIndex, true
}

// HasSeparatorsToIndex returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasSeparatorsToIndex() bool {
	if o != nil && o.SeparatorsToIndex != nil {
		return true
	}

	return false
}

// SetSeparatorsToIndex gets a reference to the given string and assigns it to the SeparatorsToIndex field.
func (o *BaseIndexSettings) SetSeparatorsToIndex(v string) *BaseIndexSettings {
	o.SeparatorsToIndex = &v
	return o
}

// GetSearchableAttributes returns the SearchableAttributes field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetSearchableAttributes() []string {
	if o == nil || o.SearchableAttributes == nil {
		var ret []string
		return ret
	}
	return o.SearchableAttributes
}

// GetSearchableAttributesOk returns a tuple with the SearchableAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetSearchableAttributesOk() ([]string, bool) {
	if o == nil || o.SearchableAttributes == nil {
		return nil, false
	}
	return o.SearchableAttributes, true
}

// HasSearchableAttributes returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasSearchableAttributes() bool {
	if o != nil && o.SearchableAttributes != nil {
		return true
	}

	return false
}

// SetSearchableAttributes gets a reference to the given []string and assigns it to the SearchableAttributes field.
func (o *BaseIndexSettings) SetSearchableAttributes(v []string) *BaseIndexSettings {
	o.SearchableAttributes = v
	return o
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetUserData() map[string]any {
	if o == nil || o.UserData == nil {
		var ret map[string]any
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetUserDataOk() (map[string]any, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]any and assigns it to the UserData field.
func (o *BaseIndexSettings) SetUserData(v map[string]any) *BaseIndexSettings {
	o.UserData = v
	return o
}

// GetCustomNormalization returns the CustomNormalization field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetCustomNormalization() map[string]map[string]string {
	if o == nil || o.CustomNormalization == nil {
		var ret map[string]map[string]string
		return ret
	}
	return *o.CustomNormalization
}

// GetCustomNormalizationOk returns a tuple with the CustomNormalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetCustomNormalizationOk() (*map[string]map[string]string, bool) {
	if o == nil || o.CustomNormalization == nil {
		return nil, false
	}
	return o.CustomNormalization, true
}

// HasCustomNormalization returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasCustomNormalization() bool {
	if o != nil && o.CustomNormalization != nil {
		return true
	}

	return false
}

// SetCustomNormalization gets a reference to the given map[string]map[string]string and assigns it to the CustomNormalization field.
func (o *BaseIndexSettings) SetCustomNormalization(v map[string]map[string]string) *BaseIndexSettings {
	o.CustomNormalization = &v
	return o
}

// GetAttributeForDistinct returns the AttributeForDistinct field value if set, zero value otherwise.
func (o *BaseIndexSettings) GetAttributeForDistinct() string {
	if o == nil || o.AttributeForDistinct == nil {
		var ret string
		return ret
	}
	return *o.AttributeForDistinct
}

// GetAttributeForDistinctOk returns a tuple with the AttributeForDistinct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseIndexSettings) GetAttributeForDistinctOk() (*string, bool) {
	if o == nil || o.AttributeForDistinct == nil {
		return nil, false
	}
	return o.AttributeForDistinct, true
}

// HasAttributeForDistinct returns a boolean if a field has been set.
func (o *BaseIndexSettings) HasAttributeForDistinct() bool {
	if o != nil && o.AttributeForDistinct != nil {
		return true
	}

	return false
}

// SetAttributeForDistinct gets a reference to the given string and assigns it to the AttributeForDistinct field.
func (o *BaseIndexSettings) SetAttributeForDistinct(v string) *BaseIndexSettings {
	o.AttributeForDistinct = &v
	return o
}

func (o BaseIndexSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AttributesForFaceting != nil {
		toSerialize["attributesForFaceting"] = o.AttributesForFaceting
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Virtual != nil {
		toSerialize["virtual"] = o.Virtual
	}
	if o.PaginationLimitedTo != nil {
		toSerialize["paginationLimitedTo"] = o.PaginationLimitedTo
	}
	if o.UnretrievableAttributes != nil {
		toSerialize["unretrievableAttributes"] = o.UnretrievableAttributes
	}
	if o.DisableTypoToleranceOnWords != nil {
		toSerialize["disableTypoToleranceOnWords"] = o.DisableTypoToleranceOnWords
	}
	if o.AttributesToTransliterate != nil {
		toSerialize["attributesToTransliterate"] = o.AttributesToTransliterate
	}
	if o.CamelCaseAttributes != nil {
		toSerialize["camelCaseAttributes"] = o.CamelCaseAttributes
	}
	if o.DecompoundedAttributes != nil {
		toSerialize["decompoundedAttributes"] = o.DecompoundedAttributes
	}
	if o.IndexLanguages != nil {
		toSerialize["indexLanguages"] = o.IndexLanguages
	}
	if o.DisablePrefixOnAttributes != nil {
		toSerialize["disablePrefixOnAttributes"] = o.DisablePrefixOnAttributes
	}
	if o.AllowCompressionOfIntegerArray != nil {
		toSerialize["allowCompressionOfIntegerArray"] = o.AllowCompressionOfIntegerArray
	}
	if o.NumericAttributesForFiltering != nil {
		toSerialize["numericAttributesForFiltering"] = o.NumericAttributesForFiltering
	}
	if o.SeparatorsToIndex != nil {
		toSerialize["separatorsToIndex"] = o.SeparatorsToIndex
	}
	if o.SearchableAttributes != nil {
		toSerialize["searchableAttributes"] = o.SearchableAttributes
	}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	if o.CustomNormalization != nil {
		toSerialize["customNormalization"] = o.CustomNormalization
	}
	if o.AttributeForDistinct != nil {
		toSerialize["attributeForDistinct"] = o.AttributeForDistinct
	}
	serialized, err := json.Marshal(toSerialize)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal BaseIndexSettings: %w", err)
	}

	return serialized, nil
}

func (o BaseIndexSettings) String() string {
	out := ""
	out += fmt.Sprintf("  attributesForFaceting=%v\n", o.AttributesForFaceting)
	out += fmt.Sprintf("  replicas=%v\n", o.Replicas)
	out += fmt.Sprintf("  virtual=%v\n", o.Virtual)
	out += fmt.Sprintf("  paginationLimitedTo=%v\n", o.PaginationLimitedTo)
	out += fmt.Sprintf("  unretrievableAttributes=%v\n", o.UnretrievableAttributes)
	out += fmt.Sprintf("  disableTypoToleranceOnWords=%v\n", o.DisableTypoToleranceOnWords)
	out += fmt.Sprintf("  attributesToTransliterate=%v\n", o.AttributesToTransliterate)
	out += fmt.Sprintf("  camelCaseAttributes=%v\n", o.CamelCaseAttributes)
	out += fmt.Sprintf("  decompoundedAttributes=%v\n", o.DecompoundedAttributes)
	out += fmt.Sprintf("  indexLanguages=%v\n", o.IndexLanguages)
	out += fmt.Sprintf("  disablePrefixOnAttributes=%v\n", o.DisablePrefixOnAttributes)
	out += fmt.Sprintf("  allowCompressionOfIntegerArray=%v\n", o.AllowCompressionOfIntegerArray)
	out += fmt.Sprintf("  numericAttributesForFiltering=%v\n", o.NumericAttributesForFiltering)
	out += fmt.Sprintf("  separatorsToIndex=%v\n", o.SeparatorsToIndex)
	out += fmt.Sprintf("  searchableAttributes=%v\n", o.SearchableAttributes)
	out += fmt.Sprintf("  userData=%v\n", o.UserData)
	out += fmt.Sprintf("  customNormalization=%v\n", o.CustomNormalization)
	out += fmt.Sprintf("  attributeForDistinct=%v\n", o.AttributeForDistinct)
	return fmt.Sprintf("BaseIndexSettings {\n%s}", out)
}
