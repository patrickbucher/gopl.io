package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type dirstats struct {
	root   string
	nfiles int64
	nbytes int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	rootStats := make(map[string]*dirstats)
	fileSizes := make(chan dirstats)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		rootStats[root] = &dirstats{root, 0, 0}
		go walkDir(root, &n, fileSizes, root)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var globalStats dirstats
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // filesSizes was closed
			}
			globalStats.nfiles++
			globalStats.nbytes += size.nbytes
			if stats, ok := rootStats[size.root]; ok {
				stats.nfiles++
				stats.nbytes += size.nbytes
			}
		case <-tick:
			printDiskUsage(globalStats, rootStats)
		}
	}
	printDiskUsage(globalStats, rootStats)
}

func printDiskUsage(global dirstats, roots map[string]*dirstats) {
	for k, v := range roots {
		fmt.Printf("%s  %d files  %.1f GB\n", k, v.nfiles, float64(v.nbytes)/1e9)
	}
	fmt.Printf("%d files  %.1f GB\n", global.nfiles, float64(global.nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- dirstats, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes, root)
		} else {
			fileSizes <- dirstats{root, 1, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
