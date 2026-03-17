package heappq

import "container/heap"

/*
Implement a simplified version of Twitter which allows users to post tweets, follow/unfollow each other, and view the 10 most recent tweets within their own news feed.

Users and tweets are uniquely identified by their IDs (integers).

Implement the following methods:

Twitter() Initializes the twitter object.
void postTweet(int userId, int tweetId) Publish a new tweet with ID tweetId by the user userId. You may assume that each tweetId is unique.
List<Integer> getNewsFeed(int userId) Fetches at most the 10 most recent tweet IDs in the user's news feed. Each item must be posted by users who the user is following or by the user themself. Tweets IDs should be ordered from most recent to least recent.
void follow(int followerId, int followeeId) The user with ID followerId follows the user with ID followeeId.
void unfollow(int followerId, int followeeId) The user with ID followerId unfollows the user with ID followeeId.
Example 1:

Input:
["Twitter", "postTweet", [1, 10], "postTweet", [2, 20], "getNewsFeed", [1], "getNewsFeed", [2], "follow", [1, 2], "getNewsFeed", [1], "getNewsFeed", [2], "unfollow", [1, 2], "getNewsFeed", [1]]

Output:
[null, null, null, [10], [20], null, [20, 10], [20], null, [10]]

Explanation:
Twitter twitter = new Twitter();
twitter.postTweet(1, 10); // User 1 posts a new tweet with id = 10.
twitter.postTweet(2, 20); // User 2 posts a new tweet with id = 20.
twitter.getNewsFeed(1);   // User 1's news feed should only contain their own tweets -> [10].
twitter.getNewsFeed(2);   // User 2's news feed should only contain their own tweets -> [20].
twitter.follow(1, 2);     // User 1 follows user 2.
twitter.getNewsFeed(1);   // User 1's news feed should contain both tweets from user 1 and user 2 -> [20, 10].
twitter.getNewsFeed(2);   // User 2's news feed should still only contain their own tweets -> [20].
twitter.unfollow(1, 2);   // User 1 unfollows user 2.
twitter.getNewsFeed(1);   // User 1's news feed should only contain their own tweets -> [10].
Constraints:

1 <= userId, followerId, followeeId <= 100
0 <= tweetId <= 1000
*/

type UserID int

type FollowingSet map[UserID]bool

type Tweet struct {
	TweetID  int
	tweetIdx int64
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].tweetIdx < h[j].tweetIdx }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return x
}

type Twitter struct {
	tweetCounter int64
	tweets       map[UserID][]*Tweet
	following    map[UserID]FollowingSet
}

func NewTwitter() Twitter {
	return Twitter{
		tweets:    make(map[UserID][]*Tweet),
		following: make(map[UserID]FollowingSet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	uid := UserID(userId)
	this.initUser(uid)
	// each user's tweets is naturally ordered from least -> most recent.
	// cap them at size 10 so that when putting the tweets into a heap, the max
	// amount of tweets we put in are 10, removing the "m" term from the O(n*m*log(n+m))
	tweets := this.tweets[uid]
	tweets = append(tweets, &Tweet{
		TweetID:  tweetId,
		tweetIdx: this.tweetCounter,
	})
	if len(tweets) > 10 {
		tweets = tweets[1:]
	}
	this.tweets[uid] = tweets
	this.tweetCounter++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	uid := UserID(userId)
	this.initUser(uid)

	h := &TweetHeap{}
	heap.Init(h)

	followings, _ := this.following[uid]
	for followeeID := range followings {
		tweets, _ := this.tweets[followeeID]
		for _, tweet := range tweets { // max size of 10
			heap.Push(h, tweet)
			if h.Len() > 10 { // keep the heap of size 10
				heap.Pop(h)
			}
		}
	}

	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = ((heap.Pop(h)).(*Tweet)).TweetID
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	followerID, followeeID := UserID(followerId), UserID(followeeId)
	this.initUser(followerID)
	this.initUser(followeeID)
	this.following[followerID][followeeID] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followerID, followeeID := UserID(followerId), UserID(followeeId)
	this.initUser(followerID)
	this.initUser(followeeID)
	if followerID == followeeID { // don't let them unfollow themself
		return
	}
	delete(this.following[followerID], followeeID)
}

func (this *Twitter) initUser(uid UserID) {
	if _, ok := this.tweets[uid]; ok { // use the initialisation of the tweets array as the source of truth
		return
	}
	this.tweets[uid] = []*Tweet{}
	this.following[uid] = FollowingSet{}
	this.following[uid][uid] = true // add the same user into their list since they "follow" themselves
}
