package mucp

import (
	"time"

	"github.com/zjllib/go-micro/v3/config/source"
	proto "github.com/zjllib/go-micro/plugins/config/source/mucp/v3/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
