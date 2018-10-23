package main

import (
	"time"
	"text/tabwriter"
	"os"
	"fmt"
	"sort"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}


var tracks = []*Track{
	{"GO","Deliah","From the Rots up",2012,length("3m38s")},
	{"GO","Moby","Moby",1992,length("3m30s")},
	{"Go Ahead","Alica Keys","AS I AM",2007,length("4m36m")},
	{"Ready to go","Martin","Samsh",2011,length("2m39s")},

}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout,0,8,2,' ',0)
	fmt.Fprintf(tw,format,"Title","Artist","Album","Year","Length")
	fmt.Fprintf(tw,format,"-----", "-----", "-----", "-----","-----")
	for _,t := range tracks {
		fmt.Fprintf(tw,format,t.Title,t.Artist,t.Album,t.Year,t.Length)
	}
	tw.Flush()
}

type byArtist []*Track
func (x byArtist) Len() int   {
	return len(x)
}

func (x byArtist) Less(i,j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i,j int) {x[i],x[j] = x[j],x[i]}

type byYear []*Track

func (x byYear) Len() int   {
	return len(x)
}

func (x byYear) Less(i,j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byYear) Swap(i,j int) {x[i],x[j] = x[j],x[i]}


type customSort struct {
	t []*Track
	less func(x,y *Track) bool
}

func (x customSort) Len() int {return len(x.t)}
func (x customSort) Less(i,j int) bool { return x.less(x.t[i],x.t[j])}
func (x customSort) Swap(i,j int) {x.t[i],x.t[j] = x.t[j],x.t[i]}

func main() {
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("ReversebyArtist:")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("byYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	//!+customcall
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	//!-customcall
	printTracks(tracks)
}