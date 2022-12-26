package cache

import "sync"

// estructura que puede tener la meorizacion
type Memory struct {
	f Function
	//basicamente el chanche es un mapa que va marcar de enteros a resultados de una funcion
	cache map[int]FunctionResult

	lock sync.Mutex
}

// reronar una interface y un error en caso que lo hay
type Function func(key int) (interface{}, error)

// practicamente contiene el tipo de retorno de typo Function de arriba
type FunctionResult struct {
	value interface{}
	err   error
}

// **************************************FUNCTIONS*****************************

// Constructor (reciver function que hace de contructor)
// basicamente recibe y asigna una funcion nueva y asi tambien tambien crea un nuevo mapa
// sino el map usaria su zero value que es nil
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// devuelve los resultador del cache
func (m *Memory) Get(key int) (interface{}, error) {
	//verifica si ya existe ese resultado previamente
	result, exists := m.cache[key]
	// si no existe lo calcula
	if !exists {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	// si exites o tuve que calcularlo igualmente  lo retorna
	return result.value, result.err
}
