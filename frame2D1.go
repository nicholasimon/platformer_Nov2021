package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//MARK:var
var (

	//INTRO
	intro          bool
	introtxt       = []string{"s", "p", "l", "a", "t", "a"}
	introtxtsize   = int32(500)
	introcircs     []introcirc
	introcircsorig []introcirc
	introtimer     = fps
	introsel       int
	//XTRAS
	xtras = make([]xtra, 0)
	vines = make([]detail, 0)
	vine  bool
	//BULLETS
	bulletmod = bullmods{}

	//PLAYER
	player    = playrec{}
	playerspd = float32(12)

	//WEAPONS
	weapons        = make([]weapon, 0)
	weapmulti      float32
	weapontilesize = float32(16)
	//OBJ
	objs = make([]obj, 0)

	//LEVEL
	innerl, innerr, innert, innerb, gravity, tilew, tileimgsize, tilemulti, bordl, bordr, bordb, bordt float32

	currentlev = 1
	pause      bool
	levelh     int32
	//FX
	ghost, scan, noise, rain bool
	scanlines                = make([]rl.Vector2, 0)
	scanlines2               = make([]rl.Vector2, 0)
	raindrops                []detail
	//SCREEN
	scrw, scrh, monh, monw, scrh1, scrw1 int32
	scrwf32, scrhf32, monhf32, monwf32   float32
	scrwint, scrhint, monhint, monwint   int
	//CAMS
	camcntr rl.Vector2
	//CORE
	imgs                          rl.Texture2D
	frames                        int
	mous, mous2d                  rl.Vector2
	camera, camintro, camintrotxt rl.Camera2D
	fps                           = int32(60)

	dev, grid, cntr, sprites bool

	//IMGS

	vine1      = rl.NewRectangle(763, 8, 16, 16)
	dinoshadow = rl.NewRectangle(35, 202, 24, 24)
	ground1    = rl.NewRectangle(0, 199, 32, 32)
	bulletimgr = rl.NewRectangle(334, 150, 16, 16)
	bulletimgl = rl.NewRectangle(351, 150, 16, 16)
	cursorimg  = rl.NewRectangle(1, 182, 12, 12)
	blankrec   = rl.NewRectangle(0, 0, 0, 0)
	tile1      = rl.NewRectangle(0, 0, 32, 32)
	dinogr     = rl.NewRectangle(5, 118, 24, 24)
	dinogl     = rl.NewRectangle(339, 86, 24, 24)

	ninemmr      = rl.NewRectangle(1, 149, 16, 16)
	silencerr    = rl.NewRectangle(17, 149, 16, 16)
	shotgunr     = rl.NewRectangle(151, 171, 28, 9)
	submachiner  = rl.NewRectangle(49, 149, 16, 16)
	uzzir        = rl.NewRectangle(65, 149, 16, 16)
	ak47r        = rl.NewRectangle(81, 149, 16, 16)
	bazookar     = rl.NewRectangle(97, 149, 16, 16)
	throwingaxer = rl.NewRectangle(115, 147, 16, 16)
	axer         = rl.NewRectangle(131, 147, 16, 16)
	bombr        = rl.NewRectangle(148, 149, 16, 16)
	ninemml      = rl.NewRectangle(315, 149, 16, 16)
	silencerl    = rl.NewRectangle(299, 149, 16, 16)
	shotgunl     = rl.NewRectangle(118, 171, 28, 9)
	submachinel  = rl.NewRectangle(267, 149, 16, 16)
	uzzil        = rl.NewRectangle(251, 149, 16, 16)
	ak47l        = rl.NewRectangle(235, 149, 16, 16)
	bazookal     = rl.NewRectangle(219, 149, 16, 16)
	throwingaxel = rl.NewRectangle(201, 147, 16, 16)
	axel         = rl.NewRectangle(185, 147, 16, 16)
	bombl        = rl.NewRectangle(168, 149, 16, 16)

	trees = make([]rl.Rectangle, 15)
	//TIMERS
	onoff1, onoff3, onoff5, onoff10, onoff15, onoff20, onoff30, onoff60, onoff120 bool
)

//MARK:struct
type introcirc struct {
	wid, rad float32
	v2       rl.Vector2
	color    rl.Color
	lr       bool
}
type xtra struct {
	activ bool
	img   rl.Rectangle
	multi float32
	v2tl  rl.Vector2
}
type detail struct {
	v2       rl.Vector2
	color    rl.Color
	radius   float32
	img, rec rl.Rectangle

	recs []rl.Rectangle
}
type weapon struct {
	name                                string
	bullimgl, bullimgr, imgr, imgl, rec rl.Rectangle
	rotat                               float32
	color                               rl.Color
}
type playrec struct {
	rec, recc, fallrec, img rl.Rectangle

	direc, v2tl rl.Vector2

	rotat float32

	weaponnum, jumptimer int

	lr, onobj, airborne bool

	currweapon weapon
}

type obj struct {
	img, rec, recc          rl.Rectangle
	direc, v2tl             rl.Vector2
	weig, rotat, rotatmulti float32

	objsin []obj

	collisrecs []rl.Rectangle

	multi int

	collistimer, nochecktimer int32

	grav, fall, bounce, collis, activ, rotating, rotatlr, bullet, complex bool
}
type bullmods struct {
	wid, spd float32
}

