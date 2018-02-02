package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	swearjar "github.com/snicol/swearjar-go"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	Log               *logrus.Logger
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "info.log",
		logrus.ErrorLevel: "error.log",
	}
	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	log := &logger{Log}
	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"raid2018", "certcoop", "virtuwind", "semiotics_eu", "cybersure_eu", "CE_Iot", "Ideal_Cities",
			"@certcoop", "@virtuwind", "@semiotics_eu", "@cybersure_eu", "@CE_Iot", "@Ideal_Cities"},
		// "follow": []string{"certcoop", "VirtuWind", "semiotics_eu", "cybersure_eu"},

	})
	swears, err := swearjar.Load()

	if err != nil {
		log.Fatal(err)
	}

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Warningf("received unexpected value of type %T", v)
			continue
		}

		if t.RetweetedStatus != nil {
			continue
		}
		profane, err := swears.Profane(t.Text)
		if err != nil {
			log.Warningf("Could not check profanity of %s because of %v", t.Text, err)
		}
		if profane {
			log.Warningf("Profanity found in tweet: %s \n\t will not retweet it", t.Text)
			continue
		}
		_, err = api.Retweet(t.Id, false)
		if err != nil {
			log.Errorf("could not retweet %d: %v", t.Id, err)
			continue
		}
		log.Infof("will retweet %d from %s", t.Id, t.User.Name)
	}
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }