package apiserver

// создаём структуру для сервера
type APIServer struct {
	config *Config // поле конфиг со ссылкой на тип конфиг с прописанными параметрами
}

// делаем конструктор-копию
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// прописываем методы для структуры
func (s *APIServer) Start() error {
	return nil
}
