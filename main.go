package main

import (
    "fmt"
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
    actorToMoviesDoneMap := readAndParseMovieFile("movies.txt")
    inputActor := os.Args[1]
    movieListPrinter(actorToMoviesDoneMap[inputActor], inputActor)
}

func readAndParseMovieFile(fileName string) (m map[string][]Movie){
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
        if scanner.Text() == "" { //blank line = end of parsing a movie
            newMovie := Movie{movieName, movieYear, directorName, actors}
            //for each actor, add movie to list of movies they have been in
            for _, actor := range actors {
                m[actor] = append(m[actor], newMovie)
            }
            loc = 0 //reset loc for next movie
        } else {
            if loc == 0 { //movie name
                movieName = scanner.Text()
            } else if loc == 1 { //year
                movieYear = scanner.Text() 
            } else if loc == 2 { //director
                directorName = scanner.Text()
            } else if loc == 3 { //actors
                actors = parseActors(scanner.Text())
            } else {
                log.Fatal("There is an issue with your input file.")
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

func movieListPrinter(movieList []Movie, actor string) {
    if len(movieList) > 1 {
        fmt.Println(len(movieList), "Movies Featuring", actor, "\n")
    }
    for _, movie := range movieList {
        fmt.Print("Title: " , movie.name, " (", movie.year, ")", "\n")
        fmt.Println("Directed By:", movie.director)
        fmt.Println("Also Starring:", actorListPrinter(movie.actors, actor), "\n")
    }
}

func actorListPrinter(actorList []string, actorToRemove string) (actorPrinter string) {
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
