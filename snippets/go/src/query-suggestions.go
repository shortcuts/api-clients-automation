package snippets

import (
	suggestions "github.com/algolia/algoliasearch-client-go/v4/algolia/query-suggestions"
)

func SnippetForCreateConfigOfSuggestions() {
	/*
	   Snippet for the createConfig method.

	   createConfig0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.CreateConfig(client.NewApiCreateConfigRequest(

		suggestions.NewEmptyQuerySuggestionsConfigurationWithIndex().SetIndexName("theIndexName").SetSourceIndices(
			[]suggestions.SourceIndex{*suggestions.NewEmptySourceIndex().SetIndexName("testIndex").SetFacets(
				[]suggestions.Facet{*suggestions.NewEmptyFacet().SetAttribute("test")}).SetGenerate(
				[][]string{
					[]string{"facetA", "facetB"},
					[]string{"facetC"}})}).SetLanguages(suggestions.ArrayOfStringAsLanguages(
			[]string{"french"})).SetExclude(
			[]string{"test"}),
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForCustomDeleteOfSuggestions() {
	/*
	   Snippet for the customDelete method.

	   allow del method for a custom path with minimal parameters
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.CustomDelete(client.NewApiCustomDeleteRequest(
		"/test/minimal",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForCustomGetOfSuggestions() {
	/*
	   Snippet for the customGet method.

	   allow get method for a custom path with minimal parameters
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.CustomGet(client.NewApiCustomGetRequest(
		"/test/minimal",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForCustomPostOfSuggestions() {
	/*
	   Snippet for the customPost method.

	   allow post method for a custom path with minimal parameters
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.CustomPost(client.NewApiCustomPostRequest(
		"/test/minimal",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForCustomPutOfSuggestions() {
	/*
	   Snippet for the customPut method.

	   allow put method for a custom path with minimal parameters
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.CustomPut(client.NewApiCustomPutRequest(
		"/test/minimal",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForDeleteConfigOfSuggestions() {
	/*
	   Snippet for the deleteConfig method.

	   deleteConfig0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.DeleteConfig(client.NewApiDeleteConfigRequest(
		"theIndexName",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForGetAllConfigsOfSuggestions() {
	/*
	   Snippet for the getAllConfigs method.

	   getAllConfigs0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.GetAllConfigs()
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForGetConfigOfSuggestions() {
	/*
	   Snippet for the getConfig method.

	   getConfig0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.GetConfig(client.NewApiGetConfigRequest(
		"theIndexName",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForGetConfigStatusOfSuggestions() {
	/*
	   Snippet for the getConfigStatus method.

	   getConfigStatus0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.GetConfigStatus(client.NewApiGetConfigStatusRequest(
		"theIndexName",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForGetLogFileOfSuggestions() {
	/*
	   Snippet for the getLogFile method.

	   getLogFile0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.GetLogFile(client.NewApiGetLogFileRequest(
		"theIndexName",
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}
func SnippetForUpdateConfigOfSuggestions() {
	/*
	   Snippet for the updateConfig method.

	   updateConfig0
	*/

	// Initialize the client with your application region, eg. suggestions.YOUR_APP_ID_REGION
	client, err := suggestions.NewClient("YOUR_APP_ID", "YOUR_API_KEY", suggestions.US)
	if err != nil {
		// The client can fail to initialize if you pass an invalid parameter.
		panic(err)
	}

	// Call the API
	resp, err := client.UpdateConfig(client.NewApiUpdateConfigRequest(
		"theIndexName",
		suggestions.NewEmptyQuerySuggestionsConfiguration().SetSourceIndices(
			[]suggestions.SourceIndex{*suggestions.NewEmptySourceIndex().SetIndexName("testIndex").SetFacets(
				[]suggestions.Facet{*suggestions.NewEmptyFacet().SetAttribute("test")}).SetGenerate(
				[][]string{
					[]string{"facetA", "facetB"},
					[]string{"facetC"}})}).SetLanguages(suggestions.ArrayOfStringAsLanguages(
			[]string{"french"})).SetExclude(
			[]string{"test"}),
	))
	if err != nil {
		// handle the eventual error
		panic(err)
	}

	// use the model directly
	print(resp)
}