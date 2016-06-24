package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "regexp"
)

func main() {
    dir, err := os.Open("/Users/gdenslow/hackathon/smallset")
    if ( err != nil ) {
        fmt.Println(err)
        os.Exit(1)
    }
    dir_contents, err := dir.Readdir(-1);
    for _, item := range dir_contents {
        var filepath string = dir.Name() + "/" + item.Name()
        fileInfo, err := os.Stat( filepath )
        if err == nil {
            if ! fileInfo.IsDir() {
                parse_email( filepath );
            }
        }
    }
}

func parse_email( file string ) {
    filehandle, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer filehandle.Close();
    scanner := bufio.NewScanner(filehandle)
        scanner.Split(bufio.ScanLines)
    var this_date string
    var this_from string
    var this_subject string
    date_regexp := regexp.MustCompile("^Date:\\s+(.+)")
    from_regexp := regexp.MustCompile("^From:\\s+(.+)")
    subject_regexp := regexp.MustCompile("^Subject:\\s+(.+)")
    for scanner.Scan() {
        line := scanner.Text()
        if ( this_date == "" ) && date_regexp.MatchString( line ) {
            this_date = ( date_regexp.FindStringSubmatch( line ) )[1]
        }
        if ( this_from == "" ) && from_regexp.MatchString( line ) {
            this_from = ( from_regexp.FindStringSubmatch( line ) )[1]
        }
        if ( this_subject == "" ) && subject_regexp.MatchString( line ) {
            this_subject = ( subject_regexp.FindStringSubmatch( line ) )[1]
        }
    }
    fmt.Println(file + "|" + this_date + "|" + this_from + "|" + this_subject);
}
