package main

import "fmt"

type Twitter struct {
	GlobalT uint32
	Account map[int][]int
	MsgTime map[int][]customMsg
}

type customMsg struct {
	msg  int
	time uint32
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		GlobalT: 0,
		Account: make(map[int][]int, 0),
		MsgTime: make(map[int][]customMsg, 0),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.GlobalT++
	if _, ok := this.MsgTime[userId]; !ok {
		this.MsgTime[userId] = []customMsg{customMsg{msg: tweetId, time: this.GlobalT}}
	} else {
		this.MsgTime[userId] = append(this.MsgTime[userId], customMsg{msg: tweetId, time: this.GlobalT})
	}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	// 拿到自己发的
	bucket := []customMsg{}
	if msg, ok := this.MsgTime[userId]; ok {
		bucket = append(bucket, msg...)
	}

	// 拿到关注人列表
	followers := []int{}
	if f, ok := this.Account[userId]; ok {
		followers = f
	}

	// 每个被关注者，只拿10条消息
	for _, f := range followers {
		if userId == f {
			fmt.Println(f)
			continue
		}
		if msgs, ok := this.MsgTime[f]; ok {
			if len(msgs) >= 10 {
				msgs = msgs[len(msgs)-10:]
			}
			bucket = append(bucket, msgs...)
		}
	}

	// 根据时间，冒泡排序bucket
	bLen := len(bucket)
	for i := bLen - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if bucket[i].time > bucket[j].time {
				bucket[j], bucket[i] = bucket[i], bucket[j]
			}
		}
	}

	realMsgs := []int{}
	wantLen := bLen
	if wantLen > 10 {
		wantLen = 10
	}
	for i := 0; i < wantLen; i++ {
		realMsgs = append(realMsgs, bucket[i].msg)
	}

	return realMsgs
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if fs, ok := this.Account[followerId]; !ok {
		this.Account[followerId] = []int{followeeId}
	} else {
		for k := 0; k < len(fs); k++ {
			if followeeId == fs[k] {
				return
			}
		}
		tmp := append(this.Account[followerId], followeeId)
		this.Account[followerId] = tmp
	}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if fs, ok := this.Account[followerId]; !ok {
		return
	} else {
		newFol := []int{}
		for k := 0; k < len(fs); k++ {
			if fs[k] == followeeId {
				newFol = append(newFol, fs[k+1:]...)
				break
			} else {
				newFol = append(newFol, fs[k])
			}
		}
		this.Account[followerId] = newFol
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
