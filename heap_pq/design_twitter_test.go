package heappq

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 10)
		tw.PostTweet(2, 20)

		assertFeed(t, tw.GetNewsFeed(1), []int{10})
		assertFeed(t, tw.GetNewsFeed(2), []int{20})

		tw.Follow(1, 2)
		assertFeed(t, tw.GetNewsFeed(1), []int{20, 10})
		assertFeed(t, tw.GetNewsFeed(2), []int{20})

		tw.Unfollow(1, 2)
		assertFeed(t, tw.GetNewsFeed(1), []int{10})
	})

	t.Run("empty feed", func(t *testing.T) {
		tw := NewTwitter()
		assertFeed(t, tw.GetNewsFeed(1), []int{})
	})

	t.Run("max 10 tweets", func(t *testing.T) {
		tw := NewTwitter()
		for i := 0; i < 15; i++ {
			tw.PostTweet(1, i)
		}
		feed := tw.GetNewsFeed(1)
		if len(feed) != 10 {
			t.Fatalf("expected 10 tweets, got %d", len(feed))
		}
		// Most recent first: 14, 13, 12, ..., 5
		assertFeed(t, feed, []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5})
	})

	t.Run("follow and unfollow", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 1)
		tw.Follow(2, 1)
		assertFeed(t, tw.GetNewsFeed(2), []int{1})

		tw.Unfollow(2, 1)
		assertFeed(t, tw.GetNewsFeed(2), []int{})
	})

	t.Run("merged feed ordering", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 10)
		tw.PostTweet(2, 20)
		tw.PostTweet(1, 30)
		tw.PostTweet(2, 40)

		tw.Follow(1, 2)
		// Most recent first across both users
		assertFeed(t, tw.GetNewsFeed(1), []int{40, 30, 20, 10})
	})

	t.Run("follow self is no-op", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 10)
		tw.Follow(1, 1)
		// Should still just see own tweet once
		assertFeed(t, tw.GetNewsFeed(1), []int{10})
	})

	t.Run("unfollow without follow", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 10)
		// Should not panic
		tw.Unfollow(1, 2)
		assertFeed(t, tw.GetNewsFeed(1), []int{10})
	})

	t.Run("multiple followers see same tweets", func(t *testing.T) {
		tw := NewTwitter()
		tw.PostTweet(1, 10)
		tw.Follow(2, 1)
		tw.Follow(3, 1)

		assertFeed(t, tw.GetNewsFeed(2), []int{10})
		assertFeed(t, tw.GetNewsFeed(3), []int{10})
	})
}

func assertFeed(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) == 0 && len(want) == 0 {
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GetNewsFeed() = %v, want %v", got, want)
	}
}
