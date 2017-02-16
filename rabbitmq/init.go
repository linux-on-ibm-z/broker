package rabbitmq

import (
	"github.com/Sirupsen/logrus"
	"github.com/rai-project/broker"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

var (
	log *logrus.Entry
)

func init() {
	config.OnInit(func() {
		log = logger.New().WithField("pkg", "rabbitmq")
	})
	config.AfterInit(func() {
		broker.Standard = New()
	})
}