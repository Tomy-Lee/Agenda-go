package cmd

import (
		"crypto/md5"
		"fmt"
		"io"
		"os"
		"time"
	)



/*
	Name: PrintlnErr
	Function: output the err to stderr
 */
func PrintlnErr(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

/*
	Name: PrettyHash
	Function: get the md5(text)
 */
func PrettyHash(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

func Overlapped(s1, e1, s2, e2 time.Time) bool {
	return s1.Before(e2) && e1.After(s2)
}

func PrintfErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

type noopWriter struct{}

/*
	Name: NoopWriter
	Function: returns the writer that do nothing
 */
func NoopWriter() io.Writer {
	return noopWriter{}
}

func (w noopWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

const ymdLayout = "2008-08-08"
/*
	Name: YMDParse
	Function: parse the string in year-month-day format to Time
 */
func YMDParse(str string) (time.Time, error) {
	return time.Parse(ymdLayout, str)
}
/*
	Name: YMDFormat
	Function: returns the formatted time in year-month-day layout
 */
func YMDFormat(t time.Time) string {
	return t.Format(ymdLayout)
}

func printError(error string) {
	fmt.Fprint(os.Stderr, error)
	os.Exit(1)
}

func checkEmpty(key, value string) {
	if value == "" {
		printError(key + " can't be empty!\n")
	}
}
