package store

import . "url-shortener/internal/types"

type StoreBackEnd struct {
	Memory *map[string]URLInfo
}

func (sbe *StoreBackEnd) IsExist(key string) bool {
	if _, ok := (*sbe.Memory)[key]; ok {
		return true
	} else {
		return false
	}
}

func (sbe *StoreBackEnd) Add(key string, urlInfo URLInfo) {
	(*sbe.Memory)[key] = urlInfo
}

func (sbe *StoreBackEnd) Get(key string) URLInfo {
	return (*sbe.Memory)[key]
}
