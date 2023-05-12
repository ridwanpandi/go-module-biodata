package go_module_biodata

type Biodata interface {
	GetName() string
	GetAddress() string
	GetAge() string
}

func GetBiodata(bio Biodata) map[string]interface{} {
	return map[string]interface{}{
		"nama":   bio.GetName(),
		"alamat": bio.GetAddress(),
		"umur":   bio.GetAge(),
	}
}
