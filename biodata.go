package go_module_biodata

type Biodata struct {
	Nama   string
	Alamat string
	Umur   int
}

func (bio Biodata) GetBiodata() map[string]interface{} {
	return map[string]interface{}{
		"nama":   bio.Nama,
		"alamat": bio.Alamat,
		"umur":   bio.Umur,
	}
}
