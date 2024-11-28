# Text Counter CLI Tool

This is a simple command-line tool built in Go to analyze text files. The tool provides options to count the number of words, lines, and characters in a given file. It also includes a mode to display all counts simultaneously.

## Features
- Count the number of words in a file.
- Count the number of lines in a file.
- Count the number of characters (grapheme clusters) in a file.
- Display all counts (words, lines, and characters) at once.

# Installation and Usage
## Step 1: Clone the Repository

``` bash
git clone https://github.com/derajohnson/go-text-counter.git
cd go-text-counter

```
## Step 2: Build the Executable
Run the following command to build the program:

``` bash
go build -o counter

```

## Step 3: Move the Executable to `/usr/local/bin`
To make the `counter` tool globally accessible:

``` bash
sudo mv counter /usr/local/bin
```

You can now use the `counter` tool from anywhere on your system.

## Usage
Basic Syntax

``` bash
counter -path <path-to-file> -type <counter-type>
```

Flags
- -path (required): The path to the file to analyze.
- -type (optional): The type of counter to use. Possible values:
  words: Count the number of words.
  lines: Count the number of lines.
  characters: Count the number of characters.
  (empty): Display all counts (words, lines, and characters).


Examples

1. Count Words
   
``` bash
counter -path ./example.txt -type words
```

Output 

``` bash
./example.txt has 100 words
```

2. Count Lines
   
``` bash
counter -path ./example.txt -type lines
```

Output 

``` bash
./example.txt has 10 lines
```

3. Count Characters
   
``` bash
counter -path ./example.txt -type characters
```

Output 

``` bash
./example.txt has 500 characters
```

4. Display All Counts
   
```bash
counter -path ./example.txt
```

Output 

```bash
./example.txt has 500 characters, 10 lines, and 100 words
```


## Support

If you encounter any issues or have suggestions, please create an issue in the repository. Thank you.






