module main

go 1.15

replace BasicHashMap => ../HashMaps/BasicHashMap

replace CuckooHashMap => ../HashMaps/CuckooHashMap

replace BinaryTreeHashMap => ../HashMaps/BinaryTreeHashMap

replace Pair => ../HashMaps/Pair

replace GoLangMap => ../HashMaps/GoLangMap

require (
	BasicHashMap v0.0.0-00010101000000-000000000000
	BinaryTreeHashMap v0.0.0-00010101000000-000000000000
	CuckooHashMap v0.0.0-00010101000000-000000000000
	GoLangMap v0.0.0-00010101000000-000000000000
	Pair v0.0.0-00010101000000-000000000000
)
