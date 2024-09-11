package fetcher

import (
	"net"

	"github.com/hirochachacha/go-smb2"
)

type Fetcher interface {
	Fetch(string, string) error
	ListDir(string) ([]string, error)
	Property(string) (FileProp, error)
}

type FileProp struct {
}

type SMBInfo struct {
	host      string
	user      string
	password  string
	sharename string
}

type SMBFetcher struct {
	Info  *SMBInfo
	share *smb2.Share
}

func (f SMBFetcher) Create(info *SMBInfo) *SMBFetcher {
	return &SMBFetcher{info, nil}
}

func (f *SMBFetcher) Connect() error {
	conn, err := net.Dial("tcp", f.Info.host+":445")
	if err != nil {
		return err
	}

	dialer := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     f.Info.user,
			Password: f.Info.password,
		},
	}
	session, err := dialer.Dial(conn)
	if err != nil {
		conn.Close()
		return err
	}

}

func main() {
	conn, err := net.Dial()
	d := &smb2.Dialer{}
	s, err := d.Dial(conn)
	s.Mount()
}
