package main

import (
    "fmt"
    //"io/ioutil"
    "bufio"
    "os"
    "log"
    "strings"
)
type Movie struct {
    name string
    year string
    director string
    actors []string
}

func main() {
	//for each movie, make movie object, then go through all the actors, add m{"actor name"] = append movie for each actor.
    m := readAndParseFile("movies.txt")
    movieList := m[os.Args[1]]
    moviePrinter(movieList, os.Args[1])
	//fmt.Println(m["Tom Hanks"])
}

func readAndParseFile(fileName string) (m map[string][]Movie){
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    loc := 0
    var movieName string
    var movieYear string
    var directorName string
    var actors []string
	m = make(map[string][]Movie)
    for scanner.Scan() {
        if scanner.Text() == "" {
            newMovie := Movie{movieName, movieYear, directorName, actors}
            //for each actor, add movie
            for _, actor := range actors {
                m[actor] = append(m[actor], newMovie)
            }
            loc = 0
        } else {
            if loc == 0 { //movie name
                movieName = scanner.Text()
            } else if loc == 1 { //year
                movieYear = scanner.Text() 
            } else if loc == 2 { //director
                directorName = scanner.Text()
            } else if loc == 3 { //actors
                actors = parseActors(scanner.Text())
            }
            loc++
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return
}

func parseActors(actorString string) (actorList []string) {
    actorList = strings.Split(actorString, ", ")
    return
}

func moviePrinter(movieList []Movie, actor string) {
    if len(movieList) > 1 {
        fmt.Println(len(movieList), "Movies Featuring", actor, "\n")
    }
    for _, movie := range movieList {
        fmt.Print("Title: " , movie.name, " (", movie.year, ")", "\n") //TODO: Fix spacing in release year
        fmt.Println("Directed By:", movie.director)
        fmt.Println("Also Starring:", actorPrinter(movie.actors, actor), "\n") //TODO: remove given actor, format better
    }
}

func actorPrinter(actorList []string, actorToRemove string) (actorPrinter string) {
    firstIteration := true
    for _, actor := range actorList {
        if actorToRemove != actor {
            if firstIteration {
                firstIteration = false
                actorPrinter = actor
            } else {
                actorPrinter += ", " + actor
            }
        }
    }
    return
}
