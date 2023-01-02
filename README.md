# Cyclone Sort
Alternative to sort (GNU coreutils) that is cross compiled for Linux, Windows & Mac. Cyclone Sort by default will sort and dedup input. See usage examples below.

Usage Examples:
- sort -help _(displays help)_
- sort -version _(displays version info)_
- sort -i input.txt -o output.txt
- cat input.txt | sort > output.txt
- cat input.txt | sort _(prints to stdout)_
- _defaults to stdin if -i flag is not specified_
- _defaults to stdout if -o flag is not specified_

## Testing

### Linux:
- $ ./sort-x64-linux -i tmp.txt -o output.txt
- 2023/01/02 09:54:10 Reading file...
- 2023/01/02 09:54:17 Lines in: 16053567
- 2023/01/02 09:54:17 Lines out: 15053568
- 2023/01/02 09:54:17 Duplicates Removed: 999999
- 2023/01/02 09:54:17 Elapsed time: 6.41967925s
- 2023/01/02 09:54:17 Lines Per Second Processed: 2500681

### Windows:
- \> .\sort-x64.exe -i tmp.txt -o output.txt
- 2023/01/02 09:59:14 Reading file...
- 2023/01/02 09:59:21 Lines in: 16053567
- 2023/01/02 09:59:21 Lines out: 15053568
- 2023/01/02 09:59:21 Duplicates Removed: 999999
- 2023/01/02 09:59:21 Elapsed time: 7.0531078s
- 2023/01/02 09:59:21 Lines Per Second Processed: 2276098

### Mac:
- _untested_
