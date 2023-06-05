package scylla

import (
	"github.com/Jonathansoufer/go-reporting/config"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	cfg config.Config
}

func NewManager(cfg config.Config) *Manager {
	return &Manager{
		cfg: cfg,
	}
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	return m.connect(m.cfg.ScyllaKeyspace, m.cfg.ScyllaHosts)
}

func (m *Manager) CreateKeyspace(keyspace string) error {
	session, err := m.connect("system", m.cfg.ScyllaHosts)
	if err != nil {
		return err
	}
	defer session.Close()

	stmt := `CREATE KEYSPACE IF NOT EXISTS reporting WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`
	return session.ExecStmt(stmt)
}

func (m *Manager) connect(keyspace string, hosts []string) (gocqlx.Session, error) {
	c := gocql.NewCluster(hosts...)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}