func nocam() { //MARK: nocam

	fx()
}
func cam() { //MARK: cam

	if !pause {
		//ground
		x := float32(0)
		y := bordb
		for {

			destrec := rl.NewRectangle(x, y, tilew*2, tilew*2)
			origin := rl.NewVector2(tilew/2, tilew/2)
			rl.DrawTexturePro(imgs, ground1, destrec, origin, 0, rl.White)

			if ghost {
				destrec.X += rFloat32(-2, 3)
				destrec.Y += rFloat32(-2, 3)
				rl.DrawTexturePro(imgs, ground1, destrec, origin, 0, rl.Fade(rl.White, 0.3))
			}

			x += tilew

			if x > monwf32 {
				break
			}
		}

		//xtras
		for a := 0; a < len(xtras); a++ {

			destrec := rl.NewRectangle(xtras[a].v2tl.X-6, xtras[a].v2tl.Y+6, xtras[a].img.Width*xtras[a].multi, xtras[a].img.Height*xtras[a].multi)

			origin := rl.NewVector2((xtras[a].img.Width*xtras[a].multi)/2, xtras[a].img.Height*xtras[a].multi/2)

			rl.DrawTexturePro(imgs, xtras[a].img, destrec, origin, 0, rl.Black)

			destrec = rl.NewRectangle(xtras[a].v2tl.X, xtras[a].v2tl.Y, xtras[a].img.Width*xtras[a].multi, xtras[a].img.Height*xtras[a].multi)

			rl.DrawTexturePro(imgs, xtras[a].img, destrec, origin, 0, rl.Fade(rl.White, 0.5))

			if ghost {
				destrec.X += rFloat32(-5, 6)
				destrec.Y += rFloat32(-5, 6)
				rl.DrawTexturePro(imgs, xtras[a].img, destrec, origin, 0, rl.Fade(rl.White, 0.2))
			}
		}
		//vines
		if vine {
			for a := 0; a < len(vines); a++ {
				for b := 0; b < len(vines[a].recs); b++ {
					origin := rl.NewVector2(tilew/2, tilew/2)
					rl.DrawTexturePro(imgs, vines[a].img, vines[a].recs[b], origin, 0, randomgreen())
					vines[a].recs[b].Y -= 2

					if vines[a].recs[b].Y < 0 {
						vines[a].recs[b].Y = scrhf32 + tilew
					}
				}

				if rl.CheckCollisionRecs(vines[a].rec, player.rec) {
					player.direc.Y = -10
				}
			}
		}

		//objs
		for a := 0; a < len(objs); a++ {
			if objs[a].activ {

				if sprites {
					if len(objs[a].objsin) > 0 {
						for b := 0; b < len(objs[a].objsin); b++ {
							if objs[a].objsin[b].activ {
								destrec := rl.NewRectangle(objs[a].objsin[b].rec.X+tilew/2, objs[a].objsin[b].rec.Y+tilew/2, tilew, tilew)
								origin := rl.NewVector2(tilew/2, tilew/2)

								rl.DrawTexturePro(imgs, objs[a].objsin[b].img, destrec, origin, objs[a].objsin[b].rotat, rl.White)

								if ghost {
									destrec.X += rFloat32(-5, 6)
									destrec.Y += rFloat32(-5, 6)
									rl.DrawTexturePro(imgs, objs[a].objsin[b].img, destrec, origin, objs[a].objsin[b].rotat, rl.Fade(rl.White, 0.3))
								}

								if objs[a].objsin[b].rotating {
									if objs[a].objsin[b].rotatlr {
										objs[a].objsin[b].rotat += objs[a].objsin[b].rotatmulti
									} else {
										objs[a].objsin[b].rotat -= objs[a].objsin[b].rotatmulti
									}
								}
							}
						}
					} else {

						if objs[a].img == blankrec {
							rl.DrawRectangleLinesEx(objs[a].rec, 1, rl.Magenta)
							rl.DrawRectangleLinesEx(objs[a].recc, 1, rl.Green)
						} else {
							destrec := rl.NewRectangle(objs[a].rec.X+tilew/2, objs[a].rec.Y+tilew/2, tilew, tilew)
							origin := rl.NewVector2(tilew/2, tilew/2)
							if objs[a].bullet {

								destrec = rl.NewRectangle(objs[a].rec.X+tilew/4, objs[a].rec.Y+tilew/4, tilew/2, tilew/2)
								origin = rl.NewVector2(tilew/4, tilew/4)

								if objs[a].direc.X < 0 {
									objs[a].img = weapons[player.weaponnum].bullimgl
								} else if objs[a].direc.X > 0 {
									objs[a].img = weapons[player.weaponnum].bullimgr
								}

							}

							rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].rotat, rl.White)

							if ghost {
								destrec.X += rFloat32(-5, 6)
								destrec.Y += rFloat32(-5, 6)
								rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].rotat, rl.Fade(rl.White, 0.3))
							}

							if objs[a].rotating {
								if objs[a].rotatlr {
									objs[a].rotat += objs[a].rotatmulti
								} else {
									objs[a].rotat -= objs[a].rotatmulti
								}
							}
						}
					}
				}

				if grid {
					rl.DrawRectangleLinesEx(objs[a].rec, 1, rl.Magenta)
					rl.DrawRectangleLinesEx(objs[a].recc, 1, rl.Green)
					if len(objs[a].objsin) > 0 {
						for b := 0; b < len(objs[a].objsin); b++ {
							if objs[a].objsin[b].activ {
								rl.DrawRectangleLinesEx(objs[a].objsin[b].rec, 1, brightyellow())
								rl.DrawRectangleLinesEx(objs[a].objsin[b].recc, 1, rl.Green)
							}
						}
					}

					if objs[a].collis {
						rl.DrawRectangleRec(objs[a].rec, rl.Fade(rl.Magenta, 0.2))
						objs[a].collistimer--
						if objs[a].collistimer == 0 {
							objs[a].collis = false
						}
					}
				}

				if objs[a].rotating {
					if objs[a].rotatlr {
						objs[a].rotat += objs[a].rotatmulti
					} else {
						objs[a].rotat -= objs[a].rotatmulti
					}
				}
			}
		}

		//player
		if sprites {

			destrecweap := rl.NewRectangle(player.rec.X+player.rec.Width+12, player.rec.Y+(player.currweapon.imgr.Height*6), player.currweapon.imgr.Width*weapmulti, player.currweapon.imgr.Height*weapmulti)
			originweap := rl.NewVector2((player.currweapon.imgr.Width*weapmulti)/2, (player.currweapon.imgr.Height*weapmulti)/2)

			if !player.lr {
				destrecweap = rl.NewRectangle(player.rec.X-12, player.rec.Y+(player.currweapon.imgr.Height*6), player.currweapon.imgr.Width*weapmulti, player.currweapon.imgr.Height*weapmulti)
				originweap = rl.NewVector2((player.currweapon.imgr.Width*weapmulti)/2, (player.currweapon.imgr.Height*weapmulti)/2)
			}

			if onoff30 {
				destrecweap.Y -= 3
			}

			player.currweapon.rec.X = destrecweap.X - ((player.currweapon.imgr.Width * weapmulti) / 2)
			player.currweapon.rec.Y = destrecweap.Y - ((player.currweapon.imgr.Height * weapmulti) / 2)

			if player.lr {
				rl.DrawTexturePro(imgs, player.currweapon.imgr, destrecweap, originweap, player.currweapon.rotat, rl.White)
			} else {
				rl.DrawTexturePro(imgs, player.currweapon.imgl, destrecweap, originweap, player.currweapon.rotat, rl.White)
			}
			if grid {
				rl.DrawRectangleLinesEx(player.currweapon.rec, 1, rl.Magenta)
			}

			destrec := rl.NewRectangle(player.rec.X+player.rec.Width/2, player.rec.Y+player.rec.Height/2, player.rec.Width, player.rec.Height)
			destrecshadow := rl.NewRectangle(destrec.X, destrec.Y, destrec.Width, destrec.Width)
			origin := rl.NewVector2(player.rec.Width/2, player.rec.Height/2)

			if rolldice()+rolldice() == 12 {
				destrec.Y -= 3
			}

			rl.DrawTexturePro(imgs, dinoshadow, destrecshadow, origin, 0, rl.White)

			rl.DrawTexturePro(imgs, player.img, destrec, origin, player.rotat, rl.White)

			//player.currweapon.rotat++

			if ghost {
				destrec.X += rFloat32(-5, 6)
				destrec.Y += rFloat32(-5, 6)

				rl.DrawTexturePro(imgs, player.img, destrec, origin, player.rotat, rl.Fade(rl.White, 0.3))
			}

		}

		if grid {

			rl.DrawRectangleLinesEx(player.rec, 1, rl.Magenta)
			rl.DrawRectangleLinesEx(player.fallrec, 1, brightyellow())
			rl.DrawRectangleLinesEx(player.recc, 1, rl.Green)
		}

	}

}
func ddev() { //MARK: ddev

	//txt := strconv.FormatBool(player.fall)
	//	rl.DrawText(txt, 100, 100, 20, rl.White)

	rl.DrawRectangle(0, 0, scrw, 30, rl.Fade(rl.Magenta, 0.2))
	rl.DrawRectangle(scrw-100, 30, 100, scrh, rl.Fade(rl.Magenta, 0.2))

	x2 := scrw - 95
	y2 := int32(30)
	txt := fmt.Sprint(camera.Target.Y)
	rl.DrawText(txt, x2, y2, 10, rl.White)
	rl.DrawText("camy", x2+40, y2, 10, rl.White)

	mosy := fmt.Sprint(mous.Y)
	mosx := fmt.Sprint(mous.X)

	rl.DrawText(mosx, int32(mous.X+20), int32(mous.Y), 10, rl.White)
	rl.DrawText(mosy, int32(mous.X+20), int32(mous.Y+12), 10, rl.White)

	x := scrwf32 - 60
	y := float32(5)

	currentTime := time.Now()
	rl.DrawText(currentTime.Format("15:04"), int32(x), int32(y), 20, rl.White)

	x = 5
	rec := rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(3)
		}
	}
	rl.DrawText("1920", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(2)
		}
	}
	rl.DrawText("1600", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(4)
		}
	}
	rl.DrawText("1440", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(5)
		}
	}
	rl.DrawText("1366", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(6)
		}
	}
	rl.DrawText("1280", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(7)
		}
	}
	rl.DrawText("2160", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55

	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			scr(8)
		}
	}
	rl.DrawText("3840", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if grid {
				grid = false
			} else {
				grid = true
			}
		}
	}
	rl.DrawText("grid", rec.ToInt32().X+10, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if sprites {
				sprites = false
			} else {
				sprites = true
			}
		}
	}
	rl.DrawText("sprites", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if cntr {
				cntr = false
			} else {
				cntr = true
			}
		}
	}
	rl.DrawText("center", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if scan {
				scan = false
			} else {
				scan = true
			}
		}
	}
	rl.DrawText("scan", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if ghost {
				ghost = false
			} else {
				ghost = true
			}
		}
	}
	rl.DrawText("ghost", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if noise {
				noise = false
			} else {
				noise = true
			}
		}
	}
	rl.DrawText("noise", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if rain {
				rain = false
			} else {
				rain = true
			}
		}
	}
	rl.DrawText("rain", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)

	x += 55
	rec = rl.NewRectangle(x, y, 50, 20)
	if rl.CheckCollisionPointRec(mous, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if intro {
				intro = false
				pause = false
			} else {
				intro = true
				pause = true
			}
		}
	}
	rl.DrawText("intro", rec.ToInt32().X+7, rec.ToInt32().Y+5, 10, rl.White)
	rl.DrawRectangleLinesEx(rec, 0.5, rl.White)
}

