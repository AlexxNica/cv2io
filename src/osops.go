package main

import (
  "io"
	"io/ioutil"
  "os"
	"strconv"
	"time"
)

//createCV2 creates the .cv2 file with data provided in the *CV2 struct.
func (cv2 *CV2) createCV2() error {
	filename := "../output/" + cv2.Title + ".cv2"
	return ioutil.WriteFile(filename, cv2.Body, 0666)
}

//removeCV2 removes created output files after 20 minutes.
func (cv2 *CV2) removeCV2() []error {
	// ATTENTION: change back to time.Minute * 20 when going live
	timer := time.NewTimer(time.Second * 50)
  <- timer.C
	err := []error{os.Remove("../output/"+cv2.Title+".cv2"), os.Remove("../output/"+cv2.Title+".html"), os.Remove("../output/"+cv2.Title+".svg"),}
	return err
}

//Copy copies files from src to dst.
func Copy(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {return err}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {return err}
	defer out.Close()

	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {return err}
	return cerr
}

//createTimestamp creates a timestamp taking the current time for unique file naming.
func createTimestamp() string {
	yy := strconv.Itoa(time.Now().Year())
	mm := strconv.Itoa(int((time.Now().Month())))
	dd := strconv.Itoa(int((time.Now().Day())))
	hh := strconv.Itoa(time.Now().Hour())
	min := strconv.Itoa(time.Now().Minute())
	sec := strconv.Itoa(time.Now().Second())
	timestamp := yy + "-"+ mm + "-" + dd + "_" + hh + "-" + min + "-" + sec
	return timestamp
}
