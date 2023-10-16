package cp

import "github.com/jinzhu/copier"

var (
	TyperConverter = []copier.TypeConverter{}
	FiledMapping   = []copier.FieldNameMapping{}
)

// Copy
// copy the value, if it has pointer, only copy the referenced address
func Copy(src, dst any) error {
	return copier.CopyWithOption(dst, src, copier.Option{
		DeepCopy:         false,
		Converters:       TyperConverter,
		FieldNameMapping: FiledMapping,
	})
}

// DeepCopy
// copy the reference of the value if it exists
func DeepCopy(src, dst any) error {
	return copier.CopyWithOption(dst, src, copier.Option{
		DeepCopy:         true,
		Converters:       TyperConverter,
		FieldNameMapping: FiledMapping,
	})
}