func initial() { //MARK: initial

	//	intro = true
	//	pause = true

	levelh = 3

	bulletmod.spd = 20
	bulletmod.wid = tilew / 2

	sprites = true
	gravity = 2
	//grid = true
	dev = true
	monh = int32(rl.GetScreenHeight())
	monw = int32(rl.GetScreenWidth())
	scrh1 = monh
	scrw1 = monw
	scrh = monh * levelh
	scrw = monw

	monhf32 = float32(monh)
	monwf32 = float32(monw)
	scrhf32 = float32(scrh)
	scrwf32 = float32(scrw)

	tileimgsize = 32
	tilemulti = 2
	weapmulti = 4
	tilew = tileimgsize * tilemulti

	innert = scrhf32 - (monhf32 / 4)
	innerb = monhf32 / 4
	innerr = scrwf32 - (tilew * 4)
	innerl = tilew * 4

	bordb = float32(scrh) - tilew
	bordt = float32(0)
	bordl = float32(0)
	bordr = float32(scrw)

	camera.Target.Y = scrhf32 - monhf32

	camintrotxt.Zoom = 1.0
	camintro.Zoom = 1.0
	camintrotxt.Target.X = scrwf32 / 2
	camintrotxt.Target.Y = scrhf32 / 2

	makeintro()
	makeobjs()
	makeweapons()
	makextras()
	makeplayer(1)

	y := float32(0)
	for {
		scanlines = append(scanlines, rl.NewVector2(0, y))
		scanlines2 = append(scanlines2, rl.NewVector2(monwf32, y))
		y += 5

		if y > scrhf32 {
			break
		}
	}

}

func makeweapons() { //MARK: makeweapons

	weapons = make([]weapon, 10)

	weapons[0].name = "9mm"
	weapons[0].imgr = ninemmr
	weapons[0].imgl = ninemml
	weapons[0].bullimgr = bulletimgr
	weapons[0].bullimgl = bulletimgl

	weapons[1].name = "silencer"
	weapons[1].imgr = silencerr
	weapons[1].imgl = silencerl

	weapons[2].name = "shotgun"
	weapons[2].imgr = shotgunr
	weapons[2].imgl = shotgunl
	weapons[2].bullimgr = bulletimgr
	weapons[2].bullimgl = bulletimgl

	weapons[3].name = "submachinegun"
	weapons[3].imgr = submachiner
	weapons[3].imgl = submachinel

	weapons[4].name = "uzzi"
	weapons[4].imgr = uzzir
	weapons[4].imgl = uzzil

	weapons[5].name = "ak47"
	weapons[5].imgr = ak47r
	weapons[5].imgl = ak47l

	weapons[6].name = "bazoooka"
	weapons[6].imgr = bazookar
	weapons[6].imgl = bazookal

	weapons[7].name = "throwing axe"
	weapons[7].imgr = throwingaxer
	weapons[7].imgl = throwingaxel

	weapons[8].name = "axe"
	weapons[8].imgr = axer
	weapons[8].imgl = axel

	weapons[9].name = "bomb"
	weapons[9].imgr = bombr
	weapons[9].imgl = bombl

	for a := 0; a < len(weapons); a++ {
		weapons[a].rec = rl.NewRectangle(0, 0, weapontilesize*weapmulti, weapontilesize*weapmulti)
		weapons[a].color = randomcolor()
	}

}
func makextras() { //MARK: makextras

	//vines
	newvine := detail{}
	newvine.rec = rl.NewRectangle(rFloat32(innerl, innerr), 0, tilew, scrhf32)
	newvine.img = vine1
	newvine.recs = make([]rl.Rectangle, 0)
	y := newvine.rec.Y
	x := newvine.rec.X

	for {
		destrec := rl.NewRectangle(x, y, tilew, tilew)
		newvine.recs = append(newvine.recs, destrec)
		y += tilew

		if y > scrhf32+tilew {
			break
		}
	}

	vines = append(vines, newvine)
	//trees
	num := 10
	for a := 0; a < num; a++ {

		newxtra := xtra{}
		newxtra.activ = true
		newxtra.img = trees[rInt(0, len(trees))]
		newxtra.multi = 8
		newxtra.v2tl.X = rFloat32(0, bordr)
		newxtra.v2tl.Y = bordb - (newxtra.img.Height * (newxtra.multi / 2))

		xtras = append(xtras, newxtra)
	}

	//rain
	num = 300
	for a := 0; a < num; a++ {
		raindrop := detail{}
		raindrop.radius = rFloat32(3, 6)
		raindrop.color = randombluelight()
		raindrop.v2.X = rFloat32(0, scrwf32)
		raindrop.v2.Y = rFloat32(0, scrhf32)
		raindrops = append(raindrops, raindrop)
	}

}

