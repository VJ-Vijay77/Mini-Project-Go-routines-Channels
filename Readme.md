### Mini Project: Concurrent File Processor

#### Problem Statement:

Create a program that reads a list of file paths from a text file, concurrently processes each file, and calculates the total word count across all files. Each file should be processed in its own Goroutine.

#### Requirements:

1. Read a text file containing a list of file paths. Each line in the file represents the path to a text file.
2. Create a Goroutine for each file path to process the file concurrently. The processing should involve reading the file, counting the number of words, and sending the word count to the main Goroutine through a channel.
3. Use a channel to aggregate the word counts from each Goroutine.
4. Print the total word count after processing all files.
