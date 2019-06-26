package app

import "fmt"

type Claim struct {
	ID             int64    `json:"id"`
	ProviderID     int64    `json:"provider_id"`
	PatientHash    []byte   `json:"patient_hash"`
	DiagnosisCodes []string `json:"diagnosis_codes"`
}

type Provider struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Facility string  `json:"facility_name"`
	ZipCode  int64   `json:"zip_code"`
	Lat      float32 `json:"lat"`
	Lon      float32 `json:"lon"`
}

// Result converts a Provider into a result
func (p *Provider) Result() Result {
	return Result{
		ID:       p.ID,
		Name:     p.Name,
		Facility: p.Facility,
		Lat:      p.Lat,
		Lon:      p.Lon,
		Count:    0,
	}
}

type Result struct {
	ID       int64
	Name     string
	Facility string
	Lat      float32
	Lon      float32
	Count    int64
}

func (r *Result) String() string {
	return fmt.Sprintf("<%v, %v, %v", r.ID, r.Name, r.Count)
}
