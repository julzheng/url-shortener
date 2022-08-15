package store

type StoreBackEnd struct {
	Memory *map[string]string
}

func (sbe *StoreBackEnd) IsExist(key string) bool {
	if _, ok := (*sbe.Memory)[key]; ok {
		return true
	} else {
		return false
	}
}

func (sbe *StoreBackEnd) Add(key string, val string) {
	(*sbe.Memory)[key] = val
}

func (sbe *StoreBackEnd) Get(key string) string {
	return (*sbe.Memory)[key]
}
