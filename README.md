# Empirical proof that all go routines exits with along main go routine
This is an empirical proof that all go routines exits along with the main go routine.

## What is this?
Here you will find a main executable go file, which trigger a goroutine which writes an autoincrement number forever (inside a infinity loop) into a file called data.txt at the same directory.
Then the main go routine (even the main process that we trigger is a goroutine) sleep 20 seconds. That give you enough time to tail to the data.txt file and stop manually the main process or wait for it to finish.
When the main goroutine stops we can see that there are no new numbers into the data.txt file. That means that `all go routines exits with the main go routine`.

## Example
Example stopping the main goroutine manually
![Empirical proof](https://drive.google.com/uc?export=view&id=1JbKaTUfLhrvS3_eCcTcSlfO5Vo8G2K_4)

## Test it
open a new shell and move to the proyect's directory

    go run main.go

open a new shell and move to the proyect's directory

    tail -f data.txt

Optional: Stop the main goroutine.
