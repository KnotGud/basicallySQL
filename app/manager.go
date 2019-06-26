package app

import (
	"sync"
)

type Manager struct {
	db *Database
}

// NewManager Constructor
func NewManager(db *Database) *Manager {
	return &Manager{db: db}
}

// GetResults returns a slice of results to render
func (m *Manager) GetResults(zip int64, codes ...string) []Result {
	var wg sync.WaitGroup
	filteredClaims := make(chan *Claim)
	if len(codes) == 0 {
		// Case there is no filter
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range m.db.Claims {
				filteredClaims <- &v
			}
		}()
	} else {
		for _, c := range codes {
			wg.Add(1)
			go m.filterCodes(c, filteredClaims, &wg)
		}
	}
	go func() {
		wg.Wait()
		close(filteredClaims)
	}()

	resultMap := m.filterProviders(zip)
	for c := range filteredClaims {
		r, ok := resultMap[c.ProviderID]
		if ok {
			r.Count++
		}
	}
	ret := make([]Result, 0)
	for _, r := range resultMap {
		ret = append(ret, *r)
	}
	return ret
}

// filterProviders runs through the server's providers to find which are in the specified zip
func (m *Manager) filterProviders(zip int64) map[int64]*Result {
	fm := make(map[int64]*Result)
	for _, p := range m.db.Providers {
		if p.ZipCode == zip {
			r := p.Result()
			fm[p.ID] = &r
		}
	}
	return fm
}

// filterCodes reads through the server's claims to find ones with matching diagnosis codes
func (m *Manager) filterCodes(code string, out chan<- *Claim, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, c := range m.db.Claims {
		for _, v := range c.DiagnosisCodes {
			if v == code {
				out <- &c
				break
			}
		}
	}
}