func upplayer() { //MARK: upplayer

	player.rec.X += player.direc.X

	if player.rec.X < bordl-tilew {
		player.rec.X = bordr + tilew
	}
	if player.rec.X > bordr+tilew {
		player.rec.X = bordl - tilew
	}

	if player.rec.Y > bordt && player.rec.Y+player.rec.Height <= bordb {
		player.rec.Y += player.direc.Y
	}

	if player.rec.Y < bordt {
		player.rec.Y = bordt + 1
	}
	if player.rec.Y+player.rec.Height > bordb {
		player.rec.Y = bordb - (player.rec.Height + 1)
	}

	player.fallrec = rl.NewRectangle(player.rec.X, player.rec.Y+player.rec.Height, player.rec.Width, 2)
	player.recc = rl.NewRectangle(player.rec.X+player.direc.X, player.rec.Y+player.direc.Y, player.rec.Width, player.rec.Height)

	if player.jumptimer != 0 {
		player.jumptimer--
		if player.jumptimer <= 0 {
			player.jumptimer = 0
			player.direc.Y = 0
		}
	}

	if player.jumptimer == 0 {
		player.airborne = true
		for a := 0; a < len(objs); a++ {
			if objs[a].activ {
				if objs[a].complex {
					for _, collisrecin := range objs[a].collisrecs {
						if rl.CheckCollisionRecs(player.fallrec, collisrecin) {
							player.airborne = false
							if objs[a].direc.X != 0 || objs[a].direc.Y != 0 {
								player.direc.Y = objs[a].direc.Y
								player.direc.X = objs[a].direc.X
							}
						}
					}
				} else {
					if rl.CheckCollisionRecs(player.fallrec, objs[a].rec) {
						player.airborne = false
						if objs[a].direc.X != 0 || objs[a].direc.Y != 0 {
							player.direc.Y = objs[a].direc.Y
							player.direc.X = objs[a].direc.X
						}
					}
				}
			}
		}

		if player.airborne && player.rec.Y < bordb-player.rec.Height {
			player.direc.Y = playerspd
		} else {
			player.direc.Y = 0
		}
	}

	if player.direc.X > 0 {
		player.lr = true
		player.img = dinogr
	} else if player.direc.X < 0 {
		player.lr = false
		player.img = dinogl
	}

	if player.direc.X != 0 {
		dinogl.X -= dinogl.Width
		if dinogl.X < 20 {
			dinogl.X = 339
		}
		dinogr.X += dinogr.Width
		if dinogr.X > 339 {
			dinogr.X = 5
		}
	}

}
func fire() { //MARK: fire

	newobj := obj{}

	newobj.activ = true
	newobj.bullet = true
	newobj.img = bulletimgr
	newobj.rec = rl.NewRectangle(player.rec.X+player.rec.Width+10, player.rec.Y+player.rec.Height/3, tilew/2, tilew/2)
	newobj.direc.X = bulletmod.spd
	newobj.collistimer = fps / 2
	if !player.lr {
		newobj.rec.X = player.rec.X - 20
		newobj.direc.X = -bulletmod.spd
		newobj.img = bulletimgl
	}

	objs = append(objs, newobj)

}
func endbull(num int) { //MARK: endbull

	objs[num].activ = false

}
func upobjs() { //MARK: upobjs

	for a := 0; a < len(objs); a++ {
		if objs[a].activ {
			if objs[a].direc.X != 0 || objs[a].direc.Y != 0 {
				if objs[a].nochecktimer == 0 {
					checkc(a)
				} else {
					objs[a].nochecktimer--
					if objs[a].nochecktimer <= 0 {
						objs[a].nochecktimer = 0
					}
				}
				if objs[a].complex {
					for b := 0; b < len(objs[a].objsin); b++ {
						objs[a].objsin[b].rec.X += objs[a].direc.X
						objs[a].objsin[b].rec.Y += objs[a].direc.Y
						objs[a].objsin[b].v2tl = rl.NewVector2(objs[a].objsin[b].rec.X, objs[a].objsin[b].rec.Y)

						objs[a].objsin[b].recc.X = objs[a].objsin[b].rec.X + objs[a].direc.X
						objs[a].objsin[b].recc.Y = objs[a].objsin[b].rec.Y + objs[a].direc.Y

						objs[a].collisrecs[b] = objs[a].objsin[b].rec

					}
				} else {
					objs[a].rec.X += objs[a].direc.X
					objs[a].rec.Y += objs[a].direc.Y
					upobjsin(a)
				}

			}
			if objs[a].grav && !objs[a].fall && objs[a].direc.Y < 6 {
				objs[a].direc.Y += gravity
			} else if objs[a].direc.Y >= 6 {
				objs[a].fall = true
			} else if objs[a].grav && objs[a].fall {
				if objs[a].direc.Y > 0 {
					objs[a].direc.Y -= rFloat32(-1, -0.5)
				} else if objs[a].direc.Y <= 0 {
					objs[a].direc.Y = 0
				}
			}
			makecrec(a)
		}
	}

}
func upobjsin(num int) { //MARK: upobjsin

	x := objs[num].rec.X
	y := objs[num].rec.Y
	count := 0

	for a := 0; a < len(objs[num].objsin); a++ {

		objs[num].objsin[a].rec = rl.NewRectangle(x, y, tilew, tilew)
		objs[num].objsin[a].v2tl = rl.NewVector2(x, y)
		x += tilew
		count++
		if count == objs[num].multi {
			count = 0
			x = objs[num].rec.X
			y += tilew
		}

	}

}
func remobjs(s []obj, index int) []obj { //MARK: remobjs
	return append(s[:index], s[index+1:]...)
}
func explode(num int) { //MARK: explode

	cntrblok := len(objs[num].objsin)
	cntrblok = cntrblok / 2
	if len(objs[num].objsin)%2 == 0 {
		cntrblok -= objs[num].multi / 2
	}

	//	halfmulti := objs[num].multi / 2

	for {
		for a := 0; a < len(objs[num].objsin); a++ {

			objs[num].objsin[a].rec.X += rFloat32(-5, 6)
			objs[num].objsin[a].rec.Y += rFloat32(-5, 6)

			objs[num].objsin[a].direc.X = rFloat32(-10, 11)
			objs[num].objsin[a].direc.Y = rFloat32(-10, 11)

			objs[num].objsin[a].nochecktimer = fps / 2
			objs[num].objsin[a].rotatmulti = rFloat32(10, 40)
			objs[num].objsin[a].rotating = true
			objs[num].objsin[a].rotat = rFloat32(0, 360)
			objs[num].objsin[a].rotatlr = flipcoin()

			objs = append(objs, objs[num].objsin[a])
			objs[num].objsin = remobjs(objs[num].objsin, a)
		}

		if len(objs[num].objsin) == 0 {
			objs = remobjs(objs, num)
			//objs[num].objsin[0].activ = false
			break
		}

	}
}

