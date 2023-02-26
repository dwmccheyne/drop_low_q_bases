# README.md
drop_low_bases is a program that filters low-quality bases from FASTQ files. It takes a quality score threshold, an input FASTQ file, and an output FASTQ file as positional arguments, and writes a new FASTQ file with bases below the quality threshold replaced with N's.

## Usage
```
drop_low_bases <quality_score> <input> <output>
```
Where:

- **quality_score** is the minimum quality score allowed for a base. Any bases with a quality score below this value will be replaced with N's.
- **input** is the input FASTQ file to filter.
- **output** is the output FASTQ file to write the filtered reads to.

## Example
To filter a FASTQ file input.fastq with a quality score threshold of 30 and write the filtered reads to a new file named output.fastq, run:

``` lua
drop_low_bases 30 input.fastq output.fastq
```

## Loop Example
To run the program with different quality score thresholds and output files, you can use a loop in Bash or another shell. For example, the following loop will run the program with quality score thresholds ranging from 12 to 42 and write the filtered reads to output files named 12_plus_output.fastq, 13_plus_output.fastq, and so on up to 42_plus_output.fastq:

``` bash
for i in $(seq 12 42); 
    do     
        ./drop_low_bases $i input.fastq ${i}_plus_output.fastq; 
    done
```

## Build
To build the program, you will need a working Go installation. Go can be downloaded from the official website at https://golang.org/doc/install. Once Go is installed, you can build the program with the following command:
``` go
go build
```
This will create an executable file named drop_low_bases in the current directory.