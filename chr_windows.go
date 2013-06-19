package vfp
/*
Returns the character associated with the specified numeric unicode code.
*/
/*Example:

SetBellTo(`C:\Kugou\Listen\tn.mp3`)
Chr(7)
Wait()

*/
func Chr(zcode int) string {
	if zcode == 7 {
		if At("window", OS()) > 0 {
			if gsoundFile != "" {
				MCISendString("play " + Md5(gsoundFile))
			} else {
				PlaySound("xxxx", 0, 0)
			}
		}

	}
	return string(zcode)
}

/*Specifies a waveform sound to play when the bell is rung.
zWAVFileName can include a path to the waveform sound.
*/
func SetBellTo(zWAVFileName string) {
	zcmd := ""
	if gsoundFile != "" {
		zcmd := "close " + Md5(gsoundFile)
		MCISendString(zcmd)
	}
	gsoundFile = zWAVFileName

	zcmd = `open "` + gsoundFile + `" alias ` + Md5(gsoundFile)
	MCISendString(zcmd)
}