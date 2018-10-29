# A simple  tool to follow and retweet the projects that i work on
You need to  have a twitter oath2 access token for more on that go to 
https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html

the usage is simple just execute 

**twiterBot executable** --accessToken **token** --accessTokenSecret **tokenSecret**  --consumerKey **consumerKey** --consumerSecret **consumerSecret**

this will start a client with the credentials provided and will look for the tweets that contain the words that are at the followlist.txt (the new tweets, not those in the past)

the followlist.txt that is in my repo are for the projects that i work and follow.

you can use your own list. If so just use at your list my twitter handle @othonass :-)

# In order to build it for yourself you need go version 1.11 since the new version of bot uses go modules

# WARNING 
The tool does some basic profanity check in English so please check to see wewhat you retweet. If you just want to retweet an user just add his ID at the follow list at the followlist.txt and LEAVE blank the track list. You can find the twitter ID from this site 
http://gettwitterid.com/