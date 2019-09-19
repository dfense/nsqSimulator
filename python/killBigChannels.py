import requests
import pprint 

baseURL="http://localhost:4151"
statsURL=baseURL+"/stats"
channelDeleteURL=baseURL+"/channel/delete"
topicDeleteURL=baseURL+"/topic/delete"

maxDepth=10


params=dict(format='json')
pp = pprint.PrettyPrinter(indent=4)

## GET stats
resp = requests.get(url=statsURL, params=params)
data = resp.json()

## wipe out all channels and the topics that own them
for topic in data['topics']:
#        pp.pprint(data)
        print("topic: %s depth: %s" % (topic['topic_name'], topic['depth']))
        if int(topic['depth']) > maxDepth:
            print("=== %s exceeds max of 10, deleting ===" % (topic['topic_name']))
            topic_params=dict(topic=topic['topic_name'])
            r = requests.post(url = topicDeleteURL, params = topic_params)
            print ("Topic Delete ResultCode: %d" % r.status_code)

            for channel in topic['channels']:
                print "\t killing ... ", channel['channel_name']
                channel_params=dict(topic=topic['topic_name'],channel=channel['channel_name'])
                r = requests.post(url = channelDeleteURL, params = channel_params)
                print "Channel Delete ResultCode: ", r.status_code

            topic_params=dict(topic=topic['topic_name'])
            r = requests.post(url = topicDeleteURL, params = topic_params)
            print "Topic Delete ResultCode: ", r.status_code
