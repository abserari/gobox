// +build jsoniterator

package json

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Marshal exports jsoniter.Marshal.
	Marshal = json.Marshal

	// Unmarshal exports jsoniter.Unmarshal.
	Unmarshal = json.Unmarshal

	// NewEncoder exports jsoniter.NewEncoder.
	NewEncoder = json.NewEncoder

	// NewDecoder exports jsoniter.NewDecoder.
	NewDecoder = json.NewDecoder
)
