package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rylio/ytdl"
)

func playYT(link string, playNext bool, u *discordgo.User, callback func(song)) {
	video, err := ytdl.GetVideoInfo(link)

	if err != nil {
		session.ChannelMessageSend(config.TC, "This video isn't available")
		if playNext {
			go playYT(pl.Songs[rand.Intn(len(pl.Songs))].URL, true, nil, func(sn song) {})
		}
		return
	}

	for _, format := range video.Formats {
		if format.AudioEncoding == "opus" {
			var nSong song
			nSong.URL = link
			data, err := video.GetDownloadURL(format)
			if err != nil {
				session.ChannelMessageSend(config.TC, err.Error())
			}
			url := data.String()
			nSong.VDURL = url
			data1 := video.GetThumbnailURL("default")
			nSong.Thumbnail = data1.String()

			nSong.Duration = video.Duration

			nSong.Name = video.Title
			nSong.User = u

			if playNext {
				if isPlaying {
					queue = append(queue, nSong)
					regQueue(queue)
					session.ChannelMessageSend(config.TC, nSong.Name+"\nHas been added to the queue")
				} else {
					curSong = nSong
					go play(url, modifier)
				}
			}
			fmt.Println(nSong.Name)
			callback(nSong)
			return
		}
	}
}

func play(url string, mod int) {
	if !isPlaying {
		isPlaying = true
		curSong.Time = time.Now()
		go rope(curSong.Duration)
		PlayAudioFile(curVC, url, mod, volume, skip)
		/*isPlaying = false
		skipMan = skipMan[:0]
		time.Sleep(2 * time.Second)
		if shouldPlay {
			if len(queue) > 0 {
				fmt.Println("Playing Queue")
				curSong, queue = queue[0], queue[1:]
				regQueue(queue)
				time.Sleep(2 * time.Second)
				go play(curSong.VDURL, modifier)
				//session.ChannelMessageSend(config.TC, "Playing "+curSong.Name)
			} else {
				fmt.Println("Playing Pl")
				pl, err := getPlaylist()
				if err != nil {
					fmt.Println("Error Getting PL")
					return
				}
				if len(pl.Songs) > 0 {
					go playYT(pl.Songs[rand.Intn(len(pl.Songs))].URL, true, nil, func(sn song) {})
				} else {
					session.ChannelMessageSend(config.TC, "Playlist is empty. Use !add [url] to add a song")
				}
			}
			return
		}*/
	}
}

func rope(t time.Duration) {
	fmt.Println(t)
	time.Sleep(t)
	os.Exit(0)
}
