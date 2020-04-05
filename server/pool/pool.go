package pool

import "net"

// ClientsPool : Pool of client connections
type ClientsPool interface {
	GetClients() []*net.Conn
	Write(data interface{}) error
}

//TODO: Gotta learn about synhc.Mutex for locks

// Pool : Pool datatructure
type Pool struct {
	conns   chan net.Conn
	minConn int
	maxConn int
}

// CreatePool : Creates a new Pool of clients.
// 							Use the method `pool.GetClients()` to get all of the client
// 							connections from a pool.
func CreatePool(minConn int, maxConn int) (Pool, error) {
	pool := Pool{
		conns:   make(chan net.Conn, maxConn),
		minConn: minConn,
		maxConn: maxConn,
	}

	err := pool.init()
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (p *Pool) init() error {
	p.cre
	return nil
}


func 