package event

// BlockPlaced is called by a client when it places a block.
type BlockPlaced struct {
	// AuxType is the metadata value of the block placed.
	AuxType int
	// Block is the block that was placed. This is a string like 'stone'.
	Block string
	// Namespace is the namespace to which the block belongs. Usually, this is 'minecraft'.
	Namespace string
	// Type is the numerical ID of the block placed. For stone, this is 1.
	Type int
	// PlacementMethod was the method that was used to place the block.
	PlacementMethod int
}