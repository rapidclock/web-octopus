/*
Package adapter contains implementations of the OutputAdapter interface
of the octopus crawler.

This package contains two types of adapters StdOpAdapter and
FileWriterAdapter. The StdOpAdapter prints the depth and url to standard output (usually the
screen). The FileWriterAdapter prints the output to a specified File.

Both can be used as an OutputAdapter as part of the octopus crawler's
CrawlOptions.
*/
package adapter
