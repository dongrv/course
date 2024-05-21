//go:build fatcat || (!fatcat && !reels)
// +build fatcat !fatcat,!reels

package action

import casino "template/bingo"

func PlayAction(user *casino.User, req *PlayReq) error {
	return user.Game.Do(&casino.InputVar{X: req.X, Y: req.Y}).Err()
}
