//go:build fatcat
// +build fatcat

package action

import casino "template/bingo"

func PlayAction(user *casino.User, req *PlayReq) error {
	return user.Game.Do(&casino.InputVar{X: req.X, Y: req.Y}).Err()
}
