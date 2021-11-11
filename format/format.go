package format

func ReturnFormattedData(searchResults string) (string, error) {

	//the input and output are string as a holding space
	// Ideas
	// Human readable format
	// would be a nice display of kv pairs
	// Look up golang formatting options for display to terminal
	// Display should include all relevant data and context
	// If nil values are not the search query do we display them or exclude them?
	// If asked for nil values we should display associated records

	// If greater than the expected display (20 lines?) - paginate results for scrolling through
	// Consider dumping into an output file
	// Make it pretty in colour and demarcation of sections. How do i even do this!
	// A challenge.
	// This section needs far more fleshing out. TBD.

	formattedResults := "the end result of all of that"
	return formattedResults, nil
}

func pagination() {

	//paginate if over a certain size
}
