# Cyclone Sort
Alternative to sort (GNU coreutils) that is cross compiled for Linux, Windows & Mac. Cyclone Sort by default will sort and dedup input. See usage examples below.

Usage Examples:
- cyclone_sort -help _(displays help)_
- cyclone_sort -version _(displays version info)_
- cyclone_sortsort -i input.txt -o output.txt
- cat input.txt | cyclone_sort > output.txt
- cat input.txt | cyclone_sort _(prints to stdout)_
- _defaults to stdin if -i flag is not specified_
- _defaults to stdout if -o flag is not specified_

## Testing

### Linux:
- $ ./cyclone_sort.bin -i tmp.txt -o output.txt
- 2023/01/02 09:54:10 Reading file...
- 2023/01/02 09:54:17 Lines in: 16053567
- 2023/01/02 09:54:17 Lines out: 15053568
- 2023/01/02 09:54:17 Duplicates Removed: 999999
- 2023/01/02 09:54:17 Elapsed time: 6.111929518s
- 2023/01/02 09:54:17 Lines Per Second Processed: 2626596

### Windows:
- \> .\cyclone_sort.exe -i tmp.txt -o output.txt
- 2023/01/02 09:59:14 Reading file...
- 2023/01/02 09:59:21 Lines in: 16053567
- 2023/01/02 09:59:21 Lines out: 15053568
- 2023/01/02 09:59:21 Duplicates Removed: 999999
- 2023/01/02 09:59:21 Elapsed time: 7.0531078s
- 2023/01/02 09:59:21 Lines Per Second Processed: 2276098

### Mac:
- _untested_
