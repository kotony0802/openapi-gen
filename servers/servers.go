package servers

type Servers []Server

type Server struct {
	Url         string
	Description string
	Variables   map[string]string // ここあとで考える
}