func makecrec(num int) { //MARK: makecrec
	objs[num].recc = rl.NewRectangle(objs[num].rec.X+(objs[num].direc.X*2), objs[num].rec.Y+(objs[num].direc.Y*2), objs[num].rec.Width, objs[num].rec.Height)

}
func checkc(num int) { //MARK: checkc
	if objs[num].complex {
		for _, collisrecin := range objs[num].collisrecs {
			if collisrecin.X+objs[num].direc.X <= bordl {
				changedirec(num, 2)
			}
			if collisrecin.X+collisrecin.Width+objs[num].direc.X >= bordr {
				changedirec(num, 4)
			}
			if collisrecin.Y+collisrecin.Height+objs[num].direc.Y >= bordb {
				changedirec(num, 1)
			}
			if collisrecin.Y+objs[num].direc.Y <= bordt {
				changedirec(num, 3)
			}
		}

	} else {
		if objs[num].direc.X < 0 {
			if objs[num].rec.X+objs[num].direc.X <= bordl {
				if objs[num].bullet && !objs[num].bounce {
					endbull(num)
				} else {
					changedirec(num, 2)
				}
			}
		} else if objs[num].direc.X > 0 {
			if objs[num].rec.X+objs[num].rec.Width+objs[num].direc.X >= bordr {
				if objs[num].bullet && !objs[num].bounce {
					endbull(num)
				} else {
					changedirec(num, 4)
				}
			}
		}
		if objs[num].direc.Y < 0 {
			if objs[num].rec.Y+objs[num].direc.Y <= bordt {
				if objs[num].bullet && !objs[num].bounce {
					endbull(num)
				} else {
					changedirec(num, 3)
				}
			}
		} else if objs[num].direc.Y > 0 {
			if objs[num].rec.Y+objs[num].rec.Height+objs[num].direc.Y >= bordb {
				if objs[num].bullet && !objs[num].bounce {
					endbull(num)
				} else {
					changedirec(num, 1)
				}
			}
		}
	}
	if objs[num].complex {

		for _, collisrecin := range objs[num].collisrecs {

			for num2, checkobj := range objs {
				if objs[num].activ && checkobj.activ {
					if rl.CheckCollisionRecs(collisrecin, checkobj.rec) && num2 != num {

						collision(num, num2)

						objs[num].collis = true
						objs[num2].collis = true
						objs[num].collistimer = fps / 2
						objs[num2].collistimer = fps / 2
					}
				}

			}

		}

	} else {

		for num2, checkobj := range objs {
			if objs[num].activ && checkobj.activ {
				if rl.CheckCollisionRecs(objs[num].recc, checkobj.rec) && num2 != num {

					collision(num, num2)

					objs[num].collis = true
					objs[num2].collis = true
					objs[num].collistimer = fps / 2
					objs[num2].collistimer = fps / 2
				}
			}

		}
	}

}
func collision(num, num2 int) { //MARK: collision

	if objs[num].direc.X == 0 && objs[num].direc.Y > 0 {
		if objs[num].rec.Y < objs[num2].rec.Y {
			spdy(num, rFloat32(-1.5, -0.5), true)
		} else {
			spdy(num, rFloat32(-1.5, -0.5), false)
		}
	}
	if objs[num].direc.X == 0 && objs[num].direc.Y < 0 {
		if objs[num].rec.Y < objs[num2].rec.Y {
			spdy(num, rFloat32(-1.5, -0.5), false)
		} else {
			spdy(num, rFloat32(0.5, 1.5), true)
		}
	}

	if objs[num].direc.X > 0 && objs[num].direc.Y == 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(-1.5, -0.5), true)
		} else {
			spdx(num, rFloat32(0.5, 1.5), false)
		}
	}

	if objs[num].direc.X < 0 && objs[num].direc.Y == 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(-1.5, -0.5), false)
		} else {
			spdx(num, rFloat32(-1.5, -0.5), true)
		}
	}

	if objs[num].direc.X > 0 && objs[num].direc.Y > 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(-1.5, -0.5), true)
			spdy(num, rFloat32(-1.5, -0.5), true)
		} else {
			spdx(num, rFloat32(0.5, 1.5), false)
			spdy(num, rFloat32(-1.5, -0.5), false)
		}
	}
	if objs[num].direc.X > 0 && objs[num].direc.Y < 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(-1.5, -0.5), true)
			spdy(num, rFloat32(-1.5, -0.5), false)
		} else {
			spdx(num, rFloat32(-1.5, -0.5), false)
			spdy(num, rFloat32(-1.5, -0.5), true)
		}
	}
	if objs[num].direc.X < 0 && objs[num].direc.Y < 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(0.5, 1.5), false)
			spdy(num, rFloat32(-1.5, -0.5), false)
		} else {
			spdx(num, rFloat32(-1.5, -0.5), true)
			spdy(num, rFloat32(-1.5, -0.5), true)
		}
	}
	if objs[num].direc.X < 0 && objs[num].direc.Y > 0 {
		if objs[num].rec.X < objs[num2].rec.X {
			spdx(num, rFloat32(-1.5, -0.5), true)
			spdy(num, rFloat32(-1.5, -0.5), true)
		} else {
			spdx(num, rFloat32(-1.5, -0.5), true)
			spdy(num, rFloat32(-1.5, -0.5), true)
		}
	}

	if objs[num].direc.X >= 10 {
		objs[num].direc.X -= rFloat32(0.5, 1.5)
	}
	if objs[num].direc.X <= -10 {
		objs[num].direc.X += rFloat32(0.5, 1.5)
	}
	if objs[num].direc.Y >= 10 {
		objs[num].direc.Y -= rFloat32(0.5, 1.5)
	}
	if objs[num].direc.Y <= -10 {
		objs[num].direc.Y += rFloat32(0.5, 1.5)
	}

	if objs[num2].direc.X >= 10 {
		objs[num2].direc.X -= rFloat32(0.5, 1.5)
	}
	if objs[num2].direc.X <= -10 {
		objs[num2].direc.X += rFloat32(0.5, 1.5)
	}
	if objs[num2].direc.Y >= 10 {
		objs[num2].direc.Y -= rFloat32(0.5, 1.5)
	}
	if objs[num2].direc.Y <= -10 {
		objs[num2].direc.Y += rFloat32(0.5, 1.5)
	}

}
func spdx(num int, amount float32, reverse bool) { //MARK: spdx
	objs[num].direc.X += amount

	if reverse {
		if objs[num].direc.X > 0 {
			objs[num].direc.X = -objs[num].direc.X
		} else {
			objs[num].direc.X = float32(math.Abs(float64(objs[num].direc.X)))
		}
	}
}
func spdy(num int, amount float32, reverse bool) { //MARK: spdy
	objs[num].direc.Y += amount

	if reverse {
		if objs[num].direc.Y > 0 {
			objs[num].direc.Y = -objs[num].direc.Y
		} else {
			objs[num].direc.Y = float32(math.Abs(float64(objs[num].direc.Y)))
		}
	}
}
func changedirec(objnum, num int) { //MARK: changedirec

	switch num {
	case 2:
		objs[objnum].direc.X = float32(math.Abs(float64(objs[objnum].direc.X)))
	case 4:
		objs[objnum].direc.X = -objs[objnum].direc.X
	case 1:
		objs[objnum].direc.Y = -objs[objnum].direc.Y
	case 3:
		objs[objnum].direc.Y = float32(math.Abs(float64(objs[objnum].direc.Y)))
	}

}
func makeobjs() { //MARK: makeobjs

	currenttile := tile1

	ylevel := scrhf32 - tilew*6

	for {

		platformtype := rInt(0, 11)

		switch platformtype {
		case 0: // square
			newobj := obj{}
			newobj.multi = rInt(2, 9)
			wid := float32(newobj.multi * int(tilew))
			newobj.rec = rl.NewRectangle(rFloat32(0, scrwf32), ylevel, wid, wid)
			newobj.activ = true
			newobj.objsin = make([]obj, newobj.multi*newobj.multi)
			x := newobj.rec.X
			y := newobj.rec.Y
			count := 0
			for b := 0; b < len(newobj.objsin); b++ {
				newobj.objsin[b].img = currenttile
				newobj.objsin[b].rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.objsin[b].v2tl = rl.NewVector2(x, y)
				newobj.objsin[b].activ = true

				x += tilew
				count++
				if count == newobj.multi {
					count = 0
					x = newobj.rec.X
					y += tilew
				}
			}

			objs = append(objs, newobj)

		case 1: //platform 1 block
			newobj := obj{}
			newobj.multi = rInt(3, 11)
			wid := float32(newobj.multi * int(tilew))
			newobj.rec = rl.NewRectangle(rFloat32(0, innerr), ylevel, wid, tilew)
			newobj.activ = true
			newobj.objsin = make([]obj, newobj.multi)
			x := newobj.rec.X
			y := newobj.rec.Y

			for b := 0; b < len(newobj.objsin); b++ {
				newobj.objsin[b].img = currenttile
				newobj.objsin[b].rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.objsin[b].v2tl = rl.NewVector2(x, y)
				newobj.objsin[b].activ = true
				x += tilew
			}

			objs = append(objs, newobj)

		case 2: //platform 2 block
			newobj := obj{}
			newobj.multi = rInt(3, 11)
			wid := float32(newobj.multi * int(tilew))
			heig := tilew * 2
			newobj.rec = rl.NewRectangle(rFloat32(0, innerr), ylevel, wid, heig)

			newobj.activ = true
			newobj.objsin = make([]obj, newobj.multi*2)

			x := newobj.rec.X
			y := newobj.rec.Y
			count := 0
			for b := 0; b < len(newobj.objsin); b++ {
				newobj.objsin[b].img = currenttile
				newobj.objsin[b].rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.objsin[b].v2tl = rl.NewVector2(x, y)
				newobj.objsin[b].activ = true
				x += tilew
				count++
				if count == newobj.multi {
					x = newobj.rec.X
					y += tilew
				}
			}

			objs = append(objs, newobj)

		case 3: //steps right
			newobj := obj{}
			x := rFloat32(0, innerr)
			y := ylevel

			num := rInt(2, 8)

			for {
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.objsin = make([]obj, 1)
				newobj.objsin[0].img = currenttile
				newobj.objsin[0].rec = newobj.rec
				newobj.objsin[0].v2tl = rl.NewVector2(x, y)
				newobj.objsin[0].activ = true
				objs = append(objs, newobj)
				x += tilew
				y -= tilew / 2
				num--
				if num == 0 {
					break
				}
			}

		case 4: //steps left
			newobj := obj{}
			x := rFloat32(innerl, innerr)
			y := ylevel

			num := rInt(2, 8)

			for {
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.objsin = make([]obj, 1)
				newobj.objsin[0].img = currenttile
				newobj.objsin[0].rec = newobj.rec
				newobj.objsin[0].v2tl = rl.NewVector2(x, y)
				newobj.objsin[0].activ = true
				objs = append(objs, newobj)
				x -= tilew
				y -= tilew / 2
				num--
				if num == 0 {
					break
				}
			}

		case 5: //island

			num := rInt(5, 10)
			if num%2 == 0 {
				num++
			}

			x := scrwf32 / 2
			xorig := x
			y := ylevel

			newobj := obj{}
			newobj.direc.X = rFloat32(1, 8)
			newobj.complex = true
			newobj.activ = true
			newobj.objsin = make([]obj, 0)
			newobj.collisrecs = make([]rl.Rectangle, 0)

			count := 0

			for {
				newobjin := obj{}
				newobjin.img = currenttile
				newobjin.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobjin.v2tl = rl.NewVector2(x, y)
				newobjin.activ = true
				newobj.objsin = append(newobj.objsin, newobjin)
				newobj.collisrecs = append(newobj.collisrecs, newobjin.rec)
				count++
				x += tilew
				if count == num {
					x = xorig
					x += tilew
					xorig = x
					y += tilew
					count = 0
					num -= 2
				}
				if num < 0 {
					break
				}
			}

			objs = append(objs, newobj)

		case 6: //2 block gaps

			newobj := obj{}

			x := float32(0)
			y := ylevel

			for {
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.img = currenttile
				objs = append(objs, newobj)
				x += tilew
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.img = currenttile
				objs = append(objs, newobj)

				x += tilew * 3

				if x > scrwf32 {
					break
				}

			}
		case 7: //random block gaps

			newobj := obj{}

			x := float32(0)
			y := ylevel

			for {

				num := rInt(2, 5)
				for a := 0; a < num; a++ {
					newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
					x += tilew
				}

				x += tilew * rFloat32(2, 5)

				if x > scrwf32 {
					break
				}

			}
		case 8: // random shape block gaps

			newobj := obj{}

			x := float32(0)
			y := ylevel

			for {
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.img = currenttile
				objs = append(objs, newobj)
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y+tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y-tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				x += tilew
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.img = currenttile
				objs = append(objs, newobj)
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y+tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y-tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				x += tilew
				newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
				newobj.activ = true
				newobj.img = currenttile
				objs = append(objs, newobj)
				if flipcoin() {
					x += tilew
					newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y+tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}
				if flipcoin() {
					newobj.rec = rl.NewRectangle(x, y-tilew, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
				}

				x += tilew * rFloat32(3, 5)

				if x > scrwf32 {
					break
				}

			}

		case 9: //random steps

			x := float32(0)
			y := ylevel
			yorig := y

			if flipcoin() {
				x += rFloat32(tilew, tilew*3)
			}
			for {
				num := rInt(2, 5)
				updown := flipcoin()
				for a := 0; a < num; a++ {
					newobj := obj{}
					newobj.rec = rl.NewRectangle(x, y, tilew, tilew)
					newobj.activ = true
					newobj.img = currenttile
					objs = append(objs, newobj)
					if flipcoin() {
						newobj.rec.Y += tilew
						objs = append(objs, newobj)
					}
					x += tilew
					if updown {
						y -= rFloat32(tilew/4, tilew/2)
					} else {
						y += rFloat32(tilew/4, tilew/2)
					}
				}
				x += rFloat32(tilew, tilew*3)
				y = yorig

				if x > scrwf32 {
					break
				}
			}

		case 10: //moving square bloks

			num := rInt(2, 5)

			for a := 0; a < num; a++ {

				newobj := obj{}
				newobj.multi = rInt(2, 5)
				wid := float32(newobj.multi * int(tilew))

				newobj.direc.X = rFloat32(-10, 11)
				if flipcoin() {
					newobj.direc.Y = rFloat32(-5, 6)
				}

				newobj.rec = rl.NewRectangle(rFloat32(0, monwf32/1.2), ylevel, wid, wid)
				newobj.activ = true
				newobj.objsin = make([]obj, newobj.multi*newobj.multi)

				//newobj.grav = true

				x := newobj.rec.X
				y := newobj.rec.Y
				count := 0
				for b := 0; b < len(newobj.objsin); b++ {
					newobj.objsin[b].img = currenttile
					newobj.objsin[b].rec = rl.NewRectangle(x, y, tilew, tilew)
					newobj.objsin[b].v2tl = rl.NewVector2(x, y)
					newobj.objsin[b].activ = true

					x += tilew
					count++
					if count == newobj.multi {
						count = 0
						x = newobj.rec.X
						y += tilew
					}
				}

				objs = append(objs, newobj)

			}

		}

		ylevel -= tilew * 5
		if ylevel < bordt {
			break
		}

	}

}

func fx() { //MARK: fx

	if rain {

		for a := 0; a < len(raindrops); a++ {
			rl.DrawCircleV(raindrops[a].v2, raindrops[a].radius, rl.Fade(raindrops[a].color, rFloat32(0.2, 0.6)))

			v22 := rl.NewVector2(raindrops[a].v2.X, raindrops[a].v2.Y-(raindrops[a].radius*2))

			rl.DrawCircleV(v22, raindrops[a].radius-2, rl.Fade(raindrops[a].color, rFloat32(0.2, 0.6)))

			raindrops[a].v2.Y += rFloat32(10, 15)

			if raindrops[a].v2.Y > scrhf32 {
				raindrops[a].v2.Y = 0
			}
		}

	}

	if noise {
		num := 200
		for a := 0; a < num; a++ {
			w := rFloat32(1, 4)
			rec := rl.NewRectangle(rFloat32(0, monwf32), rFloat32(0, monhf32), w, w)
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, rFloat32(0.5, 1.1)))
		}
	}

	if scan {

		for a := 0; a < len(scanlines); a++ {
			rl.DrawLineEx(scanlines[a], scanlines2[a], 2, rl.Fade(rl.Black, 0.2))
			scanlines[a].Y++
			scanlines2[a].Y++

			if scanlines[a].Y > monhf32 {
				scanlines[a].Y = 0
			}
			if scanlines2[a].Y > monhf32 {
				scanlines2[a].Y = 0
			}
		}

	}

}
func raylib() { //MARK: raylib
	rl.SetConfigFlags(rl.FlagMsaa4xHint) // enable 4X anti-aliasing
	rl.InitWindow(scrw, scrh, "GAME TITLE")
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window

	//rl.ToggleFullscreen()
	//rl.HideCursor()
	imgs = rl.LoadTexture("imgs.png") // load images
	makeimgs()
	initial()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++
		mous = rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		cam()

		rl.EndMode2D()

		if intro {
			dintro()
		}
		nocam()
		//centerlines
		if cntr {
			rl.DrawLine(scrw/2, 0, scrw/2, scrh, rl.Magenta)
			rl.DrawLine(0, scrh/2, scrw, scrh/2, rl.Magenta)
		}
		if dev {
			ddev()
		}
		rl.EndDrawing()
		up()
	}
	rl.CloseWindow()

}
func makeintro() { //MARK: makeintro

	num := 300

	for {

		newcirc := introcirc{}
		if flipcoin() {
			newcirc.lr = true
			newcirc.v2.X = rFloat32(scrwf32, scrwf32*2)
		} else {
			newcirc.v2.X = -(rFloat32(0, scrwf32))
		}
		newcirc.v2.Y = rFloat32(0, scrhf32)
		newcirc.rad = float32(rInt(30, 51))
		if int(newcirc.rad)%2 != 0 {
			newcirc.rad++
		}
		newcirc.wid = newcirc.rad * 2
		newcirc.color = darkred()

		introcircs = append(introcircs, newcirc)

		num--
		if num == 0 {
			break
		}
	}

}
func dintro() { //MARK: dintro ██████████████████████████████

	rl.BeginMode2D(camintro)

	for a := 0; a < len(introcircs); a++ {

		rl.DrawCircleV(introcircs[a].v2, introcircs[a].rad, introcircs[a].color)

		if introcircs[a].lr {
			if introcircs[a].v2.X > -introcircs[a].wid {
				introcircs[a].v2.X -= 10
			}
			if introcircs[a].v2.X < scrwf32 {
				rl.DrawRectangle(int32(scrwf32-(scrwf32-introcircs[a].v2.X)), int32(introcircs[a].v2.Y-introcircs[a].rad+1), int32(scrwf32-introcircs[a].v2.X), int32(introcircs[a].wid), introcircs[a].color)
			}

		} else {
			if introcircs[a].v2.X < scrwf32+introcircs[a].wid {
				introcircs[a].v2.X += 10
			}
			if introcircs[a].v2.X > 0 {
				rl.DrawRectangle(0, int32(introcircs[a].v2.Y-introcircs[a].rad+1), int32(introcircs[a].v2.X), int32(introcircs[a].wid), introcircs[a].color)
			}
		}

	}
	if introtimer > 0 {
		introtimer--
	}
	rl.EndMode2D()
	if introtimer == 0 {
		xcntr := int32(camintrotxt.Target.X + scrwf32/2)
		ycntr := int32(camintrotxt.Target.Y + float32(scrh1/3))
		ycntr2 := ycntr - 100
		v2 := rl.NewVector2(float32(xcntr), float32(ycntr))

		xtl := int32(camintrotxt.Target.X)
		//ytl := int32(camintrotxt.Target.Y)

		camintrotxt.Target = v2
		camintrotxt.Offset.X = scrwf32 / 2
		camintrotxt.Offset.Y = float32(scrh1 / 3)

		rl.BeginMode2D(camintrotxt)

		txtlen := rl.MeasureText("spl@ta", introtxtsize)
		if introtxtsize > 200 {
			introtxtsize -= 2
			camintrotxt.Rotation += 20
			if rl.IsKeyPressed(rl.KeySpace) {
				introtxtsize = 200
			}
		} else {

			camintrotxt.Rotation = 0
			rl.DrawRectangle(0, 0, int32(scrwf32), int32(scrhf32), darkred())

			switch introsel {
			case 0:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+40, int32(monwf32), 50, randomred())
			case 1:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+80, int32(monwf32), 50, randomred())
			case 2:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+120, int32(monwf32), 50, randomred())
			case 3:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+160, int32(monwf32), 50, randomred())
			case 4:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+200, int32(monwf32), 50, randomred())
			case 5:
				rl.DrawRectangle(xtl, ycntr2+introtxtsize+240, int32(monwf32), 50, randomred())

			}

			rl.DrawText("spl@ta", xcntr-(txtlen/2)-12, ycntr2+12, introtxtsize, rl.Black)
			rl.DrawText("spl@ta", xcntr-(txtlen/2)-4, ycntr2+4, introtxtsize, randomred())

			txtlen2 := rl.MeasureText("1 player", 40)
			rl.DrawText("1 player", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+46, 40, rl.Black)
			rl.DrawText("1 player", xcntr-(txtlen2/2), ycntr2+introtxtsize+44, 40, rl.White)
			txtlen2 = rl.MeasureText("2 player", 40)
			rl.DrawText("2 player", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+86, 40, rl.Black)
			rl.DrawText("2 player", xcntr-(txtlen2/2), ycntr2+introtxtsize+84, 40, rl.White)
			txtlen2 = rl.MeasureText("settings", 40)
			rl.DrawText("settings", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+126, 40, rl.Black)
			rl.DrawText("settings", xcntr-(txtlen2/2), ycntr2+introtxtsize+124, 40, rl.White)
			txtlen2 = rl.MeasureText("help", 40)
			rl.DrawText("help", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+166, 40, rl.Black)
			rl.DrawText("help", xcntr-(txtlen2/2), ycntr2+introtxtsize+164, 40, rl.White)
			txtlen2 = rl.MeasureText("credits", 40)
			rl.DrawText("credits", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+206, 40, rl.Black)
			rl.DrawText("credits", xcntr-(txtlen2/2), ycntr2+introtxtsize+204, 40, rl.White)
			txtlen2 = rl.MeasureText("exit", 40)
			rl.DrawText("exit", xcntr-(txtlen2/2)-2, ycntr2+introtxtsize+246, 40, rl.Black)
			rl.DrawText("exit", xcntr-(txtlen2/2), ycntr2+introtxtsize+244, 40, rl.White)

		}

		rl.DrawText("spl@ta", xcntr-(txtlen/2), ycntr2, introtxtsize, rl.White)
		if ghost {

			rl.DrawText("spl@ta", xcntr-(txtlen/2)+rInt32(-4, 5), ycntr2+rInt32(-4, 5), introtxtsize, rl.Fade(rl.White, 0.4))

		}

		rl.EndMode2D()
	}
}
func inp() { //MARK: inp

	if !pause {
		if rl.IsKeyPressed(rl.KeyLeftControl) {
			fire()
		}
		if rl.IsKeyDown(rl.KeyKp6) {
			player.direc.X = playerspd
		}
		if rl.IsKeyDown(rl.KeyKp4) {
			player.direc.X = -playerspd
		}
		if rl.IsKeyDown(rl.KeyKp5) {
			player.direc.X = 0
		}
		if rl.IsKeyDown(rl.KeyKp8) {
			if player.jumptimer == 0 {
				player.direc.Y = -(playerspd * 2)
				player.jumptimer = int(fps / 2)
			}
		}
	}
	if intro && introtxtsize == 200 {
		if rl.IsKeyPressed(rl.KeyKp8) {
			introsel--
		}
		if rl.IsKeyPressed(rl.KeyKp2) {
			introsel++
		}
		if introsel < 0 {
			introsel = 5
		}
		if introsel > 5 {
			introsel = 0
		}
	}

	if rl.IsKeyPressed(rl.KeyF4) {
		explode(0)
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if dev {
			dev = false
		} else {
			dev = true
		}
	}

}
func timers() { //MARK: timers
	if frames%2 == 0 {
		onoff1 = true
	} else {
		onoff1 = false
	}
	if frames%3 == 0 {
		onoff3 = true
	} else {
		onoff3 = false
	}
	if frames%5 == 0 {
		onoff5 = true
	} else {
		onoff5 = false
	}
	if frames%10 == 0 {
		onoff10 = true
	} else {
		onoff10 = false
	}
	if frames%15 == 0 {
		onoff15 = true
	} else {
		onoff15 = false
	}
	if frames%20 == 0 {
		onoff20 = true
	} else {
		onoff20 = false
	}
	if frames%30 == 0 {
		onoff30 = true
	} else {
		onoff30 = false
	}
	if frames%60 == 0 {
		onoff60 = true
	} else {
		onoff60 = false
	}
	if frames%120 == 0 {
		onoff120 = true
	} else {
		onoff120 = false
	}
}
func upcams() {
	if player.rec.Y < scrhf32-(monhf32/2) {
		camera.Target.Y = player.rec.Y
		camera.Target.Y -= monhf32 / 2
	}

	camcntr = rl.NewVector2(camera.Target.X+(scrwf32/2), camera.Target.Y+float32(scrh1/2))
}
func up() { //MARK: up

	inp()
	if !pause {
		timers()
		upobjs()
		upplayer()
		upcams()
	}

}
func makeplayer(num int) { //MARK: makeplayer
	switch num {
	case 1:
		player.rec = rl.NewRectangle(scrwf32/2, bordb-(dinogl.Height*(tilemulti*2)), dinogl.Width*(tilemulti*2), dinogl.Height*(tilemulti*2))
		player.v2tl = rl.NewVector2(player.rec.X, player.rec.Y)
		player.img = dinogr
		player.currweapon = weapons[2]
		player.weaponnum = 2
		player.lr = true
	}
}
func makeimgs() { // MARK: makeimgs

	x := float32(4)
	y := float32(242)
	count := 0
	for a := 0; a < len(trees); a++ {

		trees[a] = rl.NewRectangle(x, y, 32, 32)
		x += 32
		count++
		if count == 4 {
			count = 0
			y += 32
			x = 0
		}

	}

}
func resize() { //MARK: resize

	player.rec = rl.NewRectangle(monwf32/2, monhf32-(tilew*2), dinogl.Width*(tilemulti*2), dinogl.Height*(tilemulti*2))

	for a := 0; a < len(objs); a++ {

		wid := float32(objs[a].multi * int(tilew))
		objs[a].rec.Width = wid
		objs[a].rec.Height = wid

		if len(objs[a].objsin) != 0 {
			for b := 0; b < len(objs[a].objsin); b++ {
				objs[a].objsin[b].rec.Width = wid
				objs[a].objsin[b].rec.Width = wid
			}
		}
	}
}
func scr(num int) { //MARK: scr
	switch num {
	case 7:
		scrw = 2160
		scrw1 = scrw
		scrh1 = 1440
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1440, 900)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 3
		weapmulti = 6
		tilew = tileimgsize * tilemulti
		resize()
	case 8:
		scrw = 3840
		scrw1 = scrw
		scrh1 = 2160
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1440, 900)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 4
		weapmulti = 8
		tilew = tileimgsize * tilemulti
		resize()
	case 4:
		scrw = 1440
		scrw1 = scrw
		scrh1 = 900
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1440, 900)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 1
		weapmulti = 2
		tilew = tileimgsize * tilemulti
		resize()
	case 5:
		scrw = 1366
		scrw1 = scrw
		scrh1 = 768
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1366, 768)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 1
		weapmulti = 2
		tilew = tileimgsize * tilemulti
		resize()
	case 6:
		scrw = 1280
		scrw1 = scrw
		scrh1 = 720
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1280, 720)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 1
		weapmulti = 2
		tilew = tileimgsize * tilemulti
		resize()
	case 3:
		scrw = 1920
		scrw1 = scrw
		scrh1 = 1080
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1920, 1080)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 2
		weapmulti = 4
		tilew = tileimgsize * tilemulti
		resize()
	case 2:
		scrw = 1600
		scrw1 = scrw
		scrh1 = 900
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

		rl.SetWindowSize(1600, 900)
		rl.SetWindowPosition(int(monw-scrw)/2, int(monh-scrh)/2)

		tilemulti = 2
		weapmulti = 4
		tilew = tileimgsize * tilemulti
		resize()

	case 1:
		scrh1 = int32(rl.GetScreenHeight())
		scrw = int32(rl.GetScreenWidth())
		scrw1 = scrw
		scrh = scrh1 * levelh
		camera.Zoom = 1.0

	}
	bordb = float32(scrh) - tilew
	bordt = float32(0)
	bordl = float32(0)
	bordr = float32(scrw)
}
func main() { //MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	scr(1)
	raylib()
}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(0, 255)))
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int32) int32 {
	return (rand.Int31() * (max - min)) + min

}
func rFloat32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
