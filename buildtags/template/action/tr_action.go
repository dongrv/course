//go:build reels
// +build reels

package action

import casino "template/bingo"

func PlayAction(user *casino.User, _ *PlayReq) error {
	return user.Game.Do(nil).Err()
}
